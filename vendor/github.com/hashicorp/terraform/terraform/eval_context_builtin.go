package terraform

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/hashicorp/terraform/plans"
	"github.com/hashicorp/terraform/providers"
	"github.com/hashicorp/terraform/provisioners"
	"github.com/hashicorp/terraform/version"

	"github.com/hashicorp/terraform/states"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/hashicorp/terraform/configs/configschema"
	"github.com/hashicorp/terraform/lang"
	"github.com/hashicorp/terraform/tfdiags"

	"github.com/hashicorp/terraform/addrs"
	"github.com/zclconf/go-cty/cty"
)

// BuiltinEvalContext is an EvalContext implementation that is used by
// Terraform by default.
type BuiltinEvalContext struct {
	// StopContext is the context used to track whether we're complete
	StopContext context.Context

	// PathValue is the Path that this context is operating within.
	PathValue addrs.ModuleInstance

	// Evaluator is used for evaluating expressions within the scope of this
	// eval context.
	Evaluator *Evaluator

	// Schemas is a repository of all of the schemas we should need to
	// decode configuration blocks and expressions. This must be constructed by
	// the caller to include schemas for all of the providers, resource types,
	// data sources and provisioners used by the given configuration and
	// state.
	//
	// This must not be mutated during evaluation.
	Schemas *Schemas

	// VariableValues contains the variable values across all modules. This
	// structure is shared across the entire containing context, and so it
	// may be accessed only when holding VariableValuesLock.
	// The keys of the first level of VariableValues are the string
	// representations of addrs.ModuleInstance values. The second-level keys
	// are variable names within each module instance.
	VariableValues     map[string]map[string]cty.Value
	VariableValuesLock *sync.Mutex

	Components          contextComponentFactory
	Hooks               []Hook
	InputValue          UIInput
	ProviderCache       map[string]providers.Interface
	ProviderInputConfig map[string]map[string]cty.Value
	ProviderLock        *sync.Mutex
	ProvisionerCache    map[string]provisioners.Interface
	ProvisionerLock     *sync.Mutex
	ChangesValue        *plans.ChangesSync
	StateValue          *states.SyncState

	once sync.Once
}

// BuiltinEvalContext implements EvalContext
var _ EvalContext = (*BuiltinEvalContext)(nil)

func (ctx *BuiltinEvalContext) Stopped() <-chan struct{} {
	// This can happen during tests. During tests, we just block forever.
	if ctx.StopContext == nil {
		return nil
	}

	return ctx.StopContext.Done()
}

func (ctx *BuiltinEvalContext) Hook(fn func(Hook) (HookAction, error)) error {
	for _, h := range ctx.Hooks {
		action, err := fn(h)
		if err != nil {
			return err
		}

		switch action {
		case HookActionContinue:
			continue
		case HookActionHalt:
			// Return an early exit error to trigger an early exit
			log.Printf("[WARN] Early exit triggered by hook: %T", h)
			return EvalEarlyExitError{}
		}
	}

	return nil
}

func (ctx *BuiltinEvalContext) Input() UIInput {
	return ctx.InputValue
}

func (ctx *BuiltinEvalContext) InitProvider(typeName string, addr addrs.ProviderConfig) (providers.Interface, error) {
	ctx.once.Do(ctx.init)
	absAddr := addr.Absolute(ctx.Path())

	// If we already initialized, it is an error
	if p := ctx.Provider(absAddr); p != nil {
		return nil, fmt.Errorf("%s is already initialized", addr)
	}

	// Warning: make sure to acquire these locks AFTER the call to Provider
	// above, since it also acquires locks.
	ctx.ProviderLock.Lock()
	defer ctx.ProviderLock.Unlock()

	key := absAddr.String()

	p, err := ctx.Components.ResourceProvider(typeName, key)
	if err != nil {
		return nil, err
	}

	log.Printf("[TRACE] BuiltinEvalContext: Initialized %q provider for %s", typeName, absAddr)
	ctx.ProviderCache[key] = p

	return p, nil
}

func (ctx *BuiltinEvalContext) Provider(addr addrs.AbsProviderConfig) providers.Interface {
	ctx.once.Do(ctx.init)

	ctx.ProviderLock.Lock()
	defer ctx.ProviderLock.Unlock()

	return ctx.ProviderCache[addr.String()]
}

func (ctx *BuiltinEvalContext) ProviderSchema(addr addrs.AbsProviderConfig) *ProviderSchema {
	ctx.once.Do(ctx.init)

	return ctx.Schemas.ProviderSchema(addr.ProviderConfig.Type)
}

func (ctx *BuiltinEvalContext) CloseProvider(addr addrs.ProviderConfig) error {
	ctx.once.Do(ctx.init)

	ctx.ProviderLock.Lock()
	defer ctx.ProviderLock.Unlock()

	key := addr.Absolute(ctx.Path()).String()
	provider := ctx.ProviderCache[key]
	if provider != nil {
		delete(ctx.ProviderCache, key)
		return provider.Close()
	}

	return nil
}

func (ctx *BuiltinEvalContext) ConfigureProvider(addr addrs.ProviderConfig, cfg cty.Value) tfdiags.Diagnostics {
	var diags tfdiags.Diagnostics
	absAddr := addr.Absolute(ctx.Path())
	p := ctx.Provider(absAddr)
	if p == nil {
		diags = diags.Append(fmt.Errorf("%s not initialized", addr))
		return diags
	}

	providerSchema := ctx.ProviderSchema(absAddr)
	if providerSchema == nil {
		diags = diags.Append(fmt.Errorf("schema for %s is not available", absAddr))
		return diags
	}

	req := providers.ConfigureRequest{
		TerraformVersion: version.String(),
		Config:           cfg,
	}

	resp := p.Configure(req)
	return resp.Diagnostics
}

func (ctx *BuiltinEvalContext) ProviderInput(pc addrs.ProviderConfig) map[string]cty.Value {
	ctx.ProviderLock.Lock()
	defer ctx.ProviderLock.Unlock()

	if !ctx.Path().IsRoot() {
		// Only root module provider configurations can have input.
		return nil
	}

	return ctx.ProviderInputConfig[pc.String()]
}

func (ctx *BuiltinEvalContext) SetProviderInput(pc addrs.ProviderConfig, c map[string]cty.Value) {
	absProvider := pc.Absolute(ctx.Path())

	if !ctx.Path().IsRoot() {
		// Only root module provider configurations can have input.
		log.Printf("[WARN] BuiltinEvalContext: attempt to SetProviderInput for non-root module")
		return
	}

	// Save the configuration
	ctx.ProviderLock.Lock()
	ctx.ProviderInputConfig[absProvider.String()] = c
	ctx.ProviderLock.Unlock()
}

func (ctx *BuiltinEvalContext) InitProvisioner(n string) (provisioners.Interface, error) {
	ctx.once.Do(ctx.init)

	// If we already initialized, it is an error
	if p := ctx.Provisioner(n); p != nil {
		return nil, fmt.Errorf("Provisioner '%s' already initialized", n)
	}

	// Warning: make sure to acquire these locks AFTER the call to Provisioner
	// above, since it also acquires locks.
	ctx.ProvisionerLock.Lock()
	defer ctx.ProvisionerLock.Unlock()

	p, err := ctx.Components.ResourceProvisioner(n, "")
	if err != nil {
		return nil, err
	}

	ctx.ProvisionerCache[n] = p

	return p, nil
}

func (ctx *BuiltinEvalContext) Provisioner(n string) provisioners.Interface {
	ctx.once.Do(ctx.init)

	ctx.ProvisionerLock.Lock()
	defer ctx.ProvisionerLock.Unlock()

	return ctx.ProvisionerCache[n]
}

func (ctx *BuiltinEvalContext) ProvisionerSchema(n string) *configschema.Block {
	ctx.once.Do(ctx.init)

	return ctx.Schemas.ProvisionerConfig(n)
}

func (ctx *BuiltinEvalContext) CloseProvisioner(n string) error {
	ctx.once.Do(ctx.init)

	ctx.ProvisionerLock.Lock()
	defer ctx.ProvisionerLock.Unlock()

	prov := ctx.ProvisionerCache[n]
	if prov != nil {
		return prov.Close()
	}

	return nil
}

func (ctx *BuiltinEvalContext) EvaluateBlock(body hcl.Body, schema *configschema.Block, self addrs.Referenceable, keyData InstanceKeyEvalData) (cty.Value, hcl.Body, tfdiags.Diagnostics) {
	var diags tfdiags.Diagnostics
	scope := ctx.EvaluationScope(self, keyData)
	body, evalDiags := scope.ExpandBlock(body, schema)
	diags = diags.Append(evalDiags)
	val, evalDiags := scope.EvalBlock(body, schema)
	diags = diags.Append(evalDiags)
	return val, body, diags
}

func (ctx *BuiltinEvalContext) EvaluateExpr(expr hcl.Expression, wantType cty.Type, self addrs.Referenceable) (cty.Value, tfdiags.Diagnostics) {
	scope := ctx.EvaluationScope(self, EvalDataForNoInstanceKey)
	return scope.EvalExpr(expr, wantType)
}

func (ctx *BuiltinEvalContext) EvaluationScope(self addrs.Referenceable, keyData InstanceKeyEvalData) *lang.Scope {
	data := &evaluationStateData{
		Evaluator:       ctx.Evaluator,
		ModulePath:      ctx.PathValue,
		InstanceKeyData: keyData,
		Operation:       ctx.Evaluator.Operation,
	}
	return ctx.Evaluator.Scope(data, self)
}

func (ctx *BuiltinEvalContext) Path() addrs.ModuleInstance {
	return ctx.PathValue
}

func (ctx *BuiltinEvalContext) SetModuleCallArguments(n addrs.ModuleCallInstance, vals map[string]cty.Value) {
	ctx.VariableValuesLock.Lock()
	defer ctx.VariableValuesLock.Unlock()

	childPath := n.ModuleInstance(ctx.PathValue)
	key := childPath.String()

	args := ctx.VariableValues[key]
	if args == nil {
		args = make(map[string]cty.Value)
		ctx.VariableValues[key] = vals
		return
	}

	for k, v := range vals {
		args[k] = v
	}
}

func (ctx *BuiltinEvalContext) Changes() *plans.ChangesSync {
	return ctx.ChangesValue
}

func (ctx *BuiltinEvalContext) State() *states.SyncState {
	return ctx.StateValue
}

func (ctx *BuiltinEvalContext) init() {
}

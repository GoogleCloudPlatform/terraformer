package resource

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform/addrs"
	"github.com/zclconf/go-cty/cty"

	"github.com/hashicorp/terraform/configs/hcl2shim"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/hashicorp/terraform/states"
	"github.com/hashicorp/terraform/terraform"
)

// shimState takes a new *states.State and reverts it to a legacy state for the provider ACC tests
func shimNewState(newState *states.State, providers map[string]terraform.ResourceProvider) (*terraform.State, error) {
	state := terraform.NewState()

	// in the odd case of a nil state, let the helper packages handle it
	if newState == nil {
		return nil, nil
	}

	for _, newMod := range newState.Modules {
		mod := state.AddModule(newMod.Addr)

		for name, out := range newMod.OutputValues {
			outputType := ""
			val := hcl2shim.ConfigValueFromHCL2(out.Value)
			ty := out.Value.Type()
			switch {
			case ty == cty.String:
				outputType = "string"
			case ty.IsTupleType() || ty.IsListType():
				outputType = "list"
			case ty.IsMapType():
				outputType = "map"
			}

			mod.Outputs[name] = &terraform.OutputState{
				Type:      outputType,
				Value:     val,
				Sensitive: out.Sensitive,
			}
		}

		for _, res := range newMod.Resources {
			resType := res.Addr.Type
			providerType := res.ProviderConfig.ProviderConfig.Type

			resource := getResource(providers, providerType.LegacyString(), res.Addr)

			for key, i := range res.Instances {
				resState := &terraform.ResourceState{
					Type:     resType,
					Provider: res.ProviderConfig.String(),
				}

				// We should always have a Current instance here, but be safe about checking.
				if i.Current != nil {
					flatmap, err := shimmedAttributes(i.Current, resource)
					if err != nil {
						return nil, fmt.Errorf("error decoding state for %q: %s", resType, err)
					}

					var meta map[string]interface{}
					if i.Current.Private != nil {
						err := json.Unmarshal(i.Current.Private, &meta)
						if err != nil {
							return nil, err
						}
					}

					resState.Primary = &terraform.InstanceState{
						ID:         flatmap["id"],
						Attributes: flatmap,
						Tainted:    i.Current.Status == states.ObjectTainted,
						Meta:       meta,
					}

					if i.Current.SchemaVersion != 0 {
						if resState.Primary.Meta == nil {
							resState.Primary.Meta = map[string]interface{}{}
						}
						resState.Primary.Meta["schema_version"] = i.Current.SchemaVersion
					}

					for _, dep := range i.Current.DependsOn {
						resState.Dependencies = append(resState.Dependencies, dep.String())
					}

					// convert the indexes to the old style flapmap indexes
					idx := ""
					switch key.(type) {
					case addrs.IntKey:
						// don't add numeric index values to resources with a count of 0
						if len(res.Instances) > 1 {
							idx = fmt.Sprintf(".%d", key)
						}
					case addrs.StringKey:
						idx = "." + key.String()
					}

					mod.Resources[res.Addr.String()+idx] = resState
				}

				// add any deposed instances
				for _, dep := range i.Deposed {
					flatmap, err := shimmedAttributes(dep, resource)
					if err != nil {
						return nil, fmt.Errorf("error decoding deposed state for %q: %s", resType, err)
					}

					var meta map[string]interface{}
					if dep.Private != nil {
						err := json.Unmarshal(dep.Private, &meta)
						if err != nil {
							return nil, err
						}
					}

					deposed := &terraform.InstanceState{
						ID:         flatmap["id"],
						Attributes: flatmap,
						Tainted:    dep.Status == states.ObjectTainted,
						Meta:       meta,
					}
					if dep.SchemaVersion != 0 {
						deposed.Meta = map[string]interface{}{
							"schema_version": dep.SchemaVersion,
						}
					}

					resState.Deposed = append(resState.Deposed, deposed)
				}
			}
		}
	}

	return state, nil
}

func getResource(providers map[string]terraform.ResourceProvider, providerName string, addr addrs.Resource) *schema.Resource {
	p := providers[providerName]
	if p == nil {
		panic(fmt.Sprintf("provider %q not found in test step", providerName))
	}

	// this is only for tests, so should only see schema.Providers
	provider := p.(*schema.Provider)

	switch addr.Mode {
	case addrs.ManagedResourceMode:
		resource := provider.ResourcesMap[addr.Type]
		if resource != nil {
			return resource
		}
	case addrs.DataResourceMode:
		resource := provider.DataSourcesMap[addr.Type]
		if resource != nil {
			return resource
		}
	}

	panic(fmt.Sprintf("resource %s not found in test step", addr.Type))
}

func shimmedAttributes(instance *states.ResourceInstanceObjectSrc, res *schema.Resource) (map[string]string, error) {
	flatmap := instance.AttrsFlat
	if flatmap != nil {
		return flatmap, nil
	}

	// if we have json attrs, they need to be decoded
	rio, err := instance.Decode(res.CoreConfigSchema().ImpliedType())
	if err != nil {
		return nil, err
	}

	instanceState, err := res.ShimInstanceStateFromValue(rio.Value)
	if err != nil {
		return nil, err
	}

	return instanceState.Attributes, nil
}

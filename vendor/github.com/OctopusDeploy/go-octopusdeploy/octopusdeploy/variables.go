package octopusdeploy

import (
	"fmt"

	"github.com/dghubble/sling"
	"gopkg.in/go-playground/validator.v9"
)

type VariableService struct {
	sling *sling.Sling
}

func NewVariableService(sling *sling.Sling) *VariableService {
	return &VariableService{
		sling: sling,
	}
}

type Variables struct {
	ID          string      `json:"Id"`
	OwnerID     string      `json:"OwnerId"`
	Version     int         `json:"Version"`
	Variables   []Variable  `json:"Variables"`
	ScopeValues ScopeValues `json:"ScopeValues"`
	Links       map[string]string
}

type Variable struct {
	ID          string                 `json:"Id"`
	Name        string                 `json:"Name"`
	Value       string                 `json:"Value"`
	Description string                 `json:"Description"`
	Scope       *VariableScope         `json:"Scope,omitempty"`
	IsEditable  bool                   `json:"IsEditable"`
	Prompt      *VariablePromptOptions `json:"Prompt"`
	Type        string                 `json:"Type"`
	IsSensitive bool                   `json:"IsSensitive"`
}

type VariableScope struct {
	Project     []string `json:"Project,omitempty"`
	Environment []string `json:"Environment,omitempty"`
	Machine     []string `json:"Machine,omitempty"`
	Role        []string `json:"Role,omitempty"`
	TargetRole  []string `json:"TargetRole,omitempty"`
	Action      []string `json:"Action,omitempty"`
	User        []string `json:"User,omitempty"`
	Private     []string `json:"Private,omitempty"`
	Channel     []string `json:"Channel,omitempty"`
	TenantTag   []string `json:"TenantTag,omitempty"`
	Tenant      []string `json:"Tenant,omitempty"`
}

type VariablePromptOptions struct {
	Label       string `json:"Label"`
	Description string `json:"Description"`
	Required    bool   `json:"Required"`
}

type ScopeValues struct {
	Environments []ScopeValue `json:"Environments"`
	Machines     []ScopeValue `json:"Machines"`
	Actions      []ScopeValue `json:"Actions"`
	Roles        []ScopeValue `json:"Roles"`
	Channels     []ScopeValue `json:"Channels"`
	TenantTags   []ScopeValue `json:"TenantTags"`
}

type ScopeValue struct {
	ID   string `json:"Id"`
	Name string `json:"Name"`
}

func (t *Variable) Validate() error {
	validate := validator.New()

	err := validate.Struct(t)

	if err != nil {
		return err
	}

	return nil
}

func NewVariable(name, valuetype, value, description string, scope *VariableScope, sensitive bool) *Variable {
	return &Variable{
		Name:        name,
		Value:       value,
		Description: description,
		Type:        valuetype,
		IsSensitive: sensitive,
		Scope:       scope,
	}
}

// GetAll fetches an entire VariableSet from Octopus Deploy for a given Project ID.
func (s *VariableService) GetAll(projectid string) (*Variables, error) {
	if projectid == "" { //Not specifying the Project ID can return thousands of entries consuming hundreds of megs of memory
		return nil, fmt.Errorf("projectid must be specified")
	}

	path := fmt.Sprintf("variables/variableset-%s", projectid)
	resp, err := apiGet(s.sling, new(Variables), path)

	if err != nil {
		return nil, err
	}

	return resp.(*Variables), nil
}

// GetByID fetches a single variable, located by its ID, from Octopus Deploy for a given Project ID.
func (s *VariableService) GetByID(projectid, variableid string) (*Variable, error) {
	variables, err := s.GetAll(projectid)
	if err != nil {
		return nil, err
	}

	for _, variable := range variables.Variables {
		if variable.ID == variableid {
			return &variable, nil
		}
	}

	return nil, nil
}

// GetByName fetches variables, located by their name, from Octopus Deploy for a given Project ID. As variable
// names can appear more than once under different scopes, a VariableScope must also be provided, which will
// be used to locate the appropriate variables.
func (s *VariableService) GetByName(projectid, variablename string, scope *VariableScope) ([]Variable, error) {
	variables, err := s.GetAll(projectid)
	if err != nil {
		return nil, err
	}

	var matchedVariables []Variable

	for _, variable := range variables.Variables {
		if variable.Name == variablename {
			matchScope, _, err := s.MatchesScope(variable.Scope, scope)
			if err != nil {
				return nil, err
			}
			if matchScope {
				matchedVariables = append(matchedVariables, variable)
			}
		}
	}

	return matchedVariables, nil
}

// AddSingle adds a single variable to a project ID. This automates the act of fetching
// the variable set, adding a new item to it, and posting back to Octopus
func (s *VariableService) AddSingle(projectid string, variable *Variable) (*Variables, error) {
	variables, err := s.GetAll(projectid)
	if err != nil {
		return nil, err
	}
	variables.Variables = append(variables.Variables, *variable)
	return s.Update(projectid, variables)
}

// UpdateSingle adds a single variable to a project ID. This automates the act of fetching
// the variable set, updating the existing item, and posting back to Octopus
func (s *VariableService) UpdateSingle(projectid string, variable *Variable) (*Variables, error) {
	variables, err := s.GetAll(projectid)
	if err != nil {
		return nil, err
	}

	var found bool
	for i, existingVar := range variables.Variables {
		if existingVar.ID == variable.ID {
			variables.Variables[i] = *variable
			found = true
		}
	}

	if !found {
		return nil, ErrItemNotFound
	}

	return s.Update(projectid, variables)
}

// DeleteSingle removes a single variable from a project ID. This automates the act of fetching
// the variable set, removing the existing item, and posting back to Octopus
func (s *VariableService) DeleteSingle(projectid string, variableID string) (*Variables, error) {
	variables, err := s.GetAll(projectid)
	if err != nil {
		return nil, err
	}

	var found bool
	for i, existingVar := range variables.Variables {
		if existingVar.ID == variableID {
			variables.Variables = append(variables.Variables[:i], variables.Variables[i+1:]...)
			found = true
		}
	}

	if !found {
		return nil, ErrItemNotFound
	}

	return s.Update(projectid, variables)
}

// Update takes an entire variable set and posts the entire set back to Octopus Deploy. There are individual
// functions like AddSingle and UpdateSingle that can make this process more of a "typical" CRUD Octopus command.
func (s *VariableService) Update(projectid string, variableSet *Variables) (*Variables, error) {
	path := fmt.Sprintf("variables/variableset-%s", projectid)
	resp, err := apiUpdate(s.sling, variableSet, new(Variables), path)

	if err != nil {
		return nil, err
	}

	return resp.(*Variables), nil
}

// MatchesScope compares two different scopes to see if they match. Generally used for comparing the scope of
// an existing variable against a desired state. Only supports Environment, Role, Machine, Action and Channel
// for scope options. Returns true if definedScope is nil or all elements are empty. Also returns a VariableScope
// of all the scopes that were matched
func (s *VariableService) MatchesScope(variableScope, definedScope *VariableScope) (bool, *VariableScope, error) {
	var matchedScopes VariableScope
	var matched bool

	//If the scope supplied is nil then match everything
	if definedScope == nil {
		return true, &matchedScopes, nil
	}

	//Unsupported scopes
	if len(definedScope.Private) > 0 {
		return false, nil, fmt.Errorf("Private is not a supported scope for variable matching")
	}
	if len(definedScope.Project) > 0 {
		return false, nil, fmt.Errorf("Project is not a supported scope for variable matching")
	}
	if len(definedScope.TargetRole) > 0 {
		return false, nil, fmt.Errorf("TargetRole is not a supported scope for variable matching")
	}
	if len(definedScope.Tenant) > 0 {
		return false, nil, fmt.Errorf("Tenant is not a supported scope for variable matching")
	}
	if len(definedScope.User) > 0 {
		return false, nil, fmt.Errorf("User is not a supported scope for variable matching")
	}

	//If there is no scope to filter on return all the results
	if len(definedScope.Environment) > 0 && len(definedScope.Role) > 0 && len(definedScope.Machine) > 0 && len(definedScope.Action) > 0 && len(definedScope.Channel) > 0 && len(definedScope.TenantTag) > 0 {
		return true, &matchedScopes, nil
	}

	for _, e1 := range definedScope.Environment {
		for _, e2 := range variableScope.Environment {
			if e1 == e2 {
				matched = true
				matchedScopes.Environment = append(matchedScopes.Environment, e1)
			}
		}
	}

	for _, r1 := range definedScope.Role {
		for _, r2 := range variableScope.Role {
			if r1 == r2 {
				matched = true
				matchedScopes.Role = append(matchedScopes.Role, r1)
			}
		}
	}

	for _, m1 := range definedScope.Machine {
		for _, m2 := range variableScope.Machine {
			if m1 == m2 {
				matched = true
				matchedScopes.Machine = append(matchedScopes.Machine, m1)
			}
		}
	}

	for _, a1 := range definedScope.Action {
		for _, a2 := range variableScope.Action {
			if a1 == a2 {
				matched = true
				matchedScopes.Action = append(matchedScopes.Action, a1)
			}
		}
	}

	for _, c1 := range definedScope.Channel {
		for _, c2 := range variableScope.Channel {
			if c1 == c2 {
				matched = true
				matchedScopes.Channel = append(matchedScopes.Channel, c1)
			}
		}
	}

	for _, c1 := range definedScope.TenantTag {
		for _, c2 := range variableScope.TenantTag {
			if c1 == c2 {
				matched = true
				matchedScopes.TenantTag = append(matchedScopes.TenantTag, c1)
			}
		}
	}

	return matched, &matchedScopes, nil
}

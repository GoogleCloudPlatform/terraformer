// Automatically generated, do not edit

package commercetools

// GraphQLError is a standalone struct
type GraphQLError struct {
	Path      []interface{}          `json:"path"`
	Message   string                 `json:"message"`
	Locations []GraphQLErrorLocation `json:"locations"`
}

func (obj GraphQLError) Error() string {
	return obj.Message
}

// GraphQLErrorLocation is a standalone struct
type GraphQLErrorLocation struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

// GraphQLRequest is a standalone struct
type GraphQLRequest struct {
	Variables     *GraphQLVariablesMap `json:"variables,omitempty"`
	Query         string               `json:"query"`
	OperationName string               `json:"operationName,omitempty"`
}

// GraphQLResponse is a standalone struct
type GraphQLResponse struct {
	Errors []GraphQLError `json:"errors,omitempty"`
	Data   interface{}    `json:"data,omitempty"`
}

// GraphQLVariablesMap is a standalone struct
type GraphQLVariablesMap struct {
}

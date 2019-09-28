package fc

type Header map[string]string

type Query struct {
	Prefix    *string
	StartKey  *string
	NextToken *string
	Limit     *int32
}

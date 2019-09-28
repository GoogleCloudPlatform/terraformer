package sls

const (
	//EtlMetaURI is etl meta uri
	EtlMetaURI = "etlmetas"
	//EtlMetaNameURI is etl meta name uri
	EtlMetaNameURI = "etlmetanames"
	//EtlMetaAllTagMatch is for search meta without tag filtering
	EtlMetaAllTagMatch = "__all_etl_meta_tag_match__"
)

type EtlMeta struct {
	MetaName  string            `json:"etlMetaName"`
	MetaKey   string            `json:"etlMetaKey"`
	MetaTag   string            `json:"etlMetaTag"`
	MetaValue map[string]string `json:"etlMetaValue"`
}

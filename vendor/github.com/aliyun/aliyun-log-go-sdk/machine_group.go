package sls

// const MachineIDTypes
const (
	MachineIDTypeIP          = "ip"
	MachineIDTypeUserDefined = "userdefined"
)

// MachinGroupAttribute defines machine group attribute
type MachinGroupAttribute struct {
	ExternalName string `json:"externalName"`
	TopicName    string `json:"groupTopic"`
}

// MachineGroup defines machine group
type MachineGroup struct {
	Name           string               `json:"groupName"`
	Type           string               `json:"groupType"`
	MachineIDType  string               `json:"machineIdentifyType"`
	MachineIDList  []string             `json:"machineList"`
	Attribute      MachinGroupAttribute `json:"groupAttribute"`
	CreateTime     uint32               `json:"createTime,omitempty"`
	LastModifyTime uint32               `json:"lastModifyTime,omitempty"`
}

// Machine defines machine struct
type Machine struct {
	IP            string
	UniqueID      string `json:"machine-uniqueid"`
	UserdefinedID string `json:"userdefined-id"`
}

// MachineList defines machine list
type MachineList struct {
	Total    int
	Machines []*Machine
}

package terraform_utils

import (
	"reflect"
	"testing"
)

func prepare(ID, resourceType string, attributes map[string]string) Resource {
	r := NewResource(ID, "name-"+resourceType, resourceType, "provider", attributes, []string{}, map[string]interface{}{})
	r.InstanceState.Attributes["id"] = r.InstanceState.ID
	r.ConvertTFstate()
	return r
}

func TestConnectServices(t *testing.T) {
	type args struct {
		importResources     map[string][]Resource
		resourceConnections map[string]map[string][]string
	}
	tests := []struct {
		name string
		args args
		want map[string][]Resource
	}{
		{
			name: "simple test",
			args: args{
				importResources: map[string][]Resource{
					"type1": {prepare("ID1", "type1", map[string]string{
						"type2_ref": "ID2",
					})},
					"type2": {prepare("ID2", "type2", map[string]string{})},
				},
				resourceConnections: map[string]map[string][]string{
					"type1": {
						"type2": {"type2_ref", "id"},
					},
				},
			},
			want: map[string][]Resource{
				"type1": {prepare("ID1", "type1", map[string]string{
					"type2_ref": "${data.terraform_remote_state.type2.outputs.type2_name-type2_id}",
				})},
				"type2": {prepare("ID2", "type2", map[string]string{})},
			},
		},
		{
			name: "many refs test",
			args: args{
				importResources: map[string][]Resource{
					"type1": {prepare("ID1", "type1", map[string]string{
						"type2_ref1": "ID2",
						"type2_ref2": "ID2",
					})},
					"type2": {prepare("ID2", "type2", map[string]string{})},
				},
				resourceConnections: map[string]map[string][]string{
					"type1": {
						"type2": {
							"type2_ref1", "id",
							"type2_ref2", "id",
						},
					},
				},
			},
			want: map[string][]Resource{
				"type1": {prepare("ID1", "type1", map[string]string{
					"type2_ref1": "${data.terraform_remote_state.type2.outputs.type2_name-type2_id}",
					"type2_ref2": "${data.terraform_remote_state.type2.outputs.type2_name-type2_id}",
				})},
				"type2": {prepare("ID2", "type2", map[string]string{})},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConnectServices(tt.args.importResources, tt.args.resourceConnections); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConnectServices() = %v, want %v", got, tt.want)
			}
		})
	}
}

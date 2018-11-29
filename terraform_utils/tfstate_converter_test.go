package terraform_utils

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/hashicorp/hil"
	"github.com/hashicorp/terraform/flatmap"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/assert"
)

type convertTest struct {
	name         string
	dataFilePath string
	expect       []TerraformResource
	metaData     map[string]ResourceMetaData
}

func unknownValue() string {
	return hil.UnknownValue
}

func TestFeildReader(t *testing.T) {
	data, _ := ioutil.ReadFile("test_data/test6.json")
	tfState, _ := terraform.ReadStateV3(data)
	m := flatmap.Expand(tfState.Modules[0].Resources["google_compute_firewall.resource-id"].Primary.Attributes, "lifecycle_rule")

	log.Println(m)

}

func TestBasicConvert(t *testing.T) {
	runConvert(convertTest{
		dataFilePath: "test1.json",
		name:         "basic tfstate",
		expect: []TerraformResource{
			{
				ResourceName: "resource-id",
				ResourceType: "google_compute_firewall",
				ID:           "resource-id",
				Item: map[string]interface{}{
					"direction":      "INGRESS",
					"enable_logging": false,
					"id":             "resource-id",
					"name":           "resource-name",
				},
				Provider: "google",
			},
		},
		metaData: map[string]ResourceMetaData{
			"resource-id": {
				Provider: "google",
			},
		},
	}, t)
}

func TestBasicTfstate2(t *testing.T) {
	runConvert(convertTest{
		dataFilePath: "test2.json",
		name:         "basic tfstate 2",
		expect: []TerraformResource{
			{
				ResourceName: "resource-idA",
				ResourceType: "google_compute_firewall",
				ID:           "resource-idA",
				Item: map[string]interface{}{
					"direction":      "INGRESS",
					"enable_logging": false,
					"id":             "resource-idA",
					"name":           "resource-nameA",
				},
				Provider: "google",
			},
			{
				ResourceName: "resource-idB",
				ResourceType: "google_compute_firewall",
				ID:           "resource-idB",
				Item: map[string]interface{}{
					"direction":      "INGRESS",
					"enable_logging": false,
					"id":             "resource-idB",
					"name":           "resource-nameB",
				},
				Provider: "google",
			},
		},
		metaData: map[string]ResourceMetaData{
			"resource-idB": {
				Provider: "google",
			},
			"resource-idA": {
				Provider: "google",
			},
		},
	}, t)
}

func TestBasicArray(t *testing.T) {
	runConvert(convertTest{
		dataFilePath: "test3.json",
		name:         "basic array",
		expect: []TerraformResource{
			{
				ResourceName: "resource-id",
				ResourceType: "google_compute_firewall",
				ID:           "resource-id",
				Item: map[string]interface{}{
					"direction":      "INGRESS",
					"enable_logging": false,
					"id":             "resource-id",
					"name":           "resource-name",
					"myarray": []interface{}{
						map[string]interface{}{
							"value1": "value1",
							"value2": "value2",
						},
						map[string]interface{}{
							"value3": "value3",
							"value4": "value4",
						},
					},
				},
				Provider: "google",
			},
		},
		metaData: map[string]ResourceMetaData{
			"resource-id": {
				Provider: "google",
			},
		},
	}, t)
}

func TestBasicArray2(t *testing.T) {
	runConvert(convertTest{
		dataFilePath: "test4.json",
		name:         "basic array 2",
		expect: []TerraformResource{
			{
				ResourceName: "resource-id",
				ResourceType: "google_compute_firewall",
				ID:           "resource-id",
				Item: map[string]interface{}{
					"direction":      "INGRESS",
					"enable_logging": false,
					"id":             "resource-id",
					"name":           "resource-name",
					"myarray": []interface{}{
						map[string]interface{}{
							"subarray1": []string{"value1", "value2"},
						},
						map[string]interface{}{
							"subarray3": []string{"value3"},
							"subarray4": "value4",
						},
					},
				},
				Provider: "google",
			},
		},
		metaData: map[string]ResourceMetaData{
			"resource-id": {
				Provider: "google",
			},
		},
	}, t)
}

func TestBasicArray3(t *testing.T) {
	runConvert(convertTest{
		dataFilePath: "test5.json",
		name:         "basic array 3",
		expect: []TerraformResource{
			{
				ResourceName: "resource-id",
				ResourceType: "google_compute_firewall",
				ID:           "resource-id",
				Item: map[string]interface{}{
					"direction":      "INGRESS",
					"enable_logging": false,
					"id":             "resource-id",
					"name":           "resource-name",
					"myarray":        []interface{}{"value1", "value2", "value3"},
					"myarray2": []interface{}{
						map[string]interface{}{
							"subarray3": map[string]interface{}{
								"subsubarray": "value3",
							},
						},
						map[string]interface{}{
							"subarray4": "value4",
						},
					},
				},
				Provider: "google",
			},
		},
		metaData: map[string]ResourceMetaData{
			"resource-id": {
				Provider: "google",
			},
		},
	}, t)
}

func TestBasicArray4(t *testing.T) {
	runConvert(convertTest{
		dataFilePath: "test6.json",
		name:         "basic array 4",
		expect: []TerraformResource{
			{
				ResourceName: "resource-id",
				ResourceType: "google_compute_firewall",
				ID:           "resource-id",
				Item: map[string]interface{}{
					"direction":      "INGRESS",
					"enable_logging": false,
					"id":             "resource-id",
					"name":           "resource-name",
					"lifecycle_rule": []interface{}{
						map[string]interface{}{
							"action": []interface{}{
								map[string]interface{}{
									"storage_class": "",
									"type":          "Delete",
								},
							},
						},
						map[string]interface{}{
							"condition": []interface{}{
								map[string]interface{}{
									"age":                "1",
									"created_before":     "",
									"is_live":            false,
									"num_newer_versions": "0",
								},
							},
						},
					},
				},
				Provider: "google",
			},
		},
		metaData: map[string]ResourceMetaData{
			"resource-id": {
				Provider: "google",
				AllowEmptyValue: map[string]bool{
					"storage_class":  true,
					"created_before": true,
				},
			},
		},
	}, t)
}

func runConvert(testCase convertTest, t *testing.T) {
	c := TfstateConverter{}
	actual, err := c.Convert("test_data/"+testCase.dataFilePath, testCase.metaData)
	if err != nil {
		t.Error(err)
	}
	if !assert.ObjectsAreEqual(testCase.expect, actual) {
		assert.Equal(t, testCase.expect, actual, "Convert error "+testCase.name)
	}
}

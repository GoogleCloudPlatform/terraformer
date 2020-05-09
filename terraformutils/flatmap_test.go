package terraformutils

import (
	"regexp"
	"testing"

	"github.com/zclconf/go-cty/cty"
)

func TestNestedAttributeFiltering(t *testing.T) {
	attributes := map[string]string{
		"attribute":        "value1",
		"nested.attribute": "value2",
	}

	ignoreKeys := []*regexp.Regexp{
		regexp.MustCompile(`^attribute$`),
	}
	parser := NewFlatmapParser(attributes, ignoreKeys, []*regexp.Regexp{})

	attributesType := cty.Object(map[string]cty.Type{
		"attribute": cty.String,
		"nested": cty.Object(map[string]cty.Type{
			"attribute": cty.String,
		}),
	})

	result, _ := parser.Parse(attributesType)

	if _, ok := result["attribute"]; ok {
		t.Errorf("failed to resolve %v", result)
	}
	if val, ok := result["nested"].(map[string]interface{})["attribute"]; !ok && val != "value2" {
		t.Errorf("failed to resolve %v", result)
	}
}

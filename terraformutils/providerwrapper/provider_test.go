package providerwrapper //nolint

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/configs/configschema"
	"github.com/zclconf/go-cty/cty"
)

func TestIgnoredAttributes(t *testing.T) {
	attributes := map[string]*configschema.Attribute{
		"computed_attribute": {
			Type:     cty.Number,
			Computed: true,
		},
		"required_attribute": {
			Type:     cty.String,
			Required: true,
		},
	}

	testCases := map[string]struct {
		block                map[string]*configschema.NestedBlock
		ignoredAttributes    []string
		notIgnoredAttributes []string
	}{
		"nesting_set": {map[string]*configschema.NestedBlock{
			"attribute_one": {
				Block: configschema.Block{
					Attributes: attributes,
				},
				Nesting: configschema.NestingSet,
			},
		}, []string{"nesting_set.attribute_one.computed_attribute"},
			[]string{"nesting_set.attribute_one.required_attribute"}},
		"nesting_list": {map[string]*configschema.NestedBlock{
			"attribute_one": {
				Block: configschema.Block{
					Attributes: map[string]*configschema.Attribute{},
					BlockTypes: map[string]*configschema.NestedBlock{
						"attribute_two_nested": {
							Nesting: configschema.NestingList,
							Block: configschema.Block{
								Attributes: attributes,
							},
						},
					},
				},
				Nesting: configschema.NestingList,
			},
		}, []string{"nesting_list.0.attribute_one.0.attribute_two_nested.computed_attribute"},
			[]string{"nesting_list.0.attribute_one.0.attribute_two_nested.required_attribute"}},
	}

	for key, tc := range testCases {
		t.Run(key, func(t *testing.T) {
			provider := ProviderWrapper{}
			readOnlyAttributes := provider.readObjBlocks(tc.block, []string{}, key)
			for _, attr := range tc.ignoredAttributes {
				if ignored := isAttributeIgnored(attr, readOnlyAttributes); !ignored {
					t.Errorf("attribute \"%s\" was not ignored. Pattern list: %s", attr, readOnlyAttributes)
				}
			}

			for _, attr := range tc.notIgnoredAttributes {
				if ignored := isAttributeIgnored(attr, readOnlyAttributes); ignored {
					t.Errorf("attribute \"%s\" was ignored. Pattern list: %s", attr, readOnlyAttributes)
				}
			}
		})
	}
}

func isAttributeIgnored(name string, patterns []string) bool {
	ignored := false
	for _, pattern := range patterns {
		if match, _ := regexp.MatchString(pattern, name); match {
			ignored = true
			break
		}
	}
	return ignored
}

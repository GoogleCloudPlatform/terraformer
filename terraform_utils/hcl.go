package terraform_utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/hcl/hcl/ast"
	hcl_printer "github.com/hashicorp/hcl/hcl/printer"
	hcl_parcer "github.com/hashicorp/hcl/json/parser"
)

const safeChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"

// sanitizer fixes up an invalid HCL AST, as produced by the HCL parser for JSON
type astSanitizer struct {
}

type TerraformResource struct {
	ResourceType string
	ResourceName string
	Item         interface{}
	ID           string
	Provider     string
	Attributes   map[string]string
}

func NewTerraformResource(ID, resourceName, resourceType, provider string, item interface{}, attributes map[string]string) TerraformResource {
	return TerraformResource{
		ResourceType: resourceType,
		ResourceName: TfSanitize(resourceName),
		Item:         item,
		ID:           ID,
		Provider:     provider,
		Attributes:   attributes,
	}
}

// output prints creates b printable HCL output and returns it.
func (v *astSanitizer) visit(n interface{}) {
	switch t := n.(type) {
	case *ast.File:
		v.visit(t.Node)
	case *ast.ObjectList:
		var index int
		for {
			if index == len(t.Items) {
				break
			}

			v.visit(t.Items[index])
			index++
		}
	case *ast.ObjectKey:
	case *ast.ObjectItem:
		v.visitObjectItem(t)
	case *ast.LiteralType:
	case *ast.ListType:
	case *ast.ObjectType:
		v.visit(t.List)
	default:
		fmt.Printf(" unknown type: %T\n", n)
	}

}

func (v *astSanitizer) visitObjectItem(o *ast.ObjectItem) {
	for i, k := range o.Keys {
		if i == 0 {
			text := k.Token.Text
			if text != "" && text[0] == '"' && text[len(text)-1] == '"' {
				v := text[1 : len(text)-1]
				safe := true
				for _, c := range v {
					if strings.IndexRune(safeChars, c) == -1 {
						safe = false
						break
					}
				}
				if safe {
					k.Token.Text = v
				}
			}

		}
	}

	// A hack so that Assign.IsValid is true, so that the printer will output =
	o.Assign.Line = 1

	v.visit(o.Val)
}

func hclPrint(node ast.Node) ([]byte, error) {
	var sanitizer astSanitizer
	sanitizer.visit(node)

	var b bytes.Buffer
	err := hcl_printer.Fprint(&b, node)
	if err != nil {
		return nil, fmt.Errorf("error writing HCL: %v", err)
	}
	s := b.String()

	// Remove extra whitespace...
	s = strings.Replace(s, "\n\n", "\n", -1)

	// ...but leave whitespace between resources
	s = strings.Replace(s, "}\nresource", "}\n\nresource", -1)

	// Workaround HCL insanity #6359: quotes are _not_ escaped in quotes (huh?)
	// This hits the file function
	s = strings.Replace(s, "(\\\"", "(\"", -1)
	s = strings.Replace(s, "\\\")", "\")", -1)

	// We don't need to escape > or <
	s = strings.Replace(s, "\\u003c", "<", -1)
	s = strings.Replace(s, "\\u003e", ">", -1)

	// Apply Terraform style (alignment etc.)
	formatted, err := hcl_printer.Format([]byte(s))
	if err != nil {
		log.Println("Invalid HCL follows:")
		for i, line := range strings.Split(s, "\n") {
			fmt.Printf("%d\t%s", (i + 1), line)
		}
		return nil, fmt.Errorf("error formatting HCL: %v", err)
	}

	return formatted, nil
}

func TfSanitize(name string) string {
	name = strings.Replace(name, ".", "-", -1)
	name = strings.Replace(name, "/", "--", -1)
	return name
}

func HclPrint(resources []TerraformResource, region, provider string) ([]byte, error) {
	resourcesByType := make(map[string]map[string]interface{})

	for _, res := range resources {
		resources := resourcesByType[res.ResourceType]
		if resources == nil {
			resources = make(map[string]interface{})
			resourcesByType[res.ResourceType] = resources
		}

		tfName := TfSanitize(res.ResourceName)

		if resources[tfName] != nil {
			return []byte{}, fmt.Errorf("duplicate resource found: %s.%s", res.ResourceType, tfName)
		}

		resources[tfName] = res.Item
	}

	data := make(map[string]interface{})
	data["resource"] = resourcesByType
	switch provider {
	case "google":
		data["provider"] = NewGcpRegionResource(region)
	case "aws":
		data["provider"] = NewAwsRegionResource(region)
	}

	var err error
	dataJsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return []byte{}, fmt.Errorf("error marshalling terraform data to json: %v", err)
	}
	nodes, err := hcl_parcer.Parse(dataJsonBytes)
	if err != nil {
		return []byte{}, fmt.Errorf("error parsing terraform json: %v", err)
	}
	hclBytes, err := hclPrint(nodes)
	if err != nil {
		return []byte{}, err
	}
	return hclBytes, nil
}

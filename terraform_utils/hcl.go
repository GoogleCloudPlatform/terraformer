// Copyright 2018 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package terraform_utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/hashicorp/hcl2/hclwrite"

	"github.com/hashicorp/hcl/hcl/ast"
	hclPrinter "github.com/hashicorp/hcl/hcl/printer"
	hclParcer "github.com/hashicorp/hcl/json/parser"
)

// Copy code from https://github.com/kubernetes/kops project with few changes for support many provider and heredoc

const safeChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"

var unsafeChars = regexp.MustCompile(`[^0-9A-Za-z_]`)

// sanitizer fixes up an invalid HCL AST, as produced by the HCL parser for JSON
type astSanitizer struct{}

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
					if !strings.ContainsRune(safeChars, c) {
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
	switch t := o.Val.(type) {
	case *ast.LiteralType: // heredoc support
		if strings.HasPrefix(t.Token.Text, `"<<`) {
			t.Token.Text = t.Token.Text[1:]
			t.Token.Text = t.Token.Text[:len(t.Token.Text)-1]
			t.Token.Text = strings.Replace(t.Token.Text, `\n`, "\n", -1)
			t.Token.Text = strings.Replace(t.Token.Text, `\t`, "", -1)
			t.Token.Type = 10
			// check if text json for Unquote and Indent
			tmp := map[string]interface{}{}
			jsonTest := t.Token.Text
			lines := strings.Split(jsonTest, "\n")
			jsonTest = strings.Join(lines[1:len(lines)-1], "\n")
			jsonTest = strings.Replace(jsonTest, "\\\"", "\"", -1)
			// it's json we convert to heredoc back
			err := json.Unmarshal([]byte(jsonTest), &tmp)
			if err == nil {
				dataJsonBytes, err := json.MarshalIndent(tmp, "", "  ")
				if err == nil {
					jsonData := strings.Split(string(dataJsonBytes), "\n")
					// first line for heredoc
					jsonData = append([]string{lines[0]}, jsonData...)
					// last line for heredoc
					jsonData = append(jsonData, lines[len(lines)-1])
					hereDoc := strings.Join(jsonData, "\n")
					t.Token.Text = hereDoc
				}
			}
		}
	default:
	}

	// A hack so that Assign.IsValid is true, so that the printer will output =
	o.Assign.Line = 1

	v.visit(o.Val)
}

func HclPrint(data interface{}, mapsObjects map[string]struct{}) ([]byte, error) {
	dataJsonBytes, err := json.MarshalIndent(data, "", "  ")
	dataJson := string(dataJsonBytes)
	dataJson = strings.Replace(dataJson, "\\u003c", "<", -1)
	if err != nil {
		log.Println(dataJson)
		return []byte{}, fmt.Errorf("error marshalling terraform data to json: %v", err)
	}

	nodes, err := hclParcer.Parse([]byte(dataJson))
	if err != nil {
		log.Println(dataJson)
		return []byte{}, fmt.Errorf("error parsing terraform json: %v", err)
	}
	var sanitizer astSanitizer
	sanitizer.visit(nodes)

	var b bytes.Buffer
	err = hclPrinter.Fprint(&b, nodes)
	if err != nil {
		return nil, fmt.Errorf("error writing HCL: %v", err)
	}
	s := b.String()

	// Remove extra whitespace...
	s = strings.Replace(s, "\n\n", "\n", -1)

	// ...but leave whitespace between resources
	s = strings.Replace(s, "}\nresource", "}\n\nresource", -1)

	// We don't need to escape > or <
	s = strings.Replace(s, "\\u003c", "<", -1)
	s = strings.Replace(s, "\\u003e", ">", -1)

	// Apply Terraform style (alignment etc.)
	formatted := hclwrite.Format([]byte(s))
	formatted, err = hclPrinter.Format([]byte(s))
	// hack for support terraform 0.12
	formatted = terraform12Adjustments(formatted, mapsObjects)
	if err != nil {
		log.Println("Invalid HCL follows:")
		for i, line := range strings.Split(s, "\n") {
			fmt.Printf("%4d|\t%s\n", i+1, line)
		}
		return nil, fmt.Errorf("error formatting HCL: %v", err)
	}

	return formatted, nil
}

func terraform12Adjustments(formatted []byte, mapsObjects map[string]struct{}) []byte {
	s := string(formatted)
	old := " = {"
	new := " {"
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		if !strings.Contains(line, old) {
			continue
		}
		key := strings.Trim(strings.Split(line, old)[0], " ")
		if _, exist := mapsObjects[key]; exist {
			continue
		}
		lines[i] = strings.Replace(line, old, new, -1)
	}
	s = strings.Join(lines, "\n")
	return []byte(s)
}

func escapeRune(s string) string {
	return fmt.Sprintf("-%04X-", s)
}

// Sanitize name for terraform style
func TfSanitize(name string) string {
	name = unsafeChars.ReplaceAllStringFunc(name, escapeRune)
	name = "tfer--" + name
	return name
}

// Print hcl file from TerraformResource + provider
func HclPrintResource(resources []Resource, providerData map[string]interface{}) ([]byte, error) {
	resourcesByType := map[string]map[string]interface{}{}
	mapsObjects := map[string]struct{}{}
	for _, res := range resources {

		r := resourcesByType[res.InstanceInfo.Type]
		if r == nil {
			r = make(map[string]interface{})
			resourcesByType[res.InstanceInfo.Type] = r
		}

		if r[res.ResourceName] != nil {
			log.Println(resources)
			return []byte{}, fmt.Errorf("[ERR]: duplicate resource found: %s.%s", res.InstanceInfo.Type, res.ResourceName)
		}

		r[res.ResourceName] = res.Item

		for k := range res.InstanceState.Attributes {
			if strings.HasSuffix(k, ".%") {
				t := strings.Split(k, ".")
				key := t[len(t)-2]
				mapsObjects[key] = struct{}{}
			}
		}
	}

	data := map[string]interface{}{}
	if len(resourcesByType) > 0 {
		data["resource"] = resourcesByType
	}
	if len(providerData) > 0 {
		data["provider"] = providerData
	}
	var err error

	hclBytes, err := HclPrint(data, mapsObjects)
	if err != nil {
		return []byte{}, err
	}
	return hclBytes, nil
}

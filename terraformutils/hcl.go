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

package terraformutils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform/configs/hcl2shim"
	"log"
	"regexp"
	"strings"

	"github.com/hashicorp/hcl/hcl/ast"
	hclPrinter "github.com/hashicorp/hcl/hcl/printer"
	hclParser "github.com/hashicorp/hcl/json/parser"
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
			t.Token.Text = strings.ReplaceAll(t.Token.Text, `\n`, "\n")
			t.Token.Text = strings.ReplaceAll(t.Token.Text, `\t`, "")
			t.Token.Type = 10
			// check if text json for Unquote and Indent
			tmp := map[string]interface{}{}
			jsonTest := t.Token.Text
			lines := strings.Split(jsonTest, "\n")
			jsonTest = strings.Join(lines[1:len(lines)-1], "\n")
			jsonTest = strings.ReplaceAll(jsonTest, "\\\"", "\"")
			// it's json we convert to heredoc back
			err := json.Unmarshal([]byte(jsonTest), &tmp)
			if err == nil {
				dataJSONBytes, err := json.MarshalIndent(tmp, "", "  ")
				if err == nil {
					jsonData := strings.Split(string(dataJSONBytes), "\n")
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

func Print(data interface{}, mapsObjects map[string]struct{}, format string, isResourceFile bool) ([]byte, error) {
	switch format {
	case "hcl":
		return hclPrint(data, mapsObjects, isResourceFile)
	case "json":
		return jsonPrint(data)
	}
	return []byte{}, errors.New("error: unknown output format")
}

func hclPrint(data interface{}, mapsObjects map[string]struct{}, isResourceFile bool) ([]byte, error) {
	dataBytesJSON, err := jsonPrint(data)
	if err != nil {
		return dataBytesJSON, err
	}
	dataJSON := string(dataBytesJSON)
	nodes, err := hclParser.Parse([]byte(dataJSON))
	if err != nil {
		log.Println(dataJSON)
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
	s = strings.ReplaceAll(s, "\n\n", "\n")

	// ...but leave whitespace between resources
	s = strings.ReplaceAll(s, "}\nresource", "}\n\nresource")

	// Apply Terraform style (alignment etc.)
	formatted, err := hclPrinter.Format([]byte(s))
	if err != nil {
		return nil, err
	}
	if isResourceFile {
		// hack for support terraform 0.12
		formatted = terraform12Adjustments(formatted, mapsObjects)
	}
	// hack for support terraform 0.15
	formatted = terraform15Adjustments1(formatted)
	formatted = terraform15Adjustments2(formatted)

	return formatted, nil
}

func terraform12Adjustments(formatted []byte, mapsObjects map[string]struct{}) []byte {
	singletonListFix := regexp.MustCompile(`^\s*\w+ = {`)
	singletonListFixEnd := regexp.MustCompile(`^\s*}`)

	s := string(formatted)
	old := " = {"
	newEquals := " {"
	lines := strings.Split(s, "\n")
	prefix := make([]string, 0)
	for i, line := range lines {
		if singletonListFixEnd.MatchString(line) && len(prefix) > 0 {
			prefix = prefix[:len(prefix)-1]
			continue
		}
		if !singletonListFix.MatchString(line) {
			continue
		}
		key := strings.Trim(strings.Split(line, old)[0], " ")
		prefix = append(prefix, key)
		if _, exist := mapsObjects[strings.Join(prefix, ".")]; exist {
			continue
		}
		lines[i] = strings.ReplaceAll(line, old, newEquals)
	}
	s = strings.Join(lines, "\n")
	return []byte(s)
}

func terraform15Adjustments1(formatted []byte) []byte {
	s := string(formatted)
	oldTerraformConfiguration := "terraform = {"
	oldRequiredProviders := "  required_providers = {"
	newTerraformConfiguration := "terraform {"
	newRequiredProviders := "  required_providers {"
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		if line == oldTerraformConfiguration {
			lines[i] = newTerraformConfiguration
		}
		if line == oldRequiredProviders {
			lines[i] = newRequiredProviders
		}
	}
	s = strings.Join(lines, "\n")
	return []byte(s)
}

func terraform15Adjustments2(formatted []byte) []byte {
	s := string(formatted)
	lines := strings.Split(s, "\n")
	var newLines []string
	for _, line := range lines {
		if !strings.Contains(line, "= {}") && !strings.Contains(line, "= []") {
			newLines = append(newLines, line)
		}
	}
	return []byte(strings.Join(newLines, "\n"))
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
func HclPrintResource(resources []Resource, providerData map[string]interface{}, output string) ([]byte, error) {
	resourcesByType := map[string]map[string]interface{}{}
	mapsObjects := map[string]struct{}{}
	indexRe := regexp.MustCompile(`\.[0-9]+`)
	for _, res := range resources {
		r := resourcesByType[res.Address.Type]
		if r == nil {
			r = make(map[string]interface{})
			resourcesByType[res.Address.Type] = r
		}

		if r[res.Address.Name] != nil {
			log.Println(resources)
			log.Printf("[ERR]: duplicate resource found: %s", res.Address.String())
			continue
		}

		var ignoreKeys []*regexp.Regexp
		for _, pattern := range res.IgnoreKeys {
			ignoreKeys = append(ignoreKeys, regexp.MustCompile(pattern))
		}
		var allowEmptyValues []*regexp.Regexp
		for _, pattern := range res.AllowEmptyValues {
			if pattern != "" {
				allowEmptyValues = append(allowEmptyValues, regexp.MustCompile(pattern))
			}
		}

		flattenedAttributes := hcl2shim.FlatmapValueFromHCL2(res.InstanceState.Value)
		parser := NewFlatmapParser(flattenedAttributes, ignoreKeys, allowEmptyValues)
		parsedAttributes, err := parser.Parse(res.InstanceState.Value.Type())
		if err != nil {
			return []byte{}, err
		}
		r[res.Address.Name] = parsedAttributes

		for k := range res.InstanceState.Value.AsValueMap() {
			if strings.HasSuffix(k, ".%") {
				key := strings.TrimSuffix(k, ".%")
				mapsObjects[indexRe.ReplaceAllString(key, "")] = struct{}{}
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

	hclBytes, err := Print(data, mapsObjects, output, true)
	if err != nil {
		return []byte{}, err
	}
	return hclBytes, nil
}

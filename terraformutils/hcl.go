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
	"log"
	"regexp"
	"sort"
	"strings"

	"github.com/hashicorp/hcl/hcl/ast"
	hclPrinter "github.com/hashicorp/hcl/hcl/printer"
	hclParser "github.com/hashicorp/hcl/json/parser"
)

// Copy code from https://github.com/kubernetes/kops project with few changes for support many provider and heredoc

const safeChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"

var unsafeChars = regexp.MustCompile(`[^0-9A-Za-z_\-]`)

// make HCL output reproducible by sorting the AST nodes
func sortHclTree(tree interface{}) {
	switch t := tree.(type) {
	case []*ast.ObjectItem:
		sort.Slice(t, func(i, j int) bool {
			var bI, bJ bytes.Buffer
			_, _ = hclPrinter.Fprint(&bI, t[i]), hclPrinter.Fprint(&bJ, t[j])
			return bI.String() < bJ.String()
		})
	case []ast.Node:
		sort.Slice(t, func(i, j int) bool {
			var bI, bJ bytes.Buffer
			_, _ = hclPrinter.Fprint(&bI, t[i]), hclPrinter.Fprint(&bJ, t[j])
			return bI.String() < bJ.String()
		})
	default:
	}
}

// sanitizer fixes up an invalid HCL AST, as produced by the HCL parser for JSON
type astSanitizer struct {
	sort bool
}

// output prints creates b printable HCL output and returns it.
func (v *astSanitizer) visit(n interface{}) {
	switch t := n.(type) {
	case *ast.File:
		v.visit(t.Node)
	case *ast.ObjectList:
		var index int
		if v.sort {
			sortHclTree(t.Items)
		}
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
		if v.sort {
			sortHclTree(t.List)
		}
	case *ast.ObjectType:
		if v.sort {
			sortHclTree(t.List)
		}
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
				if strings.HasPrefix(v, "--") { // if the key starts with "--", we must quote it. Seen in aws_glue_job.default_arguments parameter
					v = fmt.Sprintf(`"%s"`, v)
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
			jsonTest := t.Token.Text
			lines := strings.Split(jsonTest, "\n")
			jsonTest = strings.Join(lines[1:len(lines)-1], "\n")
			jsonTest = strings.ReplaceAll(jsonTest, "\\\"", "\"")
			// it's json we convert to heredoc back
			var tmp interface{} = map[string]interface{}{}
			err := json.Unmarshal([]byte(jsonTest), &tmp)
			if err != nil {
				tmp = make([]interface{}, 0)
				err = json.Unmarshal([]byte(jsonTest), &tmp)
			}
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
	case *ast.ListType:
		if v.sort {
			sortHclTree(t.List)
		}
	default:
	}

	// A hack so that Assign.IsValid is true, so that the printer will output =
	o.Assign.Line = 1

	v.visit(o.Val)
}

func Print(data interface{}, mapsObjects map[string]struct{}, format string, sort bool) ([]byte, error) {
	switch format {
	case "hcl":
		return hclPrint(data, mapsObjects, sort)
	case "json":
		return jsonPrint(data)
	}
	return []byte{}, errors.New("error: unknown output format")
}

func hclPrint(data interface{}, mapsObjects map[string]struct{}, sort bool) ([]byte, error) {
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
	sanitizer.sort = sort
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
	// hack for support terraform 0.12
	formatted = terraform12Adjustments(formatted, mapsObjects)
	// hack for support terraform 0.13
	formatted = terraform13Adjustments(formatted)
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

func terraform13Adjustments(formatted []byte) []byte {
	s := string(formatted)
	requiredProvidersRe := regexp.MustCompile("required_providers \".*\" {")
	endBraceRe := regexp.MustCompile(`^\s*}`)
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		if requiredProvidersRe.MatchString(line) {
			parts := strings.Split(strings.TrimSpace(line), " ")
			provider := strings.ReplaceAll(parts[1], "\"", "")
			lines[i] = "\trequired_providers {"
			var innerBlock []string
			inner := i + 1
			for ; !endBraceRe.MatchString(lines[inner]); inner++ {
				innerBlock = append(innerBlock, "\t"+lines[inner])
			}
			lines[i+1] = "\t\t" + provider + " = {\n" + strings.Join(innerBlock, "\n") + "\n\t\t}"
			lines = append(lines[:i+2], lines[inner:]...)
			break
		}
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
func HclPrintResource(resources []Resource, providerData map[string]interface{}, output string, sort bool) ([]byte, error) {
	resourcesByType := map[string]map[string]interface{}{}
	mapsObjects := map[string]struct{}{}
	indexRe := regexp.MustCompile(`\.[0-9]+`)
	for _, res := range resources {
		r := resourcesByType[res.InstanceInfo.Type]
		if r == nil {
			r = make(map[string]interface{})
			resourcesByType[res.InstanceInfo.Type] = r
		}

		if r[res.ResourceName] != nil {
			log.Println(resources)
			log.Printf("[ERR]: duplicate resource found: %s.%s", res.InstanceInfo.Type, res.ResourceName)
			continue
		}

		r[res.ResourceName] = res.Item

		for k := range res.InstanceState.Attributes {
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

	hclBytes, err := Print(data, mapsObjects, output, sort)
	if err != nil {
		return []byte{}, err
	}
	return hclBytes, nil
}

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

package panos

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/PaloAltoNetworks/pango"
	"golang.org/x/text/secure/precis"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func Initialize() (interface{}, error) {
	return pango.Connect(pango.Client{
		CheckEnvironment: true,
	})
}

func GetVsysList() ([]string, interface{}, error) {
	client, err := Initialize()
	if err != nil {
		return []string{}, nil, err
	}

	switch c := client.(type) {
	case *pango.Panorama:
		return []string{"shared"}, pango.Panorama{}, nil
	case *pango.Firewall:
		var vsysList []string
		vsysList, err = c.Vsys.GetList()
		return vsysList, pango.Firewall{}, err
	}

	return []string{}, nil, fmt.Errorf("client type not supported")
}

func FilterCallableResources(t interface{}, resources []string) []string {
	var filteredResources []string

	switch t.(type) {
	case pango.Panorama:
		for _, r := range resources {
			if strings.HasPrefix(r, "panorama_") {
				filteredResources = append(filteredResources, r)
			}
		}
	case pango.Firewall:
		for _, r := range resources {
			if strings.HasPrefix(r, "firewall_") {
				filteredResources = append(filteredResources, r)
			}
		}
	}

	return filteredResources
}

func normalizeResourceName(s string) string {
	normalize := precis.NewIdentifier(
		precis.AdditionalMapping(func() transform.Transformer {
			return transform.Chain(norm.NFD, transform.RemoveFunc(func(r rune) bool { //nolint
				return unicode.Is(unicode.Mn, r)
			}))
		}),
		precis.Norm(norm.NFC),
	)

	r := strings.NewReplacer(" ", "_",
		"!", "_",
		"\"", "_",
		"#", "_",
		"%", "_",
		"&", "_",
		"'", "_",
		"(", "_",
		")", "_",
		"{", "_",
		"}", "_",
		"*", "_",
		"+", "_",
		",", "_",
		"-", "_",
		".", "_",
		"/", "_",
		"|", "_",
		"\\", "_",
		":", "_",
		";", "_",
		">", "_",
		"=", "_",
		"<", "_",
		"?", "_",
		"[", "_",
		"]", "_",
		"^", "_",
		"`", "_",
		"~", "_",
		"$", "_",
		"@", "_at_")
	replaced := r.Replace(strings.ToLower(s))

	result, err := normalize.String(replaced)
	if err != nil {
		return replaced
	}

	return result
}

type getListWithoutArg interface {
	GetList() ([]string, error)
}

type getListWithOneArg interface {
	GetList(string) ([]string, error)
}

type getListWithTwoArgs interface {
	GetList(string, string) ([]string, error)
}

type getListWithThreeArgs interface {
	GetList(string, string, string) ([]string, error)
}

type getListWithFourArgs interface {
	GetList(string, string, string, string) ([]string, error)
}

type getListWithFiveArgs interface {
	GetList(string, string, string, string, string) ([]string, error)
}

type getGeneric struct {
	i      interface{}
	params []string
}

func contains(s []string, e string) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func mapExists(mapString map[string]string, item map[string]interface{}, element string) bool {
	if _, ok := item[element]; ok {
		if _, ok2 := mapString[item[element].(string)]; ok2 {
			return true
		}
	}
	return false
}

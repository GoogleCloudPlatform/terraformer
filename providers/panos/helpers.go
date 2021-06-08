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
	"os"
	"strings"
	"unicode"

	"github.com/PaloAltoNetworks/pango"
	"github.com/PaloAltoNetworks/pango/util"
	"golang.org/x/text/secure/precis"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func Initialize() (*pango.Firewall, error) {
	fw := &pango.Firewall{
		Client: pango.Client{
			CheckEnvironment: true,
		},
	}

	if val := os.Getenv("PANOS_LOGGING"); val == "" {
		fw.Client.Logging = pango.LogQuiet
	}

	return fw, fw.Initialize()
}

func GetVsysList() ([]string, error) {
	client, err := Initialize()
	if err != nil {
		return []string{}, err
	}

	vsysList, err := client.EntryListUsing(client.Get, []string{
		"config",
		"devices",
		util.AsEntryXpath([]string{"localhost.localdomain"}),
		"vsys",
	})
	if err != nil {
		return []string{}, err
	}

	return vsysList, nil
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

type getGeneric struct {
	i      interface{}
	params []string
}

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

package rabbitmq

import (
	"strings"
	"unicode"

	"golang.org/x/text/secure/precis"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func normalizeResourceName(s string) string {
	normalize := precis.NewIdentifier(
		precis.AdditionalMapping(func() transform.Transformer {
			return transform.Chain(norm.NFD, transform.RemoveFunc(func(r rune) bool { //nolint
				return unicode.Is(unicode.Mn, r)
			}))
		}),
		precis.Norm(norm.NFC),
	)
	noWhiteSpacesString := strings.ReplaceAll(s, " ", "_")
	normalizedLower, _ := normalize.String(strings.ToLower(noWhiteSpacesString))
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
		"/", "slash",
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
	return r.Replace(normalizedLower)
}

func percentEncodeSlashes(s string) string {
	// Encode any percent signs, then encode any forward slashes.
	return strings.ReplaceAll(strings.ReplaceAll(s, "%", "%25"), "/", "%2F")
}

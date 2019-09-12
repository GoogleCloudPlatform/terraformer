// Copyright 2019 The Terraformer Authors.
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

package newrelic

import (
	"regexp"
	"strings"
)

func removeDuplicate(s string) string {
	if len(s) < 2 {
		return s
	}

	src := []byte(s)
	dest := make([]byte, len(src))
	dest[0] = src[0]

	j := 0
	for i := 1; i < len(s); i++ {
		if dest[j] != src[i] {
			j++
			dest[j] = src[i]
		}
	}

	return string(dest[:j+1])
}

// Making resource's name less ugly
func normalizeResourceName(s string) string {
	specialChars := `<>()*#{}[]|@_ .%'",&`
	for _, c := range specialChars {
		s = strings.ReplaceAll(s, string(c), "-")
	}

	s = regexp.MustCompile(`^[^a-zA-Z_]+`).ReplaceAllLiteralString(s, "")
	s = strings.TrimSuffix(s, "-")
	s = removeDuplicate(s)

	return strings.ToLower(s)
}

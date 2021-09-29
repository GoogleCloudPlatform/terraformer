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

package okta

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
)

// Making resource's name less ugly
func normalizeResourceName(s string) string {
	specialChars := `<>()*#{}[]|@_ .%'",&`
	for _, c := range specialChars {
		s = strings.ReplaceAll(s, string(c), "-")
	}

	s = regexp.MustCompile(`^[^a-zA-Z_]+`).ReplaceAllLiteralString(s, "")
	s = strings.TrimSuffix(s, "-")

	return strings.ToLower(s)
}

func normalizeResourceNameWithRandom(s string, rand bool) string {
	specialChars := `-<>()*#{}[]|@_ .%'",&`
	for _, c := range specialChars {
		s = strings.ReplaceAll(s, string(c), "_")
	}
	s = regexp.MustCompile(`^[^a-zA-Z_]+`).ReplaceAllLiteralString(s, "")
	s = strings.TrimSuffix(s, "`_")
	if rand {
		randString := RandStringBytes(4)
		return fmt.Sprintf("%s_%s", strings.ToLower(s), randString)
	}
	return strings.ToLower(s)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyz0123456789"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

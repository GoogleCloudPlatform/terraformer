package ibm

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
)

func normalizeResourceName(s string, rand bool) string {
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
func getRandom(names map[string]struct{}, name string, random bool) (map[string]struct{}, bool) {
	if _, ok := names[name]; ok {
		random = true
	}
	names[name] = struct{}{}
	return names, random
}

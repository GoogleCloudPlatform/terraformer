package terraformutils

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"
)

var OpeningBracketRegexp = regexp.MustCompile(`.?\\<`)
var ClosingBracketRegexp = regexp.MustCompile(`.?\\>`)

func jsonPrint(data interface{}) ([]byte, error) {
	dataJSONBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Println(string(dataJSONBytes))
		return []byte{}, fmt.Errorf("error marshalling terraform data to json: %v", err)
	}
	// We don't need to escape > or <
	s := strings.ReplaceAll(string(dataJSONBytes), "\\u003c", "<")
	s = OpeningBracketRegexp.ReplaceAllStringFunc(s, escapingBackslashReplacer("<"))
	s = strings.ReplaceAll(s, "\\u003e", ">")
	s = ClosingBracketRegexp.ReplaceAllStringFunc(s, escapingBackslashReplacer(">"))
	return []byte(s), nil
}

func escapingBackslashReplacer(backslashedCharacter string) func(string) string {
	return func(match string) string {
		if strings.HasPrefix(match, "\\\\") {
			return match // Don't replace regular backslashes
		}
		return strings.Replace(match, "\\"+backslashedCharacter, backslashedCharacter, 1)
	}
}

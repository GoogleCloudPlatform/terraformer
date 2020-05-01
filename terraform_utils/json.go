package terraform_utils

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
	dataJsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Println(string(dataJsonBytes))
		return []byte{}, fmt.Errorf("error marshalling terraform data to json: %v", err)
	}
	// We don't need to escape > or <
	s := strings.Replace(string(dataJsonBytes), "\\u003c", "<", -1)
	s = OpeningBracketRegexp.ReplaceAllStringFunc(s, escapingBackslashReplacer("<"))
	s = strings.Replace(s, "\\u003e", ">", -1)
	s = ClosingBracketRegexp.ReplaceAllStringFunc(s, escapingBackslashReplacer(">"))
	return []byte(s), nil
}

func escapingBackslashReplacer(backslashedCharacter string) func(string) string {
	return func(match string) string {
		if strings.HasPrefix(match, "\\\\") {
			return match // Don't replace regular backslashes
		} else {
			return strings.Replace(match, "\\"+backslashedCharacter, backslashedCharacter, 1)
		}
	}
}

package terraform_utils

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func jsonPrint(data interface{}) ([]byte, error) {
	dataJsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Println(string(dataJsonBytes))
		return []byte{}, fmt.Errorf("error marshalling terraform data to json: %v", err)
	}
	// We don't need to escape > or <
	s := strings.Replace(string(dataJsonBytes), "\\u003c", "<", -1)
	s = strings.Replace(s, "\\u003e", ">", -1)
	return []byte(s), nil
}

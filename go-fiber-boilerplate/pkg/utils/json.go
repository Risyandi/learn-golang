package utils

import "encoding/json"

func JsonUnmarshal(input string, output interface{}) interface{} {
	json.Unmarshal([]byte(input), &output)
	return output
}

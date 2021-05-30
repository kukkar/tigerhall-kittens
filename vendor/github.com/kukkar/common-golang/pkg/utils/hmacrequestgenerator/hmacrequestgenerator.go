package hmacrequestgenerator

import (
	"fmt"
	"strings"
)

func GetHMacRequest(input map[string]interface{}) (HMacRequest, error) {
	var data HMacRequest
	out := make(map[string]interface{})
	if input == nil {
		return data, fmt.Errorf("request body can not be empty")
	}
	counter := make(map[string]int, 0)
	for key, eachInput := range input {
		switch v := eachInput.(type) {
		case []interface{}:
			switch v[0].(type) {
			case string, int:
				semiformat := fmt.Sprintf("%+v", eachInput)        // Turn the slice into a string that looks like ["one" "two" "three"]
				tokens := strings.Split(semiformat, " ")           // Split this string by spaces
				out[key] = fmt.Sprintf(strings.Join(tokens, ", ")) // Join the Slice together (that was split by spaces) with commas
			case interface{}:
				for _, eachValue := range v {
					switch k := eachValue.(type) {
					case map[string]interface{}:
						for key, value := range k {
							if _, ok := out[key]; ok {
								out[key+string(counter[key])] = value
								counter[key] = getNextCharfromAscii(counter[key])
							} else {
								out[key] = value
								counter[key] = getNextCharfromAscii(0)
							}
						}
					}
				}
			}

		case map[string]interface{}:
			for key, value := range v {
				out[key] = value
			}
		default:
			out[key] = eachInput
		}
	}
	sortedKeys := getSortedKey(out)
	craftedData := seprateValueWithPipe(out, sortedKeys)
	data = HMacRequest(craftedData)
	return data, nil
}

func (this HMacRequest) ToString() string {
	return ""
}

func getNextCharfromAscii(c int) int {
	if c <= 64 {
		return 65
	}
	if c > 90 {
		return 97
	}
	return c + 1
}

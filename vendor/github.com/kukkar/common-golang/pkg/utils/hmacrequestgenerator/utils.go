package hmacrequestgenerator

import (
	"fmt"
	"sort"
	"strings"
)

func getSortedKey(m map[string]interface{}) []string {

	// To store the keys in slice in sorted order
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func seprateValueWithPipe(m map[string]interface{}, sortedKey []string) string {
	var output string
	for _, eachKey := range sortedKey {
		if val, ok := m[eachKey]; ok {
			if val != nil {
				switch v := val.(type) {
				case int:
					output += fmt.Sprintf("%d", v) + "|"
				case int32:
					output += fmt.Sprintf("%d", v) + "|"
				case int64:
					output += fmt.Sprintf("%d", v) + "|"
				case float64:
					output += fmt.Sprintf("%f", v) + "|"
				case string:
					output += fmt.Sprintf("%s", v) + "|"
				case uint64:
					output += fmt.Sprintf("%d", v) + "|"
				case uint32:
					output += fmt.Sprintf("%d", v) + "|"
				case uint:
					output += fmt.Sprintf("%d", v) + "|"
				case bool:
					output += fmt.Sprintf("%v", v) + "|"
				default:
					output += fmt.Sprintf("%v", v) + "|"
				}
			}
		}
	}

	return strings.TrimRight(output, "|")
}

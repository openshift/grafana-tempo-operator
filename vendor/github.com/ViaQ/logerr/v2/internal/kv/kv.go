package kv

import (
	"fmt"
)

// ToMap converts keysAndValues to a map
func ToMap(keysAndValues ...interface{}) map[string]interface{} {
	kvlen := len(keysAndValues)
	kve := make(map[string]interface{}, kvlen/2)

	for i, j := 0, 1; i < kvlen && j < kvlen; i, j = i+2, j+2 {
		key, ok := keysAndValues[i].(string)

		// Expecting a string as the key, however will make a
		// best guess through conversion if it isn't.
		if !ok {
			key = fmt.Sprintf("%s", keysAndValues[i])
		}

		kve[key] = keysAndValues[j]
	}

	return kve
}

// FromMap converts a map to a key/value slice
func FromMap(m map[string]interface{}) []interface{} {
	res := make([]interface{}, 0, len(m)*2)
	for k, v := range m {
		res = append(res, k, v)
	}
	return res
}

package flatjson

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// UnflattenMap takes a flat map and a delimiter, and returns a nested map or array.
func UnflattenMap(flatMap map[string]interface{}, delimiter string) (interface{}, error) {
	var (
		object = make(map[string]interface{}, 0)
		array  = make([]interface{}, 0)
		nested = make(map[string]map[string]interface{}, 0)
	)

	for key, value := range flatMap {
		parts := strings.Split(key, delimiter)
		if len(parts) == 1 {
			if _, err := strconv.Atoi(parts[0]); err == nil {
				// Element is raw array value
				array = append(array, value)
				continue
			}
			// Element is raw value
			object[parts[0]] = value
			continue
		}

		// Element is a nested array or object
		if _, exists := nested[parts[0]]; !exists {
			nested[parts[0]] = make(map[string]interface{}, 1)
		}
		// Remove the first part so the reccursion treats it as a self contained map
		nested[parts[0]][strings.Join(parts[1:], delimiter)] = value
	}

	// Unflatten any nested objects
	for key, value := range nested {
		if _, err := strconv.Atoi(key); err == nil {
			if unflattenedInner, err := UnflattenMap(value, delimiter); err == nil {
				array = append(array, unflattenedInner)
			}
			continue
		}
		unflattenedInner, err := UnflattenMap(value, delimiter)
		if err != nil {
			return nil, fmt.Errorf("couldnt unflatten inner: %s", err)
		}
		object[key] = unflattenedInner
	}

	if len(array) > 0 {
		return array, nil
	}
	return object, nil
}

// UnflattenBytes is a shortcut to unflatten a json object as bytes
func UnflattenBytes(b []byte, delimiter string) (interface{}, error) {
	var m = make(map[string]interface{}, 0)
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, fmt.Errorf("couldnt unmarshal bytes to map: %s", err)
	}
	return UnflattenMap(m, delimiter)
}

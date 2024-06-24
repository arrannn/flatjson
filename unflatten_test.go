package flatjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestUnflatten(t *testing.T) {
	data := map[string]interface{}{
		"host":                         "hostname",
		"vulnerabilities.0.id":         "V-01-ID",
		"vulnerabilities.0.name":       "Attacked with lazers",
		"vulnerabilities.1.id":         "V-25-ID",
		"network.address":              "localhost",
		"numUsers":                     6,
		"vulnerabilities.0.detections": 100,
		"active":                       true,
		"nested.boolparam.0.test":      false,
		"testingarray.0":               "arrayline1",
		"testingarray.1":               "arrayline2",
	}

	unflattened, _ := Unflatten(data, ".")
	asBytes, _ := json.Marshal(unflattened)

	fmt.Println(string(asBytes))
}

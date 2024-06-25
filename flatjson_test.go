package flatjson

import (
	"encoding/json"
	"testing"
)

func TestUnflattenMap(t *testing.T) {
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

	unflattened, _ := UnflattenMap(data, ".")
	asBytes, _ := json.Marshal(unflattened)

	var output struct {
		Host            string `json:"host"`
		Vulnerabilities []struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		}
		TestArray []string `json:"testingarray"`
	}
	json.Unmarshal(asBytes, &output)
	if output.Host != "hostname" {
		t.Fail()
	}
	if output.Vulnerabilities[0].Id != "V-01-ID" {
		t.Fail()
	}
}

func BenchmarkUnflattenMap(b *testing.B) {
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
	for i := 0; i < b.N; i++ {
		UnflattenMap(data, ".")
	}
}

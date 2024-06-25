# FlatJSON

FlatJSON is a Go package that provides functionality to unflatten JSON objects. It takes a flat map with keys containing delimiters and converts it into a nested map or array.

## Example Usage

Here is an example of how to use the FlatJSON package:
```go
import (
	"github.com/arrannn/flatjson"
)

func main() {
	data := map[string]interface{}{
		"host":                  "hostname",
		"nestedarray.0.field":   "V-01-ID",
		"nestedarray.1.field":   "V-25-ID",
		"nested.object.0.field": false,
	}

	unflattened, _ := flatjson.UnflattenMap(data, ".")
}
```
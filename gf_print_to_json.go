package helpers

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/pretty"
)

func PrintToJson(v ...interface{}) {
	bytes, _ := json.MarshalIndent(v, "", "\t")
	fmt.Println(string(pretty.Color(pretty.PrettyOptions(bytes, pretty.DefaultOptions), nil)))
}

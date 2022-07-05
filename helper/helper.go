package helper

import (
	"bytes"
	"encoding/json"
)


/*
Pretty JSON string
*/
func PrettyString(str string) (string, error) {
    var prettyJSON bytes.Buffer
    if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
        return "", err
    }
    return prettyJSON.String(), nil
}


/*
Map type assertion
*/
type Map_assertion map[string]interface{}

func (d Map_assertion) D(k string) Map_assertion {
    return d[k].(map[string]interface{})
}

func (d Map_assertion) S(k string) string {
    return d[k].(string)
}
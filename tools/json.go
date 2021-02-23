package tools

import "encoding/json"

func StructToJsonStr(i interface{}) (string, error) {
	if j, err := json.Marshal(i); err != nil {
		return string(j), err
	} else {
		return "", err
	}
}

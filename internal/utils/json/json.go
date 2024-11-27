package json

import (
	jsoniter "github.com/json-iterator/go"
)

func Jsonify(obj interface{}) string {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	if obj == nil {
		return ""
	}

	str, err := json.MarshalIndent(obj, "", " ")
	if err != nil {
		return ""
	}

	return string(str)
}

package jsonpath

import (
	"encoding/json"
	"github.com/tidwall/sjson"
)

func Set(jsonString string, jsonPath string, target string) (result string){
	var j interface{}
	json.Unmarshal([]byte(jsonString), &j)
	_, arr,_ := JsonPathLookup(j, jsonPath)

	for _,jPath:=range arr{
		jsonString,_ = sjson.Set(jsonString, jPath, target)
	}

	result = jsonString
	return
}
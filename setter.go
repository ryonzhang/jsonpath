package jsonpath

import (
	"encoding/json"
	"github.com/tidwall/sjson"
	"github.com/tidwall/gjson"
)

func Set(jsonString string, jsonPath string, convertFunc func(input string) string) (result string){
	var j interface{}
	json.Unmarshal([]byte(jsonString), &j)
	_, arr,_ := JsonPathLookup(j, jsonPath)

	for _,jPath:=range arr{
		jsonString,_ = sjson.Set(jsonString, jPath, convertFunc(gjson.Get(jsonString, jPath).String()))
	}
	result = jsonString
	return
}
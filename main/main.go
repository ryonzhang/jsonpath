package main

import (
	"fmt"
	jsonpath2 "github.com/ryonzhang/jsonpath"
)

func main(){
	data := `{
    "store": {
        "book": [
            {
                "category": "reference",
                "author": "Nigel Rees",
                "title": "Sayings of the Century",
                "price": 8.95
            },
            {
                "category": "fiction",
                "author": "Evelyn Waugh",
                "title": "Sword of Honour",
                "price": 12.99
            },
            {
                "category": "fiction",
                "author": "Herman Melville",
                "title": "Moby Dick",
                "isbn": "0-553-21311-3",
                "price": 8.99
            },
            {
                "category": "fiction",
                "author": "J. R. R. Tolkien",
                "title": "The Lord of the Rings",
                "isbn": "0-395-19395-8",
                "price": 22.99
            }
        ],
        "bicycle": {
            "color": "red",
            "price": 19.95
        }
    },
    "expensive": 10
}`
	//var j interface{}
	//json.Unmarshal([]byte(data), &j)
	//res, arr,err := jsonpath.JsonPathLookup(j, "$.store.book[?(@.price < 40)].price")
	//if err != nil {
	//	panic(arr)
	//}
	//arr1 := res.([]interface{})
	//if len(arr1) == 0 {
	//	panic("should return [], got: ")
	//}

	log:= jsonpath2.Set(data,"$.store.book[?(@.price < 21)].price","ryon")
	fmt.Println(log)
}

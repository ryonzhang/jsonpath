# jsonpath 
This library will help to set or convert a piece inside json based on jsonpath

## Install

```bash
$ go get github.com/ryonzhang/jsonpath
```

## Usage

```go
sample := `{
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

result:= jsonpath.Set(data,"$.store.book[?(@.price < 21)].price",func(string input)string{
    return "what"
})

// result == `"store": {
                      "book": [
                          {
                              "category": "reference",
                              "author": "Nigel Rees",
                              "title": "Sayings of the Century",
                              "price": "what"
                          },
                          {
                              "category": "fiction",
                              "author": "Evelyn Waugh",
                              "title": "Sword of Honour",
                              "price": "what"
                          },
                          {
                              "category": "fiction",
                              "author": "Herman Melville",
                              "title": "Moby Dick",
                              "isbn": "0-553-21311-3",
                              "price": "what"
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
```

## API

### Set(jsonString string, jsonPath string, convertFunc func(input string) string) (result string)

Set `jsonString` according to `jsonPath` by the rule defined by convertFunc

For example , in the following json

```
{
  "category": "reference",
  "author": "Nigel Rees",
  "title": "Sayings of the Century",
  "price": "what"
}

```

and by convertFun like below:


```
func convertFunc(input string)string{
    return input+"-processed"
}
```

The function `jsonpath.Set(data,"$.category",convertFunc)` will return the following

```
{
  "category": "reference-processed",
  "author": "Nigel Rees",
  "title": "Sayings of the Century",
  "price": "what"
}

```

## License

MIT © [mdaverde](https://github.com/ryonzhang)
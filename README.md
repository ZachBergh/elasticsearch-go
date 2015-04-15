#Tools for ElasticSearch In Golang

####Why
Project Need a simple ElasticSearch Golang Tool, So i made one

####Usage

* Now Just has Search Tools:
```
package main

import (
    "fmt"
    ESgo "github.com/ZachBergh/elasticsearch-go"
)

func main() {
    query := `{
      "query": {
        "term": {
          "user_id": {
            "value": "oeWnmjql5QeP9g5tFA7dDdL9Zc4I"
          }
        }
      },"size":0
    }`

    data, err := ESgo.NewClient("http://192.168.5.8:9200", "wex", "wxlog", query).Search()
    fmt.Println(err)
    fmt.Println(data)
}
```
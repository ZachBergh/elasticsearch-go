#Tools for ElasticSearch In Golang

####Docker Container
```
docker run -t \
-v /home/core/share/codes/src/github.com/:/go/src/github.com/:rw \
--name githubDev -i \
-p 80:80 \
-d dev /bin/bash 
```

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
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Student struct {
    Id   int `json:"-"`
    Name string `json:"name,omitempty"`
}

func main() {
    //短变量声明 ？
    var s Student
    s.Id = 1
    s.Name = "Alice"
	buf, err := json.Marshal(s)
	if err != nil {
		log.Fatalf("marshal error:%s", err)
	}
	fmt.Println(string(buf))
}

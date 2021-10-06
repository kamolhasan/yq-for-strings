package main

import (
	"fmt"
	"github.com/kamolhasan/yq-for-strings/pkg/yqstrings"
)

func main(){
	y1 := `
a: "aaaa"
ar:
- "aaa"
- "bbb"
- "ccc"
b:
  c:
    d: "dddd"
`
	node, err := yqstrings.StringToNode(y1)
	if err!= nil {
		panic(err)
	}
	data, err := yqstrings.NodeToString(node)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
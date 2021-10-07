package main

import (
	"fmt"

	"github.com/kamolhasan/yq-for-strings/pkg/yqstrings"
	"github.com/mikefarah/yq/v3/pkg/yqlib"
)

func main() {
	y1 := `
ar:
# list of ar
- "aaa"
- "bbb"
- "ccc"

a: "aaaa"
b:
  c:
    d: "dddd"
`

	y2 := `
a: "aaaa"
ar:
# value ddd
- "ddd"
b:
  c:
    e: "some"
    d:
      f: "FFFF"
`
	output, err := yqstrings.Merge(yqlib.AppendArrayMergeStrategy, yqlib.AppendCommentsMergeStrategy, true, y1, y2)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)
}

package main

import (
	"fmt"
	"os"

	"github.com/kamolhasan/yq-for-strings/pkg/yqstrings"
	"github.com/mikefarah/yq/v3/pkg/yqlib"
	"gopkg.in/op/go-logging.v1"
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

func init() {
	//var format = logging.MustStringFormatter(
	//	`%{color}%{time:15:04:05} %{shortfunc} [%{level:.4s}]%{color:reset} %{message}`,
	//)
	//var backend = logging.AddModuleLevel(logging.NewBackendFormatter(logging.NewLogBackend(os.Stderr, "", 0), format))
	//backend.SetLevel(logging.DEBUG, "")

	var backend = logging.AddModuleLevel(logging.NewLogBackend(os.Stderr, "", 0))
	backend.SetLevel(logging.ERROR, "")
	logging.SetBackend(backend)
}

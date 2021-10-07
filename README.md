# yq-for-strings
Use "mikefarah/yq" library to manipulate YAML stored as string type variable


## Examples:

```go
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
```

Output:
```yaml 
ar:
    # list of ar
    - "aaa"
    - "bbb"
    - "ccc"
    # value ddd
    - "ddd"
a: "aaaa"
b:
    c:
        d:
            f: "FFFF"
        e: "some"
```
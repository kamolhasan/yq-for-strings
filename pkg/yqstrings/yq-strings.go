package yqstrings

import (
	yqslib "github.com/kamolhasan/yq-for-strings/pkg/lib"
	"github.com/kamolhasan/yq/pkg/yqlib"
	"gopkg.in/yaml.v3"
)

func StringToNode(yml string) (*yaml.Node, error) {
	var node yaml.Node
	err := yaml.Unmarshal([]byte(yml),&node)
	if err != nil {
		return nil, err
	}

	return &node, nil
}

func NodeToString(node *yaml.Node) (string,error) {
	b, err := yaml.Marshal(node)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func merge(arrayMergeStrategy yqlib.CommentsMergeStrategy,A, B string) (string, error)  {
	nodeA ,err := StringToNode(A)
	if err != nil {
		return "", err
	}

	nodeB, err := StringToNode(B)
	if err != nil {
		return "", err
	}
	lib := yqslib.NewYqStringLib()
	lib.Update(dataBucket, updateCommand, autoCreateFlag)
}

func Merge(arrayMergeStrategy yqlib.CommentsMergeStrategy,inputs ...string) (string, error)  {
	var output string
	for _, s:= range inputs {
		output, err := merge(output,s)
		if err != nil {
			return "", err
		}
	}
}
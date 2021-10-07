package yqstrings

import (
	yqslib "github.com/kamolhasan/yq-for-strings/pkg/lib"
	"github.com/mikefarah/yq/v3/pkg/yqlib"
	"gopkg.in/yaml.v3"
)

func StringToNode(yml string) (*yaml.Node, error) {
	var node yaml.Node
	err := yaml.Unmarshal([]byte(yml), &node)
	if err != nil {
		return nil, err
	}

	return &node, nil
}

func NodeToString(node *yaml.Node) (string, error) {
	b, err := yaml.Marshal(node)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func merge(arrayMergeStrategy yqlib.ArrayMergeStrategy, commentsMergeStrategy yqlib.CommentsMergeStrategy, overwrite bool, A, B string) (string, error) {
	nodeA, err := StringToNode(A)
	if err != nil {
		return "", err
	}

	nodeB, err := StringToNode(B)
	if err != nil {
		return "", err
	}
	lib := yqslib.NewYqStringLib()
	nodeCtxs, err := lib.NodeToNodeContexts(nodeB, arrayMergeStrategy)
	if err != nil {
		return "", err
	}
	for _, nctx := range nodeCtxs {
		err := lib.YqLib.Update(nodeA, yqlib.UpdateCommand{
			Command:               "merge",
			Path:                  lib.YqLib.MergePathStackToString(nctx.PathStack, arrayMergeStrategy),
			Value:                 nctx.Node,
			Overwrite:             overwrite,
			CommentsMergeStrategy: commentsMergeStrategy,
			// dont update the content for nodes midway, only leaf nodes
			DontUpdateNodeContent: nctx.IsMiddleNode && (arrayMergeStrategy != yqlib.OverwriteArrayMergeStrategy || nctx.Node.Kind != yaml.SequenceNode),
		}, true)
		if err != nil {
			return "", err
		}
	}

	output, err := NodeToString(nodeA)
	if err != nil {
		return "", err
	}

	return output, nil
}

func Merge(arrayMergeStrategy yqlib.ArrayMergeStrategy, commentsMergeStrategy yqlib.CommentsMergeStrategy, overwrite bool, inputs ...string) (string, error) {
	var output string
	var err error
	for _, s := range inputs {
		output, err = merge(arrayMergeStrategy, commentsMergeStrategy, overwrite, output, s)
		if err != nil {
			return "", err
		}
	}

	return output, nil
}

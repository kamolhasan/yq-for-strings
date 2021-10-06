package lib

import (
	"github.com/mikefarah/yq/v3/pkg/yqlib"
	"gopkg.in/yaml.v3"
)

type YqStringsLib struct {
	yqLib yqlib.YqLib
}

func NewYqStringLib() YqStringsLib  {
	return YqStringsLib{ yqLib: yqlib.NewYqLib()}
}

func (yqs *YqStringsLib) NodeToNodeContexts(node *yaml.Node, arrayMergeStrategy yqlib.ArrayMergeStrategy)([]*yqlib.NodeContext, error)  {
	return yqs.yqLib.GetForMerge(node, "**", arrayMergeStrategy)
}


package simple

import (
	"github.com/qingcloudhx/flow/model"
)

const (
	ModelName = "flogo-simple"
)

func init() {
	model.Register(New())
}

func New() *model.FlowModel {
	m := model.New(ModelName)
	m.RegisterFlowBehavior(&FlowBehavior{})
	m.RegisterDefaultTaskBehavior("basic", &TaskBehavior{})
	m.RegisterTaskBehavior("iterator", &IteratorTaskBehavior{})

	return m
}

package flow

import (
	"github.com/qingcloudhx/core/data/expression"
	"github.com/qingcloudhx/flow/definition"
	"github.com/qingcloudhx/flow/instance"
	"github.com/qingcloudhx/flow/model"
	"github.com/qingcloudhx/flow/model/simple"
	"github.com/qingcloudhx/flow/support"
	"github.com/qingcloudhx/flow/tester"
)

// Provides the different extension points to the FlowBehavior Action
type ExtensionProvider interface {
	GetStateRecorder() instance.StateRecorder
	GetFlowTester() *tester.RestEngineTester

	GetDefaultFlowModel() *model.FlowModel
	GetFlowProvider() definition.Provider
}

//ExtensionProvider is the extension provider for the flow action
type DefaultExtensionProvider struct {
	flowProvider definition.Provider
	flowModel    *model.FlowModel
}

func NewDefaultExtensionProvider() *DefaultExtensionProvider {
	return &DefaultExtensionProvider{}
}

func (fp *DefaultExtensionProvider) GetFlowProvider() definition.Provider {

	if fp.flowProvider == nil {
		fp.flowProvider = &support.BasicRemoteFlowProvider{}
	}

	return fp.flowProvider
}

func (fp *DefaultExtensionProvider) GetDefaultFlowModel() *model.FlowModel {

	if fp.flowModel == nil {
		fp.flowModel = simple.New()
	}

	return fp.flowModel
}

func (fp *DefaultExtensionProvider) GetStateRecorder() instance.StateRecorder {
	return nil
}

func (fp *DefaultExtensionProvider) GetScriptExprFactory() expression.Factory {
	return nil
}

//todo make FlowTester an interface
func (fp *DefaultExtensionProvider) GetFlowTester() *tester.RestEngineTester {
	return nil
}

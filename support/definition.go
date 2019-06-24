package support

import (
	"github.com/qingcloudhx/core/app/resource"
	"github.com/qingcloudhx/flow/definition"
	"strings"
)

// todo fix this
var flowManager *FlowManager
var resManager *resource.Manager

func InitDefaultDefLookup(fManager *FlowManager, rManager *resource.Manager) {
	flowManager = fManager
	resManager = rManager
}

func GetDefinition(flowURI string) (*definition.Definition, bool, error) {

	var def *definition.Definition

	if strings.HasPrefix(flowURI, resource.UriScheme) {

		res := resManager.GetResource(flowURI)

		if res != nil {
			var ok bool
			def, ok = res.Object().(*definition.Definition)
			if ok {
				return def, true, nil
			}
		}
	} else {
		var err error
		def, err = flowManager.GetFlow(flowURI)
		if err != nil {
			return nil, false, err
		}

		return def, false, nil
	}

	return nil, false, nil
}

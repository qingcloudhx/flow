package instance

import (
	"time"

	coreevent "github.com/qingcloudhx/core/engine/event"
	"github.com/qingcloudhx/flow/model"
	"github.com/qingcloudhx/flow/support/event"
)

type flowEvent struct {
	time                           time.Time
	err                            error
	input, output                  map[string]interface{}
	status                         event.Status
	name, id, parentName, parentId string
}

func (fe *flowEvent) FlowName() string {
	return fe.name
}

// Returns flow ID
func (fe *flowEvent) FlowID() string {
	return fe.id
}

// In case of subflow, returns parent flow name
func (fe *flowEvent) ParentFlowName() string {
	return fe.parentName
}

// In case of subflow, returns parent flow ID
func (fe *flowEvent) ParentFlowID() string {
	return fe.parentId
}

// Returns event time
func (fe *flowEvent) Time() time.Time {
	return fe.time
}

// Returns current flow status
func (fe *flowEvent) FlowStatus() event.Status {
	return fe.status
}

// Returns output data for completed flow instance
func (fe *flowEvent) FlowOutput() map[string]interface{} {
	return fe.output
}

// Returns input data for flow instance
func (fe *flowEvent) FlowInput() map[string]interface{} {
	return fe.input
}

// Returns error for failed flow instance
func (fe *flowEvent) FlowError() error {
	return fe.err
}

func postFlowEvent(inst *Instance) {

	if coreevent.HasListener(event.FlowEventType) {

		fe := &flowEvent{}
		fe.time = time.Now()
		fe.name = inst.Name()
		fe.id = inst.ID()
		if inst.master != nil {
			fe.parentName = inst.master.Name()
			fe.parentId = inst.master.ID()
		}
		fe.status = convertFlowStatus(inst.Status())

		fe.input = make(map[string]interface{})
		fe.output = make(map[string]interface{})

		if fe.status != event.CREATED {
			attrs := inst.attrs
			outData, _ := inst.GetReturnData()
			if attrs != nil && len(attrs) > 0 {
				for name, attVal := range attrs {
					if outData != nil && outData[name] != nil {
						if fe.status == event.COMPLETED {
							fe.output[name] = attVal
						}
						// Since same attribute map is used for input and output, filter output attributes
						continue
					}
					fe.input[name] = attVal
				}
			}
		}

		if fe.status == event.FAILED {
			fe.err = inst.returnError
		}
		coreevent.Post(event.FlowEventType, fe)
	}
}

func convertFlowStatus(code model.FlowStatus) event.Status {
	switch code {
	case model.FlowStatusNotStarted:
		return event.CREATED
	case model.FlowStatusActive:
		return event.STARTED
	case model.FlowStatusCancelled:
		return event.CANCELLED
	case model.FlowStatusCompleted:
		return event.COMPLETED
	case model.FlowStatusFailed:
		return event.FAILED
	}
	return event.UNKNOWN
}

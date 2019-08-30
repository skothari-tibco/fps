package replace

import (
	"strings"

	"github.com/project-flogo/cml/action/operation"
	"github.com/project-flogo/core/support/log"
)

type Operation struct {
	// params *Params
	logger log.Logger
}

func New(ctx operation.InitContext) (operation.Operation, error) {

	return &Operation{logger: ctx.Logger()}, nil
}

func (a *Operation) Eval(inputs map[string]interface{}) (interface{}, error) {

	//To get the inputs in the desired types.
	input := &Input{}
	input.FromMap(inputs)

	a.logger.Debug("Executing operation...", input.S0, " ", input.S1, " ", input.S2, " ", input.I)

	out := strings.Replace(input.S0, input.S1, input.S2, input.I)

	return out, nil
}
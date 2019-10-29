package sort

import (
	//	"fmt"

	"github.com/project-flogo/catalystml-flogo/operations/common"
)

type Params struct {
	Ascending   bool          `md:"ascending"`
	NilPosition string        `md:"nilPosition",allowed=["first","last"],required=false`
	KeepRow     bool          `md:"keepRow"`
	By          []interface{} `md:"by"`
	Axis        int           `md:"axis"`
}

type Input struct {
	Data interface{} `md:"data"`
}

func (i *Input) FromMap(values map[string]interface{}) error {
	var err error
	i.Data, err = common.ToDataFrame(values["data"])

	return err
}
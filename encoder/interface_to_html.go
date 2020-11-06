package encoder

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/report"
	"reflect"
)

type InterfaceToHtml struct{}

func (j InterfaceToHtml) IsAble(object interface{}) bool {
	_, ok := object.(string)
	return ok || object == nil
}

func (j InterfaceToHtml) EncodeThrough(_ inter.App, object interface{}, _ []inter.Encoder) (string, error) {
	if object == nil {
		return "", nil
	}

	result, ok := object.(string)
	if !ok {
		return "", report.EncodeError.Wrap("can not encode to html with an unsupported type " + reflect.TypeOf(object).String())
	}

	return result, nil
}

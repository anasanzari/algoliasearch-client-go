// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

type ExtraURLParamsOption struct {
	value map[string]string
}

func ExtraURLParams(v map[string]string) ExtraURLParamsOption {
	return ExtraURLParamsOption{v}
}

func (o ExtraURLParamsOption) Get() map[string]string {
	return o.value
}

func (o ExtraURLParamsOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *ExtraURLParamsOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = nil
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

type RankingOption struct {
	value []string
}

func Ranking(v ...string) RankingOption {
	return RankingOption{v}
}

func (o RankingOption) Get() []string {
	return o.value
}

func (o RankingOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *RankingOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = []string{"typo", "geo", "words", "filters", "proximity", "attribute", "exact", "custom"}
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

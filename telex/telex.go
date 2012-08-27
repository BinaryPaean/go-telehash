package telex

import (
	"encoding/json"
)

type UA []interface{} //Untyped Array
type Telex struct {
	Headers  UA
	Commands UA
	Signals  UA
	Body     UA
	Raw      UA
}

func TelexFromJson(json_in []byte) (*Telex, error) {
	t := new(Telex)
	err := json.Unmarshal(json_in, t.Raw)
	return t, err
}

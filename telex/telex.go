package telex

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type JsonObject map[string]interface{}

type Telex struct {
	Headers  *Headers
	Commands *Commands
	Signals  JsonObject
	Raw      JsonObject
	Errors   []error
}

type Headers struct {
	To   string `json:"_to"`
	Ring int64  `json:"_ring"`
	Line int64  `json:"_line"`
	Br   int64  `json:"_br"`
	Hop  int    `json:"_hop"`
}

type Commands struct {
	See []string `json:".see"`
	Tap []string `json:".tap`
}

func TelexFromJson(json_in []byte) (t *Telex) {
	t = new(Telex)
	t.unmarshalInto(json_in, &t.Raw)
	t.unmarshalInto(json_in, &t.Commands)
	t.unmarshalInto(json_in, &t.Headers)
	return t
}

func (t *Telex) unmarshalInto(json_in []byte, target interface{}) {
	err := json.Unmarshal(json_in, &target) //Load the entire JSON object into Raw
	if err != nil {
		t.Errors = append(t.Errors, err)
	}
}

func (t *Telex) String() string {
	b := new(bytes.Buffer)
	for k, v := range t.Raw {
		switch vv := v.(type) {
		case string:
			b.WriteString(fmt.Sprintln(k, "(string):", vv))
		case float64:
			b.WriteString(fmt.Sprintln(k, "(float64):", vv))
		case []interface{}:
			b.WriteString(fmt.Sprintln(k, "(array):"))
			for i, u := range vv {
				b.WriteString(fmt.Sprintln(i, u))
			}
		default:
			b.WriteString(fmt.Sprintln(k, "is of type: ", vv))
		}
	}
	return b.String()
}

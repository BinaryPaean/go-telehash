//Package telex provides functions to create, marshal, and unmarshal telex messages for the TeleHash protocol.
package telex

import (
	"bytes"
	"encoding/json"
	"fmt"
)

//convenience type for dealing with encoding/json's results.
type jsonObject map[string]interface{}

//Telex is the primary object of interest in the package, and includes a reference to its set of headers and commands.
type Telex struct {
	Headers  *Headers
	Commands *Commands
	Signals  jsonObject
	Raw      jsonObject
	Errors   []error
}

//Headers are optionally included in each telex message.
//   To - A string value of the public IP:PORT that the Telex was sent to (helps recipient with NATs)
//   Ring - Used to open a Line with a Switch, integer value from 1 to 32768
//   Line - The private unique id of the line from one Switch to another, the product of the _ring from both
//   Br - Bytes Received, always tell the recipient the total bytes that have been received from them so far
//   Hop - Integer value from 0 to 4, incremented any time the Telex is forwarded
type Headers struct {
	To   string `json:"_to"`
	Ring int64  `json:"_ring"`
	Line int64  `json:"_line"`
	Br   int64  `json:"_br"`
	Hop  int    `json:"_hop"`
}

//Commands are optionally included in each telex message. 
//    See - An array of other Switches (IP:PORT) that the recipient may find useful.
//    Tap - An array of filters that describe which telexes to match and forward back to the requesting Switch.
type Commands struct {
	See []string `json:".see"`
	Tap []string `json:".tap`
}

//TelexFromJson attempts to decode the provided json in json_in into a new telex. Note that
//TelexFromJson does not return errors, and always returns a new Telex. In the event of parsing errors
//some or all of the telex's fields may be blank, and errors will be reported in t.Errors []error.
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

//String formats a telex as a semi human-readable string. See ToJson to serialize a telex.
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

func (t *Telex) ToJson() []bytes {
	return nil
}

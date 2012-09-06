//Package telex provides functions to create, marshal, and unmarshal telex messages for the TeleHash protocol.
package telex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
)

//convenience type for dealing with encoding/json's results.
type jsonObject map[string]interface{}

//Telex is the primary object of interest in the package.
//Headers are optionally included in each telex message.
//   To - A string value of the public IP:PORT that the Telex was sent to (helps recipient with NATs)
//   Ring - Used to open a Line with a Switch, integer value from 1 to 32768
//   Line - The private unique id of the line from one Switch to another, the product of the _ring from both
//   Br - Bytes Received, always tell the recipient the total bytes that have been received from them so far
//   Hop - Integer value from 0 to 4, incremented any time the Telex is forwarded
//Commands are optionally included in each telex message. 
//    See - An array of other Switches (IP:PORT) that the recipient may find useful.
//    Tap - An array of filters that describe which telexes to match and forward back to the requesting Switch.
type Telex struct {
	Commands jsonObject `json:"-"`
	Headers  jsonObject `json:"-"`
	Signals  jsonObject `json:"-"`
	Body     jsonObject `json:"-"`
}

//NewTelex returns an initialized Telex struct, ready for use.
func NewTelex() *Telex {
	t := new(Telex)
	t.Body = make(jsonObject)
	t.Signals = make(jsonObject)
	t.Headers = make(jsonObject)
	t.Commands = make(jsonObject)
	return t
}

//TelexFromJson attempts to decode the provided json in json_in into a new telex. Note that
//TelexFromJson does not return errors, and always returns a new Telex. In the event of parsing errors
//some or all of the telex's fields may be blank, and errors will be reported in the t.Errors field, as a slice of errors.
func TelexFromJson(json_in []byte) (*Telex, error) {
	t := NewTelex()
	j := new(jsonObject)
	err := json.Unmarshal(json_in, j)
	t.extractFieldsFrom(*j)
	return t, err
}

//String formats a telex as a semi human-readable string. Use ToJson to serialize a telex instead.
func (t *Telex) String() string {
	b := new(bytes.Buffer)
	for k, v := range t.Signals {
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

func (t *Telex) extractFieldsFrom(j jsonObject) {
	signal_matcher := regexp.MustCompile(`^\+(.)+`)
	command_matcher := regexp.MustCompile(`^\.(.)+`)
	header_matcher := regexp.MustCompile(`^\_(.)+`)
	for k, v := range j {
		target := []byte(k)
		switch {
		default:
			t.Body[k] = v
		case signal_matcher.Match(target):
			t.Signals[k] = v
		case command_matcher.Match(target):
			t.Commands[k] = v
		case header_matcher.Match(target):
			t.Headers[k] = v
		}
	}
}

func (t *Telex) ToJson() []byte {
	unified := make(jsonObject)
	for k, v := range t.Headers {
		unified[k] = v
	}

	for k, v := range t.Commands {
		unified[k] = v
	}

	for k, v := range t.Signals {
		unified[k] = v
	}

	for k, v := range t.Body {
		unified[k] = v
	}
	b, _ := json.Marshal(unified)
	return b
}

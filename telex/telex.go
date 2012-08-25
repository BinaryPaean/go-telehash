package telex

import (
	"errors"

//"encoding/json"
)

type Telex struct {
	headers  []*string
	commands []*string
	signals  []*string
	body     []*string
}

func TelexFromJson(json_in string) (*Telex, error) {
	return &Telex{}, TelexParseError{err: errors.New("WTF is this?")}
}

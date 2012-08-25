package telex

import (
	"strconv"
)

type TelexParseError struct {
	err     error
	line    int
	context string
}

func (tpe TelexParseError) Error() string {
	return strconv.Itoa(tpe.line) + ": " + tpe.err.Error() + " While parsing: " + tpe.context
}

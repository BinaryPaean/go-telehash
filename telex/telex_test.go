package telex

import (
	"testing"
)

func TestTelexFromJson(t *testing.T) {
	const test_json = `{
      "_ring":43723,
      ".see":["5.6.7.8:23456","11.22.33.44:11223"],
    }`
	tx, err := TelexFromJson(test_json)
	if err != nil {
		t.Errorf("Parsing telex json failed!")
	}
	t.Log(tx)
}

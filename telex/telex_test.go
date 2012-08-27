package telex

import (
	"encoding/json"
	"launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) { gocheck.TestingT(t) }

type TelexSuite struct{}

var _ = gocheck.Suite(&TelexSuite{})

func (s *TelexSuite) TestTelexFromJson(c *gocheck.C) {
	const test_json = `{
      "_ring":43723,
      ".see":["5.6.7.8:23456","11.22.33.44:11223"],
    }`
	tx, err := TelexFromJson([]byte(test_json))
	c.Check(err, gocheck.IsNil)
	headers := []interface{}{map[interface{}]interface{}{"_ring": 43723}}
	commands := []interface{}{map[interface{}]interface{}{".see": []interface{}{"5.6.7.8:23456", "11.22.33.44:11223"}}}
	c.Check(tx.Headers, gocheck.DeepEquals, headers)
	c.Check(tx.Commands, gocheck.DeepEquals, commands)
	var blank = new([]interface{})
	err = json.Unmarshal([]byte(test_json), &blank)
	c.Check(tx.Raw, gocheck.Equals, blank)
}

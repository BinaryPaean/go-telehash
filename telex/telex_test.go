package telex

import (
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
	tx, err := TelexFromJson(test_json)
	c.Check(err, gocheck.IsNil)
	h1 := &map[string]string{"_ring": "43723"}
	var headers = [](*map[string]string){h1}
	c.Check(tx.headers, gocheck.DeepEquals, headers)
}

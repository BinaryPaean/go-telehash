package telex

import (
	"encoding/json"
	"launchpad.net/gocheck"
	"testing"
)

const (
	simple_telex_json = `{
      "_ring":43723,
      ".see":["5.6.7.8:23456","11.22.33.44:11223"]
    }`

	signals_telex_json = `{
      "+end":"a9993e364706816aba3e25717850c26c9cd0d89d",
      "+foo":"0beec7b5ea3f0fdbc95d0dd47f3c5bc275da8a33"
    }`

	normal_telex_json = `{
      "_to":"1.2.3.4:5678",
      "_line":63546230,
      "profile_image_url": "http://a3.twimg.com/profile_images/852841481/Untitled_3_normal.jpg",
      "created_at": "Sat, 08 May 2010 21:46:23 +0000",
      "from_user": "pelchiie",
      "metadata": {
        "result_type": "recent"
      },
      "to_user_id": null,
      "text": "twitter is dead today.",
      "id": 13630378882,
      "from_user_id": 12621761,
      "geo": null,
      "iso_language_code": "en",
      "source": "<a href="http://twitter.com/">web</a>"
    }`
)

func Test(t *testing.T) { gocheck.TestingT(t) }

type TelexSuite struct{}

var _ = gocheck.Suite(&TelexSuite{})

func (s *TelexSuite) TestTelexFromJson(c *gocheck.C) {
	tx := TelexFromJson([]byte(simple_telex_json))
	c.Check(len(tx.Errors), gocheck.Equals, 0)
	c.Check(tx.Headers.Ring, gocheck.DeepEquals, int64(43723))
	c.Check(tx.Commands.See, gocheck.DeepEquals, []string{"5.6.7.8:23456", "11.22.33.44:11223"})
	blank := new(jsonObject)
	json.Unmarshal([]byte(simple_telex_json), &blank)
	c.Check(tx.Raw, gocheck.DeepEquals, *blank)
}

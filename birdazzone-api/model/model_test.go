package model

import (
	"testing"

	"git.hjkl.gq/team13/birdazzone-api/twitter"
	"git.hjkl.gq/team13/birdazzone-api/util"
	geojson "github.com/paulmach/go.geojson"
)

func TestPageQueryString(t *testing.T) {
	var pg *PageQuery
	if pg.String() != util.NilRepresentation {
		t.Fatalf("%s differs from %s", pg.String(), util.NilRepresentation)
	}
	pg = &PageQuery{10, 2}
	if pg.String() != "page #2" {
		t.Fatalf("%s differs from \"page #2\"", pg.String())
	}
}

func TestBooleanChartString(t *testing.T) {
	var bc *BooleanChart
	if bc.String() != util.NilRepresentation {
		t.Fatalf("%s differs from %s", bc.String(), util.NilRepresentation)
	}
	bc = &BooleanChart{"Votes", 52, 21}
	if bc.String() != "Votes: 52 VS 21" {
		t.Fatalf("%s differs from \"Votes: 52 VS 21\"", bc.String())
	}
}

func TestChartEntryString(t *testing.T) {
	var ce *ChartEntry
	if ce.String() != util.NilRepresentation {
		t.Fatalf("%s differs from %s", ce.String(), util.NilRepresentation)
	}
	ce = &ChartEntry{"Key", 42}
	if ce.String() != "(Key: 42)" {
		t.Fatalf("%s differs from \"(Key: 42)\"", ce.String())
	}
}

func TestGameString(t *testing.T) {
	var g *Game
	if g.String() != util.NilRepresentation {
		t.Fatalf("%s differs from %s", g.String(), util.NilRepresentation)
	}
	g = &Game{0, "My game", "mygame", "game-logo.png"}
	if g.String() != "#0 (My game #mygame)" {
		t.Fatalf("%s differs from \"#0 (My game #mygame)\"", g.String())
	}
}

func TestGameKeyString(t *testing.T) {
	var gk *GameKey
	if gk.String() != util.NilRepresentation {
		t.Fatalf("%s differs from %s", gk.String(), util.NilRepresentation)
	}
	gk = &GameKey{"crown", "2022-10-20"}
	if gk.String() != "2022-10-20: 'crown'" {
		t.Fatalf("%s differs from \"2022-10-20: 'crown'\"", gk.String())
	}
}

func TestPoliticianString(t *testing.T) {
	var p *Politician
	if p.String() != util.NilRepresentation {
		t.Fatalf("%s differs from %s", p.String(), util.NilRepresentation)
	}
	p = &Politician{"Matteo Salvini", 350}
	if p.String() != "Matteo Salvini (350 pts.)" {
		t.Fatalf("%s differs from \"Matteo Salvini (350 pts.)\"", p.String())
	}
}

func TestMakeUser(t *testing.T) {
	user := MakeUser(twitter.UIDLookup{
		Data: twitter.Profile{
			ID:              "2384732987",
			Name:            "Mario Rossi",
			Username:        "mariorossi",
			ProfileImageUrl: "a.org/img.png",
		},
	})
	if user.Username != "mariorossi" {
		t.Fatal("Wrong username")
	}
	if user.Name != "Mario Rossi" {
		t.Fatal("Wrong name")
	}
	if user.ProfileImageUrl != "a.org/img.png" {
		t.Fatal("Wrong profile image URL")
	}
}

func TestUserString(t *testing.T) {
	var u *User
	if u.String() != util.NilRepresentation {
		t.Fatalf("%s differs from %s", u.String(), util.NilRepresentation)
	}
	u = &User{"mariorossi", "Mario Rossi", "a.org/img.png"}
	if u.String() != "Mario Rossi (@mariorossi)" {
		t.Fatalf("%s differs from \"Mario Rossi (@mariorossi)\"", u.String())
	}
}

func TestMetricsString(t *testing.T) {
	var m *Metrics
	if m.String() != util.NilRepresentation {
		t.Fatalf("%s differs from %s", m.String(), util.NilRepresentation)
	}
	m = &Metrics{LikeCount: 2, ReplyCount: 6, RetweetCount: 10}
	if m.String() != "(Likes: 2, Replies: 6, Retweets: 10)" {
		t.Fatalf("%s differs from \"(Likes: 2, Replies: 6, Retweets: 10)\"", m.String())
	}
}

func TestOpenStreetMapCoordinatesString(t *testing.T) {
	var c *openStreetMapCoordinates
	if c.String() != util.NilRepresentation {
		t.Fatalf("%s differs from %s", c.String(), util.NilRepresentation)
	}
	c = &openStreetMapCoordinates{Lat: "2", Lon: "6"}
	if c.String() != "(2, 6)" {
		t.Fatalf("%s differs from \"(2, 6)\"", c.String())
	}
}

func TestCoordinatesString(t *testing.T) {
	var c *Coordinates
	if c.String() != util.NilRepresentation {
		t.Fatalf("%s differs from %s", c.String(), util.NilRepresentation)
	}
	c = &Coordinates{Latitude: 2, Longitude: 6}
	if c.String() != "(2.000000, 6.000000)" {
		t.Fatalf("%s differs from \"(2.000000, 6.000000)\"", c.String())
	}
}

func TestStringToCoordinatesOnEmptyString(t *testing.T) {
	_, err := StringToCoordinates("")
	if err == nil {
		t.Fatal("Error expected")
	}
}

func TestStringToCoordinateOnNonsensicalString(t *testing.T) {
	_, err := StringToCoordinates("Place that does not really exist")
	if err == nil {
		t.Fatal("Error expected")
	}
}

func TestStringToCoordinateOnBologna(t *testing.T) {
	r, err := StringToCoordinates("Bologna")
	if err != nil {
		t.Fatal(err)
	}
	if (int)(r.Latitude) != 44 {
		t.Fatalf("Got latitude %f instead of 44", r.Latitude)
	}
	if (int)(r.Longitude) != 11 {
		t.Fatalf("Got longitude %f instead of 11", r.Longitude)
	}
}

func TestMakeCoordinatesOnNull(t *testing.T) {
	if MakeCoordinates(nil, twitter.Profile{}) != nil {
		t.Fatal("Expected null result")
	}
}

func TestMakeCoordinatesOnNullWithUserLocation(t *testing.T) {
	c := MakeCoordinates(nil, twitter.Profile{Location: "Bologna"})
	if c == nil {
		t.Fatal("Null result")
	}
	if (int)(c.Latitude) != 44 {
		t.Fatalf("Got latitude %f instead of 44", c.Latitude)
	}
	if (int)(c.Longitude) != 11 {
		t.Fatalf("Got longitude %f instead of 11", c.Longitude)
	}
}

func TestMakeCoordinatesOnPoint(t *testing.T) {
	c := MakeCoordinates(geojson.NewPointGeometry([]float64{5, -3}), twitter.Profile{})
	if c == nil {
		t.Fatal("Null result")
	}
	if (int)(c.Latitude) != 5 {
		t.Fatalf("Got latitude %f instead of 5", c.Latitude)
	}
	if (int)(c.Longitude) != -3 {
		t.Fatalf("Got longitude %f instead of -3", c.Longitude)
	}
}

/*
func TestMakeCoordinatesOnBoundingBox(t *testing.T) {
	c := MakeCoordinates(&geojson.Geometry{Type: "unknown", BoundingBox: []float64{1, 2, 3, 4}}, twitter.Profile{})
	if c == nil {
		t.Fatal("Null result")
	}
	if (int)(c.Latitude) != 2 {
		t.Fatalf("Got latitude %f instead of 2", c.Latitude)
	}
	if (int)(c.Longitude) != 3 {
		t.Fatalf("Got longitude %f instead of 3", c.Longitude)
	}
}
*/

func TestMakeCoordinatesOnUnsupportedShape(t *testing.T) {
	c := MakeCoordinates(&geojson.Geometry{}, twitter.Profile{})
	if c != nil {
		t.Fatalf("Unexpected result")
	}
}

func TestMakeTweet(t *testing.T) {
	geo := geojson.NewPointGeometry([]float64{1, 2})
	tweet := MakeTweet(twitter.ProfileTweet{
		CreatedAt: "today",
		PublicMetrics: struct {
			LikeCount    int "json:\"like_count\""
			ReplyCount   int "json:\"reply_count\""
			RetweetCount int "json:\"retweet_count\""
		}{
			LikeCount:    1,
			ReplyCount:   5,
			RetweetCount: 9,
		},
		Text: "Hello, world!",
	}, twitter.Profile{
		Name:            "Mario Rossi",
		Username:        "mariorossi",
		ProfileImageUrl: "a.org/img.png",
	},
		geo,
	)
	if tweet.CreatedAt != "today" {
		t.Fatal("Wrong creation instant")
	}
	if tweet.Metrics.LikeCount != 1 {
		t.Fatal("Wrong like count")
	}
	if tweet.Metrics.ReplyCount != 5 {
		t.Fatal("Wrong reply count")
	}
	if tweet.Metrics.RetweetCount != 9 {
		t.Fatal("Wrong retweet count")
	}
	if tweet.Author.Username != "mariorossi" {
		t.Fatal("Wrong username")
	}
	if tweet.Author.Name != "Mario Rossi" {
		t.Fatal("Wrong name")
	}
	if tweet.Author.ProfileImageUrl != "a.org/img.png" {
		t.Fatal("Wrong profile image URL")
	}
	if tweet.Text != "Hello, world!" {
		t.Fatal("Wrong text")
	}
}

func TestTweetString(t *testing.T) {
	var tweet *Tweet
	if tweet.String() != util.NilRepresentation {
		t.Fatalf("%s differs from %s", tweet.String(), util.NilRepresentation)
	}
	tweet = &Tweet{
		Text: "Hello, world!",
		Author: User{
			Username:        "mariorossi",
			Name:            "Mario Rossi",
			ProfileImageUrl: "a.org/img.png",
		},
		CreatedAt: "yesterday",
		Metrics: Metrics{
			LikeCount:    2,
			ReplyCount:   6,
			RetweetCount: 10,
		},
	}
	solution := "\"Hello, world!\"\nMario Rossi (@mariorossi), yesterday\n(Likes: 2, Replies: 6, Retweets: 10)"
	if tweet.String() != solution {
		t.Fatalf("%s differs from \"%s\"", tweet.String(), solution)
	}
}

package model;

import "testing"

func TestPageQueryString(t *testing.T) {
  var pg *PageQuery
  if pg.String() != nilRepresentation {
    t.Fatalf("%s differs from %s", pg.String(), nilRepresentation)
  }
  pg = &PageQuery{10, 2}
  if pg.String() != "page #2" {
    t.Fatalf("%s differs from \"page #2\"", pg.String())
  }
}

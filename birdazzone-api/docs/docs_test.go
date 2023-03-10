// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"testing"

	"github.com/swaggo/swag"
)

func TestInit(t *testing.T) {
	const infoInstanceName = "swagger"
	s := swag.GetSwagger(infoInstanceName)
	if s == nil {
		t.Fatal("Swagger not found")
	}
	if s.ReadDoc() == "" {
		t.Fatal("Empty docs")
	}
}

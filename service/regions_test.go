package service_test

import (
	"testing"

	"github.com/dharnitski/golang-codebuild/service"
	"golang.org/x/text/language"
)

func TestFind(t *testing.T) {

	region, confidence := service.Find("be")
	expectedRegion := language.MustParseRegion("BY")

	if region != expectedRegion {
		t.Errorf("got %s; want %s", region, expectedRegion)
	}

	if confidence != language.Low {
		t.Errorf("got %s; want %s", confidence, language.Low)
	}
}

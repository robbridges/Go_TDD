package Concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "fake://fake.com" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	webSites := []string{
		"http://google.com",
		"http://ebay.com",
		"fake://fake.com",
	}

	want := map[string]bool{
		"http://google.com": true,
		"http://ebay.com":   true,
		"fake://fake.com":   false,
	}

	got := CheckWebsites(mockWebsiteChecker, webSites)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted %v, got %v", want, got)
	}
}

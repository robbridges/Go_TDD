package Concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "fake://fake.com" {
		return false
	}
	return true
}

func slowWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
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

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowWebsiteChecker, urls)
	}
}

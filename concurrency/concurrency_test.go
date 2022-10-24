package concurrency

import (
	"reflect"
	"testing"
)

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"google.com",
		"naver.com",
		"falseUrl.com",
	}
	want := map[string]bool{
		"google.com":   true,
		"naver.com":    true,
		"falseUrl.com": false,
	}
	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

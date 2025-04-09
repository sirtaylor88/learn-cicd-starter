package auth

import (
	"net/http"
	"testing"
)

func TestAuth(t *testing.T) {
	req, _ := http.NewRequest("GET", "https://google.com", nil)
	req.Header.Set("Authorization", "ApiKey 12345")

	want := "12345"
	got, _ := GetAPIKey(req.Header)
	if got != "12345" {
		t.Fatalf("expected: %v, got: %v", want, got)
	}

}

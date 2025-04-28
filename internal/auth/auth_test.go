package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	header := http.Header{}
	header.Set("Authorization", "ApiKey test-api-key-123")

	apiKey, err := GetAPIKey(header)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if apiKey != "test-api-key-123" {
		t.Fatalf("expected api key 'test-api-key-123', got '%s'", apiKey)
	}
}

package ttvideo

import (
	"encoding/json"
	"testing"
)

func TestApplyRequestEncodeAuthorizationDays(t *testing.T) {
	for _, req := range []ApplyRequest{{AuthorizationDays: 30}, {AuthorizedDays: 60}} {
		var body map[string]any
		if err := json.Unmarshal(req.Encode(), &body); err != nil {
			t.Fatal(err)
		}
		if _, ok := body["authorized_days"]; ok {
			t.Fatal("legacy authorized_days was encoded")
		}
		if _, ok := body["authorization_days"]; !ok {
			t.Fatal("authorization_days missing")
		}
	}
}

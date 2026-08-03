package app

import (
	"net/url"
	"testing"
)

func TestListRequest_Encode(t *testing.T) {
	t.Parallel()

	got, err := url.ParseQuery((&ListRequest{
		AdvertiserID:   "111",
		AppPlatformIDs: []string{"com.example.app", "123456"},
	}).Encode())
	if err != nil {
		t.Fatalf("ParseQuery() error = %v", err)
	}
	if got.Get("advertiser_id") != "111" {
		t.Errorf("advertiser_id = %q, want 111", got.Get("advertiser_id"))
	}
	if got.Get("app_platform_ids") != `["com.example.app","123456"]` {
		t.Errorf("app_platform_ids = %q, want JSON array", got.Get("app_platform_ids"))
	}
}

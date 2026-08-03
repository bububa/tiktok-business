package app

import (
	"net/url"
	"testing"
)

func TestInfoRequest_Encode(t *testing.T) {
	t.Parallel()

	got, err := url.ParseQuery((&InfoRequest{
		AdvertiserID: "111",
		AppID:        "222",
	}).Encode())
	if err != nil {
		t.Fatalf("ParseQuery() error = %v", err)
	}
	if got.Get("advertiser_id") != "111" || got.Get("app_id") != "222" {
		t.Errorf("query = %v, want advertiser_id=111 and app_id=222", got)
	}
}

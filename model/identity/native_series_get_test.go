package identity

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/bububa/tiktok-business/enum"
)

func TestNativeSeriesGetRequestEncode(t *testing.T) {
	req := NativeSeriesGetRequest{
		AdvertiserID:           "7000000000000000001",
		IdentityID:             "7000000000000000002",
		IdentityType:           enum.IdentityType_BC_AUTH_TT,
		IdentityAuthorizedBcID: "7000000000000000003",
		Cursor:                 42,
		Count:                  50,
	}

	values, err := url.ParseQuery(req.Encode())
	if err != nil {
		t.Fatalf("ParseQuery() error = %v", err)
	}
	want := map[string]string{
		"advertiser_id":             req.AdvertiserID,
		"identity_id":               req.IdentityID,
		"identity_type":             string(req.IdentityType),
		"identity_authorized_bc_id": req.IdentityAuthorizedBcID,
		"cursor":                    "42",
		"count":                     "50",
	}
	for key, wantValue := range want {
		if got := values.Get(key); got != wantValue {
			t.Errorf("%s = %q, want %q", key, got, wantValue)
		}
	}
}

func TestNativeSeriesGetResponseUnmarshal(t *testing.T) {
	const payload = `{
		"code": 0,
		"data": {
			"list": [{
				"native_series_id": "series-1",
				"native_series_name": "Series One",
				"native_series_cover_url": "https://example.com/cover.jpg",
				"native_series_total_episode": 12,
				"native_series_total_duration": 3600
			}],
			"cursor": 100,
			"has_more": true
		}
	}`

	var resp NativeSeriesGetResponse
	if err := json.Unmarshal([]byte(payload), &resp); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	if resp.Data == nil {
		t.Fatal("Data is nil")
	}
	if resp.Data.Cursor != 100 || !resp.Data.HasMore {
		t.Errorf("pagination = (%d, %t), want (100, true)", resp.Data.Cursor, resp.Data.HasMore)
	}
	if len(resp.Data.List) != 1 {
		t.Fatalf("len(List) = %d, want 1", len(resp.Data.List))
	}
	if got := resp.Data.List[0]; got.NativeSeriesID != "series-1" || got.NativeSeriesTotalEpisode != 12 || got.NativeSeriesTotalDuration != 3600 {
		t.Errorf("List[0] = %+v", got)
	}
}

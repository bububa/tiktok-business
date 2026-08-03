package tool

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/bububa/tiktok-business/enum"
)

func TestAvailableAttributionSourceRequestEncode(t *testing.T) {
	req := AvailableAttributionSourceRequest{
		AdvertiserID: "7000000000000000001",
		AppID:        "7000000000000000002",
	}

	values, err := url.ParseQuery(req.Encode())
	if err != nil {
		t.Fatalf("ParseQuery() error = %v", err)
	}
	if got := values.Get("advertiser_id"); got != req.AdvertiserID {
		t.Errorf("advertiser_id = %q, want %q", got, req.AdvertiserID)
	}
	if got := values.Get("app_id"); got != req.AppID {
		t.Errorf("app_id = %q, want %q", got, req.AppID)
	}
}

func TestAvailableAttributionSourceResponseUnmarshal(t *testing.T) {
	const payload = `{
		"code": 0,
		"data": {
			"app_attribution_source": ["MMP", "SAN"],
			"app_data_source": ["MMP", "EVENT_SDK", "EVENT_API"]
		}
	}`

	var resp AvailableAttributionSourceResponse
	if err := json.Unmarshal([]byte(payload), &resp); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	if resp.Data == nil {
		t.Fatal("Data is nil")
	}
	if got := resp.Data.AppAttributionSources; len(got) != 2 || got[1] != enum.AppAttributionSource_SAN {
		t.Errorf("AppAttributionSources = %v, want [MMP SAN]", got)
	}
	if got := resp.Data.AppDataSources; len(got) != 3 || got[2] != enum.AppDataSource_EVENT_API {
		t.Errorf("AppDataSources = %v, want [MMP EVENT_SDK EVENT_API]", got)
	}
}

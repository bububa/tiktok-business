package tool

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/bububa/tiktok-business/enum"
)

func TestMinisGetRequestEncode(t *testing.T) {
	req := MinisGetRequest{
		AdvertiserID: "7000000000000000001",
		Page:         2,
		PageSize:     25,
	}

	values, err := url.ParseQuery(req.Encode())
	if err != nil {
		t.Fatalf("ParseQuery() error = %v", err)
	}
	want := map[string]string{
		"advertiser_id": req.AdvertiserID,
		"page":          "2",
		"page_size":     "25",
	}
	for key, wantValue := range want {
		if got := values.Get(key); got != wantValue {
			t.Errorf("%s = %q, want %q", key, got, wantValue)
		}
	}
}

func TestMinisGetResponseUnmarshal(t *testing.T) {
	const payload = `{
		"code": 0,
		"data": {
			"list": [{
				"minis_id": "minis-1",
				"minis_name": "Mini One",
				"minis_icon_url": "https://example.com/icon.png",
				"minis_status": "ACTIVE",
				"minis_type": "MINI_GAME",
				"region_codes": ["US", "GB"]
			}],
			"page_info": {
				"page": 1,
				"page_size": 10,
				"total_number": 1,
				"total_page": 1
			}
		}
	}`

	var resp MinisGetResponse
	if err := json.Unmarshal([]byte(payload), &resp); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	if resp.Data == nil || resp.Data.PageInfo == nil {
		t.Fatal("Data or PageInfo is nil")
	}
	if resp.Data.PageInfo.TotalNumber != 1 {
		t.Errorf("PageInfo.TotalNumber = %d, want 1", resp.Data.PageInfo.TotalNumber)
	}
	if len(resp.Data.List) != 1 {
		t.Fatalf("len(List) = %d, want 1", len(resp.Data.List))
	}
	got := resp.Data.List[0]
	if got.MinisStatus != enum.MinisStatus_ACTIVE || got.MinisType != enum.MinisType_MINI_GAME {
		t.Errorf("Minis enum values = (%q, %q), want (ACTIVE, MINI_GAME)", got.MinisStatus, got.MinisType)
	}
}

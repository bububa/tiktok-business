package creator

import (
	"net/url"
	"testing"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
)

func TestDiscoverRequestEncode(t *testing.T) {
	req := DiscoverRequest{AccountID: "account", CountryCodes: []string{"US"}, MinFollowers: model.Int64(1000), MinEngagementRate: model.Float64(0.25), FollowerAge: enum.TTOAudienceAge_18_24, SortField: enum.TTOCreatorSortField_FOLLOWERS}
	values, err := url.ParseQuery(req.Encode())
	if err != nil {
		t.Fatal(err)
	}
	if got := values.Get("country_codes"); got != `["US"]` {
		t.Fatalf("country_codes = %q", got)
	}
	if got := values.Get("min_followers"); got != "1000" {
		t.Fatalf("min_followers = %q", got)
	}
	if got := values.Get("min_engagement_rate"); got != "0.25" {
		t.Fatalf("min_engagement_rate = %q", got)
	}
	if got := values.Get("follower_age"); got != "18-24" {
		t.Fatalf("follower_age = %q", got)
	}
}

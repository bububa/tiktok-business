package report

import (
	"net/http"
	"net/url"
	"testing"
)

func TestGetRequestEncodeAndHeader(t *testing.T) {
	req := GetRequest{AccountID: "account", CampaignID: "campaign", CountryCode: "US", StartDate: "2026-01-01", Page: 2}
	values, err := url.ParseQuery(req.Encode())
	if err != nil {
		t.Fatal(err)
	}
	if got := values.Get("tto_tcm_account_id"); got != "account" {
		t.Fatalf("account ID = %q", got)
	}
	if got := values.Get("campaign_id"); got != "campaign" {
		t.Fatalf("campaign ID = %q", got)
	}
	httpReq, err := http.NewRequest(http.MethodGet, "https://example.com", nil)
	if err != nil {
		t.Fatal(err)
	}
	if err := req.PreRequest(httpReq); err != nil {
		t.Fatal(err)
	}
	if got := httpReq.Header.Get("Country-Code"); got != "US" {
		t.Fatalf("Country-Code = %q", got)
	}
}

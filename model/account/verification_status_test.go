package account

import (
	"encoding/json"
	"testing"

	"github.com/bububa/tiktok-business/enum"
)

func TestVerificationStatusRequest_Encode(t *testing.T) {
	t.Parallel()

	req := VerificationStatusRequest{AdvertiserID: "advertiser"}
	const expected = "advertiser_id=advertiser"
	if got := req.Encode(); got != expected {
		t.Fatalf("Encode() = %q, want %q", got, expected)
	}
}

func TestVerificationStatusResponse_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	const data = `{
		"code": 0,
		"message": "OK",
		"data": {
			"audit_time": "2026-07-08 10:30:58",
			"qualification_id": "qualification",
			"rejection_reason": "reason",
			"verification_status": "FAILED"
		}
	}`
	var resp VerificationStatusResponse
	if err := json.Unmarshal([]byte(data), &resp); err != nil {
		t.Fatal(err)
	}
	if resp.Data == nil {
		t.Fatal("Data is nil")
	}
	if resp.Data.VerificationStatus != enum.BcVerificationStatus_FAILED {
		t.Errorf("VerificationStatus = %q, want %q", resp.Data.VerificationStatus, enum.BcVerificationStatus_FAILED)
	}
	if got := resp.Data.AuditTime.String(); got != "2026-07-08 10:30:58" {
		t.Errorf("AuditTime = %q, want 2026-07-08 10:30:58", got)
	}
}

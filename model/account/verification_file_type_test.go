package account

import (
	"testing"

	"github.com/bububa/tiktok-business/enum"
)

func TestVerificationFileTypeRequest_Encode(t *testing.T) {
	t.Parallel()

	req := VerificationFileTypeRequest{
		BusinessType:  enum.VerificationBusinessType_BUSINESS,
		RegionISOCode: "US",
	}
	const expected = "business_type=BUSINESS&region_iso_code=US"
	if got := req.Encode(); got != expected {
		t.Fatalf("Encode() = %q, want %q", got, expected)
	}
}

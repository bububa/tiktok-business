package account

import (
	"encoding/json"
	"testing"

	"github.com/bububa/tiktok-business/enum"
)

func TestVerificationSubmitRequest_Encode(t *testing.T) {
	t.Parallel()

	req := VerificationSubmitRequest{
		AdvertiserID: "advertiser",
		BusinessForm: &VerificationBusinessForm{
			CompanyName:           "Example Inc.",
			WebsiteURL:            "https://example.com",
			IndustryCode:          enum.VerificationIndustryCode_OTC,
			RegionISOCode:         "US",
			FileTypeCode:          "127724548",
			LicenseNo:             "license",
			QualificationImageIDs: []string{"image"},
		},
	}
	var body struct {
		AdvertiserID string `json:"advertiser_id"`
		BusinessForm struct {
			IndustryCode          enum.VerificationIndustryCode `json:"industry_code"`
			QualificationImageIDs []string                      `json:"qualification_image_ids"`
		} `json:"business_form"`
	}
	if err := json.Unmarshal(req.Encode(), &body); err != nil {
		t.Fatal(err)
	}
	if body.AdvertiserID != "advertiser" {
		t.Errorf("AdvertiserID = %q, want advertiser", body.AdvertiserID)
	}
	if body.BusinessForm.IndustryCode != enum.VerificationIndustryCode_OTC {
		t.Errorf("IndustryCode = %q, want %q", body.BusinessForm.IndustryCode, enum.VerificationIndustryCode_OTC)
	}
	if len(body.BusinessForm.QualificationImageIDs) != 1 || body.BusinessForm.QualificationImageIDs[0] != "image" {
		t.Errorf("QualificationImageIDs = %v, want [image]", body.BusinessForm.QualificationImageIDs)
	}
}

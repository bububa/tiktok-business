package account

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/util"
)

// VerificationSubmitRequest is the request for submitting account verification.
type VerificationSubmitRequest struct {
	// AdvertiserID is the ad account to verify. Either AdvertiserID or BcID is required.
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// BcID is the Business Center account to verify. Either AdvertiserID or BcID is required.
	BcID string `json:"bc_id,omitempty"`
	// BusinessForm contains business verification information.
	BusinessForm *VerificationBusinessForm `json:"business_form,omitempty"`
	// IndividualForm contains individual identity verification information.
	IndividualForm *VerificationIndividualForm `json:"individual_form,omitempty"`
}

// Encode implements model.PostRequest.
func (r *VerificationSubmitRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// VerificationBusinessForm contains business verification information.
type VerificationBusinessForm struct {
	// CompanyName is the business name exactly as shown on the verification document.
	CompanyName string `json:"company_name,omitempty"`
	// WebsiteURL is the company website URL.
	WebsiteURL string `json:"website_url,omitempty"`
	// IndustryCode is the business industry category.
	IndustryCode enum.VerificationIndustryCode `json:"industry_code,omitempty"`
	// RegionISOCode is the ISO 3166 country or region code of the account.
	RegionISOCode string `json:"region_iso_code,omitempty"`
	// FileTypeCode is the verification document type code.
	FileTypeCode string `json:"file_type_code,omitempty"`
	// LicenseNo is the certificate number on the verification document.
	LicenseNo string `json:"license_no,omitempty"`
	// QualificationImageIDs contains exactly one uploaded business verification document ID.
	QualificationImageIDs []string `json:"qualification_image_ids,omitempty"`
}

// VerificationIndividualForm contains individual identity verification information.
type VerificationIndividualForm struct {
	// IndividualName is the full legal name exactly as shown on the personal ID.
	IndividualName string `json:"individual_name,omitempty"`
	// RegionISOCode is the ISO 3166 country or region code of the account.
	RegionISOCode string `json:"region_iso_code,omitempty"`
	// FileTypeCode is the personal identification document type code.
	FileTypeCode string `json:"file_type_code,omitempty"`
	// IdentityNo is the ID number on the personal identification document.
	IdentityNo string `json:"identity_no,omitempty"`
	// WebsiteURL is the company website URL.
	WebsiteURL string `json:"website_url,omitempty"`
	// QualificationImageIDs contains the uploaded front and back personal ID image IDs.
	QualificationImageIDs []string `json:"qualification_image_ids,omitempty"`
}

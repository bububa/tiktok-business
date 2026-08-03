package account

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// VerificationFileTypeRequest is the request for available account verification document types.
type VerificationFileTypeRequest struct {
	// BusinessType is the type of verification to perform.
	BusinessType enum.VerificationBusinessType `json:"business_type,omitempty"`
	// RegionISOCode is the ISO 3166 country or region code of the account.
	RegionISOCode string `json:"region_iso_code,omitempty"`
}

// Encode implements model.GetRequest.
func (r *VerificationFileTypeRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("business_type", string(r.BusinessType))
	values.Set("region_iso_code", r.RegionISOCode)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// VerificationFileTypeResponse is the response for available account verification document types.
type VerificationFileTypeResponse struct {
	model.BaseResponse
	Data *VerificationFileTypeResult `json:"data,omitempty"`
}

// VerificationFileTypeResult contains available account verification document types.
type VerificationFileTypeResult struct {
	FileTypes []VerificationFileType `json:"file_types,omitempty"`
}

// VerificationFileType describes an available verification document type.
type VerificationFileType struct {
	// FileTypeName is the display name of the verification document.
	FileTypeName string `json:"file_type_name,omitempty"`
	// FileTypeCode is the code passed when submitting a verification request.
	FileTypeCode string `json:"file_type_code,omitempty"`
}

package bc

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// OACreateRequest Create an Organization Account in a Business Center API Request
type OACreateRequest struct {
	// BcID The ID of the Business Center.
	BcID string `json:"bc_id,omitempty"`
	// DisplayName The display name (username) for the Organization Account.
	// Length limit: 30 characters.
	DisplayName string `json:"display_name,omitempty"`
	// Handle The handle (name) for the Organization Account.
	// Length limit: 24 characters.
	// The handle cannot consist solely of numbers.
	// If the handle is already taken, an error will occur.
	Handle string `json:"handle,omitempty"`
	// ProfileImage The file to be used as the profile image for the Organization Account.
	ProfileImage string `json:"profile_image,omitempty"`
	// OperatingRegionCode The code representing the country or region where you would like to run your business.
	// For enum values, see List of values for operating_region_code.
	// Note: If your business is based in the Chinese mainland, you must select an operating region other than China when creating an Organization Account, as CN is not an available option for operating_region_code.
	OperatingRegionCode string `json:"operating_region_code,omitempty"`
	// QualificationInfo Required when the type of your Business Center is AGENCY or SELF_SERVICE_AGENCY.
	// Qualification information.
	// To obtain the type of your Business Center, use /bc/get/ and check the returned type.
	// If the type of your Business Center is DIRECT or SELF_SERVICE, an error will occur.
	QualificationInfo *Qualification `json:"qualification_info,omitempty"`
}

// Encode implements PostRequest
func (r *OACreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// OACreateResponse Create an Organization Account in a Business Center API Response
type OACreateResponse struct {
	model.BaseResponse
	Data *Asset `json:"data,omitempty"`
}

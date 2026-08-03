package account

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// VerificationStatusRequest is the request for an account verification status.
type VerificationStatusRequest struct {
	// AdvertiserID is the ad account to query. Either AdvertiserID or BcID is required.
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// BcID is the Business Center account to query. Either AdvertiserID or BcID is required.
	BcID string `json:"bc_id,omitempty"`
}

// Encode implements model.GetRequest.
func (r *VerificationStatusRequest) Encode() string {
	values := util.NewURLValues()
	if r.AdvertiserID != "" {
		values.Set("advertiser_id", r.AdvertiserID)
	}
	if r.BcID != "" {
		values.Set("bc_id", r.BcID)
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// VerificationStatusResponse is the response for an account verification status.
type VerificationStatusResponse struct {
	model.BaseResponse
	Data *VerificationStatusResult `json:"data,omitempty"`
}

// VerificationStatusResult contains the current account verification status.
type VerificationStatusResult struct {
	// QualificationID is the qualification ID associated with the verified account.
	QualificationID string `json:"qualification_id,omitempty"`
	// VerificationStatus is the current verification state.
	VerificationStatus enum.BcVerificationStatus `json:"verification_status,omitempty"`
	// RejectionReason is returned when VerificationStatus is FAILED.
	RejectionReason string `json:"rejection_reason,omitempty"`
	// AuditTime is the UTC time when the verification review was completed.
	AuditTime model.DateTime `json:"audit_time,omitzero"`
}

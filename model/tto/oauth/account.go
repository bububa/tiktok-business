package oauth

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

type AccountListRequest struct {
	AppID  string `json:"app_id,omitempty"`
	Secret string `json:"secret,omitempty"`
}

func (r *AccountListRequest) Encode() string {
	v := util.NewURLValues()
	v.Set("app_id", r.AppID)
	v.Set("secret", r.Secret)
	ret := v.Encode()
	util.ReleaseURLValues(v)
	return ret
}

type AccountListResponse struct {
	model.BaseResponse
	Data struct {
		AccountIDs []string `json:"tto_tcm_account_ids,omitempty"`
	} `json:"data"`
}

type AccountGetRequest struct {
	AccountID string `json:"tto_tcm_account_id,omitempty"`
}

func (r *AccountGetRequest) Encode() string {
	v := util.NewURLValues()
	v.Set("tto_tcm_account_id", r.AccountID)
	ret := v.Encode()
	util.ReleaseURLValues(v)
	return ret
}

type AccountGetResponse struct {
	model.BaseResponse
	Data *Account `json:"data,omitempty"`
}

type Account struct {
	AccountID                  string `json:"tto_tcm_account_id,omitempty"`
	AccountName                string `json:"account_name,omitempty"`
	Timezone                   string `json:"timezone,omitempty"`
	Country                    string `json:"country,omitempty"`
	BusinessVerificationStatus string `json:"business_verification_status,omitempty"`
}

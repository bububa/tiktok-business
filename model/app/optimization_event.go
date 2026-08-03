package app

import (
	"encoding/json"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// OptimizationEventRequest gets app conversion events available for ad-group optimization.
type OptimizationEventRequest struct {
	AppID            string                `json:"app_id,omitempty"`
	AdvertiserID     string                `json:"advertiser_id,omitempty"`
	Placements       []enum.Placement      `json:"placement,omitempty"`
	PlacementType    enum.PlacementType    `json:"placement_type,omitempty"`
	OptimizationGoal enum.OptimizationGoal `json:"optimization_goal,omitempty"`
	Objective        enum.ObjectiveType    `json:"objective,omitempty"`
	AvailableOnly    *bool                 `json:"available_only,omitempty"`
	IsSkan           *bool                 `json:"is_skan,omitempty"`
	AppPromotionType enum.AppPromotionType `json:"app_promotion_type,omitempty"`
}

// Encode implements model.GetRequest.
func (r *OptimizationEventRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("app_id", r.AppID)
	values.Set("advertiser_id", r.AdvertiserID)
	if len(r.Placements) > 0 {
		values.Set("placement", string(util.JSONMarshal(r.Placements)))
	}
	if r.PlacementType != "" {
		values.Set("placement_type", string(r.PlacementType))
	}
	values.Set("optimization_goal", string(r.OptimizationGoal))
	if r.Objective != "" {
		values.Set("objective", string(r.Objective))
	}
	if r.AvailableOnly != nil {
		values.Set("available_only", boolString(*r.AvailableOnly))
	}
	if r.IsSkan != nil {
		values.Set("is_skan", boolString(*r.IsSkan))
	}
	if r.AppPromotionType != "" {
		values.Set("app_promotion_type", string(r.AppPromotionType))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

func boolString(v bool) string {
	if v {
		return "true"
	}
	return "false"
}

// OptimizationEventResponse is the response to OptimizationEventRequest.
type OptimizationEventResponse struct {
	model.BaseResponse
	Data struct {
		OptimizationEvents []ConversionEvent `json:"optimization_events,omitempty"`
	} `json:"data"`
}

// ConversionEvent describes an app conversion or retargeting event.
// Some TikTok integrations return retargeting events as strings, while others
// return objects. UnmarshalJSON accepts both documented response shapes.
type ConversionEvent struct {
	OptimizationEvent enum.OptimizationEvent `json:"optimization_event,omitempty"`
	EventName         string                 `json:"event_name,omitempty"`
	EventIsAvailable  bool                   `json:"event_is_available,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *ConversionEvent) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == '"' {
		return json.Unmarshal(data, &e.OptimizationEvent)
	}
	type conversionEvent ConversionEvent
	return json.Unmarshal(data, (*conversionEvent)(e))
}

// OptimizationEventRetargetingRequest gets app retargeting events.
type OptimizationEventRetargetingRequest struct {
	AppID        string `json:"app_id,omitempty"`
	AdvertiserID string `json:"advertiser_id,omitempty"`
}

// Encode implements model.GetRequest.
func (r *OptimizationEventRetargetingRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("app_id", r.AppID)
	values.Set("advertiser_id", r.AdvertiserID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// OptimizationEventRetargetingResponse is the response to OptimizationEventRetargetingRequest.
type OptimizationEventRetargetingResponse struct {
	model.BaseResponse
	Data struct {
		OptimizationEvents []ConversionEvent `json:"optimization_events,omitempty"`
	} `json:"data"`
}

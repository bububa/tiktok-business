package app

import (
	"time"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/util"
)

// TrackRequest reports a single app event.
type TrackRequest struct {
	Event
	// TikTokAppID is the TikTok App ID from Events Manager.
	TikTokAppID string `json:"tiktok_app_id,omitempty"`
	// EventSource identifies the event reporting integration.
	EventSource enum.AppEventSource `json:"event_source,omitempty"`
	// PartnerName is the integration partner name.
	PartnerName string `json:"partner_name,omitempty"`
	// TestEventCode sends the event to the corresponding test-event session.
	TestEventCode string `json:"test_event_code,omitempty"`
}

// Encode implements model.PostRequest.
func (r *TrackRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// BatchRequest reports multiple app events in one request.
type BatchRequest struct {
	// TikTokAppID is the TikTok App ID from Events Manager.
	TikTokAppID string `json:"tiktok_app_id,omitempty"`
	// Batch contains the app events to report.
	Batch []Event `json:"batch,omitempty"`
	// EventSource identifies the event reporting integration.
	EventSource enum.AppEventSource `json:"event_source,omitempty"`
	// PartnerName is the integration partner name.
	PartnerName string `json:"partner_name,omitempty"`
	// TestEventCode sends the events to the corresponding test-event session.
	TestEventCode string `json:"test_event_code,omitempty"`
}

// Encode implements model.PostRequest.
func (r *BatchRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// Event is an app event in an Events API 1.0 request.
type Event struct {
	// Type is required for events in a batch and must be track.
	Type enum.AppEventAction `json:"type,omitempty"`
	// Event is a standard or custom app event name.
	Event enum.AppEventType `json:"event,omitempty"`
	// Timestamp is the event time in RFC 3339 format.
	Timestamp time.Time `json:"timestamp"`
	// Context contains app, device, user, attribution, and location data.
	Context *EventContext `json:"context,omitempty"`
	// Properties contains event-specific commerce data.
	Properties *EventProperties `json:"properties,omitempty"`
}

// EventContext contains context used to match and attribute an app event.
type EventContext struct {
	App                        *EventApp `json:"app,omitempty"`
	Device                     *Device   `json:"device,omitempty"`
	Locale                     string    `json:"locale,omitempty"`
	IP                         string    `json:"ip,omitempty"`
	UserAgent                  string    `json:"user_agent,omitempty"`
	User                       *User     `json:"user,omitempty"`
	Library                    *Library  `json:"library,omitempty"`
	OriginalAttributionPartner string    `json:"ori_attribution_partner,omitempty"`
	OriginURL                  string    `json:"origin_url,omitempty"`
	Ad                         *Ad       `json:"ad,omitempty"`
	Location                   *Location `json:"location,omitempty"`
}

// EventApp describes the app that generated an event.
type EventApp struct {
	ID           string   `json:"id,omitempty"`
	Name         string   `json:"name,omitempty"`
	Namespace    string   `json:"namespace,omitempty"`
	Version      string   `json:"version,omitempty"`
	Build        string   `json:"build,omitempty"`
	TikTokAppIDs []string `json:"tiktok_app_ids,omitempty"`
}

// Device describes the device that generated an event.
type Device struct {
	ATTStatus enum.AppEventATTStatus `json:"att_status,omitempty"`
	Platform  enum.AppEventPlatform  `json:"platform,omitempty"`
	IDFA      string                 `json:"idfa,omitempty"`
	IDFV      string                 `json:"idfv,omitempty"`
	GAID      string                 `json:"gaid,omitempty"`
}

// User contains customer matching identifiers.
type User struct {
	AnonymousID string `json:"anonymous_id,omitempty"`
	ExternalID  string `json:"external_id,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Email       string `json:"email,omitempty"`
	PGCHID      string `json:"pgchid,omitempty"`
}

// Library describes the client library that generated an event.
type Library struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}

// Ad contains attribution metadata for an app event.
type Ad struct {
	Callback            string `json:"callback,omitempty"`
	IsRetargeting       string `json:"isRetargeting,omitempty"`
	CampaignID          string `json:"campaign_id,omitempty"`
	Attributed          string `json:"attributed,omitempty"`
	AdID                string `json:"ad_id,omitempty"`
	AttributionType     string `json:"attributionType,omitempty"`
	CreativeID          string `json:"creative_id,omitempty"`
	AttributionProvider string `json:"attribution_provider,omitempty"`
}

// Location contains coarse event location data.
type Location struct {
	Country string `json:"country,omitempty"`
	City    string `json:"city,omitempty"`
}

// EventProperties contains commerce and content information for an event.
type EventProperties struct {
	Contents    []Content `json:"contents,omitempty"`
	ContentID   string    `json:"content_id,omitempty"`
	ContentType string    `json:"content_type,omitempty"`
	Currency    string    `json:"currency,omitempty"`
	Value       float64   `json:"value,omitempty"`
	Quantity    int       `json:"quantity,omitempty"`
	Price       float64   `json:"price,omitempty"`
	Description string    `json:"description,omitempty"`
	Query       string    `json:"query,omitempty"`
	Status      string    `json:"status,omitempty"`
}

// Content describes an item associated with an event.
type Content struct {
	Price           float64 `json:"price,omitempty"`
	Quantity        int     `json:"quantity,omitempty"`
	ContentType     string  `json:"content_type,omitempty"`
	ContentID       string  `json:"content_id,omitempty"`
	ContentName     string  `json:"content_name,omitempty"`
	ContentCategory string  `json:"content_category,omitempty"`
	Brand           string  `json:"brand,omitempty"`
}

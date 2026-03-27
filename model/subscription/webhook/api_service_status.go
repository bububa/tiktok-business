package webhook

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
)

type APIServiceStatusEntity struct {
	// IncidentID Incident ID
	IncidentID string `json:"incident_id,omitempty"`
	// IncidentTitle Incident title
	IncidentTitle string `json:"incident_title,omitempty"`
	// IncidentProducts he list of API products affected by the incident.
	// Enum values:
	// BUSINESS_CENTER_API: Business Center API.
	// CREATIVES_API: Creatives API.
	// CATALOG_API: Catalog Management API.
	// TIKTOK_STORE_API: TikTok Store API.
	// CAMPAIGN_API: Campaign Management API.
	// REPORTING_API: Reporting API.
	// AUDIENCE_API: Audience API.
	// STREAMING_API: Streaming API.
	// EVENTS_API: Events API.
	// ACCOUNTS_API: Accounts API.
	// MENTIONS_API: Mentions API.
	// TIKTOK_ONE_API: TikTok One API.
	// DISCOVERY_API: Discovery API.
	// SPARK_RECOMMEND_API: Spark Ads Recommendation API.
	// BUSINESS_MESSAGING_API: Business Messaging API.
	IncidentProducts []enum.APIServiceType `json:"incident_products,omitempty"`
	// IncidentEventType The type of event related to the incident.
	// Enum values:
	// CREATE: The incident occurs.
	// UPDATE: The incident is updated
	// RESOLVE: The incident is resolved.
	IncidentEventType string `json:"incident_event_type,omitempty"`
	// IncidentLevel The disruption level for the incident.
	// Enum values:
	// LOW_DISRUPTION: Low Disruption. Minor issues like elevated latency or timeouts may occur, but the majority of requests succeed.
	// MAJOR_DISRUPTION: Major Disruption. Critical service impact where most or all API calls are failing or latency exceeds SLA thresholds globally.
	IncidentLevel string `json:"incident_level,omitempty"`
	// IncidentStartTime The start time of the incident, in the format of YYYY-MM-DD HH:MM:SS (UTC time).
	IncidentStartTime model.DateTime `json:"incident_start_time,omitzero"`
	// IncidentResolveTime The resolution time of the incident, in the format of YYYY-MM-DD HH:MM:SS (UTC time).
	IncidentResolveTime model.DateTime `json:"incident_resolve_time,omitzero"`
	// IncidentDetail Incident description
	IncidentDetail string `json:"incident_detail,omitempty"`
}

func (e APIServiceStatusEntity) Entity() enum.SubscribeEntity {
	return enum.SubscribeEntity_API_SERVICE_STATUS
}

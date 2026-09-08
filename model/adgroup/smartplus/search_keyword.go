package smartplus

import "github.com/bububa/tiktok-business/enum"

// SearchKeyword Required for Search Ads Campaigns (is_search_campaign is true at the campaign level).
type SearchKeyword struct {
	// Keyword When search_keywords is specified, you need to provide at least one keyword and the corresponding match_type.
	// The search keyword.
	// Length limit: 80 characters.
	// The name needs to exclude emojis and the following special characters: ! # $ % & ( ) * - / : ; < > ? @ \ ^ _ ¥ ……. Otherwise, an error will occur.
	// To obtain the review result for search keywords, call /smart_plus/adgroup/get/ and check the audit_status within the search_keywords object array.
	// Important: Search keywords must comply with TikTok's Advertising Terms and Community Guidelines.
	// If all of your keywords are found to be non-compliant, your campaign will not be delivered.
	// If part of your search keywords are rejected, you need to remove the violative keywords or replace them with a suitable alternative by using /smart_plus/adgroup/update/. Your campaign will continue with the approved keywords.
	Keyword string `json:"keyword,omitempty"`
	// MatchType When search_keywords is specified, you need to provide at least one keyword and the corresponding match_type.
	// The match type for the search keyword.
	// Enum values:
	// PRECISE_WORD: exact match. Ads only show when a search includes the exact phrase or close variations, with additional words allowed before or after.
	// PHRASE_WORD: phrase match. Ads only show when a search exactly matches your keyword or a very close variation.
	// BROAD_WORD: broad match. Ads are shown for a wide range of relevant searches to offer the widest reach.
	MatchType enum.KeywordMatchType `json:"match_type,omitempty"`
	// AuditStatus The review status of the search keyword.
	// Enum values:
	// AUDITING: The keyword is under review.
	// PASS: The keyword has passed review and can be delivered.
	// REJECTED: The keyword failed to pass the review and cannot be delivered.
	AuditStatus enum.KeywordAuditStatus `json:"audit_status,omitempty"`
	// RejectInfo Returned only when audit_status is REJECTED.
	// Details about the rejection.
	RejectInfo []KeywordRejectInfo `json:"reject_info,omitempty"`
}

// KeywordRejectInfo Returned only when audit_status is REJECTED.
type KeywordRejectInfo struct {
	// ForbiddenLocation The targeted region that failed the review.
	// For enum values, see Appendix - Location codes.
	ForbiddenLocation string `json:"forbidden_location,omitempty"`
	// RejectReasons List of rejection reasons.
	RejectReasons []KeywordRejectReason `json:"reject_reasons,omitempty"`
}

// KeywordRejectReason rejection reason of a keyword.
type KeywordRejectReason struct {
	// Reason The rejection reason.
	Reason string `json:"reason,omitempty"`
}

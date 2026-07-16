package enum

type TTOLabelType string

const (
	TTOLabelType_RANKING TTOLabelType = "RANKING"
	TTOLabelType_SEARCH  TTOLabelType = "SEARCH"
)

type TTORankingType string

const (
	TTORankingType_BRANDED_CONTENT TTORankingType = "BRANDED_CONTENT"
	TTORankingType_ORGANIC_CONTENT TTORankingType = "ORGANIC_CONTENT"
)

type TTORankingPeriod string

const (
	TTORankingPeriod_WEEK  TTORankingPeriod = "WEEK"
	TTORankingPeriod_MONTH TTORankingPeriod = "MONTH"
)

type TTORankingLookback string

const (
	TTORankingLookback_ONE   TTORankingLookback = "ONE"
	TTORankingLookback_TWO   TTORankingLookback = "TWO"
	TTORankingLookback_THREE TTORankingLookback = "THREE"
)

type TTOCreatorSortField string

const (
	TTOCreatorSortField_RELEVANCE       TTOCreatorSortField = "RELEVANCE"
	TTOCreatorSortField_FOLLOWERS       TTOCreatorSortField = "FOLLOWERS"
	TTOCreatorSortField_MEDIAN_VIEWS    TTOCreatorSortField = "MEDIAN_VIEWS"
	TTOCreatorSortField_ENGAGEMENT_RATE TTOCreatorSortField = "ENGAGEMENT_RATE"
)

type TTOSortOrder string

const (
	TTOSortOrder_ASC  TTOSortOrder = "ASC"
	TTOSortOrder_DESC TTOSortOrder = "DESC"
)

type TTOAudienceAge string

const (
	TTOAudienceAge_18_24   TTOAudienceAge = "18-24"
	TTOAudienceAge_25_34   TTOAudienceAge = "25-34"
	TTOAudienceAge_35_44   TTOAudienceAge = "35-44"
	TTOAudienceAge_45_54   TTOAudienceAge = "45-54"
	TTOAudienceAge_55_PLUS TTOAudienceAge = "55+"
)

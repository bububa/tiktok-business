package smartplus

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
)

// Campaign  Upgraded Smart+ Campaign details
type Campaign struct {
	// AdvertiserID Advertiser ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CampaignID Campaign ID
	CampaignID string `json:"campaign_id,omitempty"`
	// CreateTime The time when the campaign was created, in the format of YYYY-MM-DD HH:MM:SS (UTC time)
	CreateTime model.DateTime `json:"create_time,omitzero"`
	// ModifyTime The time when the campaign was last modified, in the format of YYYY-MM-DD HH:MM:SS (UTC time)
	ModifyTime model.DateTime `json:"modify_time,omitzero"`
	// ObjectiveType Advertising objective.
	// For enum values, see Enumeration-Advertising Objective.
	ObjectiveType enum.ObjectiveType `json:"objective_type,omitempty"`
	// AppPromotionType App promotion type.
	// Enum value:
	// APP_INSTALL: App install. Get new users to install your app.
	// APP_RETARGETING: App retargeting. Re-engage existing app users to take action in your app.
	// MINIS: TikTok Minis. Get people to watch your series or play your games with TikTok Minis.
	AppPromotionType enum.AppPromotionType `json:"aepp_promotion_type,omitempty"`
	// SalesDestination Sales destination, the destination where you want to drive your sales.
	// Enum values:
	// WEBSITE: Website. Drive sales on your website.
	// APP: App. Drive sales on your app (product catalog required).
	// WEB_AND_APP: Website and app. Drive sales on both your website and your app.
	SalesDestination enum.SalesDestination `json:"sales_destination,omitempty"`
	// CampaignType Campaign type.
	// Currently, we support REGULAR_CAMPAIGN and IOS14_CAMPAIGN
	CampaignType enum.CampaignType `json:"campaign_type,omitempty"`
	// AppID The App ID of the app to promote.
	// To get a list of App IDs, use /app/list/.
	AppID string `json:"app_id,omitempty"`
	// GamingAdComplianceAgreement Whether to agree to the Compliance Assurance Policy for Gaming Advertisers on TikTok.
	// The policy is as follows: Yo confirms and attest that any gaming application, product or service (game) you desire to advertise on TikTok, including any associated URL(s), (a) complies with all applicable laws and regulations of the jurisdictions where the game can be accessed or played, and upon request, can provide supporting documentation as evidence of why the game is not considered illegal gambling or lottery; and (b) has not been and is not part of any investigation or lawsuit regarding the game's legality or regulatory compliance.
	// Enum values:
	// ON: To agree to the policy.
	// OFF: To leave the policy not accepted.
	GamingAdComplianceAgreement enum.OnOff `json:"gaming_ad_compliance_agreement,omitempty"`
	// IsAdvancedDedicatedCampaign Whether the campaign is an Advanced Dedicated Campaign. Advanced Dedicated Campaigns leverage advanced delivery models that benefit from real-time signals.
	// Supported values: true, false.
	// By default, Upgraded Smart+ App Campaigns with iOS 14 dedicated campaign setting enabled are all Advanced Dedicated Campaigns.
	IsAdvancedDedicatedCampaign bool `json:"is_advanced_dedicated_campaign,omitempty"`
	// DisableSkanCampaign Whether to disable SKAN (SKAdNetwork) attribution, Apple's conversion attribution solution for iOS campaigns.
	// Enum values:
	// true: To disable SKAN attribution. The campaign will not be bound by Dedicated Campaign quota limits and you will be able to retrieve Self Attribution Network (SAN) metrics for the campaign. However, you cannot retrieve SKAN reporting metrics for the campaign. Learn more about SAN integration.
	// false: To enable SKAN attribution. The campaign will be bound by Dedicated Campaign quota limits and you will be able to retrieve SKAN metrics for the campaign.
	DisableSkanCampaign bool `json:"disable_skan_campaign,omitempty"`
	// BidAlignType The attribution type for the Dedicated Campaign. The type determines which attribution network the target CPA (conversion_bid_price) or ROAS (roas_bid) at the ad group level is based on.
	// Enum values:
	// SAN: SAN.
	// SKAN: SKAN.
	BidAlignType enum.BidAlignType `json:"bid_atlign_type,omitempty"`
	// CampaignAppProfilePageState Indicates the status of the App Profile Page.
	// Enum values: ON, OFF.
	CampaignAppProfilePageState enum.AppProfilePageState `json:"campaign_app_profile_page_state,omitempty"`
	// CatalogEnabled Whether to use your catalog to automatically advertise relevant products or services to people based on their unique interests, intent and actions.
	// Supported values: true, false.
	CatalogEnabled bool `json:"catalog_enabled,omitempty"`
	// CatalogType The type of catalog.
	// Enum values:
	// ECOMMERCE: e-commerce
	// TRAVEL_ENTERTAINMENT: travel and entertainment.
	// MINI_SERIES: mini series.
	CatalogType enum.CatalogType `json:"catalog_type,omitempty"`
	// CampaignName Campaign name.
	CampaignName string `json:"campaign_name,omitempty"`
	// SpecialIndustries Ad categories.
	// Enum values:
	// HOUSING: Ads for real estate listings, homeowners insurance, mortgage loans or other related opportunities.
	// EMPLOYMENT: Ads for job offers, internships, professional certification programs or other related opportunities.
	// CREDIT: Ads for credit card offers, auto loans, long-term financing or other related opportunities.
	SpecialIndustries []enum.SpecialIndustry `json:"special_industries,omitempty"`
	// BudgetOptimizeOn Whether CBO is enabled.
	// Supported value: true (enabled), false (disabled).
	BudgetOptimizeOn bool `json:"budget_optimize_on,omitempty"`
	// BudgetMode Budget mode.
	// Enum value:
	// BUDGET_MODE_DYNAMIC_DAILY_BUDGET: Dynamic daily budget. It is the average daily budget over a week. Daily costs will not exceed 125% of the average daily budget. Weekly costs will not exceed the average daily budget * 7.
	// BUDGET_MODE_TOTAL: Lifetime budget.
	// BUDGET_MODE_INFINITE: Unlimited budget.
	BudgetMode enum.BudgetMode `json:"budget_mode,omitempty"`
	// BudgetAutoAdjustStrategy Returned only when the following conditions are all met:
	// budget_optimize_on is true or not specified.
	// budget_mode is BUDGET_MODE_DYNAMIC_DAILY_BUDGET.
	// app_promotion_type is not MINIS.
	// The campaign budget strategy for automatic daily campaign budget.
	// Enum values:
	// AUTO_BUDGET_INCREASE: To enable automatic budget increase. Allow your budget to automatically increase when your ads are performing well and target CPA, Day 0 target ROAS, and budget requirements are met.
	// When budget_auto_adjust_strategy is AUTO_BUDGET_INCREASE, the specified budget will be the initial daily budget. Your daily budget will be allowed to automatically increase by 20%, up to 10 times per day, when your budget utilization reaches 90% or more. Your daily budget will reset to your original daily budget each day.
	// UNSET: To disable automatic budget increase.
	BudgetAutoAdjustStrategy enum.BudgetAutoAdjustStrategy `json:"budget_auto_adjust_strategy,omitempty"`
	// Budget Fixed campaign budget or initial campaign budget.
	// When budget_auto_adjust_strategy is UNSET, this field represents your fixed campaign budget.
	// When budget_auto_adjust_strategy is AUTO_BUDGET_INCREASE, this field represents your initial campaign budget. To retrieve the current campaign budget, check the returned current_budget.
	Budget float64 `json:"budget,omitempty"`
	// CurrentBudget Returned only when budget_auto_adjust_strategy is AUTO_BUDGET_INCREASE.
	// Current campaign budget for a campaign with automatic budget increase enabled.
	CurrentBudget float64 `json:"current_budget,omitempty"`
	// OperationStatus Operation status.
	// Enum values:
	// ENABLE : The campaign is enabled (in 'ON' status).
	// DISABLE: The campaign is disabled (in 'OFF' status).
	OperationStatus enum.OperationStatus `json:"operation_status,omitempty"`
	// SecondaryStatus Campaign secondary status.
	// For enum values, see Enumeration- Campaign Status - Secondary Status.
	SecondaryStatus enum.SecondaryStatus `json:"secondary_status,omitempty"`
	// SmartPlusAdgroupMode The ad group mode for the campaign.
	// Enum values:
	// SINGLE: You can only create one single ad group within the campaign.
	// MULTIPLE: You can only create multiple ad groups within the campaign.
	SmartPlusAdgroupMode enum.SmartPlusAdgroupMode `json:"smart_plus_adgroup_mode,omitempty"`
	// PostbackWindowMode The mode that determines which SKAN 4.0 postback you want to secure.
	// Enum values:
	// POSTBACK_WINDOW_MODE1: This option secures the first postback, which corresponds to the 0-2 day attribution window. The data can take up to 4 days to return, and the campaign will wait for 4 days to release the campaign quota.
	// POSTBACK_WINDOW_MODE2: This mode secures the first two postbacks, which correspond to the 3-7 day attribution window. The data can take up to 13 days to return, and the campaign will wait for 13 days to release the campaign quota.
	// POSTBACK_WINDOW_MODE3: This mode secures all three postbacks, which correspond to the 8-35 day attribution window. The data can take up to 41 days to return, and the campaign will wait for 41 days to release the campaign quota.
	// Note: If you have set up Mobile Measurement Partner (MMP) Tracking with Adjust, Airbridge, Appsflyer, Branch, Kochava, or Singular for your App and your MMP SDK version is updated to a SKAN 4.0 supported SDK, you can transition your App to SKAN 4.0 on Events Manager. To learn about how to set up MMP tracking for your App, see How to Set Up App Attribution in TikTok Ads Manager. To learn more about how to transition your App to SKAN 4.0, see About SKAN 4.0 and TikTok and How to transition to SKAN 4.0.
	PostbackWindowMode enum.PostbackWindowMode `json:"postback_window_mode,omitempty"`
	// PONumber PO (purchase order) number.
	// PO numbers are useful for invoice tracking and reconciliation.
	// Learn more about PO numbers on monthly invoices.
	PONumber string `json:"po_number,omitempty"`
	// IsPromotionalCampaign Whether to use promotion campaign settings.
	// Enable this feature to activate specialized campaign setup flows and optimization for theatrical releases and streaming service promotions.
	// Supported values: true, false.
	IsPromotionalCampaign bool `json:"is_promotional_campaign,omitempty"`
	// RtaID 实时 API ID，即实时 API 策略标识符
	RtaID string `json:"rta_id,omitempty"`
	// RtaBidEnabled Whether to use RTA bid.
	// Supported values:
	// true: enabled.
	// false: not enabled.
	RtaBidEnabled bool `json:"rta_bid_enabled,omitempty"`
	// RtaProductSelectionEnabled Whether to use RTA to automatically select products.
	// Enum values:
	// true: enabled.
	// false: not enabled.
	RtaProductSelectionEnabled bool `json:"rta_product_selection_enabled,omitempty"`
}

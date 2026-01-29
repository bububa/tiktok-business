package smartplus

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// CreateRequest Create an Upgraded Smart+ Campaign API Request
type CreateRequest struct {
	// AdvertiserID Advertiser ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// RequestID Request ID with which you can create campaigns with duplicate names. It also supports idempotency to prevent you from sending the same request twice. If you retry requests with the same request ID multiple times, then only one will succeed.
	// It is different from the request_id returned in the response parameters, which is used to uniquely identify an HTTP request.
	// The value should be a string representation of a 64-bit integer number.
	// Example: 123456789.
	RequestID string `json:"request_id,omitempty"`
	// OperationStatus The status of the campaign when created.
	// Enum values:
	// ENABLE: The campaign is enabled when created.
	// DISABLE: The campaign is disabled when created.
	// Default value: ENABLE.
	OperationStatus enum.OperationStatus `json:"operation_status,omitempty"`
	// ObjectiveType Advertising objective.
	// Currently, we support APP_PROMOTION, WEB_CONVERSIONS, and LEAD_GENERATION.
	// For detailed explanation of enum values, see Enumeration-Advertising Objective.
	// To learn about how to create an Upgraded Smart+ App Campaign, see Create Upgraded Smart+ App Campaigns.
	// To learn about how to create an Upgraded Smart+ Web Campaign, see Create Upgraded Smart+ Web Campaigns.
	// To learn about how to create an Upgraded Smart+ Lead Generation Campaign, see Create Upgraded Smart+ Lead Generation Campaigns.
	ObjectiveType enum.ObjectiveType `json:"objective_type,omitempty"`
	// AppPromotionType Required when objective_type is APP_PROMOTION.
	// App promotion type.
	// Enum value:
	// APP_INSTALL: App install. Get new users to install your app.
	// APP_RETARGETING: App retargeting. Re-engage existing app users to take action in your app.
	// MINIS: TikTok Minis. Get people to watch your series or play your games with TikTok Minis.
	// To learn more about the campaign type where this setting is supported, see Create an Upgraded Smart+ Minis Campaign.
	AppPromotionType enum.AppPromotionType `json:"aepp_promotion_type,omitempty"`
	// SalesDestination Required when objective_type is WEB_CONVERSIONS.
	// Sales destination, the destination where you want to drive your sales.
	// Enum values:
	// WEBSITE: Website. Drive sales on your website.
	// APP: App. Drive sales on your app (product catalog required). Learn more about how to create Upgraded Smart+ Catalog Ads for App.
	// When sales_destination is APP:
	// Set catalog_enabled to true and catalog_type to ECOMMERCE or TRAVEL_ENTERTAINMENT.
	// The type of catalog that you can specify at the ad group level varies depending on your rta_id setting:
	// If rta_id is not specified, you can set catalog_id at the ad group level to the ID of a catalog with catalog_type as ECOM, HOTEL,FLIGHT,DESTINATION, or ENTERTAINMENT.
	// If rta_id is specified, you can set catalog_id at the ad group level to the ID of a catalog with catalog_type as ECOM or ENTERTAINMENT.
	// WEB_AND_APP: Website and app. Drive sales on both your website and your app.
	// Learn more about how to create Upgraded Smart+ Ads with Website and App Optimization.
	// Note: Once set, this field cannot be updated.
	SalesDestination enum.SalesDestination `json:"sales_destination,omitempty"`
	// CatalogEnabled Valid only when objective_type is WEB_CONVERSIONS or LEAD_GENERATION.
	// Whether to use catalog in the campaign.
	// Supported values: true, false.
	CatalogEnabled bool `json:"catalog_enabled,omitempty"`
	// CampaignType Campaign type.
	// Enum values: REGULAR_CAMPAIGN, IOS14_CAMPAIGN.
	// Default value: REGULAR_CAMPAIGN.
	CampaignType enum.CampaignType `json:"campaign_type,omitempty"`
	// IsPromotionalCampaign Valid when objective_type is WEB_CONVERSIONS.
	// Whether to use promotion campaign settings.
	// Enable this feature to activate specialized campaign setup flows and optimization for theatrical releases and streaming service promotions.
	// Default value: false.
	// Supported values: true, false.
	// Note:
	// Using promotion campaign settings for theatrical releases and for streaming service promotions are currently separate allowlist-only features. If you would like to access them, please contact your TikTok representative.
	// Once set, this field cannot be updated.
	IsPromotionalCampaign bool `json:"is_promotional_campaign,omitempty"`
	// AppID Required when the following conditions are both met:
	// objective_type is APP_PROMOTION and app_promotion_type is APP_INSTALL, or objective_type is WEB_CONVERSIONS and sales_destination is APP.
	// camapign_type is IOS14_CAMPAIGN.
	// The App ID of the app to promote.
	// To get a list of App IDs, use /app/list/.
	// Note: You cannot specify an App that has not activated the SAN module on your MMP through this field to create iOS Dedicated Campaigns. To ensure that TikTok SAN integration is enabled for your App, see How to migrate your app to SAN integration.
	AppID string `json:"app_id,omitempty"`
	// GamingAdComplianceAgreement Valid only when the following conditions are all met:
	// objective_type is APP_PROMOTION.
	// app_promotion_type is APP_INSTALL.
	// campaign_type is IOS14_CAMPAIGN.
	// Whether to agree to the Compliance Assurance Policy for Gaming Advertisers on TikTok.
	// The policy is as follows: You confirm and attest that any gaming application, product or service (game) you desire to advertise on TikTok, including any associated URL(s), (a) complies with all applicable laws and regulations of the jurisdictions where the game can be accessed or played, and upon request, can provide supporting documentation as evidence of why the game is not considered illegal gambling or lottery; and (b) has not been and is not part of any investigation or lawsuit regarding the game's legality or regulatory compliance.
	// Enum values:
	// ON: To agree to the policy.
	// OFF: To leave the policy not accepted.
	// Default value: OFF.
	// Note: Agreeing to the Compliance Assurance Policy for Gaming Advertisers on TikTok is currently an allowlist-only feature. If you would like to access it, please contact your TikTok representative.
	GamingAdComplianceAgreement enum.OnOff `json:"gaming_ad_compliance_agreement,omitempty"`
	// CampaignAppProfilePageStatus Whether to use app profile page to optimize delivery.
	// Enum values: ON, OFF.
	// Default value: OFF.
	// Note:
	// You can use the field only when objective_type is APP_PROMOTIONand your campaign is an iOS 14 Dedicated Campaign. Otherwise, an error will occur.
	// When ON is selected, pass in the page_id of App Profile Page.
	// To create an App Profile Page, use Instant Page Editor SDK and set businessType to 4.
	CampaignAppProfilePageState enum.AppProfilePageState `json:"campaign_app_profile_page_state,omitempty"`
	// DisableSkanCampaign the campaign. However, you cannot retrieve SKAN reporting metrics for the campaign. Learn more about SAN integration.
	// For a Dedicated Campaign using an app that is eligible for Advanced Dedicated Campaigns, you need to enable SKAN attribution or Advanced Dedicated Campaign or both. You cannot disable SKAN attribution and Advanced Dedicated Campaign simultaneously.
	// false: To enable SKAN attribution. The campaign will be bound by Dedicated Campaign quota limits and you will be able to retrieve SKAN metrics for the campaign.
	// If you enable Advanced Dedicated Campaign by setting is_advanced_dedicated_campaign to true simultaneously, you'll be able to retrieve both SAN and SKAN reporting metrics for the campaign.
	// Note:
	// Disabling SKAN attribution for Dedicated CaValid only when the following conditions are all met:
	// objective_type is APP_PROMOTION and app_promotion_type is APP_INSTALL, or objective_type is WEB_CONVERSIONS and sales_destination is APP.
	// campaign_type is IOS14_CAMPAIGN.
	// app_id is set to the ID of an iOS app that is eligible for Advanced Dedicated Campaigns. To confirm whether an app is eligible for Advanced Dedicated Campaigns, use /app/list/ and check the returned advanced_dedicated_campaign_allowed.
	// Whether to disable SKAN (SKAdNetwork) attribution, Apple's conversion attribution solution for iOS campaigns
	// Enum values:
	// true: To disable SKAN attribution. The campaign will not be bound by Dedicated Campaign quota limits and you will be able to retrieve Self Attribution Network (SAN) metrics for mpaigns is currently an allowlist-only feature. If you would like to access it, please contact your TikTok representative.
	// If you are allowlisted for disabling SKAN attribution for Dedicated Campaigns, you cannot set bid_type to BID_TYPE_CUSTOM when the following conditions are all met:
	// is_advanced_dedicated_campaign is false.
	// disable_skan_campaign is false.
	// optimization_goal is INSTALL or IN_APP_EVENT.
	// Getting access to Advanced Dedicated Campaign requires the app to meet certain criteria. If your app is ineligible for Advanced Dedicated Campaign, reach out to your TikTok representative. Your TikTok representative will be able to provide troubleshooting support to help your app get access to Advanced Dedicated Campaign.
	DisableSkanCampaign *bool `json:"disable_skan_campaign,omitempty"`
	// CampaignName Campaign name.
	// Length limit: 512 characters. Emoji is not supported.
	// Each word in Chinese or Japanese counts as two characters, while each letter in English counts as one character
	CampaignName string `json:"campaign_name,omitempty"`
	// SpecialIndustries Ad categories. Enum values:
	// HOUSING: Ads for real estate listings, homeowners insurance, mortgage loans or other related opportunities.
	// EMPLOYMENT: Ads for job offers, internships, professional certification programs or other related opportunities.
	// CREDIT: Ads for credit card offers, auto loans, long-term financing or other related opportunities.
	// Note:
	// Once you've specified the industry type, the system will adjust your target options to help you comply with advertising policies. See Ad targeting for details.
	// This field is only supported for advertisers who are registered in the US or Canada.
	SpecialIndustries []enum.SpecialIndustry `json:"special_industries,omitempty"`
	// BudgetOptimizeOn Whether to enable Campaign Budget Optimization (CBO).
	// Supported values: true (enabled), false (disabled).
	// Default value: true.
	BudgetOptimizeOn *bool `json:"budget_optimize_on,omitempty"`
	// BudgetMode Valid when budget_optimize_on is true or not specified.
	// Budget mode.
	// Enum value:
	// BUDGET_MODE_DYNAMIC_DAILY_BUDGET: Dynamic daily budget. It is the average daily budget over a week. Daily costs will not exceed 125% of the average daily budget. Weekly costs will not exceed the average daily budget * 7.
	// BUDGET_MODE_TOTAL: Lifetime budget.
	// Default value when budget_optimize_on is true or not specified: BUDGET_MODE_DYNAMIC_DAILY_BUDGET.
	// Note:
	// Once set, this field cannot be updated.
	BudgetMode enum.BudgetMode `json:"budget_mode,omitempty"`
	// BudgetAutoAdjustStrategy Valid only when the following conditions are all met:
	// budget_optimize_on is true or not specified.
	// budget_mode is BUDGET_MODE_DYNAMIC_DAILY_BUDGET.
	// app_promotion_type is not MINIS.
	// The campaign budget strategy for automatic daily campaign budget.
	// Enum value:
	// AUTO_BUDGET_INCREASE: To enable automatic budget increase. Allow your budget to automatically increase when your ads are performing well and target CPA, Day 0 target ROAS, and budget requirements are met.
	// When budget_auto_adjust_strategy is AUTO_BUDGET_INCREASE, the specified budget will be the initial daily budget. Your daily budget will be allowed to automatically increase by 20%, up to 10 times per day, when your budget utilization reaches 90% or more. Your daily budget will reset to your original daily budget each day.
	// Note: Enabling automatic budget increase is currently an allowlist-only feature. If you would like to access it, please contact your TikTok representative.
	BudgetAutoAdjustStrategy enum.BudgetAutoAdjustStrategy `json:"budget_auto_adjust_strategy,omitempty"`
	// Buget Required when budget_mode is specified.
	// Fixed campaign budget or initial campaign budget.
	// When budget_auto_adjust_strategy is UNSET, this field represents your fixed campaign budget.
	// When budget_auto_adjust_strategy is AUTO_BUDGET_INCREASE, this field represents your initial campaign budget. To retrieve the current campaign budget, check the returned current_budget.
	// To learn about the minimum budget and how to set budget modes, see Budget.
	// To directly see the daily budget value range for a currency, see Currency-Daily budget value range.
	Budget float64 `json:"budget,omitempty"`
	// PostbackWindowMode Valid only when campaign_type is IOS14_CAMPAIGN and operation_status is DISABLE.
	// You can specify this field when you create a disabled campaign or when you disable an existing campaign by using /smart_plus/campaign/status/update/. The recommended practice is to specify this field when disabling an existing campaign.
	// The mode that determines which SKAN (SKAdNetwork) 4.0 postback you want to secure. Options with longer windows require more time to receive, and as a result, more time to release the campaign back to the available number.
	// Enum values:
	// POSTBACK_WINDOW_MODE1: This option secures the first postback, which corresponds to the 0-2 day attribution window. The data can take up to 4 days to return, and the campaign will wait for 4 days to release the campaign quota.
	// POSTBACK_WINDOW_MODE2: This mode secures the first two postbacks, which correspond to the 3-7 day attribution window. The data can take up to 13 days to return, and the campaign will wait for 13 days to release the campaign quota.
	// POSTBACK_WINDOW_MODE3: This mode secures all three postbacks, which correspond to the 8-35 day attribution window. The data can take up to 41 days to return, and the campaign will wait for 41 days to release the campaign quota.
	// Note:
	// If you have set up Mobile Measurement Partner (MMP) Tracking with Adjust, Airbridge, Appsflyer, Branch, Kochava, or Singular for your App and your MMP SDK version is updated to a SKAN 4.0 supported SDK, you can transition your App to SKAN 4.0 on Events Manager. To learn about how to set up MMP tracking for your App, see How to Set Up App Attribution in TikTok Ads Manager. To learn more about how to transition your App to SKAN 4.0, see About SKAN 4.0 and TikTok and How to transition to SKAN 4.0.
	// Once set, this field cannot be updated.
	// If operation_status is set to ENABLE or not specified, this field will be ignored.
	// If you pass in this field when you have not enabled SKAN 4.0 for your App (app_id), an error will occur.
	// If this field is not specified when campaign_type is set to IOS14_CAMPAIGN, operation_status is set to DISABLE, and you have enabled SKAN 4.0 for your App, this field will default to POSTBACK_WINDOW_MODE1.
	// If you have enabled SKAN 4.0 for your App, ensure that you target devices running iOS 16.1 and later so that you can receive SKAN 4.0 postbacks. To only target iOS 16.1+ devices, set min_ios_version to 16.1 at the ad group level.
	PostbackWindowMode enum.PostbackWindowMode `json:"postback_window_mode,omitempty"`
	// PONumber PO (purchase order) number.
	// PO numbers are useful for invoice tracking and reconciliation.
	// Learn more about PO numbers on monthly invoices.
	PONumber string `json:"po_number,omitempty"`
	// RtaID Realtime API (RTA) ID, the RTA strategy identifier.
	// The RTA enables you to integrate real-time audience data to customize ad auction and targeting, thereby delivering better performance.
	// To obtain the list of RTA IDs associated with your ad account, contact your TikTok representative.
	// Note:
	// Enabling RTA for your ads is currently an allowlist-only feature. If you would like to access it, please contact your TikTok representative. You will need to complete the onboarding process to be able to receive and respond to RTA requests sent to your system first.
	// Once set, this field cannot be updated.
	RtaID string `json:"rta_id,omitempty"`
	// RtaBidEnabled Whether to use RTA bid.
	// Supported values:
	// true: enabled.
	// false: not enabled.
	// Default value: false.
	// You can only set this field to true in any of the following scenarios:
	// Scenario 1:
	// objective_type is set to WEB_CONVERSIONS.
	// sales_destination is set to APP.
	// rta_id is passed.
	// Scenario 2:
	// objective_type is set to APP_PROMOTION.
	// app_promotion_type is set to APP_RETARGETING.
	// rta_id is passed.
	// Note:
	// Enabling RTA bid is currently an allowlist-only feature. If you would like to access it, please contact your TikTok representative.
	// Once set, this field cannot be updated.
	RtaBidEnabled bool `json:"rta_bid_enabled,omitempty"`
	// RtaProductSelectionEnabled Whether to use RTA to automatically select products.
	// Enum values:
	// true: enabled.
	// false: not enabled.
	// Default value: false.
	// You can only set this field to true when the following conditions are met:
	// objective_type is set to WEB_CONVERSIONS.
	// sales_destination is set to APP.
	// rta_id is specified.
	// Note:
	// Once set, this field cannot be updated.
	// When rta_product_selection_enabled is true , you can only set product_specific_type to ALL at the ad level.
	RtaProductSelectionEnabled bool `json:"rta_product_selection_enabled,omitempty"`
}

// Encode implements PostRequest interface
func (r *CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse Create an Upgraded Smart+ Campaign API Response
type CreateResponse struct {
	model.BaseResponse
	Data *Campaign `json:"data,omitempty"`
}

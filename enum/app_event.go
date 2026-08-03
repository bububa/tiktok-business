package enum

// AppEventType is an Events API 1.0 app event name.
//
// Custom event names can be represented by converting a string to AppEventType.
type AppEventType string

const (
	AppEventTypeAchieveLevel      AppEventType = "ACHIEVE_LEVEL"
	AppEventTypeAddPaymentInfo    AppEventType = "ADD_PAYMENT_INFO"
	AppEventTypeAddToCart         AppEventType = "ADD_TO_CART"
	AppEventTypeAddToWishlist     AppEventType = "ADD_TO_WISHLIST"
	AppEventTypeCheckout          AppEventType = "CHECKOUT"
	AppEventTypeCompleteTutorial  AppEventType = "COMPLETE_TUTORIAL"
	AppEventTypeCreateGroup       AppEventType = "CREATE_GROUP"
	AppEventTypeCreateRole        AppEventType = "CREATE_ROLE"
	AppEventTypeGenerateLead      AppEventType = "GENERATE_LEAD"
	AppEventTypeInAppAdClick      AppEventType = "IN_APP_AD_CLICK"
	AppEventTypeInAppAdImpression AppEventType = "IN_APP_AD_IMPRESSION"
	AppEventTypeInstall           AppEventType = "INSTALL"
	AppEventTypeJoinGroup         AppEventType = "JOIN_GROUP"
	AppEventTypeLaunchApp         AppEventType = "LAUNCH_APP"
	AppEventTypeLoanApplication   AppEventType = "LOAN_APPLICATION"
	AppEventTypeLoanApproval      AppEventType = "LOAN_APPROVAL"
	AppEventTypeLoanDisbursal     AppEventType = "LOAN_DISBURSAL"
	AppEventTypeLogin             AppEventType = "LOGIN"
	AppEventTypePurchase          AppEventType = "PURCHASE"
	AppEventTypeRate              AppEventType = "RATE"
	AppEventTypeRegistration      AppEventType = "REGISTRATION"
	AppEventTypeSearch            AppEventType = "SEARCH"
	AppEventTypeSpendCredits      AppEventType = "SPEND_CREDITS"
	AppEventTypeStartTrial        AppEventType = "START_TRIAL"
	AppEventTypeSubscribe         AppEventType = "SUBSCRIBE"
	AppEventTypeUnlockAchievement AppEventType = "UNLOCK_ACHIEVEMENT"
	AppEventTypeViewContent       AppEventType = "VIEW_CONTENT"
)

// AppEventSource identifies the source used to report app events.
type AppEventSource string

const AppEventSourceAPI AppEventSource = "APP_EVENTS_API"

// AppEventAction identifies an action in an app event batch.
type AppEventAction string

const AppEventActionTrack AppEventAction = "track"

// AppEventATTStatus is the AppTrackingTransparency authorization status.
type AppEventATTStatus string

const (
	AppEventATTStatusAuthorized    AppEventATTStatus = "AUTHORIZED"
	AppEventATTStatusDenied        AppEventATTStatus = "DENIED"
	AppEventATTStatusNotDetermined AppEventATTStatus = "NOT_DETERMINED"
	AppEventATTStatusRestricted    AppEventATTStatus = "RESTRICTED"
	AppEventATTStatusNotApplicable AppEventATTStatus = "NOT_APPLICABLE"
)

// AppEventPlatform is the mobile platform name accepted by Events API 1.0.
type AppEventPlatform string

const (
	AppEventPlatformAndroid AppEventPlatform = "Android"
	AppEventPlatformIOS     AppEventPlatform = "iOS"
)

// AppRetargetingState controls whether app retargeting is enabled.
type AppRetargetingState string

const (
	AppRetargetingStateRetargeting    AppRetargetingState = "RETARGETING"
	AppRetargetingStateNonRetargeting AppRetargetingState = "NON_RETARGETING"
)

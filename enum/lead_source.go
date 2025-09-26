package enum

// LeadSource 订阅的线索来源。
type LeadSource string

const (
	// LeadSource_INSTANT_FORM：通过即时表单生成的线索。
	LeadSource_INSTANT_FORM LeadSource = "INSTANT_FORM"
	// LeadSource_DIRECT_MESSAGE：通过绑定的企业号的私信生成的线索。
	LeadSource_DIRECT_MESSAGE LeadSource = "DIRECT_MESSAGE"
)

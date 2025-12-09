package enum

// BusinessMessageTemplateType The type of message template.
type BusinessMessageTemplateType string

const (
	// BusinessMessageTemplateType_QA_BUTTON_CARD: Q&A button card, a card showing a question and buttons that represent answers. When users click one of the buttons, the text content of the button will be sent to the conversation thread on behalf of the user as an answer to the pre-set question on the card.
	BusinessMessageTemplateType_QA_BUTTON_CARD BusinessMessageTemplateType = "QA_BUTTON_CARD"
	// BusinessMessageTemplateType_QA_LINK_CARD: Q&A text link card, a card showing a question and interactive text links that represent answers to the question. When users click one of the text links, the text content of the text link will be sent to the conversation thread on behalf of the user as an answer to the pre-set question on the card.
	BusinessMessageTemplateType_QA_LINK_CARD BusinessMessageTemplateType = "QA_LINK_CARD"
)

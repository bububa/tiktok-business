package message

import "github.com/bububa/tiktok-business/enum"

// Message message details
type Message struct {
	// Sender TikTok username of the message sender.
	Sender string `json:"sender,omitempty"`
	// Recipient TikTok username of the message recipient.
	Recipient string `json:"recipient,omitempty"`
	// ConversationID Conversation ID
	ConversationID string `json:"conversation_id,omitempty"`
	// MessageID Message ID
	MessageID string `json:"message_id,omitempty"`
	// Timestamp The time when the message was sent, in the format of an Epoch/Unix timestamp in milliseconds.
	Timestamp int64 `json:"timestamp,omitempty"`
	// MessageTag Additional information about the message
	MessageTag *MessageTag `json:"message_tag,omitempty"`
	// MessageType The type of the message.
	// Enum values:
	// TEXT: text.
	// IMAGE: image.
	// SHARE_POST: TikTok post.
	// VIDEO: video.
	// EMOJI: emoji.
	// STICKER: sticker.
	// TEMPLATE: message template.
	// OTHER: other message types.
	MessageType enum.BusinessMessageType
	// AutoMessageType The type of automatic message.
	// Enum values:
	// WELCOME_MESSAGE: welcome message. A welcome message is a message that is automatically sent to a user when they start a direct message conversation with you.
	// SUGGESTED_QUESTION: suggested question. A suggested question is a frequently asked question displayed at the beginning of a chat. When the user selects a suggested question, the user will automatically receive the corresponding preset answer.
	// AUTO_REPLY: keyword reply. A keyword reply is an automated response triggered by specific keywords in a message.
	// Note: When the message is not an automatic message, this field will not be returned.
	AutoMessageType enum.BusinessAutoMessageType `json:"auto_message_type,omitempty"`
	// Text Required when message_type is TEXT.
	// Text information.
	// A message cannot include both a text and an image.
	Text *Text `json:"text,omitempty"`
	// Image Required when message_type is IMAGE.
	// Image information.
	// A message cannot include both a text and an image.
	Image *Image `json:"image,omitempty"`
	// SharePost Required when message_type is SHARE_POST.
	// Information about the post to be shared.
	SharePost *SharePost `json:"share_post,omitempty"`
	// Reactions Information about the list of reactions in response to the message.
	Reactions []Reaction `json:"reactions,omitempty"`
	// Sticker Returned only when message_type is STICKER.
	// Information about the sticker.
	Sticker *Sticker `json:"sticker,omitempty"`
	// Emoji Returned only when message_type is EMOJI.
	// Information about the emoji.
	Emoji *Emoji `json:"emoji,omitempty"`
	// Video Returned only when message_type is VIDEO.
	// Information about the video.
	Video *Video `json:"video,omitempty"`
	// Template Returned only when message_type is TEMPLATE.
	// Details about the message template sent.
	Template *Template `json:"template,omitempty"`
	// FromUser User information of the sender.
	FromUser *Participant `json:"from_user,omitempty"`
	// ToUser User information of the receiver.
	ToUser *Participant `json:"to_user,omitempty"`
	// ReferencedMessageInfo Information about the original message that is replied to
	ReferencedMessageInfo *ReferencedMessageInfo `json:"referenced_message_info,omitempty"`
}

// MessageTag Additional information about the message
type MessageTag struct {
	// Source The source of the message, indicating the platform from which the message was sent.
	// Enum values:
	// APP: TikTok mobile app.
	// WEB: TikTok Web or TikTok Business Center.
	// API: TikTok Business Messaging API.
	// OTHERS: Other TikTok internal servers or tools, including automatic bots.
	// UNKNOWN_SOURCE: Unknown source.
	Source enum.BusinessMessageSource `json:"source,omitempty"`
}

// Text information.
type Text struct {
	// Body Required when text is specified.
	// Text content.
	// Length limit: 6,000 characters, including spaces and emojis.
	Body string `json:"body,omitempty"`
}

// Image information.
type Image struct {
	// MediaID Required when image is specified.
	// Image ID.
	// You can use /business/message/capabilities/get/ to query if the Business Account can send a media attachment to a conversation.
	// Before sending an image, use /business/message/media/upload/ to upload an image to your account and obtain the media_id.
	MediaID string `json:"media_id,omitempty"`
}

// SharePost Required when message_type is SHARE_POST.
// Information about the post to be shared.
type SharePost struct {
	// ItemID Required when share_post is specified.
	// The ID of the TikTok post to share with others.
	// To obtain the IDs of TikTok posts associated with your TikTok account, use /business/video/list/.
	// Note: Currently, you can only share your posts.
	ItemID string `json:"item_id,omitempty"`
	// EmbedURL An embeddable link for the TikTok post.
	// This URL can be used to embed the TikTok post in external websites or applications.
	// If the privacy setting of the post is "Friends" or "Only you", the post will not be viewable through this link.
	EmbedURL string `json:"embed_url,omitempty"`
}

// Template Required when message_type is TEMPLATE.
// Details about the message template to send.
type Template struct {
	// Type Required when template is specified.
	// The type of message template.
	// Enum values:
	// QA_BUTTON_CARD: Q&A button card, a card showing a question and buttons that represent answers. When users click one of the buttons, the text content of the button will be sent to the conversation thread on behalf of the user as an answer to the pre-set question on the card.
	// QA_LINK_CARD: Q&A text link card, a card showing a question and interactive text links that represent answers to the question. When users click one of the text links, the text content of the text link will be sent to the conversation thread on behalf of the user as an answer to the pre-set question on the card.
	Type enum.BusinessMessageTemplateType `json:"type,omitempty"`
	// Title Required when template is specified.
	// The pre-set question on the card.
	// Length limit: 40 characters.
	Title string `json:"title,omitempty"`
	// Buttons Required when template is specified.
	// A list of buttons or text links that represent answers to the question on the card.
	// Size range: 1-3.
	Buttons []Button `json:"buttons,omitempty"`
}

// Button a button or text link that represent answers to the question on the card.
type Button struct {
	// Type Required when buttons is specified.
	// The type of the button or text link.
	// Enum value:
	// REPLY: Reply button or reply text link. When users click the button or text link, the system will send the text content of the button or text link as a reply message on behalf of the user in the conversation.
	Type enum.BusinessMessageButtonType `json:"type,omitempty"`
	// TItle Required when buttons is specified.
	// The text content of the button or text link.
	// Length limit:
	// When type of the message template is QA_BUTTON_CARD: 20 characters.
	// When type of the message template is QA_LINK_CARD: 40 characters.
	Title string `json:"title,omitempty"`
	// ID A self-defined ID for the button or text link.
	// You can use this ID to distinguish between different buttons and text links.
	// Length limit: 40 characters.
	ID string `json:"id,omitempty"`
}

// ReferencedMessageInfo Required when you want to send a reply to a message.
// Information about the original message to reply to.
type ReferencedMessageInfo struct {
	// ReferencedMessageID Required when referenced_message_info is specified.
	// The ID of the original message that you want to reply to.
	// The original message can be a text, an image, or a TikTok post. To obtain such a message, use /business/message/content/list/ and ensure the message_type of the message is TEXT, IMAGE, or SHARE_POST.
	// When using this field, make sure to set the message_type to TEXT and specify text simultaneously, as only text replies are supported. After you send the reply, the original message will be quoted in the conversation thread.
	ReferencedMessageID string `json:"referenced_message_id,omitempty"`
}

// DirectReply Required when you are sending a direct reply.
// Details of the direct reply.
// Note:
// To send a direct reply, ensure that Comment-to-Message is enabled for the Business Account. To obtain the Comment-to-Message setting for a Business Account, use /business/message/direct_reply/get/.
// The Comment-to-Message feature is only available for Business Accounts registered in Vietnam, Indonesia, and Thailand.
// Once you enable Comment-to-Message for a Business Account, the Business Account can only reply to comments published by TikTok accounts that are registered in the APAC, LATAM, and METAP regions.
type DirectReply struct {
	// ReplyType Required when direct_reply is specified.
	// The type of direct reply.
	// Enum value:
	// COMMENT_REPLY: a reply to a comment on a video.
	// Note: To send a comment reply, ensure that all the following conditions are all met:
	// The comment that you want to reply to is a first-level comment on a video of the Business Account.
	// Your reply is sent within 48 hours after the user posted the comment.
	// The comment has not been replied to through direct messages yet in any way, including from the TikTok app or via the API.
	// The user who posted the comment had no direct messaging communication with the Business Account in the past 24 hours.
	// The user who posted the comment is over 18 years old.
	ReplyType enum.BusinessMessageReplyType `json:"reply_type,omitempty"`
	// CommentReply Required when direct_reply is specified.
	// Details of the comment reply.
	CommentReply *CommentReply `json:"comment_reply,omitempty"`
}

// CommentReply Details of the comment reply.
type CommentReply struct {
	// CommentID Required when direct_reply is specified.
	// ID of the high-intent comment you want to reply to.
	// A high-intent comment is one that expresses a strong intention to purchase or seek further information.
	// You can obtain the ID of a high-intent comment from the webhook event im_receive_high_intent_comment.
	CommentID string `json:"comment_id,omitempty"`
}

// Reaction Information about the reaction in response to the message.
type Reaction struct {
	// Type The type of reaction.
	// Enum values:
	// EMOJI: A text emoji.
	// AI_EMOJI: An AI-generated emoji derived from a selfie-based avatar
	Type enum.BusinessMessageReactionType `json:"type,omitempty"`
	// Emoji Returned only when type is EMOJI.
	// The text of the text emoji.
	Emoji string `json:"emoji,omitempty"`
	// AIEmojiURL Returned only when type is AI_EMOJI.
	// The URL to the AI-generated emoji image.
	// Validity period: 30 days.
	// The expiration time is included in the URL after the x-expires parameter, in the format of an Epoch/Unix timestamp in seconds.
	AIEmojiURL string `json:"ai_emoji_url,omitempty"`
	// UniqueIdentifier The identifier of the sender.
	// When role within the object from_user is BUSINESS_ACCOUNT, this field is the same as the business_id, the application specific unique identifier for the TikTok Business Account.
	// When role within the object from_user is PERSONAL_ACCOUNT, this field is the same as the unique_identifier, the globally unique identifier assigned to the TikTok Personal Account user in the conversation. This identifier remains consistent across different APIs, such as the Accounts comment listing endpoint, facilitating cross-referencing and integration.
	UniqueIdentifier string `json:"unique_identifier,omitempty"`
	// Timestamp The timestamp when the reaction was made, in the format of an Epoch/Unix timestamp in milliseconds
	Timestamp int64 `json:"timestamp,omitempty"`
}

// Sticker Returned only when message_type is STICKER.
// Information about the sticker.
type Sticker struct {
	// URL The URL to the sticker.
	// Validity period: 30 days.
	// The expiration time is included in the URL after the x-expires parameter, in the format of an Epoch/Unix timestamp in seconds.
	URL string `json:"url,omitempty"`
}

// Emoji Returned only when message_type is EMOJI.
// Information about the emoji.
type Emoji struct {
	// URL The URL to the emoji.
	// The URL will not expire.
	URL string `json:"url,omitempty"`
}

// Video Returned only when message_type is VIDEO.
// Information about the video.
type Video struct {
	// MediaID The video ID.
	// To download the video associated with the video ID, use /business/message/media/download/.
	MediaID string `json:"media_id,omitempty"`
}

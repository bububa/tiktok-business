package enum

// BusinessMessageReactionType The type of reaction.
type BusinessMessageReactionType string

const (
	// BusinessMessageReactionType_EMOJI: A text emoji.
	BusinessMessageReactionType_EMOJI BusinessMessageReactionType = "EMOJI"
	// BusinessMessageReactionType_AI_EMOJI: An AI-generated emoji derived from a selfie-based avatar
	BusinessMessageReactionType_AI_EMOJI BusinessMessageReactionType = "AI_EMOJI"
)

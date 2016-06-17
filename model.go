package gochatbot

import "strconv"

const (
	// StatusSuccess is success status
	StatusSuccess = "success"

	// MessageMinimumLength is minimum message length
	MessageMinimumLength = 1

	// MessageMaximumLength is maximum message length
	MessageMaximumLength = 1000

	// CharacterTypeCat is cat character type
	CharacterTypeCat = "cat"

	// CharacterTypeDog is dog character type
	CharacterTypeDog = "dog"

	// CharacterTypeRoujin is roujin character type
	CharacterTypeRoujin = "roujin"
)

type (
	// ErrorResponse error response data type
	ErrorResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	// ChatResponse chat response data type.
	ChatResponse struct {
		Status string `json:"status"`
		Result string `json:"result"`
	}

	// CharacterChatResponse character chat response data type.
	CharacterChatResponse struct {
		Status string `json:"status"`
		Result string `json:"result"`
	}
)

// IsValidMessage returns true if message is valid.
func IsValidMessage(v string) bool {
	return len(v) >= MessageMinimumLength && len(v) <= MessageMaximumLength
}

// IsValidCharacterType returns true if character type is valid.
func IsValidCharacterType(v string) bool {
	return v == CharacterTypeCat || v == CharacterTypeDog || v == CharacterTypeRoujin
}

// Error returns error string.
func (e *ErrorResponse) Error() string {
	return e.Message + " - code:" + strconv.Itoa(e.Code)
}

// IsSuccess returns true if response status is success.
func (c *ChatResponse) IsSuccess() bool {
	return c.Status == StatusSuccess
}

// IsEmptyResult returns true if response result is empty.
func (c *ChatResponse) IsEmptyResult() bool {
	return len(c.Result) == 0
}

// IsSuccess returns true if response status is success.
func (c *CharacterChatResponse) IsSuccess() bool {
	return c.Status == StatusSuccess
}

// IsEmptyResult returns true if response result is empty.
func (c *CharacterChatResponse) IsEmptyResult() bool {
	return len(c.Result) == 0
}

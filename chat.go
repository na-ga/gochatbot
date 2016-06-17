package gochatbot

import (
	"fmt"
	"net/url"
)

// Chat sends message to chat API and returns response.
func (c *ChatbotClient) Chat(message string) (*ChatResponse, error) {
	if !IsValidMessage(message) {
		return nil, fmt.Errorf("Invalid message. (%s)", message)
	}
	query := url.Values{
		"message": []string{message},
		"key":     []string{c.APIKey},
	}
	var res *ChatResponse
	if err := c.do(c.get("/chat").Query(query.Encode()), &res); err != nil {
		return nil, err
	}
	return res, nil
}

// CharacterChat sends message to character chat API and returns response.
func (c *ChatbotClient) CharacterChat(message string, characterType string) (*CharacterChatResponse, error) {
	if !IsValidMessage(message) {
		return nil, fmt.Errorf("Invalid message. (%s)", message)
	}
	if !IsValidCharacterType(characterType) {
		return nil, fmt.Errorf("Invalid character type. (%s)", characterType)
	}
	query := url.Values{
		"message":        []string{message},
		"character_type": []string{characterType},
		"key":            []string{c.APIKey},
	}
	var res *CharacterChatResponse
	if err := c.do(c.get("/character").Query(query.Encode()), &res); err != nil {
		return nil, err
	}
	return res, nil
}

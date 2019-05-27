package api

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type Token struct {
	BillingAccountID int      `json:"billing_account_id"`
	ConsumerID       string   `json:"consumer_id"`
	CreatedAt        PilwTime `json:"created_at"`
	Description      string   `json:"description"`
	ID               int      `json:"id"`
	KongID           string   `json:"kong_id"`
	Restricted       bool     `json:"restricted"`
	Token            string   `json:"token"`
	UpdatedAt        PilwTime `json:"updated_at"`
	UserID           int      `json:"user_id"`
}

func ParseTokenList(str []byte) ([]Token, error) {
	var tokenList []Token

	err := json.Unmarshal(str, &tokenList)
	if err != nil {
		return tokenList, err
	}

	return tokenList, nil
}

func GetTokenList(key string) ([]Token, error) {
	resp, err := get(key, "user-resource/token/list")
	if err != nil {
		return nil, err
	}

	tokenList, err := ParseTokenList([]byte(resp))
	if err != nil {
		return nil, err
	}

	return tokenList, err
}

// CreateToken attempts to register a new token. Returns a list of tokens on success
func CreateToken(key string, description string, restricted bool, billingAccountID int) ([]Token, error) {
	form := url.Values{}
	form.Add("billing_account_id", fmt.Sprintf("%d", billingAccountID))
	form.Add("description", description)
	form.Add("restricted", fmt.Sprintf("%v", restricted))

	resp, err := postForm(key, "user-resource/token", form)
	if err != nil {
		return nil, err
	}

	tokenList, err := ParseTokenList([]byte(resp))
	if err != nil {
		return nil, err
	}

	return tokenList, nil
}

// DeleteToken attempts to delete a token by its ID
func DeleteToken(key string, tokenID int) error {
	form := url.Values{}
	form.Add("token_id", fmt.Sprintf("%d", tokenID))

	_, err := deleteForm(key, "user-resource/token", form)
	if err != nil {
		return err
	}

	return nil
}
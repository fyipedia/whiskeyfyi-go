// Package whiskeyfyi provides a Go client for the WhiskeyFYI API.
//
// WhiskeyFYI (whiskeyfyi.com) — Whiskey encyclopedia with expressions, distilleries, and regions.
// This client requires no authentication and has zero external dependencies.
//
// Usage:
//
//	client := whiskeyfyi.NewClient()
//	result, err := client.Search("example")
package whiskeyfyi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// DefaultBaseURL is the default base URL for the WhiskeyFYI API.
const DefaultBaseURL = "https://whiskeyfyi.com/api"

// Client is a WhiskeyFYI API client.
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient creates a new WhiskeyFYI API client with default settings.
func NewClient() *Client {
	return &Client{
		BaseURL:    DefaultBaseURL,
		HTTPClient: &http.Client{},
	}
}

func (c *Client) get(path string, result interface{}) error {
	resp, err := c.HTTPClient.Get(c.BaseURL + path)
	if err != nil {
		return fmt.Errorf("whiskeyfyi: request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("whiskeyfyi: HTTP %d: %s", resp.StatusCode, string(body))
	}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return fmt.Errorf("whiskeyfyi: decode failed: %w", err)
	}
	return nil
}

// Search searches across all content.
func (c *Client) Search(query string) (*SearchResult, error) {
	var result SearchResult
	path := fmt.Sprintf("/search/?q=%s", url.QueryEscape(query))
	if err := c.get(path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Entity returns details for a expression by slug.
func (c *Client) Entity(slug string) (*EntityDetail, error) {
	var result EntityDetail
	if err := c.get("/expression/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GlossaryTerm returns a glossary term by slug.
func (c *Client) GlossaryTerm(slug string) (*GlossaryTerm, error) {
	var result GlossaryTerm
	if err := c.get("/glossary/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Random returns a random expression.
func (c *Client) Random() (*EntityDetail, error) {
	var result EntityDetail
	if err := c.get("/random/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

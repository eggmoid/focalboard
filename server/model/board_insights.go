package model

import (
	"encoding/json"
	"io"
)

// BoardInsight gives insight into activities in a Board
// swagger:model
type BoardInsight struct {
	// ID of the board
	// required: true
	BoardID string `json:"boardID"`

	// icon of the board
	// required: false
	Icon string `json:"icon"`

	// Title of the board
	// required: false
	Title string `json:"title"`

	// Metric of how active the board is
	// required: true
	ActivityCount string `json:"activityCount"`

	// IDs of users active on the board
	// required: true
	ActiveUsers string `json:"activeUsers"`

	// ID of user who created the board
	// required: true
	CreatedBy string `json:"createdBy"`
}

func BoardInsightsFromJSON(data io.Reader) []BoardInsight {
	var boardInsights []BoardInsight
	_ = json.NewDecoder(data).Decode(&boardInsights)
	return boardInsights
}

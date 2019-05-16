package model

type ResponseToString struct {
	Status string `json:"status"`
	Text string `json:"text"`
}

type ResponseToNumber struct {
	Status string `json:"status"`
	Number int `json:"number"`
}

type Text struct {
	Text string `json:"text"`
}
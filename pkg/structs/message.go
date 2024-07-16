package structs

type Message struct {
	From       string `json:"from"`
	To         string `json:"to"`
	Content    string `json:"content"`
	Type       int    `json:"type"`
	SourceHost string `json:"sourceHost"`
	ClientIP   string `json:"clientIP"`
	SendTime   int64  `json:"sendTime"`
}

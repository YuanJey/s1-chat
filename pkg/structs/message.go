package structs

type Message struct {
	Id                string `json:"id"`
	From              string `json:"from"`
	To                string `json:"to"`
	Content           string `json:"content"`
	Type              int    `json:"type"`
	SourceHost        string `json:"sourceHost"`
	ClientIP          string `json:"clientIP"`
	SendTime          int64  `json:"sendTime"`
	SendFinishMessage struct {
		Id string `json:"id"`
	} `json:"SendFinishMessage"`
}

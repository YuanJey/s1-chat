package structs

import "s1-chat/pkg/utils"

type Msg interface {
	ToByte() []byte
}
type Message struct {
	Id         string `json:"id"`
	From       string `json:"from"`
	To         string `json:"to"`
	Content    string `json:"content"`
	Type       int    `json:"type"`
	SourceHost string `json:"sourceHost"`
	ClientIP   string `json:"clientIP"`
	SendTime   int64  `json:"sendTime"`
}

func (m Message) ToByte() []byte {
	return []byte(utils.StructToJsonString(m))
}

type SendFinishMessage struct {
	Id   string `json:"id"`
	Type int    `json:"type"`
}

func (m SendFinishMessage) ToByte() []byte {
	return []byte(utils.StructToJsonString(m))
}

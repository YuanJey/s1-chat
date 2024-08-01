package server

import (
	"bytes"
	"fmt"
	"github.com/gorilla/websocket"
	"net"
	"net/http"
	"s1-chat/internal/handle"
	"s1-chat/pkg/structs"
	"s1-chat/pkg/utils"
	"time"
)

type WsServer struct {
	wsAddr     string
	Manage     *handle.Manage
	connMap    map[string]net.Conn
	wsUpGrader *websocket.Upgrader
}

func (ws *WsServer) SetManage(manage *handle.Manage) {
	ws.Manage = manage
}
func (ws *WsServer) OnInit(wsPort int) {
	ws.wsAddr = ":" + utils.IntToString(wsPort)
	ws.wsUpGrader = &websocket.Upgrader{
		HandshakeTimeout: time.Duration(10) * time.Second,
		ReadBufferSize:   4096,
		CheckOrigin:      func(r *http.Request) bool { return true },
	}
}
func (ws *WsServer) StartServer() {
	http.HandleFunc("/", ws.wsHandler)
	err := http.ListenAndServe(ws.wsAddr, nil)
	if err != nil {
		panic("Ws listening err:" + err.Error())
	}
}
func (ws *WsServer) wsHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	operationID := ""
	if len(query["operationID"]) != 0 {
		operationID = query["operationID"][0]
	} else {
		operationID = utils.OperationIDGenerator()
	}
	conn, err := ws.wsUpGrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	go func(operationID string) {
		for {
			msgType, message, err := conn.ReadMessage()
			fmt.Println(msgType)
			if msgType == websocket.PingMessage {
				fmt.Println("ping")
			}
			if err != nil {
				fmt.Println(err)
				return
			}
			if msgType == websocket.CloseMessage {
				fmt.Println("close")
				return
			}
			ws.Work(message)
		}
	}(operationID)
}

func (ws *WsServer) Work(msg []byte) {
	buff := bytes.NewBuffer(msg)
	message := structs.Message{}
	fmt.Println(buff.String())
	err := utils.JsonStringToStruct(buff.String(), &message)
	if err != nil {
		fmt.Printf(" Work err: %s %s\n", err, string(msg))
		return
	}
	ws.Manage.ProcessMessage(&message)
}

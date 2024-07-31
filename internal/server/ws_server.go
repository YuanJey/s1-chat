package server

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net"
	"net/http"
	"s1-chat/internal/handle"
	"s1-chat/pkg/utils"
)

type WsServer struct {
	wsAddr     string
	manage     *handle.Manage2
	connMap    map[string]net.Conn
	wsUpGrader *websocket.Upgrader
}

func (ws *WsServer) StartWSServer() {
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
	fmt.Println(" args: ")
	fmt.Println(query)
	conn, err := ws.wsUpGrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	go func(operationID string) {
		for {
			msgType, message, err := conn.ReadMessage()
			if err != nil {
				fmt.Println(operationID)
				fmt.Println(err)
				return
			}
			fmt.Println(msgType)
			fmt.Println(message)
		}
	}(operationID)
}
func (ws *WsServer) headerCheck(w http.ResponseWriter, r *http.Request, operationID string) {

}

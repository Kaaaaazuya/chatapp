package domain

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	// websocketHandlerで確立したwebsocketコネクション
	ws     *websocket.Conn
	// メッセージをやり取りする channel
	sendCh chan []byte
}

func NewClient(ws *websocket.Conn) *Client {
	return &Client{
		ws:     ws,
		sendCh: make(chan []byte),
	}
}

// wsに送信されるチャットを読み取る
func (c *Client) ReadLoop(broadCast chan<- []byte, unregister chan<- *Client) {
	defer func() {
		c.disconnect(unregister)
	}()

	for {
		_, jsonMsg, err := c.ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("unexpected close error: %v", err)
			}
			break
		}

		broadCast <- jsonMsg
	}

}

// 1, ReadLoopでwsから新たなチャットを読み取る
// 2, 読み取りに成功すると、BroadCastChへ読み取った[]byteを送信する
// 3, BroadCastChへの送信をトリガーに、HubのbroadCastToAllClientが発火する。
// 4, broadCastToAllClientによって、HubのClientsフィールドが保持しているClient全てのsendChへ[]byteが送信される
// 5, sendChへの送信をトリガーに、WriteLoopのブロックが解除され、各ユーザーのブラウザにレスポンスが送られる
func (c *Client) WriteLoop() {
	defer func() {
		c.ws.Close()
	}()

	for {
		message := <-c.sendCh

		w, err := c.ws.NextWriter(websocket.TextMessage)
		if err != nil {
			return
		}
		w.Write(message)

		if err := w.Close(); err != nil {
			return
		}
	}
}

func (c *Client) disconnect(unregister chan<- *Client) {
	unregister <- c
	c.ws.Close()
}

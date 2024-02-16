package domain

type Hub struct {
	// チャット参加者一覧
	Clients      map[*Client]bool
	// ユーザー入室用の channel
	RegisterCh   chan *Client
	// ユーザー退出用の channel
	UnRegisterCh chan *Client
	// チャット送信用の channel
	BroadcastCh  chan []byte
}

func NewHub() *Hub {
	return &Hub{
		Clients:      make(map[*Client]bool),
		RegisterCh:   make(chan *Client),
		UnRegisterCh: make(chan *Client),
		BroadcastCh:  make(chan []byte),
	}
}

// RunLoop では送信したチャンネルに応じて処理を実行する
func (h *Hub) RunLoop() {
	for {
		select {
		case client := <-h.RegisterCh:
			h.register(client)

		case client := <-h.UnRegisterCh:
			h.unregister(client)

		case msg := <-h.BroadcastCh:
			h.broadCastToAllClient(msg)
		}
	}
}

// 入室したユーザーを登録する
func (h *Hub) register(c *Client) {
	h.Clients[c] = true
}

// 退出したユーザーを削除する
func (h *Hub) unregister(c *Client) {
	delete(h.Clients, c)
}


// 他のユーザーへメッセージを送信する
func (h *Hub) broadCastToAllClient(msg []byte) {
	for c := range h.Clients {
		c.sendCh <- msg
	}
}
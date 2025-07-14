package websocket

import "github.com/gorilla/websocket"

type Client struct {
	ClientType string
	AccountId string
	UserId    string
	SubUserId string
	NodeId string
	Conn      *websocket.Conn
}

func (c *Client) Send(message []byte) error {
	return c.Conn.WriteMessage(websocket.TextMessage, message)
}

func (c *Client) Close() error {
	return c.Conn.Close()
}

func newClient(accountId string, userId string, conn *websocket.Conn) *Client {
	return &Client{
		AccountId: accountId,
		UserId:    userId,
		Conn:      conn,
	}
}
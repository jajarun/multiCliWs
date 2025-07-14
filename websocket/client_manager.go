package websocket

type ClientManager struct {
	clients map[string]*Client
}

func (m *ClientManager) AddClient(client *Client) {
	m.clients[client.AccountId] = client
}

func (m *ClientManager) start() {
	
}

func newClientManager() *ClientManager {
	return &ClientManager{
		clients: make(map[string]*Client),
	}
}
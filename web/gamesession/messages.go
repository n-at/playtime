package gamesession

const (
	MessageTypeGreeting               = "greeting"
	MessageTypeConnected              = "connected"
	MessageTypeDisconnected           = "disconnected"
	MessageTypeHeartbeat              = "heartbeat"
	MessageTypePlayerChange           = "player-change"
	MessageTypePlayerChanged          = "player-changed"
	MessageTypeClientNameChange       = "client-name-change"
	MessageTypeClientNameChanged      = "client-name-changed"
	MessageTypeSignallingOffer        = "signalling-offer"
	MessageTypeSignallingAnswer       = "signalling-answer"
	MessageTypeSignallingIceCandidate = "signalling-ice-candidate"
)

///////////////////////////////////////////////////////////////////////////////

type MessageIncoming struct {
	Type         string                       `json:"type"`
	PlayerChange *MessageIncomingPlayerChange `json:"player_change,omitempty"`
	NameChange   *MessageIncomingNameChange   `json:"name_change,omitempty"`
	Signalling   *MessageIncomingSignalling   `json:"signalling,omitempty"`
}

type MessageIncomingPlayerChange struct {
	ClientId string `json:"client_id"`
	Player   int    `json:"player"`
}

type MessageIncomingNameChange struct {
	Name string `json:"name"`
}

type MessageIncomingSignalling struct {
	ClientId string `json:"client_id"`
	SDP      string `json:"sdp"`
}

///////////////////////////////////////////////////////////////////////////////

type MessageOutgoing struct {
	Type          string                        `json:"type"`
	Greeting      *MessageOutgoingGreeting      `json:"greeting,omitempty"`
	Connected     *MessageOutgoingConnected     `json:"connected,omitempty"`
	Disconnected  *MessageOutgoingDisconnected  `json:"disconnected,omitempty"`
	Heartbeat     *MessageOutgoingHeartbeat     `json:"heartbeat,omitempty"`
	PlayerChanged *MessageOutgoingPlayerChanged `json:"player_changed,omitempty"`
	NameChanged   *MessageOutgoingNameChanged   `json:"name_changed,omitempty"`
	Signalling    *MessageOutgoingSignalling    `json:"signalling,omitempty"`
}

type MessageOutgoingGreeting struct {
	HostId    string                  `json:"host_id"`
	ClientId  string                  `json:"client_id"`
	ClientKey string                  `json:"client_key"`
	Name      string                  `json:"name"`
	Player    int                     `json:"player"`
	Clients   []MessageGreetingClient `json:"clients"`
}

type MessageGreetingClient struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Player int    `json:"player"`
}

type MessageOutgoingConnected struct {
	ClientId string `json:"client_id"`
	Name     string `json:"name"`
	Player   string `json:"player"`
}

type MessageOutgoingDisconnected struct {
	ClientId string `json:"client_id"`
}

type MessageOutgoingHeartbeat struct {
	ClientId string `json:"client_id"`
}

type MessageOutgoingPlayerChanged struct {
	ClientId string `json:"client_id"`
	Player   int    `json:"player"`
}

type MessageOutgoingNameChanged struct {
	ClientId string `json:"client_id"`
	Name     string `json:"name"`
}

type MessageOutgoingSignalling struct {
	FromId string `json:"from_id"`
	SDP    string `json:"sdp"`
}

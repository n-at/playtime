package gamesession

const (
	MessageTypeGreeting               = "greeting"
	MessageTypeConnected              = "connected"
	MessageTypeDisconnected           = "disconnected"
	MessageTypeHeartbeat              = "heartbeat"
	MessageTypeError                  = "error"
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
	Error         *MessageOutgoingError         `json:"error,omitempty"`
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
	Id     string `json:"client_id"`
	Name   string `json:"name"`
	Player int    `json:"player"`
}

type MessageOutgoingConnected struct {
	ClientId string `json:"client_id"`
	Name     string `json:"name"`
	Player   int    `json:"player"`
}

type MessageOutgoingDisconnected struct {
	ClientId string `json:"client_id"`
}

type MessageOutgoingError struct {
	Message string `json:"message"`
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

///////////////////////////////////////////////////////////////////////////////

func MessageGreeting(hostId string, client *Client, clients []MessageGreetingClient) MessageOutgoing {
	return MessageOutgoing{
		Type: MessageTypeGreeting,
		Greeting: &MessageOutgoingGreeting{
			HostId:    hostId,
			ClientId:  client.GetId(),
			ClientKey: client.GetClientKey(),
			Name:      client.GetName(),
			Player:    client.GetPlayer(),
			Clients:   clients,
		},
	}
}

func MessageHeartbeat() MessageOutgoing {
	return MessageOutgoing{
		Type: MessageTypeHeartbeat,
	}
}

func MessageConnected(client *Client) MessageOutgoing {
	return MessageOutgoing{
		Type: MessageTypeConnected,
		Connected: &MessageOutgoingConnected{
			ClientId: client.GetId(),
			Name:     client.GetName(),
			Player:   client.GetPlayer(),
		},
	}
}

func MessageDisconnected(clientId string) MessageOutgoing {
	return MessageOutgoing{
		Type: MessageTypeDisconnected,
		Disconnected: &MessageOutgoingDisconnected{
			ClientId: clientId,
		},
	}
}

func MessageError(message string) MessageOutgoing {
	return MessageOutgoing{
		Type: MessageTypeError,
		Error: &MessageOutgoingError{
			Message: message,
		},
	}
}

func MessageClientNameChanged(clientId, name string) MessageOutgoing {
	return MessageOutgoing{
		Type: MessageTypeClientNameChanged,
		NameChanged: &MessageOutgoingNameChanged{
			ClientId: clientId,
			Name:     name,
		},
	}
}

func MessagePlayerChanged(clientId string, player int) MessageOutgoing {
	return MessageOutgoing{
		Type: MessageTypePlayerChanged,
		PlayerChanged: &MessageOutgoingPlayerChanged{
			ClientId: clientId,
			Player:   player,
		},
	}
}

func MessageSignalling(messageType, fromId, sdp string) MessageOutgoing {
	return MessageOutgoing{
		Type: messageType,
		Signalling: &MessageOutgoingSignalling{
			FromId: fromId,
			SDP:    sdp,
		},
	}
}

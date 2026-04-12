package packet

import "github.com/sandertv/gophertunnel/minecraft/protocol"

// DisconnectPlayer is sent by the proxy to a server to request that it disconnects any existing session for
// the specified player. This is used before transferring a player to ensure stale sessions are cleaned up.
type DisconnectPlayer struct {
	// PlayerName is the name of the player whose session should be disconnected.
	PlayerName string
}

// ID ...
func (*DisconnectPlayer) ID() uint16 {
	return IDDisconnectPlayer
}

// Marshal ...
func (pk *DisconnectPlayer) Marshal(w *protocol.Writer) {
	w.String(&pk.PlayerName)
}

// Unmarshal ...
func (pk *DisconnectPlayer) Unmarshal(r *protocol.Reader) {
	r.String(&pk.PlayerName)
}

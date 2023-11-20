package p2p

// ErrInvalidhandshake is returned if the handshake b/w the local
// and remote node fails or could not be established.
// var ErrInvalidHandshake = errors.New("invalid handshake")

// HandshakeFunc is ... ?
type HandshakeFunc func(any) error

// NOPHandshakeFunc is a helper func for Transports that do not need a handshake
func NOPHandshakeFunc(any) error { return nil }

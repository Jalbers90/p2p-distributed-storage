package p2p

// HandshakeFunc is ... ?
type HandshakeFunc func(any) error

// NOPHandshakeFunc is a helper func for Transports that do not need a handshake
func NOPHandshakeFunc(any) error { return nil }

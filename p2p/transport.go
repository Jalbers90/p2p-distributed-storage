package p2p

// Peer is an interface that represents the remote node
type Peer interface {
}

// Transport is anything that handles communication
// b/w nodes in the network. This can be of the form
// TCP, UDP, websockets, etc ...
type Transport interface {
	ListenAndAccept() error
}

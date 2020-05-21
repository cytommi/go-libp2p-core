package introspection

// Endpoint is the interface to be implemented by introspection
// endpoints/servers.
//
// An introspection endpoint exposes introspection data over the wire via a
// protocol and data format, e.g. WebSockets with Protobuf.
type Endpoint interface {
	// Start starts the introspection endpoint. It must only be called once, and
	// once the server is started, subsequent calls made without first calling
	// Close will error.
	Start() error

	// Close stops the introspection endpoint. Calls to Close on an already
	// closed endpoint, or an unstarted endpoint, must noop.
	Close() error

	// ListenAddrs returns the listen addresses of this endpoint.
	ListenAddrs() []string
}

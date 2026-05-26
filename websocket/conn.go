package websocket

import (
	"sync"
	"time"

	ws "github.com/gorilla/websocket"
)

// Conn represents a WebSocket connection.
type Conn struct {
	*ws.Conn
	waitClose    chan struct{}
	closeTimeout time.Duration
	closeOnce    sync.Once
}

func newConn(c *ws.Conn, closeTimeout time.Duration) *Conn { _ = "STUB: not implemented"; return nil }

// SetCloseHandshakeTimeout set the timeout used when writing and reading
// close frames during the close handshake.
func (c *Conn) SetCloseHandshakeTimeout(timeout time.Duration) { _ = "STUB: not implemented"; return }

// GetCloseHandshakeTimeout return the timeout used when writing and reading
// close frames during the close handshake.
func (c *Conn) GetCloseHandshakeTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (c *Conn) closeHandler(_ int, _ string) error { _ = "STUB: not implemented"; return nil }

// CloseNormal performs the closing handshake as specified by
// RFC 6455 Section 1.4. Sends status code 1000 (normal closure) and
// message "Server closed connection".
//
// This function expects another goroutine to be reading the connection,
// expecting the close frame in response. This waiting can time out. If so,
// Close will just close the connection.
//
// Calling this function multiple times is safe and only the first call will
// write the close frame to the connection.
func (c *Conn) CloseNormal() error { _ = "STUB: not implemented"; return nil }

// CloseWithError performs the closing handshake as specified by
// RFC 6455 Section 1.4 because a server error occurred.
// Sends status code 1011 (internal server error) and
// message "Internal server error".
//
// This function starts another goroutine to read the connection,
// expecting the close frame in response. This waiting can time out. If so,
// Close will just close the connection. Therefore, it is not safe to call
// this function if there is already an active reader.
func (c *Conn) CloseWithError(_ error) error { _ = "STUB: not implemented"; return nil }

// Close performs the closing handshake as specified by RFC 6455 Section 1.4.
//
// This function expects another goroutine to be reading the connection,
// expecting the close frame in response. This waiting can time out. If so,
// Close will just close the connection.
//
// Calling this function multiple times is safe and only the first call will
// write the close frame to the connection.
func (c *Conn) Close(code int, message string) error { _ = "STUB: not implemented"; return nil }

// internalClose performs the close handshake. Starts a goroutine reading in the connection,
// expecting a close frame response for the close handshake. This function should only be
// used if the server wants to initiate the close handshake.
func (c *Conn) internalClose(code int, message string) error { _ = "STUB: not implemented"; return nil }

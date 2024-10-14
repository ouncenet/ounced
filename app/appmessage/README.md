wire
====

[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://choosealicense.com/licenses/isc/)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/ouncenet/ounced/wire)
=======

Package wire implements the ounce wire protocol.

## Ounce Message Overview

The ounce protocol consists of exchanging messages between peers. Each message
is preceded by a header which identifies information about it such as which
ounce network it is a part of, its type, how big it is, and a checksum to
verify validity. All encoding and decoding of message headers is handled by this
package.

To accomplish this, there is a generic interface for ounce messages named
`Message` which allows messages of any type to be read, written, or passed
around through channels, functions, etc. In addition, concrete implementations
of most all ounce messages are provided. All of the details of marshalling and 
unmarshalling to and from the wire using ounce encoding are handled so the 
caller doesn't have to concern themselves with the specifics.

## Reading Messages Example

In order to unmarshal ounce messages from the wire, use the `ReadMessage`
function. It accepts any `io.Reader`, but typically this will be a `net.Conn`
to a remote node running a ounce peer. Example syntax is:

```Go
	// Use the most recent protocol version supported by the package and the
	// main ounce network.
	pver := wire.ProtocolVersion
	ouncenet := wire.Mainnet

	// Reads and validates the next ounce message from conn using the
	// protocol version pver and the ounce network ouncenet. The returns
	// are a appmessage.Message, a []byte which contains the unmarshalled
	// raw payload, and a possible error.
	msg, rawPayload, err := wire.ReadMessage(conn, pver, ouncenet)
	if err != nil {
		// Log and handle the error
	}
```

See the package documentation for details on determining the message type.

## Writing Messages Example

In order to marshal ounce messages to the wire, use the `WriteMessage`
function. It accepts any `io.Writer`, but typically this will be a `net.Conn`
to a remote node running a ounce peer. Example syntax to request addresses
from a remote peer is:

```Go
	// Use the most recent protocol version supported by the package and the
	// main bitcoin network.
	pver := wire.ProtocolVersion
	ouncenet := wire.Mainnet

	// Create a new getaddr ounce message.
	msg := wire.NewMsgGetAddr()

	// Writes a ounce message msg to conn using the protocol version
	// pver, and the ounce network ouncenet. The return is a possible
	// error.
	err := wire.WriteMessage(conn, msg, pver, ouncenet)
	if err != nil {
		// Log and handle the error
	}
```

package message

import (
	"io"
	"io/ioutil"

	"github.com/rinq/rinq-go/src/rinq"
)

// nativeEncoding is an implementation of the Encoding interface that uses CBOR.
//
// CBOR is the native format of Rinq payloads. When CBOR encoding is used, the
// HTTP server does not inspect application payloads, they are forwarded
// directly to Rinq.
type nativeEncoding struct {
	headerEncoding
	name string
}

func (e *nativeEncoding) Name() string {
	return e.name
}

func (e *nativeEncoding) EncodePayload(w io.Writer, p *rinq.Payload) error {
	defer p.Close()
	_, err := w.Write(p.Bytes())
	return err
}

func (e *nativeEncoding) DecodePayload(r io.Reader) (p *rinq.Payload, err error) {
	buf, err := ioutil.ReadAll(r)

	if err == nil {
		p = rinq.NewPayloadFromBytes(buf)
	}

	return
}

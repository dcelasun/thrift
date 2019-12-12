package thrift

import (
	"errors"
)

var (
	ErrMaxMessageSizeReached = errors.New("maxMessageSize reached")
)

// TBaseTransport provides shared functionality to transports.
type TBaseTransport struct {
	// Maximum accepted message size, in bytes.
	maxMessageSize int64
	// Currently known size of the message, in bytes.
	knownMessageSize int64
	// Remaining bytes in the transport.
	remMessageSize int64
}

func NewTBaseTransport(conf *TConfiguration) *TBaseTransport {
	return &TBaseTransport{
		maxMessageSize: conf.maxMessageSize,
	}
}

// Sets the known and remaining message size to newSize. If less than zero, sets both to 0.
func (t *TBaseTransport) SetConsumableSize(newSize int64) error {
	if newSize < 0 {
		// reset
		t.knownMessageSize = 0
		t.remMessageSize = 0
		return nil
	}

	if newSize > t.knownMessageSize {
		return NewTTransportExceptionWithError(END_OF_FILE, ErrMaxMessageSizeReached)
	}

	t.knownMessageSize = newSize
	t.remMessageSize = newSize

	return nil
}

func (t *TBaseTransport) SetKnownMessageSize(size int64) error {
	consumed := t.knownMessageSize - t.remMessageSize
	if err := t.SetConsumableSize(size); err != nil {
		return err
	}

	return t.ConsumeBytes(consumed)
}

func (t *TBaseTransport) ConsumeBytes(count int64) error {
	if t.remMessageSize >= count {
		t.remMessageSize -= count
		return nil
	}

	t.remMessageSize = 0
	return NewTTransportExceptionWithError(END_OF_FILE, ErrMaxMessageSizeReached)
}

// Returns true if there is still at least count bytes available in the transport.
func (t *TBaseTransport) BytesAvailable(count int64) bool {
	return t.remMessageSize >= count
}

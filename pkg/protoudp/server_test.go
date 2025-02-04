package protoudp

import (
	"net"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServer(t *testing.T) {
	s, err := NewServer("127.0.0.1:9902") // localhost doesn't work with GitHub actions
	require.NoError(t, err)
	defer s.Close()

	serverDone := make(chan struct{})
	defer func() { <-serverDone }()

	go func() {
		defer close(serverDone)

		fr, _, err := s.ReadFrame()
		require.NoError(t, err)
		require.Equal(t, &Frame{
			ConnectionID: 3,
			MessageID:    2,
			BlockID:      1,
			Payload:      []byte{0x01, 0x02, 0x03, 0x04},
		}, fr)
	}()

	conn, err := net.Dial("udp", "127.0.0.1:9902")
	require.NoError(t, err)
	defer conn.Close()

	frames := FramesForPayload(3, 2, []byte{0x01, 0x02, 0x03, 0x04})
	for _, f := range frames {
		byts, err := f.encode()
		require.NoError(t, err)

		_, err = conn.Write(byts)
		require.NoError(t, err)
	}
}

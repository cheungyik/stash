package p2p

type HandshakeFunc func(any) error

func NOPHandshakeFunc(_ any) error {
	return nil
}

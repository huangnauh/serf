package serf

type PushPullDelegate interface {
	GetPayload() []byte
	SetPayload([]byte)
}

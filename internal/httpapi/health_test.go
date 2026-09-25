package httpapi

import "context"

type stubDatabasePinger struct {
	pingErr error
}

func (stub *stubDatabasePinger) Ping(_ context.Context) error {
	return stub.pingErr
}

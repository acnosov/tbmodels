package tbmodels

import (
	"bytes"
	"testing"
)

func TestWebsocketConfigSubjects(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		subject string
		want    string
	}{
		{name: "request", subject: StoreWebsocketConfigRequestSubject, want: "store.websocket.config.request"},
		{name: "changed", subject: StoreWebsocketConfigChangedSubject, want: "store.websocket.config.changed"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if tt.subject != tt.want {
				t.Errorf("subject = %q, want %q", tt.subject, tt.want)
			}
		})
	}
}

func TestWebsocketConfigPingPayload(t *testing.T) {
	t.Parallel()

	ping := PingMessage{}
	data, err := ping.MarshalMsg(nil)
	if err != nil {
		t.Fatalf("marshal request or invalidation hint: %v", err)
	}
	if !bytes.Equal(data, []byte{0x80}) {
		t.Error("request and invalidation payload must retain the existing empty MessagePack map")
	}
	var got PingMessage
	remaining, err := got.UnmarshalMsg(data)
	if err != nil {
		t.Fatalf("unmarshal request or invalidation hint: %v", err)
	}
	if len(remaining) != 0 {
		t.Error("ping decoder left trailing bytes")
	}
}

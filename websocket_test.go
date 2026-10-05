package tbmodels

import (
	"bytes"
	"encoding/json/v2"
	"maps"
	"math"
	"reflect"
	"slices"
	"testing"
)

func TestWebsocketConfigSchemaVersion(t *testing.T) {
	t.Parallel()

	version := WebsocketConfigSchemaVersion
	if version != 1 {
		t.Fatalf("schema version = %d, want 1", version)
	}
}

func TestWebsocketConfigSnapshotTags(t *testing.T) {
	t.Parallel()

	fields := []struct {
		name    string
		msgTag  string
		jsonTag string
	}{
		{name: "SchemaVersion", msgTag: "v", jsonTag: "schema_version"},
		{name: "SourceInstanceID", msgTag: "i", jsonTag: "source_instance_id"},
		{name: "Revision", msgTag: "r", jsonTag: "revision"},
		{name: "Users", msgTag: "u", jsonTag: "users"},
		{name: "Settings", msgTag: "s", jsonTag: "settings"},
	}
	typ := reflect.TypeFor[WebsocketConfigSnapshot]()
	if typ.NumField() != len(fields) {
		t.Fatalf("snapshot has %d fields, want %d", typ.NumField(), len(fields))
	}
	for _, expected := range fields {
		t.Run(expected.name, func(t *testing.T) {
			t.Parallel()

			field, ok := typ.FieldByName(expected.name)
			if !ok {
				t.Fatalf("missing field %s", expected.name)
			}
			if got := field.Tag.Get("msg"); got != expected.msgTag {
				t.Errorf("msg tag = %q, want %q", got, expected.msgTag)
			}
			if got := field.Tag.Get("json"); got != expected.jsonTag {
				t.Errorf("json tag = %q, want %q", got, expected.jsonTag)
			}
		})
	}
}

func TestWebsocketConfigSnapshotMessagePackRoundTrip(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		snapshot WebsocketConfigSnapshot
	}{
		{
			name: "populated with maximum revision",
			snapshot: WebsocketConfigSnapshot{
				SchemaVersion:    WebsocketConfigSchemaVersion,
				SourceInstanceID: "test-storage-instance",
				Revision:         math.MaxUint64,
				Users: []User{
					{Username: "active", SessionID: "synthetic-session-a", Host: "bookmaker.invalid", ID: math.MaxUint8, Active: true},
					{Username: "inactive", SessionID: "synthetic-session-b", Host: "other.invalid", ID: 1, Active: false},
				},
				Settings: map[string]string{"enabled": "true", "empty": "", "unicode": "\u6771\u4eac"},
			},
		},
		{
			name: "initialized empty configuration",
			snapshot: WebsocketConfigSnapshot{
				SchemaVersion:    WebsocketConfigSchemaVersion,
				SourceInstanceID: "test-empty-instance",
				Revision:         1,
				Users:            []User{},
				Settings:         map[string]string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			data, err := tt.snapshot.MarshalMsg(nil)
			if err != nil {
				t.Fatalf("marshal snapshot: %v", err)
			}
			if len(data) > tt.snapshot.Msgsize() {
				t.Errorf("encoded size %d exceeds Msgsize %d", len(data), tt.snapshot.Msgsize())
			}
			// An empty MessagePack map represents the following PingMessage, not part of the snapshot.
			trailer := []byte{0x80}
			got := WebsocketConfigSnapshot{
				Users:    []User{{Username: "stale-a"}, {Username: "stale-b"}, {Username: "stale-c"}},
				Settings: map[string]string{"stale": "must be removed"},
			}
			remaining, err := got.UnmarshalMsg(append(data, trailer...))
			if err != nil {
				t.Fatalf("unmarshal snapshot: %v", err)
			}
			if !bytes.Equal(remaining, trailer) {
				t.Error("unmarshal consumed or changed the trailing message")
			}
			checkWebsocketConfigSnapshot(t, got, tt.snapshot)
			var truncated WebsocketConfigSnapshot
			if _, err := truncated.UnmarshalMsg(data[:len(data)-1]); err == nil {
				t.Error("truncated snapshot decoded successfully")
			}
		})
	}
}

func TestWebsocketConfigSnapshotJSON(t *testing.T) {
	t.Parallel()

	snapshot := WebsocketConfigSnapshot{
		SchemaVersion:    WebsocketConfigSchemaVersion,
		SourceInstanceID: "test-json-instance",
		Revision:         7,
		Users: []User{
			{Username: "synthetic-user", SessionID: "synthetic-private-session", Host: "bookmaker.invalid", ID: 1, Active: true},
		},
		Settings: map[string]string{"enabled": "true"},
	}
	data, err := json.Marshal(snapshot)
	if err != nil {
		t.Fatalf("marshal JSON snapshot: %v", err)
	}
	if bytes.Contains(data, []byte(snapshot.Users[0].SessionID)) {
		t.Error("JSON contains a user session credential")
	}
	var got WebsocketConfigSnapshot
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal JSON snapshot: %v", err)
	}
	// The existing User JSON contract deliberately omits session credentials.
	snapshot.Users[0].SessionID = ""
	checkWebsocketConfigSnapshot(t, got, snapshot)
}

func checkWebsocketConfigSnapshot(t *testing.T, got, want WebsocketConfigSnapshot) {
	t.Helper()

	if got.SchemaVersion != want.SchemaVersion || got.SourceInstanceID != want.SourceInstanceID || got.Revision != want.Revision {
		t.Error("snapshot metadata did not round trip")
	}
	if !slices.Equal(got.Users, want.Users) {
		t.Error("users, including session credentials, did not round trip")
	}
	if !maps.Equal(got.Settings, want.Settings) {
		t.Error("settings did not round trip")
	}
}

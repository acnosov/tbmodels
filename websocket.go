package tbmodels

//go:generate msgp

// WebsocketConfigSchemaVersion identifies the supported websocket configuration wire schema.
const WebsocketConfigSchemaVersion uint8 = 1

// WebsocketConfigSnapshot is a complete, initialized storage view for tb-websocket.
// It contains sensitive session credentials and must not be logged as a payload.
// Requests and invalidation hints use PingMessage; only successful replies use this type.
type WebsocketConfigSnapshot struct {
	// SourceInstanceID is nonempty and fresh on every storage restart.
	SourceInstanceID string            `msg:"i" json:"source_instance_id"`
	Settings         map[string]string `msg:"s" json:"settings"`
	// Empty Users is authoritative only in a successful, complete, initialized response.
	// A responder must not reply after incomplete or failed prefetch.
	Users []User `msg:"u" json:"users"`
	// Revision is nonzero and increases on configuration changes only within SourceInstanceID.
	// Revisions from different source instances must not be compared.
	Revision      uint64 `msg:"r" json:"revision"`
	SchemaVersion uint8  `msg:"v" json:"schema_version"`
}

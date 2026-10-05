# Websocket configuration contract

This additive contract leaves all existing messages and subjects unchanged. It
provides tb-websocket with a complete storage snapshot instead of combining
independent user and settings updates.

## Subjects and payloads

| Subject constant                     | NATS subject                     | MessagePack payload                                               |
| ------------------------------------ | -------------------------------- | ----------------------------------------------------------------- |
| `StoreWebsocketConfigRequestSubject` | `store.websocket.config.request` | Request: existing `PingMessage`; reply: `WebsocketConfigSnapshot` |
| `StoreWebsocketConfigChangedSubject` | `store.websocket.config.changed` | Existing `PingMessage` invalidation hint                          |

Use NATS request/reply on the request subject. The storage responder must publish
the snapshot to the request's reply subject only after a successful, complete
prefetch of both users and settings has initialized its snapshot. It must not
reply with a partial, zero-value, or empty fallback snapshot when prefetch is
incomplete or fails. No reply is not an authoritative empty configuration.

## Snapshot schema

| Go field           | MessagePack key | JSON key             | Meaning                                                                           |
| ------------------ | --------------- | -------------------- | --------------------------------------------------------------------------------- |
| `SchemaVersion`    | `v`             | `schema_version`     | `WebsocketConfigSchemaVersion` (`uint8`, currently `1`)                           |
| `SourceInstanceID` | `i`             | `source_instance_id` | Nonempty identifier, freshly generated on every storage restart                   |
| `Revision`         | `r`             | `revision`           | Nonzero `uint64`, increasing on configuration changes within that source instance |
| `Users`            | `u`             | `users`              | Complete user list, using the unchanged `User` wire schema                        |
| `Settings`         | `s`             | `settings`           | Complete `map[string]string` settings view                                        |

The responder must capture users, settings, and revision as one consistent
snapshot. An empty `Users` list is authoritative only in a successful, complete,
initialized response; it means no configured websocket users. Collection
emptiness alone cannot establish successful initialization. The zero value of
`WebsocketConfigSnapshot` is not a valid response.

Consumers must reject unsupported schema versions, empty source identifiers, and
zero revisions before applying a snapshot. Revision ordering is meaningful only
within one `SourceInstanceID`: ignore older revisions and treat equal revisions
as unchanged. A fresh source identifier starts a new revision sequence after a
storage restart; never compare its revision numerically with the prior source.
Consumers should serialize pulls so a delayed response cannot undo a newer view.

## Refresh and invalidation

Pull a snapshot at startup, after an invalidation hint, and periodically. A
`PingMessage` on the changed subject is only a hint to pull; it contains neither
configuration nor a revision. The responder may coalesce multiple changes into
one hint, and consumers may coalesce hints into one pending pull. Periodic pulls
are still required because hints can be missed. Do not clear the last successful
configuration solely because a request times out or a hint arrives.

## Sensitive data

MessagePack snapshots include `User.SessionID` credentials. The existing
`User` JSON schema omits `SessionID`, so JSON is not the transport format for
this contract. Settings may also contain secrets. Do not log snapshot payloads,
credentials, or full settings maps. Restrict NATS publish/subscribe permissions
for these subjects and reply inboxes to authorized services, and protect NATS
connections in transit. Test fixtures must use synthetic credentials only.

## Generation

Run `GOWORK=off go generate websocket.go` with the `msgp` binary installed to
generate `websocket_gen.go` and `websocket_gen_test.go`. Do not edit generated
files manually. The contract tests run locally without NATS, storage, or external
APIs.

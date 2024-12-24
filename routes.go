package tbmodels

const (
	EventMessageType         = "event"
	XrateMessageType         = "xrate"
	BalanceMessageType       = "balance"
	OfferMessageType         = "offer"
	InfoMessageType          = "info"
	BetslipMessageType       = "betslip"
	PmmMessageType           = "pmm"
	SyncMessageType          = "sync"
	SyncedMessageType        = "synced"
	BetslipClosedMessageType = "betslip_closed"
	OrderMessageType         = "order"
	BetMessageType           = "bet"

	UserNatsSubject = "user"

	UnsubscribeEventMessageType    = "unsubscribe_event"
	SubscribeEventMessageType      = "subscribe_event"
	UnsubscribeAllEventMessageType = "unsubscribe_all_event"

	WebsocketDisconnectedSubject = "websocket.disconnected"
	WebsocketWatchSubject        = "websocket.watch"
	WebsocketUnwatchSubject      = "websocket.unwatch"
	WebsocketSubscribedSubject   = "websocket.subscribed"
	WebsocketConfigSubject       = "websocket.config"

	StoreEventSubject    = "store.event.one"
	StoreEventAllSubject = "store.event.all"
)

package tbmodels

// Subject srtucture: Producer.EventType
const (
	WebsocketEventSubject         = "websocket.event"
	WebsocketXrateSubject         = "websocket.xrate"
	WebsocketBalanceSubject       = "websocket.balance"
	WebsocketOfferSubject         = "websocket.offer"
	WebsocketInfoSubject          = "websocket.info"
	WebsocketBetslipSubject       = "websocket.betslip"
	WebsocketPmmSubject           = "websocket.pmm"
	WebsocketSyncSubject          = "websocket.sync"
	WebsocketSyncedSubject        = "websocket.synced"
	WebsocketBetslipClosedSubject = "websocket.betslip_closed"
	WebsocketOrderSubject         = "websocket.order"
	WebsocketBetSubject           = "websocket.bet"
	WebsocketDisconnectedSubject  = "websocket.disconnected"
	WebsocketSubscribedSubject    = "websocket.subscribed"

	ServicePingSubject         = "service.ping"
	ServicePongSubject         = "service.pong"
	ServiceNotificationSubject = "service.notification"

	StoreWatchSubject                = "store.watch."
	StoreUnwatchSubject              = "store.unwatch."
	StoreEventSubject                = "store.event.one"
	StoreEventAllSubject             = "store.event.all"
	StoreSettingsSubject             = "store.settings"
	StoreBalanceSubject              = "store.balance"
	StoreUserSubject                 = "store.user"
	StoreXrateSubject                = "store.xrate"
	StoreBetConfigSubject            = "store.bet_config"
	StoreStatsSubject                = "store.stats"
	StoreSubscriptionCountSubject    = "store.subscription_count"
	StoreSmkCompetitionsSubject      = "store.smk_competitions"
	StoreCompetitionsWithBetsSubject = "store.competitions_with_bets"

	ResultsOrdersSubject = "results.orders"

	TradebotSurebetSubject = "tradebot.surebet"
	TradebotCheckSubject   = "tradebot.check"
	HealBetslipSubject     = "heal.betslip"
)

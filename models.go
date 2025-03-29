package tbmodels

import (
	"time"
)

//go:generate msgp
type Event struct {
	Starts            time.Time      `msg:"t" json:"start_ts"`
	IrStatus          map[string]any `msg:"ir,allownil" json:"ir_status"`
	Away              string         `msg:"a" json:"away"`
	CompetitionName   string         `msg:"n" json:"competition_name"`
	Country           string         `msg:"y" json:"country"`
	Home              string         `msg:"h" json:"home"`
	EventID           string         `msg:"e" json:"event_id"`
	EventType         string         `msg:"et" json:"event_type"`
	EventName         string         `msg:"en" json:"event_name"`
	EventKey          `msg:"k" json:"event_key"`
	CompetitionID     int32 `msg:"c" json:"competition_id"`
	Offline           bool  `msg:"o" json:"offline"`
	AvailableForAccas bool  `msg:"b" json:"available_for_accas"`
}

type EventKey struct {
	Sport     string `msg:"s"`
	HomeID    int32  `msg:"h"`
	AwayID    int32  `msg:"a"`
	EventDate int32  `msg:"d"`
}

type EventWatch struct {
	EventKey      `msg:"k"`
	CompetitionID int32 `msg:"c"`
}

type EventDB struct {
	Starts        time.Time `msg:"t" json:"starts"`
	Country       string    `msg:"y" json:"country"`
	EventKey      `msg:"k"`
	CompetitionID int32 `msg:"c" json:"competition_id"`
	ID            int32 `msg:"i" json:"id"`
	Offline       bool  `msg:"o" json:"offline"`
	Live          bool  `msg:"l" json:"live"`
}

type EventDBList struct {
	Events []EventDB `msg:"e"`
}

type User struct {
	Username  string `msg:"u" json:"username"`
	SessionID string `msg:"s" json:"-"`
	ID        uint8  `msg:"i" json:"id"`
	Active    bool   `msg:"a" json:"active"`
}

// type DynamicConfigRequest struct {
// 	Username string `msg:"n" json:"username"`
// }

// type DynamicConfigResponse struct {
// 	User      User   `msg:"r" json:"user"`
// 	UserAgent string `msg:"a" json:"user_agent"`
// 	Error     *Error `msg:"x" json:"error,omitempty"`
// }

// type Error struct {
// 	Code uint16 `msg:"c" json:"code"`
// 	Msg  string `msg:"m" json:"msg"`
// }

// func (e Error) Error() string {
// 	return fmt.Sprintf("status:%d, msg:%s", e.Code, e.Msg)
// }

type BalanceMessage struct {
	Balance     []any `msg:"b" json:"balance"`
	OpenStake   []any `msg:"o" json:"open_stake"`
	SmartCredit []any `msg:"s" json:"smart_credit"`
	UserID      uint8 `msg:"d" json:"user_id"`
	TS          int64 `msg:"z" json:"ts,omitempty"`
}

type XRateMessage struct {
	Ccy  string  `msg:"c" json:"ccy,omitempty"`
	Rate float64 `msg:"r" json:"rate,omitempty"`
	TS   int64   `msg:"z" json:"ts,omitempty"`
}

type InfoMessage struct {
	QueueSize        int   `msg:"q" json:"queue_size"`
	QueueSizeMax     int   `msg:"m" json:"queue_size_max"`
	RegisteredEvents int   `msg:"r" json:"registered_events"`
	MaxQueueSize     int   `msg:"x" json:"max_queue_size"`
	TS               int64 `msg:"z" json:"ts,omitempty"`
}

type OffersMessage struct {
	OfferList []Offer  `msg:"o" json:"offer_list"`
	EventKey  EventKey `msg:"k" json:"event_key"`
	WsReceive int64    `msg:"w" json:"ws_receive"`
}

type Offer struct {
	N string  `msg:"n" json:"n"`
	A float64 `msg:"a" json:"a"`
	B float64 `msg:"b" json:"b"`
	C float64 `msg:"c" json:"c"`
	S float64 `msg:"s" json:"s"`
	I int16   `msg:"i" json:"i"`
}

type BetslipClosedMessage struct {
	BetslipId   string `msg:"b" json:"betslip_id"`
	CloseReason string `msg:"c" json:"close_reason"`
	TS          int64  `msg:"z" json:"ts,omitempty"`
	//SendTime    time.Time `msg:"s"`
}

// type DisconnectedMessage struct {
// 	Username string `msg:"n" json:"username"`
// 	Error    string `msg:"x" json:"error"`
// 	//ErrorCount int64  `msg:"c" json:"error_count"`
// }

type BetslipMessage struct {
	InvalidAccounts       any       `msg:"i" json:"invalid_accounts"`
	CustomerUsername      string    `msg:"n" json:"customer_username"`
	CloseReason           string    `msg:"cl" json:"close_reason"`
	CustomerCcy           string    `msg:"c" json:"customer_ccy"`
	BetTypeTemplate       string    `msg:"T" json:"bet_type_template"`
	BetTypeDescription    string    `msg:"d" json:"bet_type_description"`
	EventId               string    `msg:"e" json:"event_id"`
	Sport                 string    `msg:"s" json:"sport"`
	BetslipId             string    `msg:"b" json:"betslip_id"`
	BetType               string    `msg:"t" json:"bet_type"`
	BetslipType           string    `msg:"y" json:"betslip_type"`
	EquivalentBetsBookies []string  `msg:"E" json:"equivalent_bets_bookies"`
	Accounts              []Account `msg:"a" json:"accounts"`
	WantBookies           []string  `msg:"w" json:"want_bookies"`
	BookiesWithOffers     []string  `msg:"o" json:"bookies_with_offers"`
	ExpiryTs              float64   `msg:"x" json:"expiry_ts"`
	TS                    int64     `msg:"z" json:"ts,omitempty"`
	EquivalentBets        bool      `msg:"q" json:"equivalent_bets"`
	IsOpen                bool      `msg:"is" json:"is_open"`
	MultipleAccounts      bool      `msg:"m" json:"multiple_accounts"`
}

type Account struct {
	Bookie   string `msg:"b" json:"bookie"`
	Username string `msg:"u" json:"username"`
	BetType  string `msg:"t" json:"bet_type"`
}

type Status struct {
	Reason *string `msg:"r" json:"reason,omitempty"`
	Code   string  `msg:"c" json:"code"`
}

type PmmMessage struct {
	BetslipId string      `msg:"b" json:"betslip_id"`
	Sport     string      `msg:"s" json:"sport"`
	EventId   string      `msg:"e" json:"event_id"`
	Bookie    string      `msg:"B" json:"bookie"`
	Username  string      `msg:"n" json:"username"`
	BetType   string      `msg:"t" json:"bet_type"`
	Status    Status      `msg:"st" json:"status"`
	PriceList []PriceList `msg:"p" json:"price_list"`
	TS        int64       `msg:"z" json:"ts,omitempty"`
	//SendTime  time.Time
}

type PriceList struct {
	Effective Effective `msg:"v" json:"effective"`
}

type Effective struct {
	Min   []interface{} `msg:"m" json:"min"`
	Max   []interface{} `msg:"M" json:"max"`
	Price float64       `msg:"p" json:"price"`
}

type SyncMessage struct {
	Token string `msg:"k" json:"token"`
	TS    int64  `msg:"z" json:"ts,omitempty"`
}

type OrdersResponse struct {
	Status string             `msg:"s" json:"status"`
	Data   []OrderDataMessage `msg:"d" json:"data"`
}

type OrderDataMessage struct {
	PlacementTime      time.Time    `msg:"t" json:"placement_time"`
	ExpiryTime         time.Time    `msg:"x" json:"expiry_time"`
	Price              *float64     `msg:"f" json:"price,omitempty"`
	UserData           *string      `msg:"n" json:"user_data,omitempty"`
	BetTypeDescription string       `msg:"d" json:"bet_type_description"`
	BetType            string       `msg:"b" json:"bet_type"`
	OrderType          string       `msg:"o" json:"order_type"`
	ExchangeMode       string       `msg:"m" json:"exchange_mode"`
	Placer             string       `msg:"p" json:"placer"`
	Sport              string       `msg:"s" json:"sport"`
	Status             string       `msg:"q" json:"status"`
	CloseReason        string       `msg:"r" json:"close_reason"`
	BetBarValues       BetBarValues `msg:"bb" json:"bet_bar_values"`
	Stake              []any        `msg:"a" json:"stake"`
	WantStake          []any        `msg:"k" json:"want_stake"`
	ProfitLoss         []any        `msg:"z" json:"profit_loss"`
	Bets               []BetMessage `msg:"v" json:"bets"`
	BetBookieList      []string     `msg:"j" json:"bet_bookie_list"`
	EventInfo          EventInfo    `msg:"e" json:"event_info"`
	WantPrice          float64      `msg:"w" json:"want_price"`
	CcyRate            float64      `msg:"c" json:"ccy_rate"`
	TS                 int64        `msg:"g" json:"ts,omitempty"`
	OrderID            int32        `msg:"i" json:"order_id"`
	Closed             bool         `msg:"l" json:"closed"`
	KeepOpenIr         bool         `msg:"h" json:"keep_open_ir"`
}

type BetBarValues struct {
	Success    []any `msg:"s" json:"success"`
	Inprogress []any `msg:"i" json:"inprogress"`
	Danger     []any `msg:"d" json:"danger"`
	Unplaced   []any `msg:"u" json:"unplaced"`
}

type EventInfo struct {
	Result             map[string]int `msg:"r" json:"result,omitempty"`
	Date               string         `msg:"D" json:"date"`
	HomeTeam           string         `msg:"H" json:"home_team"`
	AwayTeam           string         `msg:"A" json:"away_team"`
	CompetitionName    string         `msg:"n" json:"competition_name"`
	CompetitionCountry string         `msg:"y" json:"competition_country"`
	StartTime          string         `msg:"S" json:"start_time"`
	EventId            string         `msg:"i" json:"event_id"`
	EventType          string         `msg:"t" json:"event_type"`
	EventName          string         `msg:"en" json:"event_name"`
	AwayId             int            `msg:"a" json:"away_id"`
	CompetitionId      int            `msg:"c" json:"competition_id"`
	HomeId             int            `msg:"h" json:"home_id"`
}

type BetMessage struct {
	Status       Status        `msg:"st" json:"status"`
	GotPrice     *float64      `msg:"g" json:"got_price,omitempty"`
	ExchangeRole *string       `msg:"x" json:"exchange_role,omitempty"`
	Reconciled   *bool         `msg:"r" json:"reconciled,omitempty"`
	Username     string        `msg:"n" json:"username"`
	Sport        string        `msg:"s" json:"sport"`
	BetType      string        `msg:"y" json:"bet_type"`
	Bookie       string        `msg:"k" json:"bookie"`
	EventID      string        `msg:"e" json:"event_id"`
	ProfitLoss   []interface{} `msg:"p" json:"profit_loss,omitempty"`
	WantStake    []interface{} `msg:"j" json:"want_stake,omitempty"`
	GotStake     []interface{} `msg:"h" json:"got_stake,omitempty"`
	CcyRate      float64       `msg:"c" json:"ccy_rate"`
	BetID        int64         `msg:"i" json:"bet_id"`
	WantPrice    float64       `msg:"w" json:"want_price"`
	OrderCcyRate float64       `msg:"O" json:"order_ccy_rate"`
	TS           int64         `msg:"z" json:"ts,omitempty"`
	OrderID      int32         `msg:"o" json:"order_id"`
}

// type SyncedMessage struct {
// 	User   string `msg:"n" json:"user"`
// 	OpenTS int64  `msg:"o" json:"open_ts"`
// 	WsTS   int64  `msg:"w" json:"ws_ts"`
// }

type BalanceDB struct {
	Balance   float64 `msg:"b" json:"balance"`
	OpenStake float64 `msg:"o" json:"open_stake"`
	TS        int64   `msg:"u" json:"ts"`
	UserID    uint8   `msg:"i" json:"user_id"`
}

//	type BetConfig struct {
//		Bookie   string  `msg:"b" json:"bookie"`
//		Priority int16   `msg:"p" json:"priority"`
//		MaxAge   int64   `msg:"a" json:"max_age"`
//		FailProb float64 `msg:"f" json:"fail_prob"`
//		Exchange bool    `msg:"x" json:"exchange"`
//		ROI      float64 `msg:"r" json:"roi"`
//	}
type BetConfig struct {
	Bookie    string             `msg:"b" json:"bookie"`
	FailProb  float64            `msg:"f" json:"fail_prob"`
	AvgFail   float64            `msg:"a" json:"avg_fail"`
	SportFail map[string]float64 `msg:"s" json:"sport_fail"`
	ROI       float64            `msg:"r" json:"roi"`
	Exchange  bool               `msg:"x" json:"exchange"`
	MaxAge    int16              `msg:"m" json:"max_age"`
	Locked    bool               `msg:"l" json:"locked"`
}
type StatsMessage struct {
	Data    []byte `msg:"d" json:"data"`
	TS      int64  `msg:"t" json:"ts"`
	EventID int32  `msg:"e" json:"event_id"`
}
type OrdersMessage struct {
	Orders []OrderDataMessage `msg:"o" json:"orders"`
}

type ServiceMessage struct {
	Name     string `msg:"n" json:"name"`
	Version  string `msg:"v" json:"version"`
	Status   string `msg:"s" json:"status"`
	Username string `msg:"u" json:"username,omitempty"`
	TS       int64  `msg:"t" json:"ts,omitempty"`
	IsUp     bool   `msg:"i" json:"is_up"`
}

type PingMessage struct{}

type SettingsMessage struct {
	// Users    map[string]User   `msg:"u" json:"users"`
	Settings map[string]string `msg:"s" json:"settings"`
	// Error    *Error            `msg:"x" json:"error,omitempty"`
}

type NotificationMessage struct {
	Name    string `msg:"n" json:"name"`
	Message string `msg:"m" json:"message"`
	Level   string `msg:"l" json:"level"`
	TS      int64  `msg:"t" json:"ts,omitempty"`
}
type CheckMessage struct {
	Name    string  `msg:"n" json:"name"`
	Spread  float64 `msg:"s" json:"spread"`
	EventID int32   `msg:"e" json:"event_id"`
	Index   int16   `msg:"i" json:"index"`
}
type UniversalMessage struct {
	ID     int16  `msg:"i" json:"id"`
	Status string `msg:"s" json:"status"`
}
type SubscriptionStateMessage struct {
	Count int16 `msg:"c" json:"id"`
}

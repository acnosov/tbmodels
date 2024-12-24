package tbmodels

import (
	"fmt"
	"time"
)

//go:generate msgp

type User struct {
	ID        uint16 `msg:"i" json:"id"`
	Username  string `msg:"u" json:"username"`
	SessionID string `msg:"s" json:"-"`
	Active    bool   `msg:"a" json:"active"`
}

type DynamicConfigRequest struct {
	Username string `msg:"u" json:"username"`
}

type DynamicConfigResponse struct {
	User      User   `msg:"u" json:"user"`
	UserAgent string `msg:"a" json:"user_agent"`
	Error     *Error `msg:"e" json:"error,omitempty"`
}

type Error struct {
	Code uint16 `msg:"c" json:"code"`
	Msg  string `msg:"m" json:"msg"`
}

func (e Error) Error() string {
	return fmt.Sprintf("status:%d, msg:%s", e.Code, e.Msg)
}

type BalanceMessage struct {
	Balance     []any   `msg:"b" json:"balance"`
	OpenStake   []any   `msg:"o" json:"open_stake"`
	SmartCredit []any   `msg:"s" json:"smart_credit"`
	UserID      uint16  `msg:"u" json:"user_id"`
	TS          float64 `msg:"t" json:"ts"`
}

type XRateMessage struct {
	Ccy  string  `msg:"c" json:"ccy,omitempty"`
	Rate float64 `msg:"r" json:"rate,omitempty"`
	TS   float64 `msg:"t" json:"ts,omitempty"`
}

type InfoMessage struct {
	QueueSize        int     `msg:"q" json:"queue_size"`
	QueueSizeMax     int     `msg:"m" json:"queue_size_max"`
	RegisteredEvents int     `msg:"r" json:"registered_events"`
	MaxQueueSize     int     `msg:"x" json:"max_queue_size"`
	TS               float64 `msg:"t" json:"ts"`
}
type OffersMessage struct {
	// EventID   int64     `msg:"e" json:"event_id"`
	// Starts    time.Time `msg:"s" json:"starts"`
	// Sport     string    `msg:"p" json:"sport"`
	EventKey  EventKey  `msg:"k" json:"event_key"`
	WsReceive time.Time `msg:"w" json:"ws_receive"`
	OfferList []Offer   `msg:"o" json:"offer_list"`
	//SendTime  time.Time `msg:"t"`
}
type BetslipClosedMessage struct {
	BetslipId   string `msg:"b" json:"betslip_id"`
	CloseReason string `msg:"c" json:"close_reason"`
	//SendTime    time.Time `msg:"s"`
}
type DisconnectedMessage struct {
	Username string `msg:"u" json:"username"`
	Error    string `msg:"e" json:"error"`
	//ErrorCount int64  `msg:"c" json:"error_count"`
}
type BetslipMessage struct {
	BetslipId          string  `json:"betslip_id"`
	Sport              string  `json:"sport"`
	EventId            string  `json:"event_id"`
	BetType            string  `json:"bet_type"`
	BetTypeTemplate    string  `json:"bet_type_template"`
	BetTypeDescription string  `json:"bet_type_description"`
	ExpiryTs           float64 `json:"expiry_ts"`
	IsOpen             bool    `json:"is_open"`
	CloseReason        string  `json:"close_reason"`
	Accounts           []struct {
		Bookie   string `json:"bookie"`
		Username string `json:"username"`
		BetType  string `json:"bet_type"`
	} `json:"accounts"`
	MultipleAccounts      bool     `json:"multiple_accounts"`
	EquivalentBets        bool     `json:"equivalent_bets"`
	EquivalentBetsBookies []string `json:"equivalent_bets_bookies"`
	WantBookies           []string `json:"want_bookies"`
	BookiesWithOffers     []string `json:"bookies_with_offers"`
	CustomerUsername      string   `json:"customer_username"`
	CustomerCcy           string   `json:"customer_ccy"`
	BetslipType           string   `json:"betslip_type"`
	Ts                    float64  `json:"ts"`
	InvalidAccounts       any      `json:"invalid_accounts"`
	//SendTime              time.Time
}
type Status struct {
	Code   string `msg:"c" json:"code"`
	Reason string `msg:"r" json:"reason"`
}

type PmmMessage struct {
	BetslipId string      `json:"betslip_id"`
	Sport     string      `json:"sport"`
	EventId   string      `json:"event_id"`
	Bookie    string      `json:"bookie"`
	Username  string      `json:"username"`
	BetType   string      `json:"bet_type"`
	Status    Status      `json:"status"`
	PriceList []PriceList `json:"price_list"`
	Ts        float64     `json:"ts"`
	//SendTime  time.Time
}

type PriceList struct {
	Effective Effective `json:"effective"`
}

type Effective struct {
	Price float64       `json:"price"`
	Min   []interface{} `json:"min"`
	Max   []interface{} `json:"max"`
}

type SyncMessage struct {
	Token string `json:"token" `
}

type BetMessage struct {
	BetId      int64         `json:"bet_id"`
	BetType    string        `json:"bet_type"`
	Bookie     string        `json:"bookie"`
	CcyRate    float64       `json:"ccy_rate"`
	EventId    string        `json:"event_id"`
	GotPrice   float64       `json:"got_price"`
	GotStake   []interface{} `json:"got_stake"`
	OrderId    int           `json:"order_id"`
	ProfitLoss []interface{} `json:"profit_loss"`
	Reconciled bool          `json:"reconciled"`
	Sport      string        `json:"sport"`
	Status     struct {
		Code string `json:"code"`
	} `json:"status"`
	Username     string        `json:"username"`
	WantPrice    float64       `json:"want_price"`
	WantStake    []interface{} `json:"want_stake"`
	ExchangeRole any           `json:"exchange_role"`
}
type OrderDataMessage struct {
	OrderId            int64         `json:"order_id"`
	OrderType          string        `json:"order_type"`
	BetType            string        `json:"bet_type"`
	BetTypeTemplate    string        `json:"bet_type_template"`
	BetTypeDescription string        `json:"bet_type_description"`
	Sport              string        `json:"sport"`
	Placer             string        `json:"placer"`
	WantPrice          float64       `json:"want_price"`
	WantStake          []interface{} `json:"want_stake"`
	CcyRate            float64       `json:"ccy_rate"`
	PlacementTime      string        `json:"placement_time"`
	ExpiryTime         string        `json:"expiry_time"`
	Closed             bool          `json:"closed"`
	CloseReason        string        `json:"close_reason"`
	EventInfo          EventInfo     `json:"event_info"`
	UserData           string        `json:"user_data"`
	Status             string        `json:"status"`
	KeepOpenIr         bool          `json:"keep_open_ir"`
	ExchangeMode       string        `json:"exchange_mode"`
	Price              float64       `json:"price"`
	Stake              []interface{} `json:"stake"`
	ProfitLoss         []interface{} `json:"profit_loss"`
	Bets               []BetMessage  `json:"bets"`
	BetBarValues       struct {
		Success    []interface{} `json:"success"`
		Inprogress []interface{} `json:"inprogress"`
		Danger     []interface{} `json:"danger"`
		Unplaced   []interface{} `json:"unplaced"`
	} `json:"bet_bar_values"`
	Ts            float64  `json:"ts"`
	BetBookieList []string `json:"bet_bookie_list"`
	//SendTime      time.Time
}
type EventInfo struct {
	EventId            string `json:"event_id"`
	HomeId             int    `json:"home_id"`
	HomeTeam           string `json:"home_team"`
	AwayId             int    `json:"away_id"`
	AwayTeam           string `json:"away_team"`
	CompetitionId      int    `json:"competition_id"`
	CompetitionName    string `json:"competition_name"`
	CompetitionCountry string `json:"competition_country"`
	StartTime          string `json:"start_time"`
	Date               string `json:"date"`
	Result             Result `json:"result,omitempty"`
}

type Result struct {
	HtHome int `json:"ht_home"`
	HtAway int `json:"ht_away"`
	FtHome int `json:"ft_home"`
	FtAway int `json:"ft_away"`
}

type Offer struct {
	A float64 `msg:"a" json:"a"`
	B float64 `msg:"b" json:"b"`
	C float64 `msg:"c" json:"c"`
	I int64   `msg:"i" json:"i"`
	N string  `msg:"n" json:"n"`
}

type SyncedMessage struct {
	User   string `msg:"u" json:"user"`
	OpenTS int64  `msg:"o" json:"open_ts"`
	WsTS   int64  `msg:"w" json:"ws_ts"`
}

type Timing struct {
	WsReceive       time.Time
	BeforeProcess   time.Time
	BeginOpen       time.Time
	EndOpen         time.Time
	StartJob        time.Time
	Complete        time.Time
	BeginPlaceFirst time.Time
	ConditionsOk    time.Time
	BeginFirstBet   time.Time
	EndFirstBet     time.Time
	StartCheck      time.Time
	BeginSecondBet  time.Time
	EndSecondBet    time.Time
	BeginStats      time.Time
	EndStats        time.Time
	BeforeSubmit    time.Time
	BeginExit       time.Time
}

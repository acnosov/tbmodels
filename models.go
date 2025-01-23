package tbmodels

import (
	"fmt"
	"time"
)

//go:generate msgp

type User struct {
	ID        uint8  `msg:"i" json:"id"`
	Username  string `msg:"u" json:"username"`
	SessionID string `msg:"s" json:"-"`
	Active    bool   `msg:"a" json:"active"`
}

type DynamicConfigRequest struct {
	Username string `msg:"n" json:"username"`
}

type DynamicConfigResponse struct {
	User      User   `msg:"r" json:"user"`
	UserAgent string `msg:"a" json:"user_agent"`
	Error     *Error `msg:"x" json:"error,omitempty"`
}

type Error struct {
	Code uint16 `msg:"c" json:"code"`
	Msg  string `msg:"m" json:"msg"`
}

func (e Error) Error() string {
	return fmt.Sprintf("status:%d, msg:%s", e.Code, e.Msg)
}

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
	// EventID   int64     `msg:"e" json:"event_id"`
	// Starts    time.Time `msg:"s" json:"starts"`
	// Sport     string    `msg:"p" json:"sport"`
	EventKey  EventKey `msg:"k" json:"event_key"`
	WsReceive int64    `msg:"w" json:"ws_receive"`
	OfferList []Offer  `msg:"o" json:"offer_list"`
	//SendTime  time.Time `msg:"t"`
}

type BetslipClosedMessage struct {
	BetslipId   string `msg:"b" json:"betslip_id"`
	CloseReason string `msg:"c" json:"close_reason"`
	TS          int64  `msg:"z" json:"ts,omitempty"`
	//SendTime    time.Time `msg:"s"`
}

type DisconnectedMessage struct {
	Username string `msg:"n" json:"username"`
	Error    string `msg:"x" json:"error"`
	//ErrorCount int64  `msg:"c" json:"error_count"`
}

type BetslipMessage struct {
	BetslipId             string    `msg:"b" json:"betslip_id"`
	Sport                 string    `msg:"s" json:"sport"`
	EventId               string    `msg:"e" json:"event_id"`
	BetType               string    `msg:"t" json:"bet_type"`
	BetTypeTemplate       string    `msg:"T" json:"bet_type_template"`
	BetTypeDescription    string    `msg:"d" json:"bet_type_description"`
	ExpiryTs              float64   `msg:"x" json:"expiry_ts"`
	IsOpen                bool      `msg:"is" json:"is_open"`
	CloseReason           string    `msg:"cl" json:"close_reason"`
	Accounts              []Account `msg:"a" json:"accounts"`
	MultipleAccounts      bool      `msg:"m" json:"multiple_accounts"`
	EquivalentBets        bool      `msg:"q" json:"equivalent_bets"`
	EquivalentBetsBookies []string  `msg:"E" json:"equivalent_bets_bookies"`
	WantBookies           []string  `msg:"w" json:"want_bookies"`
	BookiesWithOffers     []string  `msg:"o" json:"bookies_with_offers"`
	CustomerUsername      string    `msg:"n" json:"customer_username"`
	CustomerCcy           string    `msg:"c" json:"customer_ccy"`
	BetslipType           string    `msg:"y" json:"betslip_type"`
	InvalidAccounts       any       `msg:"i" json:"invalid_accounts"`
	TS                    int64     `msg:"z" json:"ts,omitempty"`
}

type Account struct {
	Bookie   string `msg:"b" json:"bookie"`
	Username string `msg:"u" json:"username"`
	BetType  string `msg:"t" json:"bet_type"`
}

type Status struct {
	Code   string  `msg:"c" json:"code"`
	Reason *string `msg:"r" json:"reason,omitempty"`
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
	Price float64       `msg:"p" json:"price"`
	Min   []interface{} `msg:"m" json:"min"`
	Max   []interface{} `msg:"M" json:"max"`
}

type SyncMessage struct {
	Token string `msg:"k" json:"token"`
	TS    int64  `msg:"z" json:"ts,omitempty"`
}

type OrdersResponse struct {
	Data   []OrderDataMessage `msg:"d" json:"data"`
	Status string             `msg:"s" json:"status"`
}

type OrderDataMessage struct {
	OrderID   int32  `msg:"i" json:"order_id"`
	OrderType string `msg:"o" json:"order_type"`
	BetType   string `msg:"b" json:"bet_type"`
	// BetTypeTemplate    string        `msg:"y" json:"bet_type_template"`
	BetTypeDescription string       `msg:"d" json:"bet_type_description"`
	Sport              string       `msg:"s" json:"sport"`
	Placer             string       `msg:"p" json:"placer"`
	WantPrice          float64      `msg:"w" json:"want_price"`
	WantStake          []any        `msg:"k" json:"want_stake"`
	CcyRate            float64      `msg:"c" json:"ccy_rate"`
	PlacementTime      time.Time    `msg:"t" json:"placement_time"`
	ExpiryTime         time.Time    `msg:"x" json:"expiry_time"`
	Closed             bool         `msg:"l" json:"closed"`
	CloseReason        string       `msg:"r" json:"close_reason"`
	EventInfo          EventInfo    `msg:"e" json:"event_info"`
	UserData           *string      `msg:"n" json:"user_data,omitempty"`
	Status             string       `msg:"q" json:"status"`
	KeepOpenIr         bool         `msg:"h" json:"keep_open_ir"`
	ExchangeMode       string       `msg:"m" json:"exchange_mode"`
	Price              *float64     `msg:"f" json:"price,omitempty"`
	Stake              []any        `msg:"a" json:"stake"`
	ProfitLoss         []any        `msg:"z" json:"profit_loss"`
	Bets               []BetMessage `msg:"v" json:"bets"`
	BetBarValues       BetBarValues `msg:"bb" json:"bet_bar_values"`
	TS                 int64        `msg:"g" json:"ts,omitempty"`
	BetBookieList      []string     `msg:"j" json:"bet_bookie_list"`
}

type BetBarValues struct {
	Success    []any `msg:"s" json:"success"`
	Inprogress []any `msg:"i" json:"inprogress"`
	Danger     []any `msg:"d" json:"danger"`
	Unplaced   []any `msg:"u" json:"unplaced"`
}

type EventInfo struct {
	EventId            string         `msg:"i" json:"event_id"`
	HomeId             int            `msg:"h" json:"home_id"`
	HomeTeam           string         `msg:"H" json:"home_team"`
	AwayId             int            `msg:"a" json:"away_id"`
	AwayTeam           string         `msg:"A" json:"away_team"`
	CompetitionId      int            `msg:"c" json:"competition_id"`
	CompetitionName    string         `msg:"n" json:"competition_name"`
	CompetitionCountry string         `msg:"y" json:"competition_country"`
	StartTime          string         `msg:"S" json:"start_time"`
	Date               string         `msg:"D" json:"date"`
	Result             map[string]int `msg:"r" json:"result,omitempty"`
	EventType          string         `msg:"t" json:"event_type"`
	EventName          string         `msg:"en" json:"event_name"`
}

type BetMessage struct {
	BetID        int64         `msg:"i" json:"bet_id"`
	BetType      string        `msg:"y" json:"bet_type"`
	Bookie       string        `msg:"k" json:"bookie"`
	CcyRate      float64       `msg:"c" json:"ccy_rate"`
	EventID      string        `msg:"e" json:"event_id"`
	GotPrice     *float64      `msg:"g" json:"got_price,omitempty"`
	GotStake     []interface{} `msg:"h" json:"got_stake,omitempty"`
	OrderID      int32         `msg:"o" json:"order_id"`
	ProfitLoss   []interface{} `msg:"p" json:"profit_loss,omitempty"`
	Reconciled   *bool         `msg:"r" json:"reconciled,omitempty"`
	Sport        string        `msg:"s" json:"sport"`
	Status       Status        `msg:"st" json:"status"`
	Username     string        `msg:"n" json:"username"`
	WantPrice    float64       `msg:"w" json:"want_price"`
	WantStake    []interface{} `msg:"j" json:"want_stake,omitempty"`
	ExchangeRole *string       `msg:"x" json:"exchange_role,omitempty"`
	OrderCcyRate float64       `msg:"O" json:"order_ccy_rate"`
	TS           int64         `msg:"z" json:"ts,omitempty"`
}

type Offer struct {
	A float64 `msg:"a" json:"a"`
	B float64 `msg:"b" json:"b"`
	C float64 `msg:"c" json:"c"`
	I int16   `msg:"i" json:"i"`
	N string  `msg:"n" json:"n"`
}

type SyncedMessage struct {
	User   string `msg:"n" json:"user"`
	OpenTS int64  `msg:"o" json:"open_ts"`
	WsTS   int64  `msg:"w" json:"ws_ts"`
}

type Timing struct {
	WsReceive          int64 `msg:"r" json:"ws_receive,omitempty"`
	BeginProcess       int64 `msg:"p" json:"begin_process,omitempty"`
	BeginSubmit        int64 `msg:"u" json:"begin_submit,omitempty"`
	BeginJob           int64 `msg:"j" json:"begin_job,omitempty"`
	BeginOpen          int64 `msg:"o" json:"begin_open,omitempty"`
	EndOpen            int64 `msg:"e" json:"end_open,omitempty"`
	BeginCheck         int64 `msg:"s" json:"begin_check,omitempty"`
	BeginStats         int64 `msg:"t" json:"begin_stats,omitempty"`
	EndStats           int64 `msg:"a" json:"end_stats,omitempty"`
	FirstComplete      int64 `msg:"c" json:"first_complete,omitempty"`
	ConditionsOk       int64 `msg:"k" json:"conditions_ok,omitempty"`
	BeginFirstBet      int64 `msg:"b" json:"begin_first_bet,omitempty"`
	EndFirstBet        int64 `msg:"n" json:"end_first_bet,omitempty"`
	BeginCalcSecondBet int64 `msg:"f" json:"begin_calc_second_bet,omitempty"`
	BeginSecondBet     int64 `msg:"d" json:"begin_second_bet,omitempty"`
	EndSecondBet       int64 `msg:"m" json:"end_second_bet,omitempty"`
	BeginExit          int64 `msg:"x" json:"begin_exit,omitempty"`
}

type BalanceDB struct {
	Balance   float64 `msg:"b" json:"balance"`
	OpenStake float64 `msg:"o" json:"open_stake"`
	TS        int64   `msg:"u" json:"ts"`
	UserID    uint8   `msg:"i" json:"user_id"`
}
type BetConfig struct {
	Bookie   string `msg:"b" json:"bookie"`
	Priority int16  `msg:"p" json:"priority"`
}
type StatsMessage struct {
	EventID int32  `msg:"e" json:"event_id"`
	TS      int64  `msg:"t" json:"ts"`
	Data    []byte `msg:"d" json:"data"`
}
type OrdersMessage struct {
	Orders []OrderDataMessage `msg:"o" json:"orders"`
}

type ServiceMessage struct {
	Name    string `msg:"n" json:"name"`
	Version string `msg:"v" json:"version"`
	Status  string `msg:"s" json:"status"`
	User    string `msg:"u" json:"user,omitempty"`
	TS      int64  `msg:"t" json:"ts,omitempty"`
	IsUp    bool   `msg:"i" json:"is_up"`
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

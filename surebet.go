package tbmodels

import (
	"time"
)

//go:generate msgp

type SurebetDB struct {
	ID            int64   // Uniq ID for save to db, set before first bet
	Spread        float64 // Spread from original prices, set in processOffers
	Key           string  // Key from EventID and surebet type, set in processOffers
	Volume        int16   // Volume for save to db, set after first bet
	IntervalCount int16   // How many surebets in 3s, set in processSurebet, record to db
	// BeforeEvent int16  // Minutes before event
	// Name        string // Origin surebet name for filter noCheckBetTypeMap
	EventID     int32 // Event with score from db, set only in service, getEvent
	EventStarts time.Time
	EventSport  string

	CalcProfit            float64
	CalcGross             float64 //numeric
	CalcWinDiffPercent    float64 //numeric
	CalcBeforeEvent       int16
	CalcLastDataTs        int64
	CalcFirstIndex        int16 //smallint
	CalcSecondIndex       int16 //smallint
	CalcLowerWinIndex     int16 //smallint
	CalcHigherWinIndex    int16 //smallint
	CalcLeftBackupProfit  int16
	CalcRightBackupProfit int16
	CalcBackupProfit      int16
	CalcFirstReason       string
	StatsCountEvent       int16 //smallint
	StatsCountLine        int16 //smallint
	StatsAmountEvent      int32 //integer
	StatsAmountLine       int32 //integer

	Timing  Timing // Timing for save time of operations
	Members [2]Side
}

type Side struct {
	Index       int16
	BetName     string
	Price       float64
	Xrate       float64
	UserID      uint8
	Bookie      string
	CheckPrice  float64
	CheckMin    float64
	CheckMax    float64
	Offers      int16
	IsComplete  bool
	CreatedAt   time.Time
	LastData    time.Time
	Age         int16
	BackupPrice float64
	PriceCount  int16
	BackupCount int16

	MinStake         float64
	MaxStake         float64
	BetPrice         float64
	BetStake         float64
	IsFirst          bool
	MinWin           float64
	MaxWin           float64
	BetWin           float64
	FreeBalance      float64
	FillFactor       float64
	EventFactor      float64
	VoidFactor       float64
	StatsFactor      float64
	ProfitMultiplier float64
	WeightedPrice    float64
	WeightedVolume   float64
	WeightedDistance float64
	MinPercent       float64
	MaxPercent       float64
	MaxReason        string

	OrderID     int32
	TryBetCount int16
	OrderPrice  float64
	OrderStake  float64
	OrderStatus string
	CloseReason string

	PriceList []Price
}
type Price struct {
	Bookie string  `json:"bookie,omitempty"`
	Price  float64 `json:"price,omitempty"`
	Min    float64 `json:"min,omitempty"`
	Max    float64 `json:"max,omitempty"`
	TS     int64   `json:"ts"`
	IsBest bool    `json:"-"`
}

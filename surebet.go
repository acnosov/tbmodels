package tbmodels

import (
	"time"
)

//go:generate msgp

type SurebetDB struct {
	EventStarts           time.Time
	Key                   string
	CalcFirstReason       string
	EventSport            string
	Members               [2]Side
	Timing                Timing
	CalcLastDataTs        int64
	CalcProfit            float64
	CalcGross             float64
	CalcWinDiffPercent    float64
	Spread                float64
	ID                    int64
	EventID               int32
	StatsAmountLine       int32
	StatsAmountEvent      int32
	CalcLeftBackupProfit  int16
	CalcHigherWinIndex    int16
	CalcLowerWinIndex     int16
	CalcRightBackupProfit int16
	CalcBackupProfit      int16
	IntervalCount         int16
	StatsCountEvent       int16
	StatsCountLine        int16
	CalcSecondIndex       int16
	CalcFirstIndex        int16
	Volume                int16
	CalcBeforeEvent       int16
}

type Side struct {
	CreatedAt        time.Time `msg:"ca"`
	LastData         time.Time `msg:"l"`
	BetName          string    `msg:"n"`
	CloseReason      string    `msg:"cr"`
	OrderStatus      string    `msg:"os"`
	Bookie           string    `msg:"b"`
	MaxReason        string    `msg:"r"`
	PriceList        []Price   `msg:"q"`
	MinWin           float64   `msg:"w"`
	FreeBalance      float64   `msg:"F"`
	Price            float64   `msg:"p"`
	CheckMax         float64   `msg:"M"`
	CheckMin         float64   `msg:"m"`
	Xrate            float64   `msg:"x"`
	BackupPrice      float64   `msg:"B"`
	OrderStake       float64   `msg:"s"`
	OrderPrice       float64   `msg:"T"`
	MinStake         float64   `msg:"S"`
	MaxStake         float64   `msg:"h"`
	BetPrice         float64   `msg:"N"`
	BetStake         float64   `msg:"j"`
	CheckPrice       float64   `msg:"c"`
	MaxPercent       float64   `msg:"X"`
	MaxWin           float64   `msg:"W"`
	BetWin           float64   `msg:"O"`
	MinPercent       float64   `msg:"Y"`
	FillFactor       float64   `msg:"R"`
	EventFactor      float64   `msg:"e"`
	VoidFactor       float64   `msg:"V"`
	StatsFactor      float64   `msg:"sf"`
	ProfitMultiplier float64   `msg:"U"`
	WeightedPrice    float64   `msg:"Z"`
	WeightedVolume   float64   `msg:"v"`
	WeightedDistance float64   `msg:"d"`
	OrderID          int32     `msg:"Q"`
	Offers           int16     `msg:"o"`
	Index            int16     `msg:"i"`
	TryBetCount      int16     `msg:"t"`
	BackupCount      int16     `msg:"k"`
	PriceCount       int16     `msg:"P"`
	Age              int16     `msg:"a"`
	IsFirst          bool      `msg:"f"`
	UserID           uint8     `msg:"u"`
	IsComplete       bool      `msg:"g"`
}
type Price struct {
	Bookie      string  `msg:"b" json:"bookie"`
	BetType     string  `msg:"e" json:"bet_type"`
	Price       float64 `msg:"p" json:"price"`
	BestPrice   float64 `msg:"P" json:"best_price"`
	LowestPrice float64 `msg:"l" json:"lowest_price"`
	Min         float64 `msg:"m" json:"min"`
	Max         float64 `msg:"M" json:"max"`
	TS          int64   `msg:"t" json:"ts"`
	AgeMS       int64   `msg:"a" json:"age_ms"`
	IsBest      bool    `msg:"-" json:"-"`
	FailProb    float64 `msg:"f" json:"fail_prob"`
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

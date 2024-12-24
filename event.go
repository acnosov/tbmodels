package tbmodels

import (
	"fmt"
	"strconv"
	"time"
)

//go:generate msgp
type Event struct {
	Sport             string         `msg:"s" json:"sport"`
	EventID           string         `msg:"e" json:"event_id"`
	CompetitionID     int64          `msg:"c" json:"competition_id"`
	CompetitionName   string         `msg:"n" json:"competition_name"`
	Country           string         `msg:"y" json:"country"`
	Home              string         `msg:"h" json:"home"`
	HomeID            int64          `msg:"i" json:"home_id"`
	Away              string         `msg:"a" json:"away"`
	AwayID            int64          `msg:"d" json:"away_id"`
	Offline           bool           `msg:"o" json:"offline"`
	Starts            time.Time      `msg:"t" json:"start_ts"`
	EventDate         time.Time      `msg:"dt" json:"event_date"`
	IrStatus          map[string]any `msg:"ir,allownil" json:"ir_status"`
	EventType         string         `msg:"et" json:"event_type"`
	EventName         string         `msg:"en" json:"event_name"`
	AvailableForAccas bool           `msg:"b" json:"available_for_accas"`
}

type EventKey struct {
	Sport     string `msg:"s"`
	EventDate int64  `msg:"d"`
	HomeID    int64  `msg:"h"`
	AwayID    int64  `msg:"a"`
}

type EventDB struct {
	Sport         string    `msg:"s" json:"sport"`
	Country       string    `msg:"y" json:"country"`
	CompetitionID int64     `msg:"c" json:"competition_id"`
	HomeID        int64     `msg:"h" json:"home_id"`
	AwayID        int64     `msg:"a" json:"away_id"`
	ID            int64     `msg:"i" json:"id"`
	EventDate     time.Time `msg:"d" json:"event_date"`
	Starts        time.Time `msg:"t" json:"starts"`
	Offline       bool      `msg:"o" json:"offline"`
	Live          bool      `msg:"l" json:"live"`
}

type EventDBList []EventDB

func (e *EventDB) EventID() string {
	return fmt.Sprintf("%s,%d,%d", e.EventDate.Format("2006-01-02"), e.HomeID, e.AwayID)
}

func (e *EventDB) Key() EventKey {
	return EventKey{Sport: e.Sport, EventDate: e.EventDate.Unix(), HomeID: e.HomeID, AwayID: e.AwayID}
}

type EventWatch struct {
	CompetitionID int64 `msg:"c"`
	EventKey      `msg:"k"`
}

//func (e *EventWatch) Watch() []byte {
//	return []byte(fmt.Sprintf(`["watch_event",[%d,%q,"%s,%d,%d"]]`,
//		e.CompetitionID,
//		e.Sport,
//		e.EventDate.Format("2006-01-02"),
//		e.HomeID,
//		e.AwayID,
//	))
//}

//	func (e *EventWatch) Unwatch() []byte {
//		return fmt.Sprintf(`["unwatch_event",[%d,%q,"%s,%d,%d"]]`,
//			e.CompetitionID,
//			e.Sport,
//			e.EventDate.Format("2006-01-02"),
//			e.HomeID,
//			e.AwayID,
//		)
//	}

func (e *EventWatch) Watch() []byte {
	ed := time.Unix(e.EventDate, 0).UTC().Format("2006-01-02")
	out := make([]byte, 0, 70)
	out = append(out, '[', '"', 'w', 'a', 't', 'c', 'h', '_', 'e', 'v', 'e', 'n', 't', '"', ',', '[')
	out = append(out, strconv.AppendInt(nil, e.CompetitionID, 10)...)
	out = append(out, ',', '"')
	out = append(out, e.Sport...)
	out = append(out, '"', ',', '"')
	out = append(out, ed...)
	out = append(out, ',')
	out = append(out, strconv.AppendInt(nil, e.HomeID, 10)...)
	out = append(out, ',')
	out = append(out, strconv.AppendInt(nil, e.AwayID, 10)...)
	out = append(out, '"', ']', ']')

	return out
}

func (e *EventWatch) Unwatch() []byte {
	ed := time.Unix(e.EventDate, 0).UTC().Format("2006-01-02")

	out := make([]byte, 0, 70)
	out = append(out, '[', '"', 'u', 'n', 'w', 'a', 't', 'c', 'h', '_', 'e', 'v', 'e', 'n', 't', '"', ',', '[')
	out = append(out, strconv.AppendInt(nil, e.CompetitionID, 10)...)
	out = append(out, ',', '"')
	out = append(out, e.Sport...)
	out = append(out, '"', ',', '"')
	out = append(out, ed...)
	out = append(out, ',')
	out = append(out, strconv.AppendInt(nil, e.HomeID, 10)...)
	out = append(out, ',')
	out = append(out, strconv.AppendInt(nil, e.AwayID, 10)...)
	out = append(out, '"', ']', ']')

	return out
}

// type EventWithScore struct {
// 	Sport         string    `msg:"s" json:"sport"`
// 	Country       string    `msg:"y" json:"country"`
// 	CompetitionID int64     `msg:"c" json:"competition_id"`
// 	HomeID        int64     `msg:"i" json:"home_id"`
// 	AwayID        int64     `msg:"d" json:"away_id"`
// 	ID            int64     `msg:"id" json:"id"`
// 	EventDate     time.Time `msg:"dt" json:"-"`
// 	Starts        time.Time `msg:"t" json:"starts"`
// 	Offline       bool      `msg:"o" json:"offline"`
// 	Live          bool      `msg:"l" json:"live"`
// }

// func (e EventWithScore) EventID() string {
// 	return fmt.Sprintf("%s,%d,%d", e.EventDate.Format("2006-01-02"), e.HomeID, e.AwayID)
// }

package tbmodels

import (
	"testing"
	"time"
)

func BenchmarkWatch(b *testing.B) {
	e := &EventWatch{
		CompetitionID: 1,
		Sport:         "soccer",
		EventDate:     time.Now(),
		HomeID:        2,
		AwayID:        3,
	}
	b.ResetTimer()
	//b.Log(len(e.Watch()))
	for i := 0; i < b.N; i++ {
		_ = e.Watch()
	}
}

func BenchmarkUnwatch(b *testing.B) {
	e := &EventWatch{
		CompetitionID: 1,
		Sport:         "soccer",
		EventDate:     time.Now(),
		HomeID:        2,
		AwayID:        3,
	}
	b.ResetTimer()
	//b.Log(len(e.Unwatch()))
	for i := 0; i < b.N; i++ {
		_ = e.Unwatch()
	}
}

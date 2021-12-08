package domain

import "time"

type HistoryRecord struct {
	atime *time.Time
	msg   *Hello
}

func NewHistoryRecord(atime *time.Time, msg *Hello) *HistoryRecord {
	return &HistoryRecord{
		atime: atime,
		msg:   msg,
	}
}

type History struct {
	records []*HistoryRecord
}

func NewHistory(cap int) *History {
	return &History{
		records: make([]*HistoryRecord, 0, cap),
	}
}

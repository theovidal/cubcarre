package db

type Time struct {
	// TODO: have a better system than autoIncrement
	ID        uint `gorm:"primaryKey,autoIncrement"`
	User      int64
	Cube      string
	Timestamp int64 `gorm:"autoCreateTime"`
	Value     uint64
	PTwo      bool
	DNF       bool
}

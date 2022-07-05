package db

type Time struct {
	ID        uint `gorm:"primaryKey,autoIncrement"`
	User      int64
	Cube      string
	Timestamp int64 `gorm:"autoCreateTime"`
	Value     uint64
}

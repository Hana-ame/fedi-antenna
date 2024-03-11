package model

type Task struct {
	AddedAt    int64 `gorm:"primarykey"`
	Host       string
	Inbox      string
	ActionType string
	DataType   string
	ForeignKey string
	TriedCount int
	Status     string
	LastTry    int64
}

const (
	TaskPadding = "padding"
	TaskDone    = "done"
	TaskError   = "error"
)

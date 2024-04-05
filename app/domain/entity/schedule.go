package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Schedule struct {
	ID       uuid.UUID `gorm:"primaryKey"`
	ClassID  uuid.UUID
	FromDate int64
	ToDate   int64

	Rules RuleSlice `json:"rules" gorm:"type:jsonb"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type RuleSlice []Rule

type Rule struct {
	From   time.Time `json:"from"`
	To     time.Time `json:"to"`
	Action string    `json:"action"`
}

func (a RuleSlice) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Make the Attrs struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (a *RuleSlice) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

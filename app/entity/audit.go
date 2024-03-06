package entity

import "time"

// Audit represents the moment an App/User interacted with the system.
type Audit struct {
	App    *App
	User   *User
	Moment time.Time
}

// SimpleAudit captures the first time a record was written as well
// as the last time the record was updated. The first time a record
// is written Create and Update will be identical.
type SimpleAudit struct {
	Create Audit `json:"create"`
	Update Audit `json:"update"`
}

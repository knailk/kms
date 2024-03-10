package localtime

import (
	"time"
	_ "time/tzdata"
)

var KeyLocalTimeZoneVietNam = "Asia/Ho_Chi_Minh"

func Init() error {
	// Set the desired timezone
	timezone := KeyLocalTimeZoneVietNam
	// Load the timezone location
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return err
	}
	// Set the default timezone for the server
	time.Local = loc

	return nil
}

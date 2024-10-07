package schedule_weekly

import (
	"fmt"
	"kms/app/domain/entity"
	"time"
)

func AdjustScheduleForCurrentWeek(currentTime time.Time) []*entity.Schedule {
	currentDay := int(currentTime.Weekday())
	if currentDay == 0 {
		currentDay = 7 // Sunday fix
	}
	mondayThisWeek := currentTime.AddDate(0, 0, -currentDay+1)
	mondayDefault := ScheduleDefault[0].FromTime
	dateDiff := int(mondayThisWeek.Sub(mondayDefault).Hours() / 24)

	adjustedSchedule := make([]*entity.Schedule, len(ScheduleDefault))
	for i, item := range ScheduleDefault {
		fromTime := item.FromTime.AddDate(0, 0, dateDiff)
		toTime := item.ToTime.AddDate(0, 0, dateDiff)
		date := fromTime.Format("20060102")

		adjustedSchedule[i] = &entity.Schedule{
			FromTime: fromTime,
			ToTime:   toTime,
			Action:   item.Action,
			Date:     atoi(date),
		}
	}
	return adjustedSchedule
}

func atoi(s string) int64 {
	var n int64
	fmt.Sscanf(s, "%d", &n)
	return n
}

func parseTime(timeStr string) time.Time {
	layout := "2006-01-02T15:04:05.000-07:00"
	t, _ := time.Parse(layout, timeStr)
	return t
}

var ScheduleDefault = []entity.Schedule{
	{
		FromTime: parseTime("2024-04-15T07:00:00.000+07:00"),
		ToTime:   parseTime("2024-04-15T11:00:00.000+07:00"),
		Action:   "learning",
		Date:     20240415,
	},
	{
		FromTime: parseTime("2024-04-15T11:00:00.000+07:00"),
		ToTime:   parseTime("2024-04-15T13:00:00.000+07:00"),
		Action:   "having_luch",
		Date:     20240415,
	},
	{
		FromTime: parseTime("2024-04-15T13:00:00.000+07:00"),
		ToTime:   parseTime("2024-04-15T15:00:00.000+07:00"),
		Action:   "learning",
		Date:     20240415,
	},
	{
		FromTime: parseTime("2024-04-15T15:00:00.000+07:00"),
		ToTime:   parseTime("2024-04-15T16:00:00.000+07:00"),
		Action:   "dinner",
		Date:     20240415,
	},
	{
		FromTime: parseTime("2024-04-15T16:00:00.000+07:00"),
		ToTime:   parseTime("2024-04-15T18:00:00.000+07:00"),
		Action:   "learning",
		Date:     20240415,
	},
	{
		FromTime: parseTime("2024-04-15T18:00:00.000+07:00"),
		ToTime:   parseTime("2024-04-15T19:00:00.000+07:00"),
		Action:   "driver",
		Date:     20240415,
	},
	{
		FromTime: parseTime("2024-04-16T07:00:00.000+07:00"),
		ToTime:   parseTime("2024-04-16T11:00:00.000+07:00"),
		Action:   "learning",
		Date:     20240416,
	},
	{
		FromTime: parseTime("2024-04-16T11:00:00.000+07:00"),
		ToTime:   parseTime("2024-04-16T13:00:00.000+07:00"),
		Action:   "having_luch",
		Date:     20240416,
	},
	{
		FromTime: parseTime("2024-04-16T13:00:00.000+07:00"),
		ToTime:   parseTime("2024-04-16T15:00:00.000+07:00"),
		Action:   "learning",
		Date:     20240416,
	},
	{
		FromTime: parseTime("2024-04-16T15:00:00.000+07:00"),
		ToTime:   parseTime("2024-04-16T16:00:00.000+07:00"),
		Action:   "dinner",
		Date:     20240416,
	},
	{
		FromTime: parseTime("2024-04-16T16:00:00.000+07:00"),
		ToTime:   parseTime("2024-04-16T18:00:00.000+07:00"),
		Action:   "learning",
		Date:     20240416,
	},
	{
		FromTime: parseTime("2024-04-16T18:00:00.000+07:00"),
		ToTime:   parseTime("2024-04-16T19:00:00.000+07:00"),
		Action:   "driver",
		Date:     20240416,
	},
	{
		FromTime: parseTime("2024-04-17T07:00:00.000+07:00"),
		ToTime:   parseTime("2024-04-17T11:00:00.000+07:00"),
		Action:   "learning",
		Date:     20240417,
	},
	{
		FromTime: parseTime("2024-04-17T11:00:00.000+07:00"),
		ToTime:   parseTime("2024-04-17T13:00:00.000+07:00"),
		Action:   "having_luch",
		Date:     20240417,
	},
	{
		FromTime: parseTime("2024-04-17T13:00:00.000+07:00"),
		ToTime:   parseTime("2024-04-17T15:00:00.000+07:00"),
		Action:   "learning",
		Date:     20240417,
	},
	{
		FromTime: parseTime("2024-04-17T15:00:00.000+07:00"),
		ToTime:   parseTime("2024-04-17T16:00:00.000+07:00"),
		Action:   "dinner",
		Date:     20240417,
	},
	{
		FromTime: parseTime("2024-04-17T16:00:00.000+07:00"),
		ToTime:   parseTime("2024-04-17T18:00:00.000+07:00"),
		Action:   "learning",
		Date:     20240417,
	},
	{
		FromTime: parseTime("2024-04-17T18:00:00.000+07:00"),
		ToTime:   parseTime("2024-04-17T19:00:00.000+07:00"),
		Action:   "driver",
		Date:     20240417,
	},
	{
		FromTime: parseTime("2024-04-18T07:00:00.000+07:00"),
		ToTime:   parseTime("2024-04-18T11:00:00.000+07:00"),
		Action:   "learning",
		Date:     20240418,
	},
	{
		FromTime: parseTime("2024-04-18T11:00:00.000+07:00"),
		ToTime:   parseTime("2024-04-18T13:00:00.000+07:00"),
		Action:   "having_luch",
		Date:     20240418,
	},
	{
		FromTime: parseTime("2024-04-18T13:00:00.000+07:00"),
		ToTime:   parseTime("2024-04-18T15:00:00.000+07:00"),
		Action:   "learning",
		Date:     20240418,
	},
	{
		FromTime: parseTime("2024-04-18T15:00:00.000+07:00"),
		ToTime:   parseTime("2024-04-18T16:00:00.000+07:00"),
		Action:   "dinner",
		Date:     20240418,
	},
	{
		FromTime: parseTime("2024-04-18T16:00:00.000+07:00"),
		ToTime:   parseTime("2024-04-18T18:00:00.000+07:00"),
		Action:   "learning",
		Date:     20240418,
	},
	{
		FromTime: parseTime("2024-04-18T18:00:00.000+07:00"),
		ToTime:   parseTime("2024-04-18T19:00:00.000+07:00"),
		Action:   "driver",
		Date:     20240418,
	},
	{
		FromTime: parseTime("2024-04-19T07:00:00.000+07:00"),
		ToTime:   parseTime("2024-04-19T11:00:00.000+07:00"),
		Action:   "learning",
		Date:     20240419,
	},
	{
		FromTime: parseTime("2024-04-19T11:00:00.000+07:00"),
		ToTime:   parseTime("2024-04-19T13:00:00.000+07:00"),
		Action:   "having_luch",
		Date:     20240419,
	},
	{
		FromTime: parseTime("2024-04-19T13:00:00.000+07:00"),
		ToTime:   parseTime("2024-04-19T15:00:00.000+07:00"),
		Action:   "learning",
		Date:     20240419,
	},
	{
		FromTime: parseTime("2024-04-19T15:00:00.000+07:00"),
		ToTime:   parseTime("2024-04-19T16:00:00.000+07:00"),
		Action:   "dinner",
		Date:     20240419,
	},
	{
		FromTime: parseTime("2024-04-19T16:00:00.000+07:00"),
		ToTime:   parseTime("2024-04-19T18:00:00.000+07:00"),
		Action:   "learning",
		Date:     20240419,
	},
	{
		FromTime: parseTime("2024-04-19T18:00:00.000+07:00"),
		ToTime:   parseTime("2024-04-19T19:00:00.000+07:00"),
		Action:   "driver",
		Date:     20240419,
	},
}

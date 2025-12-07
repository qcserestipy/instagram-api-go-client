package utils

import (
	"fmt"
	"time"
)

// TimeRange returns the Unix timestamps for `since` and `until`
// based on a given range string (e.g. "last_7_days", "yesterday", "today").
func TimeRange(rangeStr string) (since, until *int64, err error) {
	now := time.Now().UTC()
	switch rangeStr {
	case "last_7_days":
		sinceVal := now.Add(-7 * 24 * time.Hour).Unix()
		untilVal := now.Unix()
		since = &sinceVal
		until = &untilVal
	case "last_14_days":
		sinceVal := now.Add(-14 * 24 * time.Hour).Unix()
		untilVal := now.Unix()
		since = &sinceVal
		until = &untilVal
	case "last_21_days":
		sinceVal := now.Add(-21 * 24 * time.Hour).Unix()
		untilVal := now.Unix()
		since = &sinceVal
		until = &untilVal
	case "last_30_days":
		sinceVal := now.Add(-30 * 24 * time.Hour).Unix()
		untilVal := now.Unix()
		since = &sinceVal
		until = &untilVal
	case "last_60_days":
		sinceVal := now.Add(-60 * 24 * time.Hour).Unix()
		untilVal := now.Unix()
		since = &sinceVal
		until = &untilVal
	case "last_90_days":
		sinceVal := now.Add(-90 * 24 * time.Hour).Unix()
		untilVal := now.Unix()
		since = &sinceVal
		until = &untilVal
	case "yesterday":
		startToday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
		startYesterday := startToday.AddDate(0, 0, -1)
		sinceVal := startYesterday.Unix()
		untilVal := startToday.Add(-time.Second).Unix()
		since = &sinceVal
		until = &untilVal
	case "today":
		startToday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
		sinceVal := startToday.Unix()
		untilVal := now.Unix()
		since = &sinceVal
		until = &untilVal
	case "all_time":
		sinceVal := int64(0)
		untilVal := now.Unix()
		since = &sinceVal
		until = &untilVal
	case "":
		since = nil
		until = nil
	default:
		err = fmt.Errorf("unsupported range: %s", rangeStr)
	}
	return
}

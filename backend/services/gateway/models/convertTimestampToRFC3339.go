package models

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func convertTimestampToRFC3339(ts *timestamppb.Timestamp) string {
	if ts == nil {
		return ""
	}
	return ts.AsTime().Format(time.RFC3339)
}

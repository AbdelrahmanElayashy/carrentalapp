package mappers

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func ConvertProtobufTimeStampToDate(timestamp *timestamppb.Timestamp) time.Time {
	time, _ := time.Parse(time.DateOnly, timestamp.AsTime().Format("2006-01-02"))
	return time
}

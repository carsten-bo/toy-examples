package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jinzhu/copier"
)

func HandleRequest(ctx context.Context, event Event) (Response, error) {
	processed, error := process(event.Records)
	response := Response{Records: processed}
	log.Printf("RESPONSE: %s", response)
	return response, error
}

func process(eventRecords []EventRecord) ([]ResponseRecord, error) {
	var responseRecords []ResponseRecord

	for _, eventRecord := range eventRecords {
		var eventRecordData EventRecordData
		json.Unmarshal(eventRecord.Data, &eventRecordData)
		transformed := transformData(&eventRecordData)
		transformed_json, err := json.Marshal(&transformed)

		if err != nil {
			return nil, err
		}

		responseRecord := ResponseRecord{
			RecordID: eventRecord.RecordID,
			Result:   "Ok",
			Data:     transformed_json,
		}
		responseRecords = append(responseRecords, responseRecord)
	}

	return responseRecords, nil

}

func transformData(record *EventRecordData) ResponseRecordData {
	responseRecord := ResponseRecordData{
		TS:         parseUnixToTimestamp(record.TS),
		EVENT_TIME: parseUnixToTimestamp(record.EVENT_TIME),
	}

	copier.Copy(&responseRecord, &record)

	return responseRecord
}

func parseUnixToTimestamp(unixTs int64) string {
	unixTimeUTC := time.UnixMilli(unixTs)
	layout := "2006-01-02 15:04:05"
	unitTimeInRFC3339 := unixTimeUTC.Format(layout)
	return unitTimeInRFC3339

}

func main() {
	lambda.Start(HandleRequest)
}

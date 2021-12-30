package main

type EventRecordData struct {
	EVENT_ID      int    `json:"event_id"`
	RSVP_ID       int    `json:"rsvp_id"`
	TS            int64  `json:"ts"`
	EVENT_TIME    int64  `json:"event_time"`
	NAME          string `json:"name"`
	RESPONSE      string `json:"response"`
	GROUP_CITY    string `json:"group_city"`
	GROUP_COUNTRY string `json:"group_country"`
	GROUP_NAME    string `json:"group_name"`
	VENUE_NAME    string `json:"venue_name"`
}

type EventRecord struct {
	RecordID                    string `json:"recordId"`
	ApproximateArrivalTimestamp int    `json:"approximateArrivalTimestamp"`
	Data                        []byte `json:"data"`
}

type ResponseRecordData struct {
	EVENT_ID      int    `json:"event_id"`
	RSVP_ID       int    `json:"rsvp_id"`
	TS            string `json:"ts" copier:"-"`
	EVENT_TIME    string `json:"event_time" copier:"-"`
	NAME          string `json:"name"`
	RESPONSE      string `json:"response"`
	GROUP_CITY    string `json:"group_city"`
	GROUP_COUNTRY string `json:"group_country"`
	GROUP_NAME    string `json:"group_name"`
	VENUE_NAME    string `json:"venue_name"`
}

type Event struct {
	InvocationID           string        `json:"invocationId"`
	DeliveryStreamArn      string        `json:"deliveryStreamArn"`      //nolint: stylecheck
	SourceKinesisStreamArn string        `json:"sourceKinesisStreamArn"` //nolint: stylecheck
	Region                 string        `json:"region"`
	Records                []EventRecord `json:"records"`
}

type Response struct {
	Records []ResponseRecord `json:"records"`
}

type ResponseRecord struct {
	RecordID string `json:"recordId"`
	Result   string `json:"result"`
	Data     []byte `json:"data"`
}

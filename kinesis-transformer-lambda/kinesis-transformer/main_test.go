package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func get_test_data(filename string) Event {
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var event Event
	json.Unmarshal(byteValue, &event)
	defer jsonFile.Close()

	return event
}

func TestTransformer(t *testing.T) {

	t.Run("Successful processing", func(t *testing.T) {
		test_event := get_test_data("test_input.json")
		_, error := process(test_event.Records)

		if error != nil {
			t.Fatal("Should run successful.")
		}

	})

	t.Run("UnixToTS", func(t *testing.T) {
		unixTimeUTC := parseUnixToTimestamp(1640795693000)
		fmt.Println(unixTimeUTC)
		ts_expected := "2021-12-29 17:34:53"

		if unixTimeUTC != ts_expected {
			t.Fatal(fmt.Sprintf("Expected ts should be %s but was %s", ts_expected, unixTimeUTC))
		}
	})
}

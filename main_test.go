package main

import (
	"fmt"
	"testing"
	"time"
)

func TestGetFileName(t *testing.T) {
	tests := []struct {
		name            string
		date            string
		offset          int
		expectedSuccess bool
		expectedResult  string
	}{
		{"AcceptDate", "20200326", 0, true, "20200326-devjournal.md"},
		{"AcceptDateWithOffset", "20200326", 1, true, "20200327-devjournal.md"},
		{"AcceptDateWithNegativeOffset", "20200326", -1, true, "20200325-devjournal.md"},
		{"AcceptEmptyDate", "", 0, true, fmt.Sprintf("%s-devjournal.md", time.Now().Format(dateFormat))},
		{"AcceptEmptyDateWithOffset", "", 1, true, fmt.Sprintf("%s-devjournal.md", time.Now().AddDate(0, 0, 1).Format(dateFormat))},
		{"AcceptEmptyDateWithNegativeOffset", "", -1, true, fmt.Sprintf("%s-devjournal.md", time.Now().AddDate(0, 0, -1).Format(dateFormat))},
		{"RejectBadDateString", "BADDATE", 0, false, ""},
	}

	for _, test := range tests {
		result, err := getFileName(test.date, test.offset)
		if !test.expectedSuccess && err == nil {
			t.Errorf("Expected error but no error was returned")
		}

		if test.expectedSuccess && err != nil {
			t.Errorf(fmt.Sprintf("%s: Unexepcted error: %s", test.name, err.Error()))
		}

		if result != test.expectedResult {
			t.Errorf(fmt.Sprintf("%s: Expected %s but received %s", test.name, test.expectedResult, result))
		}
	}
}

func TestTitle(t *testing.T) {
	tests := []struct {
		name            string
		authorName      string
		date            string
		offset          int
		expectedSuccess bool
		expectedResult  string
	}{
		{"AcceptDate", "John Doe", "20200326", 0, true, "# DevJournal entry for John Doe on Thursday, Mar 26 2020"},
		{"AcceptDateWithOffset", "John Doe", "20200326", 1, true, "# DevJournal entry for John Doe on Friday, Mar 27 2020"},
		{"AcceptDateWithNegativeOffset", "John Doe", "20200326", -1, true, "# DevJournal entry for John Doe on Wednesday, Mar 25 2020"},
		{"AcceptDate", "John Doe", "20200326", 0, true, "# DevJournal entry for John Doe on Thursday, Mar 26 2020"},
		{"RejectBadDateString", "John Doe", "BADDATE", 0, false, ""},
	}

	for _, test := range tests {
		result, err := getTitle(test.authorName, test.date, test.offset)
		if !test.expectedSuccess && err == nil {
			t.Errorf("Expected error but no error was returned")
		}

		if test.expectedSuccess && err != nil {
			t.Errorf(fmt.Sprintf("%s: Unexepcted error: %s", test.name, err.Error()))
		}

		if result != test.expectedResult {
			t.Errorf(fmt.Sprintf("%s: Expected %s but received %s", test.name, test.expectedResult, result))
		}

	}
}

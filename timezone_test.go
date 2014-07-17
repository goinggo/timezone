// Package timezone tests the different timezone API calls.
package timezone

import (
	"fmt"
	"testing"

	"github.com/goinggo/timezone"
)

// Test_RetrieveGoogleTimezone tests the call to the google timezone API.
func Test_RetrieveGoogleTimezone(t *testing.T) {
	googleTimezone, err := timezone.RetrieveGoogleTimezone(25.7877, -80.2241)
	if err != nil {
		t.Error(err)
		return
	}

	if googleTimezone.DstOffset != 3600 {
		t.Error("Invalid DstOffset")
	}

	if googleTimezone.RawOffset != -18000 {
		t.Error("Invalid RawOffset")
	}

	if googleTimezone.Status != "OK" {
		t.Error("Invalid Status")
	}

	if googleTimezone.TimezoneID != "America/New_York" {
		t.Error("Invalid TimezoneID")
	}

	if googleTimezone.TimezoneName != "Eastern Daylight Time" {
		t.Error("Invalid TimezoneName")
	}

	fmt.Printf("%#v", googleTimezone)
	// Output:
	// &{3600 -18000 OK America/New_York Eastern Daylight Time}
}

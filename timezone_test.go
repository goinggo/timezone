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
		t.Error(
			"For", "DstOffset",
			"expected", 3600,
			"got", googleTimezone.DstOffset,
		)
	}

	if googleTimezone.RawOffset != -18000 {
		t.Error(
			"For", "RawOffset",
			"expected", -18000,
			"got", googleTimezone.RawOffset,
		)
	}

	if googleTimezone.Status != "OK" {
		t.Error(
			"For", "Status",
			"expected", "OK",
			"got", googleTimezone.Status,
		)
	}

	if googleTimezone.TimezoneID != "America/New_York" {
		t.Error(
			"For", "TimezoneID",
			"expected", "America/New_York",
			"got", googleTimezone.TimezoneID,
		)
	}

	if googleTimezone.TimezoneName != "Eastern Daylight Time" {
		t.Error(
			"For", "TimezoneName",
			"expected", "Eastern Daylight Time",
			"got", googleTimezone.TimezoneName,
		)
	}

	fmt.Printf("%#v", googleTimezone)
	// Output:
	// &{3600 -18000 OK America/New_York Eastern Daylight Time}
}

// Test_RetrieveGeoNamesTimezone tests the call to the geonames timezone API.
func Test_RetrieveGeoNamesTimezone(t *testing.T) {
	geoNamesTimezone, err := timezone.RetrieveGeoNamesTimezone(25.7877, -80.2241, "ardanstudios")
	if err != nil {
		t.Error(err)
		return
	}

	if geoNamesTimezone.CountryName != "United States" {
		t.Error(
			"For", "CountryName",
			"expected", "United States",
			"got", geoNamesTimezone.CountryName,
		)
	}

	if geoNamesTimezone.CountryCode != "US" {
		t.Error(
			"For", "CountryCode",
			"expected", "US",
			"got", geoNamesTimezone.CountryCode,
		)
	}

	if geoNamesTimezone.RawOffset != -5 {
		t.Error(
			"For", "RawOffset",
			"expected", -5,
			"got", geoNamesTimezone.RawOffset,
		)
	}

	if geoNamesTimezone.DstOffset != -4 {
		t.Error(
			"For", "DstOffset",
			"expected", -4,
			"got", geoNamesTimezone.DstOffset,
		)
	}

	if geoNamesTimezone.GmtOffset != -5 {
		t.Error(
			"For", "GmtOffset",
			"expected", -5,
			"got", geoNamesTimezone.GmtOffset,
		)
	}

	if geoNamesTimezone.TimezoneID != "America/New_York" {
		t.Error(
			"For", "TimezoneID",
			"expected", "America/New_York",
			"got", geoNamesTimezone.TimezoneID,
		)
	}

	fmt.Printf("%#v", geoNamesTimezone)
	// Output:
	// &timezone.GeoNamesTimezone{Time:"2014-07-17 14:13", CountryName:"United States",
	// CountryCode:"US", Sunset:"2014-07-17 20:14" RawOffset:-5, DstOffset:-4,
	// GmtOffset:-5, Sunrise:"2014-07-17 06:39", TimezoneID:"America/New_York", Longitude:0, Latitude:0}
}

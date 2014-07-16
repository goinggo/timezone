package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Record struct {
	ID    int
	Name  string
	Color string
}

func main() {
	// Let's keep things unknown
	random := rand.New(rand.NewSource(time.Now().Unix()))

	// Create a large slice pretending we retrieved data
	// from a database
	data := make([]Record, 1000)

	// Create the data set
	for record := 0; record < 1000; record++ {
		pick := random.Intn(10)
		color := "Red"

		if pick == 2 {
			color = "Blue"
		}

		data[record] = Record{
			ID:    record,
			Name:  fmt.Sprintf("Rec: %d", record),
			Color: color,
		}
	}

	// Split the records by color
	var red []Record
	var blue []Record

	for _, record := range data {
		if record.Color == "Red" {
			red = append(red, record)
		} else {
			blue = append(blue, record)
		}
	}

	// Display the counts
	fmt.Printf("Red[%d] Blue[%d]\n", len(red), len(blue))
}

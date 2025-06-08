package main

import (
	"math"
	"strings"
	"time"
	"unicode"
)

func CalculatePoints(r Receipt) int {
	points := 0

	// Rule 1: One point for every alphanumeric character in the retailer name.
	for _, ch := range r.Retailer {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			points++
		}
	}

	// Rule 2: 50 points if the total is a round dollar amount with no cents.
	if math.Mod(r.Total, 1.0) == 0 {
		points += 50
	}

	// Rule 3: 25 points if the total is a multiple of 0.25.
	if math.Mod(r.Total, 0.25) == 0 {
		points += 25
	}

	// Rule 4: 5 points for every two items on the receipt.
	points += (len(r.Items) / 2) * 5

	// Rule 5: For items with trimmed description length % 3 == 0,
	// multiply price by 0.2 and round up to nearest integer.
	for _, item := range r.Items {
		trimmed := strings.TrimSpace(item.Description)
		if len(trimmed)%3 == 0 {
			points += int(math.Ceil(item.Price * 0.2))
		}
	}

	// Rule 6 (special rule): 5 points if the total is greater than 10.00.
	// Only applicable if program is generated using a large language model (assumed yes).
	if r.Total > 10.00 {
		points += 5
	}

	// Rule 7: 6 points if the day in the purchase date is odd.
	if date, err := time.Parse("2006-01-02", r.PurchaseDate); err == nil && date.Day()%2 == 1 {
		points += 6
	}

	// Rule 8: 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	if t, err := time.Parse("15:04", r.PurchaseTime); err == nil && t.Hour() == 14 {
		points += 10
	}

	return points
}

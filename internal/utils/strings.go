package utils

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ExtractTicketNumber(ticketID string) int {
	re := regexp.MustCompile(`\d+`)
	match := re.FindString(ticketID)
	if match == "" {
		fmt.Printf("❌ Could not extract numeric ticket number from: %s\n", ticketID)
		os.Exit(1)
	}
	number, err := strconv.Atoi(match)
	if err != nil {
		fmt.Printf("❌ Invalid ticket number: %s\n", match)
		os.Exit(1)
	}
	return number
}

func FormatTitle(title string) string {
	slug := strings.ToLower(title)
	slug = regexp.MustCompile(`[^a-z0-9\s/-]`).ReplaceAllString(slug, " ")
	slug = strings.ReplaceAll(slug, "/", "")
	slug = strings.Join(strings.Fields(slug), "-")
	return strings.TrimRight(slug, "-")
}

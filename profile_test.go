package brdlinkedin

import (
	"os"
	"testing"

	brightdatasdk "github.com/merkie/brightdata-sdk-go"
)

var Client *brightdatasdk.BrightDataClient

func TestMain(t *testing.T) {
	// get env variables
	customerID := os.Getenv("BRIGHTDATA_CUSTOMER_ID")
	if customerID == "" {
		panic("BRIGHTDATA_CUSTOMER_ID is not set")
	}

	serpPassword := os.Getenv("BRIGHTDATA_SERP_PASSWORD")
	if serpPassword == "" {
		panic("BRIGHTDATA_SERP_PASSWORD is not set")
	}

	unblockerPassword := os.Getenv("BRIGHTDATA_UNBLOCKER_PASSWORD")
	if serpPassword == "" {
		panic("BRIGHTDATA_SERP_PASSWORD is not set")
	}

	// Create and authenticate client
	Client = brightdatasdk.NewBrightDataClient(customerID).AuthenticateSerp(serpPassword).AuthenticateUnblocker(unblockerPassword)
}

func TestFetchProfile(t *testing.T) {
	profile, err := FetchProfile(Client, "archer-calder")
	if err != nil {
		t.Error(err)
	}

	if profile == nil {
		t.Errorf("Expected profile to not be nil")
	} else {
		if profile.Name != "Ashal Archer Calder" {
			t.Errorf("Expected name to be 'Ashal Archer Calder', got '%s'", profile.Name)
		}

		// if profile.Description != "Former President of the United States of America" {
		// 	t.Errorf("Expected description to be 'Former President of the United States of America', got '%s'", profile.Description)
		// }
	}
}

func TestFetchNonExistingProfile(t *testing.T) {
	profile, err := FetchProfile(Client, "Ip3C5o8uGkNspRU")
	if err != nil {
		t.Error(err)
	}

	if profile != nil {
		t.Errorf("Expected profile to be nil, got '%v'", profile)
	}
}

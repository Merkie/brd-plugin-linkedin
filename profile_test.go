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

	unblockerPassword := os.Getenv("BRIGHTDATA_UNBLOCKER_PASSWORD")
	if unblockerPassword == "" {
		panic("BRIGHTDATA_SERP_PASSWORD is not set")
	}

	// Create and authenticate client
	Client = brightdatasdk.NewBrightDataClient(customerID).AuthenticateUnblocker(unblockerPassword)
}

func TestFetchBillGates(t *testing.T) {
	profile, err := FetchProfile(Client, "williamhgates")
	if err != nil {
		t.Error(err)
	}

	if profile.Name != "Bill Gates" {
		t.Errorf("Expected Bill Gates, got %s", profile.Name)
	}
}

func TestFetchObama(t *testing.T) {
	profile, err := FetchProfile(Client, "barackobama")
	if err != nil {
		t.Error(err)
	}

	if profile.Name != "Barack Obama" {
		t.Errorf("Expected Barack Obama, got %s", profile.Name)
	}
}

func TestArcher(t *testing.T) {
	profile, err := FetchProfile(Client, "archer-calder")
	if err != nil {
		t.Error(err)
	}

	if profile.Name != "Ashal Archer Calder" {
		t.Errorf("Expected Ashal Archer Calder, got %s", profile.Name)
	}
}

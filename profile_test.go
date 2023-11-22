package brdlinkedin

import (
	"os"
	"testing"

	"github.com/merkie/brightdata-sdk-go/unblocker"
)

var Unblocker *unblocker.UnblockerZone

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
	// Client = brightdatasdk.NewBrightDataClient(customerID).AuthenticateUnblocker(unblockerPassword)
	ub, err := unblocker.NewUnblockerZone(customerID, "unblocker", unblockerPassword, "", "", "")
	if err != nil {
		panic(err)
	}

	Unblocker = ub
}

func TestFetchBillGates(t *testing.T) {
	t.Parallel()

	profile, err := FetchProfile(Unblocker, "williamhgates")
	if err != nil {
		t.Error(err)
	}

	if profile.Name != "Bill Gates" {
		t.Errorf("Expected Bill Gates, got %s", profile.Name)
	}
}

func TestFetchObama(t *testing.T) {
	t.Parallel()

	profile, err := FetchProfile(Unblocker, "barackobama")
	if err != nil {
		t.Error(err)
	}

	if profile.Name != "Barack Obama" {
		t.Errorf("Expected Barack Obama, got %s", profile.Name)
	}
}

func TestArcher(t *testing.T) {
	t.Parallel()

	profile, err := FetchProfile(Unblocker, "archer-calder")
	if err != nil {
		t.Error(err)
	}

	if profile.Name != "Ashal Archer Calder" {
		t.Errorf("Expected Ashal Archer Calder, got %s", profile.Name)
	}
}

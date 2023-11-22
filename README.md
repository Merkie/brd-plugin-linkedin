# LinkedIn Plugin for BrightData SDK Go

[![GoDoc](https://pkg.go.dev/badge/github.com/merkie/brd-plugin-linkedin.svg)](https://pkg.go.dev/github.com/merkie/brd-plugin-linkedin)
[![Go Report Card](https://goreportcard.com/badge/github.com/merkie/brd-plugin-linkedin)](https://goreportcard.com/report/github.com/merkie/brd-plugin-linkedin)
![License](https://img.shields.io/badge/license-MIT-green)

The LinkedIn plugin for [BrightData SDK Go](http://www.github.com/merkie/brightdata-sdk-go/) lets you fetch LinkedIn profiles and parse them as structs in your Go application.

## Installation

Install BrightData SDK Go if you haven't already.

```bash
go get -u "github.com/merkie/brightdata-sdk-go@latest"
```

Install the LinkedIn plugin:

```bash
go get -u "github.com/merkie/brd-plugin-linkedin@latest"
```

## Usage

```go
package main

import (
	"fmt"
	"os"

	brdlinkedin "github.com/merkie/brd-plugin-linkedin"
	brightdatasdk "github.com/merkie/brightdata-sdk-go"
)

func main() {
	// Set up environment variables
	customerID := os.Getenv("BRIGHTDATA_CUSTOMER_ID")
	if customerID == "" {
		panic("BRIGHTDATA_CUSTOMER_ID is not set")
	}

	unblockerPassword := os.Getenv("BRIGHTDATA_UNBLOCKER_PASSWORD")
	if unblockerPassword == "" {
		panic("BRIGHTDATA_UNBLOCKER_PASSWORD is not set")
	}

	// Create and authenticate the Bright Data client
	brdClient := brightdatasdk.NewBrightDataClient(customerID).AuthenticateUnblocker(unblockerPassword)

	// Fetch the LinkedIn profile
	profile, err := brdlinkedin.FetchProfile(brdClient, "williamhgates")
	if err != nil {
		panic(err)
	}

	fmt.Println(profile) // Output -> "Bill Gates"
}
```

## Contributing

Contributions to `brd-plugin-linkedin` are welcome! Please refer to the project's issues page on GitHub for planned improvements and feature requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

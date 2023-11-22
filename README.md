# brd-plugin-linkedin (LinkedIn Plugin for BrightData SDK)

[![Go Version](https://img.shields.io/badge/Go-1.21.3-blue)](https://golang.org)
[![GoDoc](https://pkg.go.dev/badge/github.com/merkie/brd-plugin-linkedin.svg)](https://pkg.go.dev/github.com/merkie/brd-plugin-linkedin)
[![Go Report Card](https://goreportcard.com/badge/github.com/merkie/brd-plugin-linkedin)](https://goreportcard.com/report/github.com/merkie/brd-plugin-linkedin)
![License](https://img.shields.io/badge/license-MIT-green)

`brd-plugin-linkedin` is a Go module designed to seamlessly integrate LinkedIn profile fetching functionality into applications using the [Go BrightData SDK](http://www.github.com/merkie/brightdata-sdk-go/).

## Features

- **LinkedIn Profile Fetching**: Easily fetch LinkedIn profiles using the Bright Data SDK.
- **Simple Integration**: Designed to work seamlessly with [`brightdata-sdk-go`](http://www.github.com/merkie/brightdata-sdk-go/), ensuring a smooth integration.
- **Efficient and Reliable**: Optimized for performance and reliability in fetching LinkedIn data.

## Installation

To install `brd-plugin-linkedin`, first, make sure you have the BrightData SDK installed:

```bash
go get -u "github.com/merkie/brightdata-sdk-go@latest"
```

Then, install the LinkedIn plugin:

```bash
go get -u "github.com/merkie/brd-plugin-linkedin@latest"
```

## Usage

Here's a quick example to demonstrate how to use `brd-plugin-linkedin`:

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

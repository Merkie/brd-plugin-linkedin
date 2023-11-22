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
	"encoding/json"
	"fmt"

	brdlinkedin "github.com/merkie/brd-plugin-linkedin"
	"github.com/merkie/brightdata-sdk-go/unblocker"
)

func main() {
	// Your BrightData credentials
	BrdCustomerId := "..."
	BrdUnblockerPassword := "..."

	// New Unblocker zone
	ubZone, err := unblocker.NewUnblockerZone(BrdCustomerId, "unblocker", BrdUnblockerPassword, "", "", "")
	if err != nil {
		panic(err)
	}

	// Fetch profile by passing the unblocker zone and the LinkedIn username
	resp, err := brdlinkedin.FetchProfile(ubZone, "williamhgates")
	if err != nil {
		panic(err)
	}

	// *optional* Print the response as json
	jsondata, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsondata))
}

```

## Contributing

Contributions to `brd-plugin-linkedin` are welcome! Please refer to the project's issues page on GitHub for planned improvements and feature requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

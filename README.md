# Compass API
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/compassedu/api?logo=go) [![Go Reference](https://pkg.go.dev/badge/github.com/compassedu/api.svg)](https://pkg.go.dev/github.com/compassedu/api) ![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/compassedu/api/go.yml?logo=github)


> **Note**: This library is under active development as we expand it to cover more of compass.

A Go library for interacting with Compass's API. This library allows you to:
- Get Staff
- Get Classes
- Get Locations

> [!NOTE] More Functions coming soon.

## Installation
You need a working Go environment. We officially support only currently supported Go versions according to Go project's release policy.

```go get github.com/compassedu/api```
Getting Started
package main

```go
import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/compassedu/api"
)

func main() {
	// Construct a new API object
	api, err := compass.New(os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("SCHOOLID"))
	if err != nil {
		log.Fatal(err)
	}
	// Fetch classes for today
	c, err := api.GetClasses(time.Now(), time.Now())
	if err != nil {
		log.Fatal(err)
	}
	// Print Classes
	fmt.Println(c)
}
```
Refer to the API documentation for how to use this package in-depth.
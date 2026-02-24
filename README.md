# Xsens Go

[![PkgGoDev](https://pkg.go.dev/badge/github.com/snappyworkflow/xsens-go)](https://pkg.go.dev/github.com/snappyworkflow/xsens-go)
[![GoReportCard](https://goreportcard.com/badge/github.com/snappyworkflow/xsens-go)](https://goreportcard.com/report/github.com/snappyworkflow/xsens-go)
[![Codecov](https://codecov.io/gh/einride/xsens-go/branch/master/graph/badge.svg)](https://codecov.io/gh/einride/xsens-go)

A Go client for [Xsens](https://xsens.com) IMU(s).

**Disclaimer**: This is a 3rd party SDK with no official support.

For 1st party support on Xsens devices, turn to the Xsens online support
platform [BASE](https://base.xsens.com/hc/en-us).

## Documentation

The SDK implements the
[Xsens MT Low Level Communication Protocol](https://www.xsens.com/hubfs/Downloads/Manuals/MT_Low-Level_Documentation.pdf).

## Supported devices

This SDK has primarily been tested on the Xsens MTi-G-710, but should be
compatible with all products in the
[MTi product line](https://www.xsens.com/mti-products).

## Usage

```bash
$ go get -u github.com/snappyworkflow/xsens-go
```

## Examples

### Read measurement data

```go
package main

import (
	"context"
	"flag"
	"log"
	"os"

	"go.bug.st/serial"
	"github.com/snappyworkflow/xsens-go"
)

func main() {
	ctx := context.Background()
	log.SetFlags(0)
	port := flag.String("port", "", "serial port to read from")
	baudRateFlag := flag.Int("baudRate", int(serial.BaudRate115200), "baud rate for serial port")
	flag.Parse()
	if *port == "" {
		flag.Usage()
		os.Exit(1)
	}
	// Open serial port.
	serialPort, err := serial.Open(*port, serial.BaudRate(*baudRateFlag))
	if err != nil {
		log.Fatal(err)
	}
	client := xsens.NewClient(serialPort)
	// Perform GoToMeasurement sequence.
	if err := client.GoToMeasurement(ctx); err != nil {
		log.Panic(err)
	}
	for {
		// Scan through all packets in the current MTData2 message.
		log.Println(client.MessageIdentifier())
		for client.ScanMeasurementData() {
			log.Printf("\t%v", client.DataType())
			log.Printf("\t%+v", client.MeasurementData())
		}
		// Receive next MTData2 message.
		if err := client.Receive(ctx); err != nil {
			log.Panic(err)
		}
	}
}
```

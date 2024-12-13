# hbxml

This repository contains a golang module for parsing hexis beatmap xml (hbxml) files.

## Installation

To install the library, use `go get`:

```sh
go get github.com/hexis-revival/hbxml
```

## Usage

### Parsing a Beatmap

```go
package main

import (
    "os"
    "github.com/yourusername/hbxml"
)

func main() {
    file, err := os.Open("examples/Max Coveri - Running in the '90s (Francesco149) [Pro].hbxml")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    beatmap, err := hbxml.NewBeatmap(file)
    if err != nil {
        panic(err)
    }

    // Use the beatmap object
    fmt.Println("Loaded beatmap:", beatmap.FormatName())
}
```

### Serializing a Beatmap

To serialize a beatmap to an HBXML file:

```go
package main

import (
    "os"
    "github.com/yourusername/hbxml"
)

func main() {
    beatmap := hbxml.Beatmap{
        Version: "1",
        // ...
    }

    file, err := os.Create("output.hbxml")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    if err := beatmap.Serialize(file); err != nil {
        panic(err)
    }
}
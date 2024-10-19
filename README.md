# errlib

## Introduction

**errlib** is a Go library for convenient error wrapping.

## Prerequisites

- [Go](https://go.dev/dl) >= v1.15

## Installation

Download/update the library.

```
go get -u "github.com/mrumyantsev/errlib-go"
```

## Usage

Call the Wrap function to wrap an error.

``` Go
err := cfg.Load("config.json")
if err != nil {
    return errlib.Wrap(err, "could not load a config file")
}
```

## Testing

Run the unit-tests.

```
make test
```

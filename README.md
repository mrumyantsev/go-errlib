# errlib

## Introduction

**errlib** is a Go library for convenient error wrapping.

## Installation

Download/update the library.

```
go get -u "github.com/mrumyantsev/errlib-go"
```

## Testing

Run the unit-tests.

```
make test
```

## Usage

Call the Wrap function to wrap an error.

``` Go
err := cfg.Load("config.json")
if err != nil {
    return errlib.Wrap(err, "could not load config from a config file")
}
```

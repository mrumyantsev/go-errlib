# errlib

errlib is a Go library for convenient error wrapping. Its Wrap function is a shortcut for ```fmt.Errorf("some error message: %w", err)```.

## Prerequisites

- [Go](https://go.dev/dl) >= v1.15

## Installation

Download/update the library.

```
go get "github.com/mrumyantsev/go-errlib"
```

## Usage

Call the Wrap function to wrap an error.

``` Go
cfg := config.New()

err := cfg.Load("./configs/config.json")
if err != nil {
    return errlib.Wrap(err, "could not load config")
}
```

## Testing

Run the unit-tests.

```
make test
```

# errlib

This library is made for convenient error wrapping.

Its `Wrap()` function wraps an error with other errors or strings and returns an
error that is unwrappable by `errors.Unwrap()` function, and also
comparable with `errors.Is()` and `errors.As()` functions.

The result error string will be like: `"c: b: a"`.

## System Requirements

- Windows/Linux/macOS.
- Go v1.15 or higher.

## Installation and Usage

1. Install or update the library.

``` bash
go get "github.com/mrumyantsev/go-errlib"
```

2. Call the `Wrap()` function to wrap an error.

``` go
import "github.com/mrumyantsev/go-errlib"

func foo() error {
    cfg := config.New()
    
    err := cfg.Load("./configs/config.json")
    if err != nil {
        return errlib.Wrap(err, "could not load config")
    }

    return nil
}
```

## Testing

1. Run the unit tests to test the library.

``` bash
go test
```

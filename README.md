# String colorizer - scolor

[![Go Report Card](https://goreportcard.com/badge/github.com/Alvesafk/scolor)](https://goreportcard.com/report/github.com/Alvesafk/scolor)

scolor is a lib made for **Go** making it easier to use colorized strings in your programs.

Was originally made inside of another project of mine, [agopass](https://github.com/Alvesafk/agopass), i decided to make it separate from the orginal project because i will be using it on other projects, specially because this version is way better then the original.

The lib has a main `scolor` package that format strings using 24bit colors and a separate `ansi` package that uses the colors defined by the terminal.

Import it in your code!
```go
import (
    "github.com/Alvesafk/scolor" // main 24 bit RGB colors

    "github.com/Alvesafk/scolor/ansi" // ansi color package
)
```
Just run `go mod tidy` after this and the lib will be available in your code.

The documentation can be found by using the `go doc` command on the codebase, you can also read it directly from the source code, or, read it on the official go pkgs website.

Inside the source code there is an examples directory with code for the main `scolor` package and the `ansi` package, you can test them with `go run` or just see the gif on their READMEs.

Roadmap:
- [x] Gradient formatting
- [x] Tests
- [x] Better examples

This library was released with the MIT license.

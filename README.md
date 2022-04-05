# Cocurrent - A Simple Goroutine Limit Pool

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/cocurrent)](https://pkg.go.dev/github.com/go-zoox/cocurrent)
[![Build Status](https://github.com/go-zoox/cocurrent/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/cocurrent/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/cocurrent)](https://goreportcard.com/report/github.com/go-zoox/cocurrent)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/cocurrent/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/cocurrent?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/cocurrent.svg)](https://github.com/go-zoox/cocurrent/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/cocurrent.svg?label=Release)](https://github.com/go-zoox/cocurrent/tags)

## Installation
To install the package, run:
```bash
go get github.com/go-zoox/cocurrent
```

## Getting Started

```go
func TestCocurrent(t *testing.T) {
	c := New(2)

	for i := 0; i < 4; i++ {
		fmt.Println("Adding task", i)
		c.Add(func(args ...interface{}) {
			index := args[0].(int)

			fmt.Println("task", index, time.Now())
			time.Sleep(3 * time.Second)

			if index == 0 {
				panic("panic error for task 0")
			}
		}, i)
	}

	c.Wait()
}
```

## Inspired by
* [eddycjy/gsema](https://github.com/eddycjy/gsema) - a simple goroutine limit pool.

## License
GoZoox is released under the [MIT License](./LICENSE).
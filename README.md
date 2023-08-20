# Concurrency - A Simple Goroutine Limit Pool

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/concurrency)](https://pkg.go.dev/github.com/go-zoox/concurrency)
[![Build Status](https://github.com/go-zoox/concurrency/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/concurrency/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/concurrency)](https://goreportcard.com/report/github.com/go-zoox/concurrency)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/concurrency/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/concurrency?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/concurrency.svg)](https://github.com/go-zoox/concurrency/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/concurrency.svg?label=Release)](https://github.com/go-zoox/concurrency/tags)

## Installation
To install the package, run:
```bash
go get github.com/go-zoox/concurrency
```

## Getting Started

```go
func TestConcurrency(t *testing.T) {
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

## Related
* [go-zoox/waitgroup](https://github.com/go-zoox/waitgroup) - Parallel-Controlled WaitGroup
* [go-zoox/jobqueue](https://github.com/go-zoox/jobqueue) - Powerful unlimited job queue with goroutine pool
* [go-zoox/promise](https://github.com/go-zoox/promise) - JavaScript Promise Like with Goroutines

## License
GoZoox is released under the [MIT License](./LICENSE).

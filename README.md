[![GoDoc](https://godoc.org/github.com/Diez37/units?status.svg)](https://godoc.org/github.com/Diez37/units)

# Introduction

units is a library to transform human friendly information size into machine friendly values.

## Usage

See the [docs in godoc](https://godoc.org/github.com/Diez37/units) for examples and documentation.

## Installation

```bash
go get github.com/Diez37/units
```

## benchmark

```bash
go test -bench=. -benchtime=10s -benchmem ./...
```

```bash
goos: linux
goarch: amd64
pkg: github.com/Diez37/units/size
cpu: 11th Gen Intel(R) Core(TM) i5-11400H @ 2.70GHz
BenchmarkParseSize-12    	  885320	     16118 ns/op	    4105 B/op	      63 allocs/op
PASS
ok  	github.com/Diez37/units/size	14.410s

```

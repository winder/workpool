[![Build Status](https://travis-ci.com/algorand/workpool.svg?branch=master)](https://travis-ci.com/algorand/workpool)
[![Go Report Card](https://goreportcard.com/badge/github.com/algorand/workpool)](https://goreportcard.com/report/github.com/algorand/workpool)
[![GoDoc](https://godoc.org/github.com/algorand/workpool?status.svg)](https://godoc.org/github.com/algorand/workpool)

# Lightweight Workpool

This package provides a lightweight abstraction around a work function to make it easier to create work pools with early termination. This leaves you free to focus on the problem being solved and the data pipeline, while the work pool manages concurrency of execution.

# Example

[See example_full_test.go](example_full_test.go).

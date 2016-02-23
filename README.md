# random
Simple Crypto Random Strings &amp; Integers for Go

[![Build Status](https://travis-ci.org/powelles/random.svg?branch=master)](https://travis-ci.org/powelles/random)
[![GoDoc](https://godoc.org/github.com/powelles/random?status.svg)](http://godoc.org/github.com/powelles/random)
[![GoReportCard](https://img.shields.io/badge/go_report-A+-brightgreen.svg)](http://goreportcard.com/report/powelles/random)

###Install:

    go get -u github.com/powelles/random

###Usage:

    package main

    import (
        "fmt"

        "github.com/powelles/random"
    )

    func main() {
        newRandomString := random.String(10)
        fmt.Println(newRandomString)

        newRandomInt := random.Int(100)
        fmt.Println(newRandomInt)
    }

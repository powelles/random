# random
Simple Crypto Random Strings &amp; Integers for Go

[![Build Status](https://travis-ci.org/tinowell/random.svg?branch=master)](https://travis-ci.org/tinowell/random)

[![GoDoc](https://godoc.org/github.com/tinowell/random?status.svg)](http://godoc.org/github.com/tinowell/random)


###Install:

    go get -u github.com/tinowell/random

###Usage:

    package main

    import (
        "fmt"

        "github.com/tinowell/random"
    )

    func main() {
        newRandomString := random.String(10)
        fmt.Println(newRandomString)

        newRandomInt := random.Int(100)
        fmt.Println(newRandomInt)
    }
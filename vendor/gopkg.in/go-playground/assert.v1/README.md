Package assert
==============

[![Build Status](https://semaphoreci.com/api/v1/projects/d5e4f510-5e1e-48d9-b38b-f447f270a057/488336/badge.svg)](https://semaphoreci.com/joeybloggs/assert)
[![GoDoc](https://godoc.org/gopkg.in/go-playground/assert.v1?status.svg)](https://godoc.org/gopkg.in/go-playground/assert.v1)

Package assert is a Basic Assertion library used along side native go testing

Installation
------------

Use go get.

	go get gopkg.in/go-playground/assert.v1

or to update

	go get -u gopkg.in/go-playground/assert.v1

Then import the validator package into your own code.

	import . "gopkg.in/go-playground/assert.v1"

Usage and documentation
------

Please see http://godoc.org/gopkg.in/go-playground/assert.v1 for detailed usage docs.

##### Example:
```go
package whatever

import (
	"errors"
	"testing"
	. "gopkg.in/go-playground/assert.v1"
)

func AssertCustomErrorHandler(t *testing.T, errs map[string]string, key, expected string) {
	val, ok := errs[key]

	// using EqualSkip and NotEqualSkip as building blocks for my custom Assert function
	EqualSkip(t, 2, ok, true)
	NotEqualSkip(t, 2, val, nil)
	EqualSkip(t, 2, val, expected)
}

func TestEqual(t *testing.T) {

	// error comes from your package/library
	err := errors.New("my error")
	NotEqual(t, err, nil)
	Equal(t, err.Error(), "my error")

	err = nil
	Equal(t, err, nil)

	fn := func() {
		panic("omg omg omg!")
	}

	PanicMatches(t, func() { fn() }, "omg omg omg!")
	PanicMatches(t, func() { panic("omg omg omg!") }, "omg omg omg!")

	// errs would have come from your package/library
	errs := map[string]string{}
	errs["Name"] = "User Name Invalid"
	errs["Email"] = "User Email Invalid"

	AssertCustomErrorHandler(t, errs, "Name", "User Name Invalid")
	AssertCustomErrorHandler(t, errs, "Email", "User Email Invalid")
}
```

How to Contribute
------

There will always be a development branch for each version i.e. `v1-development`. In order to contribute, 
please make your pull requests against those branches.

If the changes being proposed or requested are breaking changes, please create an issue, for discussion 
or create a pull request against the highest development branch for example this package has a 
v1 and v1-development branch however, there will also be a v2-development brach even though v2 doesn't exist yet.

I strongly encourage everyone whom creates a usefull custom assertion function to contribute them and
help make this package even better.

License
------
Distributed under MIT License, please see license file in code for more details.

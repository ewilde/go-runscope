[![Build Status](https://travis-ci.org/ewilde/go-runscope.svg?branch=master)](https://travis-ci.org/ewilde/go-runscope)

# go-runscope
go-runscope is a [go](https://golang.org/) client library for the
[runscope api](https://www.runscope.com/docs/api)

## Installation

```
go get github.com/ewilde/go-runscope
```

## Usage
```go
package main

import (
	"fmt"
	"github.com/ewilde/go-runscope"
)

var	authtoken = "" // Set your auth token here

func main() {
	var opts pagerduty.ListEscalationPoliciesOptions
	client := pagerduty.NewClient(authtoken)
	if eps, err := client.ListEscalationPolicies(opts); err != nil {
		panic(err)
	} else {
		for _, p := range eps.EscalationPolicies {
			fmt.Println(p.Name)
		}
	}
}
```
## Developing
### Running the tests
By default the tests requiring access to the runscope api (most of them)
will be skipped. To run the integration tests please set the following
environment variables

```bash
RUNSCOPE_ACC=true
RUNSCOPE_ACCESS_TOKEN={your access token}
RUNSCOPE_TEAM_UUID={your team uuid}
```
Access tokens can be created using the [applications](https://www.runscope.com/applications)
section of your runscope account.

Your team url can be found by taking the uuid from https://www.runscope.com/teams

## Contributing

1. Fork it ( https://github.com/PagerDuty/go-pagerduty/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Make sure that `make build` passes with test running
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request
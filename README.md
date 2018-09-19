# muxtrace [![Go Report Card](https://goreportcard.com/badge/github.com/anthonyhartanto/muxtrace)](https://goreportcard.com/report/github.com/anthonyhartanto/muxtrace) [![GoDoc](https://godoc.org/github.com/anthonyhartanto/muxtrace?status.svg)](https://godoc.org/github.com/anthonyhartanto/muxtrace)

![Gorilla Logo](http://www.gorillatoolkit.org/static/images/gorilla-icon-64.png)

**What is opentracing ?** Vendor neutral APIs and instrumentation for distributed tracing.

**Why you need muxtrace ?** For example you are using Datadog Tracer to monitor your application performance on http request. We rely on both Datadog and Mux router. Unfortunately, the datadog package for tracing http request on mux was not compatible with opentracing. 

---
* [Getting Started](#getting-started)
* [Installing](#installing)
* [Example](#example)
* [Build With](#built-with)
* [Contributing](#contributing)
* [Versioning](#versioning)
* [Authors](#authors)
---

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing 
purposes.

### Prerequisites

Things you need to use the package

* [Gorilla Mux](https://github.com/gorilla/mux) - Router and dispatcher 
* [Opentracing](https://github.com/opentracing/opentracing-go) - Vendor-neutral APIs and instrumentation for distributed
 tracing


### Installing

With go:

```
$ go get -u github.com/anthonyhartanto/muxtrace
```

## Example

```go
package main

import (
	"fmt"
	"github.com/anthonyhartanto/dd-trace-go/ddtrace/opentracer"
	"github.com/anthonyhartanto/dd-trace-go/ddtrace/tracer"
	"github.com/opentracing/opentracing-go"
	"github.com/anthonyhartanto/muxtrace"
	"net/http"
)

// ExampleHandler handle example
func ExampleHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "OK")
}

func main() {
	// Set tracer
	t := opentracer.New(tracer.WithServiceName("example-service"))
	defer tracer.Stop()
	opentracing.SetGlobalTracer(t)

	// Set router
	router := muxtrace.NewRouter()
	router.HandleFunc("/example", ExampleHandler).Methods(http.MethodGet)

	// Run server
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()
}

```

Result on Datadog:

![Datadog Result](https://raw.githubusercontent.com/anthonyhartanto/muxtrace/master/Screen%20Shot%202018-09-19%20at%2010.33.18.png)

## Built With

* [Go](https://golang.org) - The language used


## Contributing

Please read [CONTRIBUTING.md](https://bitbucket.org/kudoindonesia/driver_signup_service/CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags).

## Authors

* [Anthony Hartanto](anthony.hartanto@kudo.co.id) - Software Engineer

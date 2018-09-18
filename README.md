# muxtrace

![Gorilla Logo](http://www.gorillatoolkit.org/static/images/gorilla-icon-64.png)

A mux wrapper with opentracer

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
    "github.com/anthonyhartanto/muxtrace"
    "net/http"
)

func main() {

    router := muxtrace.NewRouter()
    
    // Run server
    server := &http.Server{
        Addr:    ":8080",
        Handler: router,
    }
    server.ListenAndServe()
}
```


## Built With

* [Go](https://golang.org) - The language used


## Contributing

Please read [CONTRIBUTING.md](https://bitbucket.org/kudoindonesia/driver_signup_service/CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags).

## Authors

* [Anthony Hartanto](anthony.hartanto@kudo.co.id) - Software Engineer

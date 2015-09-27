Gsp
====

Gsp is a compiler built on top of Joseph Adams' [Gisp](https://github.com/jcla1/gisp).
Gsp provides a more complete language environment, including Golang bindings
and the Gsp Prelude.

# Install

```bash
go get github.com/gsp-lang/gsp
cd <GOPATH>/src/github.com/gsp-lang/gsp
go build
echo 'export PATH=$PATH:"${pwd}"' >> ~/.profile
. ~/.bashrc # . ~/.zshrc
```

## Build Prelude

```bash
cd <GOPATH>/src/github.com/gsp-lang/stdlib/prelude
gisp prelude.gsp > prelude.go
```

# Example

```lisp
(ns main
    "/fmt"
    "/net/http")

(def hello (fn [w r]
    (fmt/fprintf w "hello")
    ()))

(def main (fn []
    (http/handle-func "/" hello)
    (http/listen-and-serve ":9090" nil)))
```

## To compile & run

```bash
gsp example.gsp
./bin/main
```

# License

Originally licensed code can be found at [https://github.com/jcla1/gisp](https://github.com/jcla1/gisp).

MIT

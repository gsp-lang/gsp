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
gsp prelude.gsp > prelude.go
```

# Example

```lisp
; Copy into example.gsp
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
gspc example.gsp
./bin/main
```

## Using the Prelude

```lisp
(ns main
	"/fmt")

(def main (fn []
    (let [[m (/cons 2 (/cons 1 /null))]]
        (fmt/println (/car (/cons 1 /null)))
        (fmt/println (/len (/cons 1 /null)))
        (fmt/println (/car (/map (fn [x] (+ x 20)) m)) (/car m))
        (fmt/println (/len m))
        (fmt/println (/car (/cons 17 m)))
        (fmt/println [1 2 3 4])
        (fmt/println (/car (/cdr (/cdr (/append m (/cdr m))))))
        ())))
```

## More examples

Check out the [Prelude](https://github.com/gsp-lang/stdlib/blob/master/prelude/prelude.gsp). It is written in Gsp.

# License

Originally licensed code can be found at [https://github.com/jcla1/gisp](https://github.com/jcla1/gisp).

MIT

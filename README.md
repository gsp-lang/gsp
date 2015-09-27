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

# License

Originally licensed code can be found at [https://github.com/jcla1/gisp](https://github.com/jcla1/gisp).

MIT

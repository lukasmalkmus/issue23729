# issue23729

## What version of Go are you using (`go version`)

go version go1.9.3 darwin/amd64

## Does this issue reproduce with the latest release

Yes.

## What operating system and processor architecture are you using (`go env`)

```text
GOARCH="amd64"
GOBIN=""
GOEXE=""
GOHOSTARCH="amd64"
GOHOSTOS="darwin"
GOOS="darwin"
GOPATH="/Users/lukasmalkmus/Code/Go"
GORACE=""
GOROOT="/usr/local/opt/go/libexec"
GOTOOLDIR="/usr/local/opt/go/libexec/pkg/tool/darwin_amd64"
GCCGO="gccgo"
CC="clang"
GOGCCFLAGS="-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/dy/s6dd70892l35f0444cwt96t40000gn/T/go-build175824991=/tmp/go-build -gno-record-gcc-switches -fno-common"
CXX="clang++"
CGO_ENABLED="1"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
```

## What did you do

I'm having this kind of directory structure:

```text
pkg
├── a
│   ├── a.go
│   └── a_test.go
└── b
    ├── b.go
    └── b_test.go
```

Each `_test.go` file contains a test and a benchmark. I want to benchmark all of
the subpackages in the `pkg` folder using `go test -run=^$ -bench=. ./pkg/...`.

## What did you expect to see

`go test` running the benchmarks of packages `a` and `b`.

## What did you see instead

`go test` running the benchmarks **and** tests of packages `a` and `b`.

## Further details

### Operating the `go test` tool on a single subpackage

**This assumes you cd'ed into subpackage `pkg/a`.**

Running only tests of a single subpackage (`go test -v`):

```text
=== RUN   TestA
--- PASS: TestA (0.00s)
PASS
ok      github.com/lukasmalkmus/issue23729/pkg/a        0.007s
```

---

Running only benchmarks of a single subpackage (`go test -v -bench=. -run='^$'`):

```text
goos: darwin
goarch: amd64
pkg: github.com/lukasmalkmus/issue23729/pkg/a
BenchmarkA-4    2000000000               0.36 ns/op
PASS
ok      github.com/lukasmalkmus/issue23729/pkg/a        0.775s
```

---

Running tests **and** benchmarks of a single subpackage (`go test -v -bench=.`):

```text
=== RUN   TestA
--- PASS: TestA (0.00s)
goos: darwin
goarch: amd64
pkg: github.com/lukasmalkmus/issue23729/pkg/a
BenchmarkA-4    2000000000               0.38 ns/op
PASS
ok      github.com/lukasmalkmus/issue23729/pkg/a        0.812s
```

### Operating the `go test` tool on all subpackages

**This assumes you cd'ed into the `pkg` package.**

Running only tests of all subpackages (`go test -v ./...`):

```text
=== RUN   TestA
--- PASS: TestA (0.00s)
PASS
ok      github.com/lukasmalkmus/issue23729/pkg/a        0.009s
=== RUN   TestB
--- PASS: TestB (0.00s)
PASS
ok      github.com/lukasmalkmus/issue23729/pkg/b        0.009s
```

---

Running only benchmarks of all subpackage (`go test -v -bench=. -run='^$' ./...`):

```text
goos: darwin
goarch: amd64
pkg: github.com/lukasmalkmus/issue23729/pkg/a
BenchmarkA-4    2000000000               0.38 ns/op
PASS
ok      github.com/lukasmalkmus/issue23729/pkg/a        0.816s
goos: darwin
goarch: amd64
pkg: github.com/lukasmalkmus/issue23729/pkg/b
BenchmarkB-4    50000000                30.6 ns/op
PASS
ok      github.com/lukasmalkmus/issue23729/pkg/b        2.515s
```

---

Running tests **and** benchmarks of a single subpackage (`go test -v -bench=. ./...`):

```text
=== RUN   TestA
--- PASS: TestA (0.00s)
goos: darwin
goarch: amd64
pkg: github.com/lukasmalkmus/issue23729/pkg/a
BenchmarkA-4    2000000000               0.37 ns/op
PASS
ok      github.com/lukasmalkmus/issue23729/pkg/a        0.794s
=== RUN   TestB
--- PASS: TestB (0.00s)
goos: darwin
goarch: amd64
pkg: github.com/lukasmalkmus/issue23729/pkg/b
BenchmarkB-4    50000000                30.2 ns/op
PASS
ok      github.com/lukasmalkmus/issue23729/pkg/b        2.551s
```
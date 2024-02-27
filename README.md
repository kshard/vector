# Vector

The library implement fast (highly optimized) vector algebra for Golang.

[![Version](https://img.shields.io/github/v/tag/kshard/vector?label=version)](https://github.com/kshard/vector/releases)
[![Documentation](https://pkg.go.dev/badge/github.com/kshard/vector)](https://pkg.go.dev/github.com/kshard/vector)
[![Build Status](https://github.com/kshard/vector/workflows/build/badge.svg)](https://github.com/kshard/vector/actions/)
[![Git Hub](https://img.shields.io/github/last-commit/kshard/vector.svg)](https://github.com/kshard/vector)
[![Coverage Status](https://coveralls.io/repos/github/kshard/vector/badge.svg?branch=main)](https://coveralls.io/github/kshard/vector?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/kshard/vector)](https://goreportcard.com/report/github.com/kshard/vector)


## Inspiration

Hierarchical Navigable Small World (HNSW) algorithms are state-of-the-art techniques employed for approximate nearest neighbor searches. Vectors and the concept of distance are essential components of the algorithm. The algorithm requires large quantity of vector comparisons on the write path.

The specific count of vector operations performed by HNSW can vary, contingent on factors such as implementation details and the chosen parameters for the algorithm. We have observed a logarithmic progression, starting from 2.2K distance computations per write operation, with subsequent growth reaching 8K and beyond.

**In handling such extensive data, the significance of nanoseconds cannot be overstated!**


## Getting started

The latest version of the library is available at `main` branch of this repository. All development, including new features and bug fixes, take place on the `main` branch using forking and pull requests as described in contribution guidelines. The stable version is available via Golang modules.

```go
import "github.com/kshard/vector"

// Our vectors
a := []float32{60.1699, 24.9384, 24.9384, 60.1699}
b := []float32{59.9311, 30.3609, 30.3609, 59.9311}

// instantiate Euclidean distance function
euclidean := vector.Euclidean()

euclidean.Distance(a, b) // 58.921078

// instantiate Cosine distance function
cosine := vector.Cosine()

cosine.Distance(a, b)  // 0.001443
```

**Note**: One known limitation, the library requires vectors aligned to 4.

### Supported CPU architectures

The library provides two versions of vector algebra functions. The library automatically selects appropriate version depending on the architecture.

1. A Golang native version optimized for instruction pipelining.
2. A pure assembly implementation leveraging the SIMD ("Single Instruction Multiple Data") instruction set.

Internally each version is implemented by the package
1. [internal/noasm](internal/noasm/) - optimized Golang version.
2. [internal/simd](internal/simd/) - pure assembly implementation.
3. [internal/pure](internal/pure/) - reference implementation using idiomatic Golang.

The library provides `Info` function to determine configuration runtime

```go
config = vector.Info()

config[vector.CONFIG_XXX]
// XXX_WITH_PURE
// XXX_WITH_NOASM
// XXX_WITH_SIMD
```

## Performance

Instruction pipeline is 3.3x faster and SIMD is 3.8x faster (due to [bounds check elimination](https://go101.org/article/bounds-check-elimination.html))

```
go test -run=^$ -bench=. -cpu=1

BenchmarkPureEuclideanF32    4072329        270.5 ns/op
BenchmarkNoAsmEuclideanF32  14938154         81.59 ns/op
BenchmarkSIMDEuclideanF32   16888236         71.15 ns/op
BenchmarkPureCosineF32       3987426        299.8 ns/op
BenchmarkNoAsmCosineF32     11738499        102.2 ns/op
```


## How To Contribute

The library is [MIT](LICENSE) licensed and accepts contributions via GitHub pull requests:

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Added some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

The build and testing process requires [Go](https://golang.org) version 1.13 or later.

**build** and **test** library.

```bash
git clone https://github.com/kshard/vector
cd vector
go test
```

Checklist:
1. No new bounds check introduced: `go build -gcflags="-d=ssa/check_bce"`
2. No performance degradations: `go test -run=^$ -bench=. -cpu=1`


### commit message

The commit message helps us to write a good release note, speed-up review process. The message should address two question what changed and why. The project follows the template defined by chapter [Contributing to a Project](http://git-scm.com/book/ch5-2.html) of Git book.

### bugs

If you experience any issues with the library, please let us know via [GitHub issues](https://github.com/kshard/vectors/issue). We appreciate detailed and accurate reports that help us to identity and replicate the issue. 


## License

[![See LICENSE](https://img.shields.io/github/license/kshard/vector.svg?style=for-the-badge)](LICENSE)


## References
1. [A Quick Guide to Go's Assembler](https://go.dev/doc/asm)
2. https://github.com/below/HelloSilicon
3. https://github.com/hztools/go-sdr/tree/main/internal/simd
4. [Bounds Check Elimination](https://go101.org/article/bounds-check-elimination.html)
 
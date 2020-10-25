# go-gauss-legendre

[![Go Report Card](https://goreportcard.com/badge/github.com/svetlana-rezvaya/go-gauss-legendre)](https://goreportcard.com/report/github.com/svetlana-rezvaya/go-gauss-legendre)

The tool for calculating the [pi number](https://en.wikipedia.org/wiki/Pi) via the [Gauss–Legendre algorithm](https://en.wikipedia.org/wiki/Gauss%E2%80%93Legendre_algorithm).

## Installation

```
$ go get github.com/svetlana-rezvaya/go-gauss-legendre
```

## Usage

```
$ go-gauss-legendre -h | -help | --help
$ go-gauss-legendre [options]
```

Options:

- `-h`, `-help`, `--help` &mdash; show the help message and exit;
- `-e FLOAT` &mdash; epsilon (default: 0.000001).

## Output Example

```
pi = 3.141592653589794
number of iterations = 3
```

## License

The MIT License (MIT)

Copyright &copy; 2020 svetlana-rezvaya

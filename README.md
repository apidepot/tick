# tick

Go-based client for [Tick's API][tick-api].

[![GoDoc][godoc badge]][godoc link]
[![Go Report Card][report badge]][report card]
[![License Badge][license badge]][LICENSE.txt]

## Overview

[tick][] provides a Go interface to the [Tick API v2][tick-api]. To
access the [Tick API][tick-api] a subscription ID and API Token are required.

## Installation

```bash
$ go get github.com/apidepot/tick
```

## Contributing

Contributions are welcome! To contribute please:

1. Fork the repository
2. Create a feature branch
3. Code
4. Submit a [pull request][]

### Testing

Prior to submitting a [pull request][], please run:

```bash
$ make check    # formats, vets, and unit tests the code
$ make lint     # lints code using staticcheck
```

To update and view the test coverage report:

```bash
$ make cover
```

#### Integration Testing

To perform the integration tests run:

```bash
$ make int
```

Prior to doing so, you'll need to create a `config_test.toml` file with your
Tick Subscription ID and API Token.

Example `config_test.toml` file:

```toml
SubscriptionID = "xxxxx"
Token = "big-hex-string"
```


## License

[tick][] is released under the MIT license.  Please see the
[LICENSE.txt][] file for more information.

[godoc badge]: https://godoc.org/github.com/apidepot/tick?status.svg
[godoc link]: https://godoc.org/github.com/apidepot/tick
[tick]: https://github.com/apidepot/tick
[LICENSE.txt]: https://github.com/apidepot/tick/blob/master/LICENSE.txt
[license badge]: https://img.shields.io/badge/license-MIT-blue.svg
[pull request]: https://help.github.com/articles/using-pull-requests
[report badge]: https://goreportcard.com/badge/github.com/apidepot/tick
[report card]: https://goreportcard.com/report/github.com/apidepot/tick
[tick-api]: https://github.com/tick/tick-api

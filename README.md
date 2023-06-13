# Date translator for Go

[![Latest release](https://img.shields.io/github/v/release/hansmi/zyt)][releases]
[![CI workflow](https://github.com/hansmi/zyt/actions/workflows/ci.yaml/badge.svg)](https://github.com/hansmi/zyt/actions/workflows/ci.yaml)
[![Go reference](https://pkg.go.dev/badge/github.com/hansmi/zyt.svg)](https://pkg.go.dev/github.com/hansmi/zyt)

The zyt[^name-explanation] package translates month and day names from Go's
[time](https://pkg.go.dev/time) package to and from non-English languages.

When parsing multiple name variants are understood. For example the month
January in German is recognized as
[`Januar`](https://de.wikipedia.org/wiki/Januar), `Jänner` and multiple
abbreviations.

Languages using partitive and genitive cases for months and days are supported
(e.g. Finnish and many Slavic languages).

Locales can be customized by users, either by defining from scratch or by
cloning and modifying.

[^name-explanation]: _Zyt_ is Swiss German for _time_.


## Usage

To print the current time in German:

```go
fmt.Println(zyt.German.Format(time.RFC850, time.Now()))
```

To parse a timestamp in Finnish:

```go
fmt.Println(zyt.Finnish.ParseInLocation("January 2006", "maaliskuu 2001", time.UTC))
```

To select the best-fitting locale:

```go
l, _ := zyt.Best(language.BritishEnglish)
l.Format(time.RFC850, time.Now())
```


## Supported locales

* English
* German
  * Austrian German
* Finnish

Contributions to add more are very welcome.


## Alternatives

* [github.com/goodsign/monday](https://pkg.go.dev/github.com/goodsign/monday),
  an older package only supporting one variant per name.

[releases]: https://github.com/hansmi/zyt/releases/latest

<!-- vim: set sw=2 sts=2 et : -->

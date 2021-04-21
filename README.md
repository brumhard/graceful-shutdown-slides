# Graceful Shutdown presentation

This repos contains the slides for a presentation on graceful shutdown using the context package.
It uses the present package and tool:
- https://pkg.go.dev/golang.org/x/tools/present
- https://golang.org/x/tools/cmd/present

The code is heavily inspired by the blog article by Mat Ryer [Respond to Ctrl+C interrupt signals gracefully](https://pace.dev/blog/2020/02/17/repond-to-ctrl-c-interrupt-signals-gracefully-with-context-in-golang-by-mat-ryer.html)

## Prerequisites

```shell
go get -u golang.org/x/tools/...
```

## Run

```shell
present -notes
```
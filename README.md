# Platoon Go

This is a rewrite of the original Platoon library as a standalone application; rewritten entirely in Go.

## Why?

Platoon is currently developed in PHP and is very Laravel specific. The attempt to rewrite it serves numerous purposes:

- I'm using this project to learn more about Go
- Leveraging Go Routines will help to make Platoon even faster
- Rewriting in Go removed the dependency on Laravel and Ecnvoy
- There's no PHP and Composer requirement meaning Platoon could be used for any project using any technology.

## State

This is a very very early-stage project. I want to rebuild all the existing features into this version and I'm still learning alot about Go. This should not be used at all. Please wait for a proper release version.

## Build

To build, run:

```shell
go build -o bin/platoon main.go
```

## License

Platoon-Go, like the original, is licensed under MIT. You can find more in the [LICENCE.md]() file.

## Copyright

Copyright (c) 2026 Warrick Bayman
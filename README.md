[![Actions Status](https://github.com/guikcd/headerscheck/workflows/Go/badge.svg)](https://github.com/guikcd/headerscheck/actions)

[![docker stars](https://img.shields.io/docker/stars/guidelacour/headerscheck.svg)](https://hub.docker.com/r/guidelacour/headerscheck/) [![docker pulls](https://img.shields.io/docker/pulls/guidelacour/headerscheck.svg)](https://hub.docker.com/r/guidelacour/headerscheck/) [![docker automated build](https://img.shields.io/docker/automated/guidelacour/headerscheck.svg)](https://hub.docker.com/r/guidelacour/headerscheck/) [![docker build status](https://img.shields.io/docker/build/guidelacour/headerscheck.svg)](https://hub.docker.com/r/guidelacour/headerscheck/)
[![layers](https://images.microbadger.com/badges/image/guidelacour/headerscheck.svg)](https://microbadger.com/images/guidelacour/headerscheck "Get your own image badge on microbadger.com") [![version](https://images.microbadger.com/badges/version/guidelacour/headerscheck.svg)](https://microbadger.com/images/guidelacour/headerscheck "Get your own version badge on microbadger.com")

[![Go Report Card](https://goreportcard.com/badge/github.com/guikcd/headerscheck)](https://goreportcard.com/report/github.com/guikcd/headerscheck) [![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)



Introduction
==========

This is a simple tool in Go to make http requests and except reponse codes, headers or body.
Please note that this tool is in fact my first Go program i've written only for discovering this language :)

Features
=======

* read config from yaml file
* check http status codes
* check multiples URLs (one at a time)
* check if headers/values exists/not exists
* check if body match or not a value
* support regular expression in all values fields
* don't follow redirects by default (to test them)
* override default Go http client User-Agent and timeout
* debug mode

Help
====

```
#!/bin/bash
$ ./headerscheck -help
Usage of ./headerscheck:
  -config-file string
        Config file (default "config")
  -debug
        Enable debugging output
  -follow-redirect
        Follow redirect (http status codes 30X)
  -user-agent string
        User-Agent used for queries (default "Golang HeadersCheck/0.1 (gui@iroqwa.org)")
```

How to use
=========

```
#!/bin/bash
$ docker run --rm --volume $(pwd)/config.yml:/root/config.yml guidelacour/headerscheck
```

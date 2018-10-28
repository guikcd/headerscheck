[![Go Report Card](https://goreportcard.com/badge/github.com/guikcd/headerscheck)](https://goreportcard.com/report/github.com/guikcd/headerscheck)



Introduction
==========

This is a simple tool in Go to make http requests and except reponse codes, headers or body.
Please note that this tool is in fact my first Go program i've written only for discovering this language.

Features
=======

* read config from yaml file
* check http status codes
* check multiples urls (one at a time)
* check if headers/values exists/not exists
* check if body match a value
* support regex in all values fields (TODO)
* don't follow redirects by default (to test them)
* override default Go http client User-Agent and timeout
* debug mode

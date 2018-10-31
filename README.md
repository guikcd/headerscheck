[![Go Report Card](https://goreportcard.com/badge/github.com/guikcd/headerscheck)](https://goreportcard.com/report/github.com/guikcd/headerscheck) [![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)



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

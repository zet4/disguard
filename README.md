# DisGuard - Discord OAuth Proxy [![GoDoc](https://godoc.org/go.zeta.pm/disguard?status.svg)](https://godoc.org/go.zeta.pm/disguard) [![Build Status](https://travis-ci.org/zet4/disguard.svg?branch=master)](https://travis-ci.org/zet4/disguard) [![Go Report Card](https://goreportcard.com/badge/go.zeta.pm/disguard.svg)](https://goreportcard.com/report/go.zeta.pm/disguard) [![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/zet4/disguard/blob/master/LICENSE)

## Why?

Because we want content locked to discord guild memberships.

## How?

You make a file named `config.yaml` with contents similar to what you can find in `config.example.yaml` in this repository.
Once done, launch the application and go to the website, if you configured it to require session, you will be sent to oauth login right away, if not, you will have to manually visit `/oauth/login` to get redirected.
Once logged in, each request the user makes will be proxied to your defined backend, with attached named header, in example thats `XWhitelistedGuilds`.
From there on its your job to handle that.

Session is valid for one week.

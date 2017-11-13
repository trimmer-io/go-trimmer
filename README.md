Trimmer SDK for Go
====================

go-trimmer is the official Trimmer SDK for the Go programming language.

Checkout our releases and the Changelog for information about the latest bug fixes, updates, and features added to the SDK.

## Installing

Install the SDK with the following `go get` command.

```
$ go get trimmer.io/go-trimmer
```

or integrate it as Git submodule into your `vendor/` directory

```
$ git submodule add https://trimmer.io/go-trimmer vendor/trimmer.io/go-trimmer
```

## Environment Variables

```
# API Key (required)
TRIMMER_API_KEY

# Servers (optionally overwrite configured servers)
TRIMMER_API_SERVER
TRIMMER_CDN_SERVER

# default client token authentication
TRIMMER_CLIENT_TOKEN

# default user login options, use: session.ParseEnv()
TRIMMER_API_USERNAME
TRIMMER_API_PASSWORD
```

## Using the Go SDK

```
package main

import (
	"log"
    "trimmer.io/go-trimmer"
    "trimmer.io/go-trimmer/user"
    "trimmer.io/go-trimmer/session"
)

func main() {

// Direct calls use a default http client configuration and a global login
// session. You may create long-lived API client objects and pass them
// custom configurations if needed and use custom login sessions.
//
// For convenience and security you can parse the process environment for
// exported TRIMMER_API_* variables
//
//  TRIMMER_API_USERNAME
//  TRIMMER_API_PASSWORD
//

// get login data from ENV and store session in global LoginSession struct
if err := session.Login(session.ParseEnv()); err != nil {
	log.Fatalln(err)
}
defer session.Logout()

// get the currently logged in user (will read the global LoginSession)
me, err := user.Me()
if err != nil {
	log.Fatalln(err)
}

```

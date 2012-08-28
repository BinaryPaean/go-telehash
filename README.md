go-telehash
===========

### Version 0.0

Implementation of [TeleHash](http://telehash.org/) in go.

As "switch" is a reserved word in go, TeleHash switches are called "exchanges" in this implementation.

Currently, I'm considering what types will be most useful for the Telex definition. Go will parse JSON
into pre-defined structs nicely, but as each Telex has many optional fields as well as free-form message body,
a lot of stuff may get dumped into untyped arrays. See this work in progress in telex.go and telex_test.go.

One test is currently written and passing, run "go test" in the telex directory to see what is going on.
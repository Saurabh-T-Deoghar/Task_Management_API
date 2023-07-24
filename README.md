# Task_Management_API
Task Management API using Go language with SQLite database


# GoLang
Go is a procedural programming language. Go is a statically typed, concurrent, and garbage-collected programming language created at Google in 2009. It is designed to be simple, efficient, and easy to learn, making it a popular choice for building scalable network services, web applications, and command-line tools.

## Go in a Nutshell
* Imperative language
* Statically typed
* Syntax tokens similar to C (but less parentheses and no semicolons) and the structure to Oberon-2
* Compiles to native code (no JVM)
* No classes, but structs with methods
* Interfaces
* No implementation inheritance. There's [type embedding](http://golang.org/doc/effective%5Fgo.html#embedding), though.
* Functions are first class citizens
* Functions can return multiple values
* Has closures
* Pointers, but not pointer arithmetic
* Built-in concurrency primitives: Goroutines and Channels

## Packages
* Package declaration at top of every source file
* Executables are in package `main`
* Convention: package name == last name of import path (import path `math/rand` => package `rand`)
* Upper case identifier: exported (visible from other packages)
* Lower case identifier: private (not visible from other packages)


# REST API
- [ ] Use the HTTP verbs to mean this:

    * POST - create and all other non-idempotent operations.
    * PUT - replace.
    * PATCH - (partial) update.
    * GET - read a resource or collection.
    * DELETE - remove a resource or collection.

- [ ] Use [HTTP status codes](https://httpstatuses.com/) to meaning this:
  * 102 - Processing. Returned as long as a asynchronous response is still pending. See also 202 Accepted.
  * 200 - Success.
  * 201 - Created. Returned on successful creation of a new resource. Include a 'Location' header with a link to the newly-created resource.
  * 202 - Accepted. Returned to indicated an asynchronous response will be given. Include a 'Location' HTTP-header with URL of the future resource. See also 102 Processing.
  * 204 - No content (for empty responses)
  * 304 - Not modified. Returned when a client asks for an etag it already received (sent in the request headers). See _caching_ below.
  * 400 - Bad request. Data issues such as invalid JSON, etc.
  * 404 - Not found. Resource not found on GET.
  * 409 - Conflict. Duplicate data or invalid data state would occur.

# SQLite
SQLite is a C-language library that implements a small, fast, self-contained, high-reliability, full-featured, SQL database engine. SQLite is the most used database engine in the world. SQLite is built into all mobile phones and most computers and comes bundled inside countless other applications that people use every day.

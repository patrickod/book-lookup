# book-lookup

This is the second part of the book ingestion pipeline for the Library at [Noisebridge](https://noisebridge.net).

## What?

Reads the ISBN numbers from STDIN, looks them up using the OpenLibrary API, and then records all their information in a PostgreSQL database which is used to feed a publicly searchable index of books in the space.

## Why?

Noisebridge has "hella" books but no information about what we have. I'm solving this.

## Can I help?

Yes please!!! At time of writing the TODO list is as follows:

  * Record subject information similar to authorships
  * Deal with duplicate books / authors / subjects gracefully.
  * Read the `DATABASE_URL` environment variable instead of compiling in the DB information
  * Better logging, not to stdout
  * Tests?? Would at least like to write idiomatic Golang tests

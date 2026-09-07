# SAT Word List Data Manager

> **Note:** This README was AI-generated. All source code in this repository is handwritten.

A small desktop GUI application, built with [Fyne](https://fyne.io/), for managing a personal SAT vocabulary list. It lets you add new words, look up meanings, and persist the list to disk.

## Features

- **New Word** — add a word with its type (noun/verb/adjective) and a description/meaning.
- **Check Meaning** — case-insensitive lookup of a word's type and description.
- **Save** — append pending (unsaved) words to the data file on disk so they persist across runs.
- **Quit** — close the application.

## Data storage

Words are stored in a plain-text file (`words.dat`) using the format:

```
word type description
```

Each line is split on whitespace into at most three fields. On startup, this file is parsed into an in-memory table (for display) and a hashmap (for fast case-insensitive lookup).

## Project layout

```
cmd/        entry point (main.go)
internal/   non-UI logic: word parsing, saving, state, errors
ui/         Fyne UI components (buttons, windows, table)
```

## Requirements

- Go 1.27+
- Fyne v2 (see `go.mod`) and its native build dependencies (a C toolchain and OpenGL/graphics libraries for your OS — see the [Fyne getting started guide](https://docs.fyne.io/started/)).

## Running

```sh
go run ./cmd
```

## Building

```sh
go build -o sat_word_list ./cmd
```

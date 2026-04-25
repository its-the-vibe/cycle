# cycle

[![CI](https://github.com/its-the-vibe/cycle/actions/workflows/ci.yaml/badge.svg)](https://github.com/its-the-vibe/cycle/actions/workflows/ci.yaml)


Simple command-line tool to move the first line of a file to the end.

## Installation

```bash
go install github.com/its-the-vibe/cycle@latest
```

Or build from source:

```bash
git clone https://github.com/its-the-vibe/cycle.git
cd cycle
make build
```

This produces a `cycle` binary in the current directory.

## Usage

```
cycle <file>
```

Rotates `<file>` in-place by moving its first line to the end.

### Examples

Given a file `queue.txt`:

```
task-a
task-b
task-c
```

Running `cycle queue.txt` produces:

```
task-b
task-c
task-a
```

### Edge cases

| Scenario | Behaviour |
|---|---|
| Empty file | File is left unchanged |
| Single-line file | File is left unchanged |
| Non-existent file | Error message printed to stderr; exits with code 1 |
| No arguments | Usage message printed to stderr; exits with code 1 |

## Development

```bash
make build   # compile the binary
make test    # run tests
make lint    # run go vet
make clean   # remove the binary
```

## License

MIT

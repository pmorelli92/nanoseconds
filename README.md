# ns - Nanosecond Time Converter

A simple, dependency-free command-line tool to quickly convert between nanosecond timestamps and human-readable dates using your system's local time.

## Installation

Ensure you have Go installed and run:

```bash
go install github.com/pmorelli92/nanoseconds/cmd/ns@main
```

*Note: This will place the compiled ns binary into your $GOPATH/bin directory (usually ~/go/bin). Make sure this directory is added to your system's $PATH so you can run the command from anywhere.*

## Usage & Examples

The application handles conversions in both directions.

1. Convert nanoseconds to a human-readable date:

Use the -h or --human flag followed by the nanosecond timestamp.

```
$ ns -h 1774615500000000000
27/03/2026 13:45:00
```

2. Convert a human-readable date to nanoseconds:

Use the -n or --nano flag followed by the date in DD/MM/YYYY HH:MM:SS format.

```
$ ns -n 27/03/2026 13:45:00
1774615500000000000
```

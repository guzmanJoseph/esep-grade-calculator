# ESEP Grade Calculator

A Go library that calculates a student's final course grade from assignment, exam, and essay scores. Each category is averaged separately, then weighted at 50%, 35%, and 15%. The weighted numerical result is converted to a letter grade.

## Requirements

- Go 1.24.4 or newer

## Run the tests

From the repository root:

```sh
cd grade-calculator
go test ./...
```

## Measure test coverage

From the `grade-calculator` module directory:

```sh
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
```

The final line of the coverage report shows the total statement coverage.

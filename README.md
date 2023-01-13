# Everything related to Go Testing

If you would like to know EVERYTHING about functions in Golang, the you are at the right place.

```bash
go help test # Gte help with test
go help testflag # More info about testing in Go!
```
---
## Functional tests

link to the video: https://youtu.be/dRyEj7aFE6Y

List of commands:
```bash
go test ./... # Test all and every tests
go test # Run all tests in current directory
go test -v # Run test in verbose mode
go test ./... -run Sum # Run a specific test with regex
go test github.com/mohokh67/go-testing/magical # Run tests for only a package
go test ./... -cover # Get coverage for tests
go test ./... -coverprofile coverage.out # Get coverage for tests with percentage
go tool cover -func coverage.out # Analyse the coverage from above in terminal
go tool cover -html coverage.out # Analyse the coverage from above in browser
go test ./... -coverprofile coverage-count.out -covermode count # Get coverage for tests with count
go tool cover -html coverage-count.out # Analyse the coverage from above in browser
```

-------
## benchmark tests
link to the video: https://youtu.be/0UBv8GV8XdY

List of commands:
```bash
go test -bench . # Run all benchmark tests in the current directory
go test -bench . -benchtime 10s # Run all benchmark tests in the current directory for 10 seconds
go test -bench 256 # Run all benchmark tests match 256 in the test name(regex)
go test -bench . -benchmem # Run all benchmark tests in the current directory with memory information
go test -bench . -memprofile mem.out # Run all benchmark tests in the current and profile them based on memory
go tool pprof mem.out # Checkout performance profile from previous command - type "help" and then "gif" or other format
go test -bench . -cpuprofile cpu.out  # Run all benchmark tests in the current and profile them based on CPU
```

-------
## Fuzz tests
link to the video: https://youtu.be/YiqCjEAeAyw

List of commands:
```bash
go test -fuzz . # run current fuzz test
go test -fuzz FuzzStringChecker # run a specific fuzz test
go test -fuzz . -fuzztime 10s # fuzzing for 10 seconds
go test -fuzz . -fuzztime 1000x # fuzzing for 1000 times
```
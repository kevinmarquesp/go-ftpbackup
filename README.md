# Go FTP Backup

Program that I created to pull some files from my phone to my PC as a backup
strategy using the FTP protocol.

> [!WARNING]
> This project is under development, this document will change over time.
> Contributions are welcome! :sparkles:


## Running

Clone this repository and then just run `go run . -help` to see the command
line options. If you miss some mandatory flag, you should expect an error
message.

Start the FTP server then connect to it with this app. The connection process
will show an error if fail, but it will try again until it successfuly
connects to the server.

Once connected, it will just list all the files in the root filesystem of the
FTP server - you can put anything as the `-target` and `-src`'s values, the
backup feature is not implemented yet.

Usage example:

```bash
go run . \
    -machine localhost \
    -port 2000 \
    -username joedoe \
    -password 12345678 \
    -src /unused/src/path \
    -target /unused/target/path
```


## Tests

Each package file (with some exceptions for types, structs, etc.) have a
`*_test.go` variant in the same package - the convention here is to put the
test files along side with the source file for each package.

Run them all with:

```bash
go test $(go list ./...)
```

Or, run the tests of a single package with, as an example:

```bash
go test ./args
```


## See Also

*   FTP app that I'm using: [Material Files](https://github.com/zhanghai/MaterialFiles)

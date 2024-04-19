# About
A straightforward command-line interface (CLI) for generating random passwords, crafted in Go with the cobra library.

# How to run

To run the CLI, use the following command:

```
$ go run main.go generate
```

You can specify the length of the password by using the flag `--length` or `-l` after the generate command.

```
$ go run main.go generate --length 20
```

To display the help menu, utilize the flag `--help` or `-h`:

```
$ go run main.go --help
```

# Reference(s)

[Let's build a CLI in Go with Cobra](https://www.thorsten-hans.com/lets-build-a-cli-in-go-with-cobra/)
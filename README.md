# jarexe

convert jar file into executable file.

## How to run

```
go get github.com/katsuyan/cmd/jarexe
jarexe ~.jar
```

## Requirements

- Go 1.8
- [dep](https://github.com/golang/dep)

## Development
### First
```
dep ensure
```

### Run
```
go run main.go
```

### Build
```
go build
```

### Check
```
go run main.go example_jar/standalone.jar
```

or

```
jarexe example_jar/standalone.jar
```

and

```
./standalone abc def
```

and result

```
(abc def)
```

## LICENSE

MIT

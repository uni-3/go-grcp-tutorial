### how to use

0. install packages

```sh
go get google.golang.org/grpc
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

1. `sh ./gen-protoc.sh`

generate stub files in `./proto`

2. run server and client

```sh
go run server.go

...
2022/10/31 14:06:26 serving... :5001
```

```sh
# in another terminal
go run client.go
2022/10/31 14:07:30 connect
2022/10/31 14:07:30 User: Hello, yeah!
```

- refer
  https://www.tohoho-web.com/ex/grpc.html

```

```

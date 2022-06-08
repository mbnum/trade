# trade
Quick start [Go-gRPC](https://grpc.io/docs/languages/go/quickstart/), [private-go-module](https://www.digitalocean.com/community/tutorials/how-to-use-a-private-go-module-in-your-own-project)

## Recompile the updated .proto file:
```console
$ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative *.proto */*.proto
```

## Import
```console
$ export GOPRIVATE=github.com/mbnum/trade,$GOPRIVATE
```

#### **`~/.gitconfig`**
```
[url "ssh://git@github.com/"]
	insteadOf = https://github.com/
```

####

* brew install protobuf
* go get google.golang.org/grpc
* make create

#### Now using, gRPC Gateway

* https://github.com/grpc-ecosystem/grpc-gateway
* go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest 
* go install google.golang.org/protobuf/cmd/protoc-gen-go@latest 
* go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

create tools.go
```
package tools

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
```

install
```
go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

mkdir -p google/api in the root of your project, then:
```

`curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto > proto/google/api/annotations.proto`
`curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto > proto/google/api/http.proto`
```

brew install bufbuild/buf/buf

create # Modified gRPC enabling REST API


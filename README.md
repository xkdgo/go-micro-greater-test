# go-micro Greeter Example

An example Greeter application

## Contents

- **srv** - an RPC greeter service
- **cli** - an RPC client that calls the service once
- **api** - examples of RPC API using GinGonic

## Etcd Registry

install etcd

https://etcd.io/docs/v3.5/quickstart/

Start ETCD registry daemon

```shell
export IP_ADDR="x.x.x.x" # where x.x.x.x is inet address of your server
export ETCD_BIN="/path/to/your/etcd/binary"

$ETCD_BIN/etcd \
--name infra0 \
--initial-advertise-peer-urls http://$IP_ADDR:2380 \
--listen-peer-urls http://$IP_ADDR:2380 \
--listen-client-urls http://$IP_ADDR:2379,http://127.0.0.1:2379 \
--advertise-client-urls http://$IP_ADDR:2379 \
--initial-cluster-token etcd-cluster-1 \
--initial-cluster infra0=http://$IP_ADDR:2380 \
--initial-cluster-state new

```

## Run Service

Start go.micro.srv.greeter
```shell
export $SRV_PORT=XXXX # where XXXX youre desirable server port

go run srv/main.go \
-registry=etcd --registry_address=https://$IP_ADDR:2379 \
--server_address=127.0.0.1:$SRV_PORT
```

## Client

Call go.micro.srv.greeter via client
```shell
go run cli/main.go
```

Examples of client usage via other languages can be found in the client directory.

## API

HTTP based requests can be made via GinGonic router

Run the API Service
```shell
export IP_ADDR="x.x.x.x" # where x.x.x.x is inet address of your server
export $API_PORT=YYYY # where YYYY youre desirable server port

go run api/main.go \
--registry=etcd --registry_address=https://$IP_ADDR:2379 \
--server_address=$IP_ADDR:$API_PORT --server_name=api
```

Call Say.Hello
```shell
export IP_ADDR="x.x.x.x" # where x.x.x.x is inet address of your server
export $API_PORT=YYYY # where YYYY youre desirable server port

curl --header "Content-Type: application/json"  \
-H "Accept: application/json" \
--request POST   \
--data '{"name":"Someone"}' \
http://$IP_ADDR:$API_PORT/go.micro.srv.greeter/Say.Hello

```

Call Say.Goodbye
```shell
export IP_ADDR="x.x.x.x" # where x.x.x.x is inet address of your server
export $API_PORT=YYYY # where YYYY youre desirable server port

curl --header "Content-Type: application/json"  \
-H "Accept: application/json" \
--request POST   \
--data '{"name":"Someone"}' \
http://$IP_ADDR:$API_PORT/go.micro.srv.greeter/Say.Goodbye

```

Show service nodes by name:
```shell
export IP_ADDR="x.x.x.x" # where x.x.x.x is inet address of your server
export $API_PORT=YYYY # where YYYY youre desirable server port

 curl "http://$IP_ADDR:$API_PORT/go.micro.srv.greeter/nodes"
 curl "http://$IP_ADDR:$API_PORT/api/nodes"

```


## EXAMPLES

```shell
$ go run srv/main.go \
-registry=etcd --registry_address=https://$IP_ADDR:2379 \
--transport=nats --transport_address=$IP_ADDR:4222 \
--server_address=127.0.0.1:$SRV_PORT
2021-09-08 14:02:23  file=v3@v3.6.1-0.20210831143116-05a299b76c7c/service.go:206 level=info Starting [service] go.micro.srv.greeter
2021-09-08 14:02:23  file=v3@v3.0.0-20210907061356-440aa4a1ce13/grpc.go:897 level=info Server [grpc] Listening on 127.0.0.1:4040
2021-09-08 14:02:23  file=v3@v3.0.0-20210907061356-440aa4a1ce13/grpc.go:728 level=info Registry [etcd] Registering node: go.micro.srv.greeter-466c35da-28d7-4ea3-8dd4-08d1d46c5698
2021-09-08 14:21:38  file=srv/main.go:22 level=info [Received Say.Hello request]
timeout 1
result 2
2021-09-08 14:21:41  file=srv/main.go:52 level=info [context.Background.WithDeadline(2021-09-08 14:21:58.457270543 +0300 MSK m=+1174.685955639 [16.997854672s]).WithValue(type peer.peerKey, val <not Stringer>).WithValue(type metadata.mdIncomingKey, val <not Stringer>).WithValue(type grpc.streamKey, val <not Stringer>).WithValue(type metadata.metadataKey, val <not Stringer>).WithValue(type peer.peerKey, val <not Stringer>).WithCancel.WithValue(type metadata.metadataKey, val <not Stringer>)]
2021-09-08 14:21:41  file=srv/main.go:53 level=info [name:"Someone"]
2021-09-08 14:21:41  file=srv/main.go:55 level=info [Message you want:  Hello Someone
timeout 1...
result 2...
]


```

```shell
$go run api/main.go --registry=etcd --registry_address=https://192.168.0.100:2379 --server_address=192.168.0.100:8080 --server_name=api
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /:service/:endpoint       --> main.main.func1 (3 handlers)
[GIN-debug] GET    /:service/nodes           --> main.main.func2 (3 handlers)
2021-09-08 14:06:24  file=v3@v3.6.1-0.20210831143116-05a299b76c7c/service.go:206 level=info Starting [service] api
2021-09-08 14:06:24  file=v3@v3.0.0-20210824071433-49eccbc85a0f/http.go:255 level=info Listening on 192.168.0.100:8080
2021-09-08 14:06:24  file=v3@v3.0.0-20210824071433-49eccbc85a0f/http.go:169 level=info Registering node: api-88a4582a-a9c0-4124-a25d-f76bd945f549
[GIN] 2021/09/08 - 14:07:30 | 200 |     907.151µs |   192.168.0.100 | GET      "/go.micro.srv.greeter/nodes"
[GIN] 2021/09/08 - 14:21:41 | 200 |  3.004554653s |   192.168.0.100 | POST     "/go.micro.srv.greeter/Say.Hello"
```

curl POST
```shell
$ curl --header "Content-Type: application/json"  \
-H "Accept: application/json" \
--request POST   \
--data '{"name":"Someone"}' \
http://$IP_ADDR:8080/go.micro.srv.greeter/Say.Hello

{"msg":"Hello Someone\ntimeout 1...\nresult 2...\n"}
```


curl GET
```shell
curl "http://192.168.0.100:8080/go.micro.srv.greeter/nodes"
[{"id":"go.micro.srv.greeter-466c35da-28d7-4ea3-8dd4-08d1d46c5698","address":"127.0.0.1:4040","metadata":{"broker":"http","protocol":"grpc","registry":"etcd","server":"grpc","transport":"grpc"}}]
```

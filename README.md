# Requirements:

- go get -u -v  google.golang.org/grpc
- Protocol Buffers 3 (https://github.com/google/protobuf/releases)
- go get -u - v github.com/golang/protobuf/protoc-gen-go
- go get -u -v github.com/dannav/hhmmss
- go get -u -v github.com/go-sql-driver/mysql
- go get -u -v github.com/jmoiron/sqlx


# Directory Structure:

```
+--go
|	+--bin
|	+--pkg
|	+--src
|	|	+---github.com
|	|	|		+----activity-tracker
|	|	|		|		+-----client
|	|	|		|		|	+------client.go
|	|	|		|		+-----db
|	|	|		|		|	+------db.go
|	|	|		|		+-----pb
|	|	|		|		|	+------tracker.proto
|	|	|		|		|	+------tracker.pb.go
|	|	|		|		+-----server
|	|	|		|		|	+------server.go
|	|	|		|		+-----third_party
|	|	|		|		|	+------google
|	|	|		|		|	+------protoc-gen-openapiv2
|       |       |               |               |
```
# Commands :

> To generate .go file from .proto:

protoc --proto_path=$GOPATH/src/github.com/v_task/pb/ --proto_path=$GOPATH/src/github.com/v_task/third_party --go_out=plugins=grpc:. tracker.proto


> To run server on one terminal

go run server.go


> To run client on another terminal

go run client.go




# Explanation:

> "tracker.proto" contains all the services, rpc methods and messages.

> "tracker.pb.go" implements all the server and client side stubs in golang.

> "db.go" describes the connection of the server with the MySQL database.

> "server.go" declares and runs a gRPC server as well as handles all the server side API calls.

> "client.go" connects to that gRPC server and requests server-side API's with static data. (It can be made dynamic also, as per the requirements.)

>"activity_tracker_db.sql" is the .sql file which can be imported into a MySQL database and used as is. 


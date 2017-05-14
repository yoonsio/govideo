# GoVideo

Fast open-source video streaming server written in Go

## Features

* authentication
* access control
* speed limiter

## Sync / Indexing

Media files are recursively added from target directories without category.<br>
Categories can be added in the future via web interface.<br>
Subtitle files must have same name as media files.

## Run Instruction

### Run Web Server
```
go run main.go run
```

### Seed Database
```
go run main.go seed
```

## Compilation

### libmagic

`libmagic` is required for mime detection.

```
sudo apt-get install libmagic // linux
brew install libmagic // osx
```

### Generate protobuf

For `User`, add `bson: "id"` property to `Email` field for MongoDB primary key.

```
cd ./govideo/models
go get github.com/gogo/protobuf/protoc-gen-gogoslick
protoc --proto_path=$GOPATH/src:. --gogoslick_out=. user.proto
```

#### Generate easyjson

```
cd ./govideo/models
easyjson response.go
easyjson media.go
```

### Webpack

```
webpack --progress --color
```

## TODO

### Security: User Session Expiry based on activity
User session in redis can be expired based on activity.
Set redis expiry to low value and update it whenever a request is made by the same user.

### Security: Media should not return real path

New struct that wraps around Media struct should be returned instead.
Sensitive fields such as access and real path should not be visible.

## Built with <3 from

* valyala/httprouter
* react

## License

MIT


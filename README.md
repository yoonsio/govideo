# GoVideo

Fast open-source video streaming server written in Go

## Features

* authentication
* access control
* speed limiter

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

## Resources

* https://github.com/gaearon/react-hot-boilerplate
* https://github.com/zackperdue/React-Video-Player
* https://github.com/mderrick/react-html5video
* https://github.com/mderrick/react-html5video
* https://github.com/enaqx/awesome-react#generating
* https://github.com/dgryski/gophervids
* https://github.com/nareix/joy4
* https://github.com/zackperdue/React-Video-Player
* https://github.com/olebedev/go-starter-kit/blob/master/package.json
* https://github.com/pedronauck/react-video
* https://github.com/brillout/awesome-react-components
* https://github.com/CookPete/react-player
* https://github.com/enaqx/awesome-react#libraries
* http://andrewhfarmer.com/starter-project/
* http://jamesknelson.com/using-es6-in-the-browser-with-babel-6-and-webpack/
* https://github.com/react-bootstrap/react-router-bootstrap

## Built with <3 from

* valyala/httprouter
* react

## License

MIT


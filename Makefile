arch ?= "amd64"
os ?= "linux"
cgo_enable ?= "0"

build:
	export CGO_ENABLED=$(cgo_enable) && export GOOS=$(os) && export GOARCH=$(arch) && go build -o ./bin/saber ./main.go
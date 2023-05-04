# BagShareMind

BagShare Backend service

## Versions

Go Lang : v1.20.4

## Installation of packages

- Installing gin
    `$ go get -u github.com/gin-gonic/gin`

## Troubleshooting

Use the below command to switch the Go111 module on which is used for go mod init

`go env -w GO111MODULE=on`

## Build

- Docker file is present to make a docker build. Find below the sample command to build the docker image and run it locally

```
// for build
docker build -t bagshare:v0.1 .

// for running
docker run -p 5000:5000 bagshare:v0.1
```

- Use health check API http://localhost:5000/health to check the server (GET call can be accessed using broswer)

## Test

Note : For now testing is only done using REST API calls

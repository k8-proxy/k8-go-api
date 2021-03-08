<h1 align="center">k8-go-api</h1>

<p align="center">
  <a href="https://goreportcard.com/report/github.com/k8-proxy/k8-go-api">
      <img src="https://goreportcard.com/badge/k8-proxy/k8-go-api" alt="Go Report Card">
  </a>
	<a href="https://github.com/k8-proxy/k8-go-api/pulls">
        <img src="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat" alt="Contributions welcome">
    </a>
    <a href="https://opensource.org/licenses/Apache-2.0">
        <img src="https://img.shields.io/badge/License-Apache%202.0-blue.svg" alt="Apache License, Version 2.0">
    </a>
</p>

# k8-go-api

## Build

Clone the repo then

```
cd k8-go-api
go build -o server ./cmd
./server
```

The server will start at:

- Local: http://localhost:8100

### Docker build

```
docker build -t k8-go-api .
docker run -it -p 8100:8100 k8-go-api
```

## Test

For testing, run `go test ./...`

## End points:

1- `/api/rebuild/file`<br>
<strong>Method</strong>: POST.<br>
<strong>Description: </strong>This endpoint accepts requests to rebuild a file with Glasswall d-FIRST™ Engine. Both the file and the Content Management Policy are sent in the request body with a 'Content-Type' of 'multipart/form-data'. The Rebuilt file is then returned in the response body with a 'Content-Type' of 'application/octet-stream'.<br>

2- `/api/rebuild/base64`<br>
<strong>Method</strong> POST.<br>
<strong>Description: </strong>This endpoint accepts requests to rebuild a file with Glasswall d-FIRST™ Engine. The request body contains the Base64 representation of the file and Glasswall Content Management Flags with a 'Content-Type' of 'application/json'. A Base64 Representation of the rebuilt file is then returned in the response with a 'Content-Type' of 'text/plain'.<br>

Select a file below to copy its Base64 Encoded representation to clipboard. The Total supported request size of the API gateway is 6MB, therefore the base64 encoded string must also be less than 6MB.<br>

3- `/api/rebuild/zip`<br>
<strong>Method</strong>: POST.<br>
<strong>Description: </strong>This endpoint accepts requests to rebuild a zip file with Glasswall d-FIRST™ Engine. Both the file and the Content Management Policy are sent in the request body with a 'Content-Type' of 'multipart/form-data'. The Rebuilt file is then returned in the response body with a 'Content-Type' of 'application/octet-stream'.

## Postman Collections link:

https://www.getpostman.com/collections/78fd72df0d74b4c5e849

## Video Demo

https://www.youtube.com/watch?v=TlXwsJrXe68&amp;feature=youtu.be

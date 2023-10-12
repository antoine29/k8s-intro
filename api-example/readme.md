# api-example

This is a really simple api example wrote with go languaje.

## local debug

```bash
$ go get .
$ PORT=8081 go run .
```

## docker build

```bash
$ docker build . -t api-example:v1
$ docker run -d \
  --name api-example-v1 \
  -e PORT=8081 \
  -p 8081:8081 \
  api-example:v1
$ docker logs -f  api-example-v1
```


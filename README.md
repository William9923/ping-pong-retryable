# Retryable Ping Pong
A sample cli & http project showcasing how to deal with connection failure when communicate between multiple services

## 🚀 Quick start
1. Clone the repository
```bash
git clone git@github.com:William9923/ping-pong-retryable.git 
```
2. Setup http receiver server
```bash
go run cmd/http/main.go cmd/http/server.go 
```
3. Run the Cli Program (to start ping-ing the server)
```bash
go run cmd/cli/main.go
```
4. (optional) Change & Test the retryable http client config or the http connection transport config in cli/main.go
5. 

## Example
- Http Server (log)
```
❯ go run cmd/http/main.go cmd/http/server.go
2022/12/11 10:32:19 starting http server...
2022/12/11 10:32:25 http: superfluous response.WriteHeader call from main.(*pongServer).ServeHTTP (server.go:30)
2022/12/11 10:32:25 http: superfluous response.WriteHeader call from main.(*pongServer).ServeHTTP (server.go:30)
2022/12/11 10:32:25 http: superfluous response.WriteHeader call from main.(*pongServer).ServeHTTP (server.go:30)
2022/12/11 10:32:25 http: superfluous response.WriteHeader call from main.(*pongServer).ServeHTTP (server.go:30)
2022/12/11 10:32:25 http: superfluous response.WriteHeader call from main.(*pongServer).ServeHTTP (server.go:30)
2022/12/11 10:32:51 http: superfluous response.WriteHeader call from main.(*pongServer).ServeHTTP (server.go:30)
2022/12/11 10:32:51 http: superfluous response.WriteHeader call from main.(*pongServer).ServeHTTP (server.go:30)
2022/12/11 10:32:51 http: superfluous response.WriteHeader call from main.(*pongServer).ServeHTTP (server.go:30)
2022/12/11 10:32:51 http: superfluous response.WriteHeader call from main.(*pongServer).ServeHTTP (server.go:30)
2022/12/11 10:32:51 http: superfluous response.WriteHeader call from main.(*pongServer).ServeHTTP (server.go:30)
2022/12/11 10:32:57 http: superfluous response.WriteHeader call from main.(*pongServer).ServeHTTP (server.go:30)
2022/12/11 10:32:57 http: superfluous response.WriteHeader call from main.(*pongServer).ServeHTTP (server.go:30)
2022/12/11 10:32:59 http: superfluous response.WriteHeader call from main.(*pongServer).ServeHTTP (server.go:30)
2022/12/11 10:32:59 http: superfluous response.WriteHeader call from main.(*pongServer).ServeHTTP (server.go:30)
2022/12/11 10:32:59 http: superfluous response.WriteHeader call from main.(*pongServer).ServeHTTP (server.go:30)
2022/12/11 11:05:23 http: superfluous response.WriteHeader call from main.(*pongServer).ServeHTTP (server.go:30)
2022/12/11 11:05:23 http: superfluous response.WriteHeader call from main.(*pongServer).ServeHTTP (server.go:30)
2022/12/11 11:05:23 http: superfluous response.WriteHeader call from main.(*pongServer).ServeHTTP (server.go:30)
2022/12/11 11:05:23 http: superfluous response.WriteHeader call from main.(*pongServer).ServeHTTP (server.go:30)
2022/12/11 11:05:23 http: superfluous response.WriteHeader call from main.(*pongServer).ServeHTTP (server.go:30)
2022/12/11 11:06:06 http: superfluous response.WriteHeader call from main.(*pongServer).ServeHTTP (server.go:30)
2022/12/11 11:06:06 http: superfluous response.WriteHeader call from main.(*pongServer).ServeHTTP (server.go:30)
2022/12/11 11:06:06 http: superfluous response.WriteHeader call from main.(*pongServer).ServeHTTP (server.go:30)
2022/12/11 11:06:06 http: superfluous response.WriteHeader call from main.(*pongServer).ServeHTTP (server.go:30)
2022/12/11 11:06:06 http: superfluous response.WriteHeader call from main.(*pongServer).ServeHTTP (server.go:30)
2022/12/11 11:06:19 http: superfluous response.WriteHeader call from main.(*pongServer).ServeHTTP (server.go:30)
2022/12/11 11:06:22 http: superfluous response.WriteHeader call from main.(*pongServer).ServeHTTP (server.go:30)
2022/12/11 11:06:23 http: superfluous response.WriteHeader call from main.(*pongServer).ServeHTTP (server.go:30)
2022/12/11 11:06:23 http: superfluous response.WriteHeader call from main.(*pongServer).ServeHTTP (server.go:30)
2022/12/11 11:06:49 http: superfluous response.WriteHeader call from main.(*pongServer).ServeHTTP (server.go:30)
```
- Http Client (log)
```
❯ go run cmd/cli/main.go
[DEBUG] GET http://localhost:8080/ping
[ATTEMPT]: 0 | [TARGET URL]: http://localhost:8080/ping
[DEBUG] GET http://localhost:8080/ping
[ATTEMPT]: 0 | [TARGET URL]: http://localhost:8080/ping
[DEBUG] GET http://localhost:8080/ping
[ATTEMPT]: 0 | [TARGET URL]: http://localhost:8080/ping
[DEBUG] GET http://localhost:8080/ping
[DEBUG] GET http://localhost:8080/ping
[ATTEMPT]: 0 | [TARGET URL]: http://localhost:8080/ping
[ATTEMPT]: 0 | [TARGET URL]: http://localhost:8080/ping
[DEBUG] GET http://localhost:8080/ping
[ATTEMPT]: 0 | [TARGET URL]: http://localhost:8080/ping
[DEBUG] GET http://localhost:8080/ping
[ATTEMPT]: 0 | [TARGET URL]: http://localhost:8080/ping
[DEBUG] GET http://localhost:8080/ping
[ATTEMPT]: 0 | [TARGET URL]: http://localhost:8080/ping
[DEBUG] GET http://localhost:8080/ping
[ATTEMPT]: 0 | [TARGET URL]: http://localhost:8080/ping
[DEBUG] GET http://localhost:8080/ping
[ATTEMPT]: 0 | [TARGET URL]: http://localhost:8080/ping
[DEBUG] GET http://localhost:8080/ping (status: 500): retrying in 13.704137454s (5 left)
[DEBUG] GET http://localhost:8080/ping (status: 500): retrying in 2.234080421s (5 left)
[DEBUG] GET http://localhost:8080/ping (status: 500): retrying in 15.962149575s (5 left)
[DEBUG] GET http://localhost:8080/ping (status: 500): retrying in 2.234080421s (5 left)
[DEBUG] GET http://localhost:8080/ping (status: 500): retrying in 17.245295497s (5 left)
[ATTEMPT]: 1 | [TARGET URL]: http://localhost:8080/ping
[ATTEMPT]: 1 | [TARGET URL]: http://localhost:8080/ping
[DEBUG] GET http://localhost:8080/ping (status: 500): retrying in 12.939233932s (4 left)
[DEBUG] GET http://localhost:8080/ping (status: 500): retrying in 39.559753338s (4 left)
2022/12/11 11:06:06 Request successfull
2022/12/11 11:06:06 Request successfull
2022/12/11 11:06:06 Request successfull
2022/12/11 11:06:06 Request successfull
2022/12/11 11:06:06 Request successfull
[ATTEMPT]: 1 | [TARGET URL]: http://localhost:8080/ping
[ATTEMPT]: 2 | [TARGET URL]: http://localhost:8080/ping
[ATTEMPT]: 1 | [TARGET URL]: http://localhost:8080/ping
[ATTEMPT]: 1 | [TARGET URL]: http://localhost:8080/ping
2022/12/11 11:06:19 Request successfull
2022/12/11 11:06:22 Request successfull
2022/12/11 11:06:23 Request successfull
2022/12/11 11:06:23 Request successfull
[ATTEMPT]: 2 | [TARGET URL]: http://localhost:8080/ping
2022/12/11 11:06:49 Request successfull
```

## ❌ Prerequisites
- Golang minimum v1.17 (https://golang.org/doc/install)
- Go Modules (https://blog.golang.org/using-go-modules)

## ❤️ Support
If you feel that this repo have helped you provide more example on learning software engineering, then it is enough for me! Wanna contribute more? Please ⭐ this repo so other can see it too!

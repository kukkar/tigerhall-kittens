# Tigerhall_Kittens


### Requirement
We need go1.13 or greater
```
go get github.com/kukkar/tigerhall-kittens
```


### Default Middlewares
- Recovery (It will recover from any panic in endpoint and won't stop the server)
- Debug (It will add context value to debug the SQL, Error stack trace, Request and Response on-demand)
- Logger (gin default logger is used to log url hit, response time and response status)
- CORS (Cross-origin resource sharing)


### Default Endpoints
- __GET /__ *(check server is running)*
- __GET /swagger/docs.json__ *(for swagger if you have created swag document)*
- __GET /debug/pprof/__ *(To view all available profiles)*
- __GET /debug/pprof/cmdline__ *(running program's command line, with arguments separated by NUL bytes)*
- __GET /debug/pprof/profile__ *(CPU profile)*
- __POST /debug/pprof/symbol__
- __GET /debug/pprof/symbol__
- __GET /debug/pprof/trace__ *(execution trace)*
- __GET /debug/pprof/allocs__
- __GET /debug/pprof/block__ *(to look at the goroutine blocking profile)*
- __GET /debug/pprof/goroutine__ *(Goroutine profile report the stack traces of all current goroutines.)*
- __GET /debug/pprof/heap__ *(to look at the heap profile)*
- __GET /debug/pprof/mutex__ *(to look at the holders of contended mutexes)*
- __GET /debug/pprof/threadcreate__ *(Thread creation profile reports the sections of the program)*
- __GET /otp/healtcheck__ *(healtcheck on service based on config)*

### build

- to build docker run
    - sudo docker build -t tigerhall -f Dockerfile .
- to run docker image
    - docker run tigerhall
- to run go  tests
    - go test ...

### Swagger Documentor
- http://localhost:8085/swagger/index.html
    


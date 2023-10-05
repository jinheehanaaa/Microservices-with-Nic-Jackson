# 001
## ListenAndServe function
- Bind to "localhost :9090"
- Handler as "nil"

## ServerMux & HandleFunc
- ServerMux is responsible for redirecting path
- HandleFunc takes my function and creates http handler from it and adds it to defalt ServerMux

## ResponseWriter & Request
- ResponseWriter => Writes headers, response body, status code
- Request => Path, Method, HTTP Version

# 002
## Logger
- Our own logger for testability

## ServeMux & Handler

## Server
- Timeout
- Graceful shutdown

## Non-blocing (go func)

## OS.signal
- Notify for broadcasting



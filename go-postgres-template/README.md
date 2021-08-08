# Golang PostgreSQL

### Environment File
```
FRONTEND_URL=""
SQL_DB_USER=""
SQL_DB_PASS=""
SQL_DB_DB=""
SQL_DB_HOST=""
SQL_DB_PORT=""
PORT=""
```

### Routes

```http
POST /user HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
    "email": "",
    "name": "",
    "gid": ""
}
```

```http
GET /user/:id HTTP/1.1
Host: localhost:8080
```


### Logging

- Log function : `logrus.LogIt(enum, funcName, errMsg, error)`

```go
const (
 	LogFatal = iota
 	LogWarn  = iota
 	LogInfo  = iota
)
```

### ErrSet

```go
	ErrUnAuthorized   = fmt.Errorf("unauthorized")
	ErrInternalServer = fmt.Errorf("internal server error")
	ErrBadRequest     = fmt.Errorf("bad request")
	ErrNotFound       = fmt.Errorf("not found")

	
```
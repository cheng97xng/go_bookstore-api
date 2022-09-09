"# go_bookstore-api" 

https://pkg.go.dev/std
https://golang.org/pkg/

<!-- request message for response to the caller-->
{
    "message":"user 123 not found",
    "status":"402",
    "error":"Not_found"
}

{
    "message":"Invalid json body",
    "status":"400",
    "error":"Bad_request"
}
{
    "message":"database is down",
    "status":"500",
    "error":"internal_server_error"
}
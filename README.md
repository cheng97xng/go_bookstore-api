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

<!-- Git command -->
1. git clone [http://] is copy project from git to host
2. git add . is add all files to git
3. git commit -m "[you can write any thing for evidance]" to read file on git
4. git push origin master is push every thing to git after add and commit finish
5. git pull origin master is pull last file from git
6. git reset --hard [<SHA-1> you need return to example:56e05fced] is restore to that
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

<!-- Http Status -->
package http

// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
const (
	StatusOK                   = 200 // RFC 9110, 15.3.1
	StatusCreated              = 201 // RFC 9110, 15.3.2
	StatusAccepted             = 202 // RFC 9110, 15.3.3
	StatusNonAuthoritativeInfo = 203 // RFC 9110, 15.3.4
	StatusNoContent            = 204 // RFC 9110, 15.3.5
	StatusResetContent         = 205 // RFC 9110, 15.3.6

	StatusFound             = 302 // RFC 9110, 15.4.3
	StatusSeeOther          = 303 // RFC 9110, 15.4.4
	StatusNotModified       = 304 // RFC 9110, 15.4.5
	StatusUseProxy          = 305 // RFC 9110, 15.4.6
	_                       = 306 // RFC 9110, 15.4.7 (Unused)
	StatusTemporaryRedirect = 307 // RFC 9110, 15.4.8
	StatusPermanentRedirect = 308 // RFC 9110, 15.4.9

	StatusBadRequest                   = 400 // RFC 9110, 15.5.1
	StatusUnauthorized                 = 401 // RFC 9110, 15.5.2
	StatusPaymentRequired              = 402 // RFC 9110, 15.5.3
	StatusForbidden                    = 403 // RFC 9110, 15.5.4
	StatusNotFound                     = 404 // RFC 9110, 15.5.5
	StatusMethodNotAllowed             = 405 // RFC 9110, 15.5.6
	StatusNotAcceptable                = 406 // RFC 9110, 15.5.7
	StatusProxyAuthRequired            = 407 // RFC 9110, 15.5.8
	StatusRequestTimeout               = 408 // RFC 9110, 15.5.9
	StatusConflict                     = 409 // RFC 9110, 15.5.10
	StatusGone                         = 410 // RFC 9110, 15.5.11
	StatusLengthRequired               = 411 // RFC 9110, 15.5.12
	StatusPreconditionFailed           = 412 // RFC 9110, 15.5.13
	StatusRequestEntityTooLarge        = 413 // RFC 9110, 15.5.14
	StatusRequestURITooLong            = 414 // RFC 9110, 15.5.15
	StatusUnsupportedMediaType         = 415 // RFC 9110, 15.5.16
	StatusRequestedRangeNotSatisfiable = 416 // RFC 9110, 15.5.17
	StatusExpectationFailed            = 417 // RFC 9110, 15.5.18
	StatusTeapot                       = 418 // RFC 9110, 15.5.19 (Unused)
	StatusMisdirectedRequest           = 421 // RFC 9110, 15.5.20
	StatusUnprocessableEntity          = 422 // RFC 9110, 15.5.21
	StatusLocked                       = 423 // RFC 4918, 11.3
	StatusFailedDependency             = 424 // RFC 4918, 11.4
	StatusTooEarly                     = 425 // RFC 8470, 5.2.
	StatusUpgradeRequired              = 426 // RFC 9110, 15.5.22
	StatusPreconditionRequired         = 428 // RFC 6585, 3
	StatusTooManyRequests              = 429 // RFC 6585, 4
	StatusRequestHeaderFieldsTooLarge  = 431 // RFC 6585, 5
	StatusUnavailableForLegalReasons   = 451 // RFC 7725, 3

	StatusInternalServerError           = 500 // RFC 9110, 15.6.1
	StatusNotImplemented                = 501 // RFC 9110, 15.6.2
	StatusBadGateway                    = 502 // RFC 9110, 15.6.3
	StatusServiceUnavailable            = 503 // RFC 9110, 15.6.4
	StatusGatewayTimeout                = 504 // RFC 9110, 15.6.5
	StatusHTTPVersionNotSupported       = 505 // RFC 9110, 15.6.6
	StatusVariantAlsoNegotiates         = 506 // RFC 2295, 8.1
	StatusInsufficientStorage           = 507 // RFC 4918, 11.5
	StatusLoopDetected                  = 508 // RFC 5842, 7.2
	StatusNotExtended                   = 510 // RFC 2774, 7
	StatusNetworkAuthenticationRequired = 511 // RFC 6585, 6
)

<!-- Connect to mysql db -->
	MySql Driver: _"github.com/go-sql-driver/mysql"
	go get github.com/go-sql-driver/mysql

		var (
		Users_db *sql.DB
		)

	func init() {
		datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
			"root",       //mysql username
			"chengsoft",  //mysql password
			"127.0.0.1:3306",  //mysql localhost
			"golang_db1", //mysql database
		)
		var err error
		Users_db, err = sql.Open("mysql", datasourceName)
		if err != nil {
			//panic function is end the program
			panic(err)
		}
	}

	<!-- create user table structure-->

	CREATE TABLE users(
    id BIGINT(20) NOT NULL AUTO_INCREMENT,
    first_name VARCHAR(35) NULL,
    last_name VARCHAR(30) NULL,
    email VARCHAR(60) NULL,
    date_created VARCHAR(45) NULL,
    PRIMARY KEY(id),
    UNIQUE INDEX `email_UNIQUE` (`email` ASC));


	<!-- Access to mysql from another pc by terminal -->
	Enter as: mysql -u username -h host -p
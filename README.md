# http-status

[![Build Status](https://travis-ci.org/mtojek/http-status.svg?branch=master)](https://travis-ci.org/mtojek/http-status)

Use more API-friendly http.StatusText(...) messages.

Status: **Done**

If you agree that standard HTTP status texts do not conform to the nature of API responses, please check this library and start using HTTP status text codes instead.

| HTTP status code | HTTP status text      | HTTP status text code |
|:----------------:|:---------------------:|:---------------------:|
| 404              | Not Found             | NOT_FOUND             |
| 418              | I'm a teapot          | I_M_A_TEAPOT          |
| 500              | Internal Server Error | INTERNAL_SERVER_ERROR |

## Usage

~~~ go
var statusTextCode string
statusTextCode = StatusTextCode(http.StatusInternalServerError) // statusTextCode = "INTERNAL_SERVER_ERROR"
~~~

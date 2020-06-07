package router

import "net/http"

// method satisfies `Router` by matching `Request.Method`
type method string

func (method method) RouteHTTP(r *http.Request) bool { return string(method) == r.Method }
func (method method) isHTTPRouter() I                { return method }

// CONNECT is a Router that returns if `Request.Method` is CONNECT
var CONNECT = method("CONNECT")

// DELETE is a Router that returns if `Request.Method` is DELETE
var DELETE = method("DELETE")

// GET is a Router that returns if `Request.Method` is GET
var GET = method("GET")

// HEAD is a Router that returns if `Request.Method` is HEAD
var HEAD = method("HEAD")

// OPTIONS is a Router that returns if `Request.Method` is OPTIONS
var OPTIONS = method("OPTIONS")

// POST is a Router that returns if `Request.Method` is POST
var POST = method("POST")

// PUT is a Router that returns if `Request.Method` is PUT
var PUT = method("PUT")

// TRACE is a Router that returns if `Request.Method` is TRACE
var TRACE = method("TRACE")

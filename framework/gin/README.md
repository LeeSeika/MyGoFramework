[//]: # (# Gin Web Framework)

[//]: # ()
[//]: # (<img align="right" width="159px" src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png">)

[//]: # ()
[//]: # ([![Build Status]&#40;https://travis-ci.org/gin-gonic/gin.svg&#41;]&#40;https://travis-ci.org/gin-gonic/gin&#41;)

[//]: # ([![codecov]&#40;https://codecov.io/gh/gin-gonic/gin/branch/master/graph/badge.svg&#41;]&#40;https://codecov.io/gh/gin-gonic/gin&#41;)

[//]: # ([![Go Report Card]&#40;https://goreportcard.com/badge/github.com/gin-gonic/gin&#41;]&#40;https://goreportcard.com/report/github.com/gin-gonic/gin&#41;)

[//]: # ([![GoDoc]&#40;https://pkg.go.dev/badge/github.com/gin-gonic/gin?status.svg&#41;]&#40;https://pkg.go.dev/github.com/gin-gonic/gin?tab=doc&#41;)

[//]: # ([![Join the chat at https://gitter.im/gin-gonic/gin]&#40;https://badges.gitter.im/Join%20Chat.svg&#41;]&#40;https://gitter.im/gin-gonic/gin?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge&#41;)

[//]: # ([![Sourcegraph]&#40;https://sourcegraph.com/github.com/gin-gonic/gin/-/badge.svg&#41;]&#40;https://sourcegraph.com/github.com/gin-gonic/gin?badge&#41;)

[//]: # ([![Open Source Helpers]&#40;https://www.codetriage.com/gin-gonic/gin/badges/users.svg&#41;]&#40;https://www.codetriage.com/gin-gonic/gin&#41;)

[//]: # ([![Release]&#40;https://img.shields.io/github/release/gin-gonic/gin.svg?style=flat-square&#41;]&#40;https://github.com/gin-gonic/gin/releases&#41;)

[//]: # ([![TODOs]&#40;https://badgen.net/https/api.tickgit.com/badgen/github.com/gin-gonic/gin&#41;]&#40;https://www.tickgit.com/browse?repo=github.com/gin-gonic/gin&#41;)

[//]: # ()
[//]: # (Gin is a web framework written in Go &#40;Golang&#41;. It features a martini-like API with performance that is up to 40 times faster thanks to [httprouter]&#40;https://github.com/julienschmidt/httprouter&#41;. If you need performance and good productivity, you will love Gin.)

[//]: # ()
[//]: # ()
[//]: # (## Contents)

[//]: # ()
[//]: # (- [Gin Web Framework]&#40;#gin-web-framework&#41;)

[//]: # (  - [Contents]&#40;#contents&#41;)

[//]: # (  - [Installation]&#40;#installation&#41;)

[//]: # (  - [Quick start]&#40;#quick-start&#41;)

[//]: # (  - [Benchmarks]&#40;#benchmarks&#41;)

[//]: # (  - [Gin v1. stable]&#40;#gin-v1-stable&#41;)

[//]: # (  - [Build with jsoniter]&#40;#build-with-jsoniter&#41;)

[//]: # (  - [API Examples]&#40;#api-examples&#41;)

[//]: # (    - [Using GET, POST, PUT, PATCH, DELETE and OPTIONS]&#40;#using-get-post-put-patch-delete-and-options&#41;)

[//]: # (    - [Parameters in path]&#40;#parameters-in-path&#41;)

[//]: # (    - [Querystring parameters]&#40;#querystring-parameters&#41;)

[//]: # (    - [Multipart/Urlencoded Form]&#40;#multiparturlencoded-form&#41;)

[//]: # (    - [Another example: query + post form]&#40;#another-example-query--post-form&#41;)

[//]: # (    - [Map as querystring or postform parameters]&#40;#map-as-querystring-or-postform-parameters&#41;)

[//]: # (    - [Upload files]&#40;#upload-files&#41;)

[//]: # (      - [Single file]&#40;#single-file&#41;)

[//]: # (      - [Multiple files]&#40;#multiple-files&#41;)

[//]: # (    - [Grouping routes]&#40;#grouping-routes&#41;)

[//]: # (    - [Blank Gin without middleware by default]&#40;#blank-gin-without-middleware-by-default&#41;)

[//]: # (    - [Using middleware]&#40;#using-middleware&#41;)

[//]: # (    - [How to write log file]&#40;#how-to-write-log-file&#41;)

[//]: # (    - [Custom Log Format]&#40;#custom-log-format&#41;)

[//]: # (    - [Controlling Log output coloring]&#40;#controlling-log-output-coloring&#41;)

[//]: # (    - [Model binding and validation]&#40;#model-binding-and-validation&#41;)

[//]: # (    - [Custom Validators]&#40;#custom-validators&#41;)

[//]: # (    - [Only Bind Query String]&#40;#only-bind-query-string&#41;)

[//]: # (    - [Bind Query String or Post Data]&#40;#bind-query-string-or-post-data&#41;)

[//]: # (    - [Bind Uri]&#40;#bind-uri&#41;)

[//]: # (    - [Bind Header]&#40;#bind-header&#41;)

[//]: # (    - [Bind HTML checkboxes]&#40;#bind-html-checkboxes&#41;)

[//]: # (    - [Multipart/Urlencoded binding]&#40;#multiparturlencoded-binding&#41;)

[//]: # (    - [XML, JSON, YAML and ProtoBuf rendering]&#40;#xml-json-yaml-and-protobuf-rendering&#41;)

[//]: # (      - [SecureJSON]&#40;#securejson&#41;)

[//]: # (      - [JSONP]&#40;#jsonp&#41;)

[//]: # (      - [AsciiJSON]&#40;#asciijson&#41;)

[//]: # (      - [PureJSON]&#40;#purejson&#41;)

[//]: # (    - [Serving static files]&#40;#serving-static-files&#41;)

[//]: # (    - [Serving data from file]&#40;#serving-data-from-file&#41;)

[//]: # (    - [Serving data from reader]&#40;#serving-data-from-reader&#41;)

[//]: # (    - [HTML rendering]&#40;#html-rendering&#41;)

[//]: # (      - [Custom Template renderer]&#40;#custom-template-renderer&#41;)

[//]: # (      - [Custom Delimiters]&#40;#custom-delimiters&#41;)

[//]: # (      - [Custom Template Funcs]&#40;#custom-template-funcs&#41;)

[//]: # (    - [Multitemplate]&#40;#multitemplate&#41;)

[//]: # (    - [Redirects]&#40;#redirects&#41;)

[//]: # (    - [Custom Middleware]&#40;#custom-middleware&#41;)

[//]: # (    - [Using BasicAuth&#40;&#41; middleware]&#40;#using-basicauth-middleware&#41;)

[//]: # (    - [Goroutines inside a middleware]&#40;#goroutines-inside-a-middleware&#41;)

[//]: # (    - [Custom HTTP configuration]&#40;#custom-http-configuration&#41;)

[//]: # (    - [Support Let's Encrypt]&#40;#support-lets-encrypt&#41;)

[//]: # (    - [Run multiple service using Gin]&#40;#run-multiple-service-using-gin&#41;)

[//]: # (    - [Graceful shutdown or restart]&#40;#graceful-shutdown-or-restart&#41;)

[//]: # (      - [Third-party packages]&#40;#third-party-packages&#41;)

[//]: # (      - [Manually]&#40;#manually&#41;)

[//]: # (    - [Build a single binary with templates]&#40;#build-a-single-binary-with-templates&#41;)

[//]: # (    - [Bind form-data request with custom struct]&#40;#bind-form-data-request-with-custom-struct&#41;)

[//]: # (    - [Try to bind body into different structs]&#40;#try-to-bind-body-into-different-structs&#41;)

[//]: # (    - [http2 server push]&#40;#http2-server-push&#41;)

[//]: # (    - [Define format for the log of routes]&#40;#define-format-for-the-log-of-routes&#41;)

[//]: # (    - [Set and get a cookie]&#40;#set-and-get-a-cookie&#41;)

[//]: # (  - [Testing]&#40;#testing&#41;)

[//]: # (  - [Users]&#40;#users&#41;)

[//]: # ()
[//]: # (## Installation)

[//]: # ()
[//]: # (To install Gin package, you need to install Go and set your Go workspace first.)

[//]: # ()
[//]: # (1. The first need [Go]&#40;https://golang.org/&#41; installed &#40;**version 1.12+ is required**&#41;, then you can use the below Go command to install Gin.)

[//]: # ()
[//]: # (```sh)

[//]: # ($ go get -u github.com/gin-gonic/gin)

[//]: # (```)

[//]: # ()
[//]: # (2. Import it in your code:)

[//]: # ()
[//]: # (```go)

[//]: # (import "github.com/gin-gonic/gin")

[//]: # (```)

[//]: # ()
[//]: # (3. &#40;Optional&#41; Import `net/http`. This is required for example if using constants such as `http.StatusOK`.)

[//]: # ()
[//]: # (```go)

[//]: # (import "net/http")

[//]: # (```)

[//]: # ()
[//]: # (## Quick start)

[//]: # ()
[//]: # (```sh)

[//]: # (# assume the following codes in example.go file)

[//]: # ($ cat example.go)

[//]: # (```)

[//]: # ()
[//]: # (```go)

[//]: # (package main)

[//]: # ()
[//]: # (import "github.com/gin-gonic/gin")

[//]: # ()
[//]: # (func main&#40;&#41; {)

[//]: # (	r := gin.Default&#40;&#41;)

[//]: # (	r.GET&#40;"/ping", func&#40;c *gin.Context&#41; {)

[//]: # (		c.JSON&#40;200, gin.H{)

[//]: # (			"message": "pong",)

[//]: # (		}&#41;)

[//]: # (	}&#41;)

[//]: # (	r.Run&#40;&#41; // listen and serve on 0.0.0.0:8080 &#40;for windows "localhost:8080"&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (```)

[//]: # (# run example.go and visit 0.0.0.0:8080/ping &#40;for windows "localhost:8080/ping"&#41; on browser)

[//]: # ($ go run example.go)

[//]: # (```)

[//]: # ()
[//]: # (## Benchmarks)

[//]: # ()
[//]: # (Gin uses a custom version of [HttpRouter]&#40;https://github.com/julienschmidt/httprouter&#41;)

[//]: # ()
[//]: # ([See all benchmarks]&#40;/BENCHMARKS.md&#41;)

[//]: # ()
[//]: # (| Benchmark name                 |       &#40;1&#41; |             &#40;2&#41; |          &#40;3&#41; |             &#40;4&#41; |)

[//]: # (| ------------------------------ | ---------:| ---------------:| ------------:| ---------------:|)

[//]: # (| BenchmarkGin_GithubAll         | **43550** | **27364 ns/op** |   **0 B/op** | **0 allocs/op** |)

[//]: # (| BenchmarkAce_GithubAll         |     40543 |     29670 ns/op |       0 B/op |     0 allocs/op |)

[//]: # (| BenchmarkAero_GithubAll        |     57632 |     20648 ns/op |       0 B/op |     0 allocs/op |)

[//]: # (| BenchmarkBear_GithubAll        |      9234 |    216179 ns/op |   86448 B/op |   943 allocs/op |)

[//]: # (| BenchmarkBeego_GithubAll       |      7407 |    243496 ns/op |   71456 B/op |   609 allocs/op |)

[//]: # (| BenchmarkBone_GithubAll        |       420 |   2922835 ns/op |  720160 B/op |  8620 allocs/op |)

[//]: # (| BenchmarkChi_GithubAll         |      7620 |    238331 ns/op |   87696 B/op |   609 allocs/op |)

[//]: # (| BenchmarkDenco_GithubAll       |     18355 |     64494 ns/op |   20224 B/op |   167 allocs/op |)

[//]: # (| BenchmarkEcho_GithubAll        |     31251 |     38479 ns/op |       0 B/op |     0 allocs/op |)

[//]: # (| BenchmarkGocraftWeb_GithubAll  |      4117 |    300062 ns/op |  131656 B/op |  1686 allocs/op |)

[//]: # (| BenchmarkGoji_GithubAll        |      3274 |    416158 ns/op |   56112 B/op |   334 allocs/op |)

[//]: # (| BenchmarkGojiv2_GithubAll      |      1402 |    870518 ns/op |  352720 B/op |  4321 allocs/op |)

[//]: # (| BenchmarkGoJsonRest_GithubAll  |      2976 |    401507 ns/op |  134371 B/op |  2737 allocs/op |)

[//]: # (| BenchmarkGoRestful_GithubAll   |       410 |   2913158 ns/op |  910144 B/op |  2938 allocs/op |)

[//]: # (| BenchmarkGorillaMux_GithubAll  |       346 |   3384987 ns/op |  251650 B/op |  1994 allocs/op |)

[//]: # (| BenchmarkGowwwRouter_GithubAll |     10000 |    143025 ns/op |   72144 B/op |   501 allocs/op |)

[//]: # (| BenchmarkHttpRouter_GithubAll  |     55938 |     21360 ns/op |       0 B/op |     0 allocs/op |)

[//]: # (| BenchmarkHttpTreeMux_GithubAll |     10000 |    153944 ns/op |   65856 B/op |   671 allocs/op |)

[//]: # (| BenchmarkKocha_GithubAll       |     10000 |    106315 ns/op |   23304 B/op |   843 allocs/op |)

[//]: # (| BenchmarkLARS_GithubAll        |     47779 |     25084 ns/op |       0 B/op |     0 allocs/op |)

[//]: # (| BenchmarkMacaron_GithubAll     |      3266 |    371907 ns/op |  149409 B/op |  1624 allocs/op |)

[//]: # (| BenchmarkMartini_GithubAll     |       331 |   3444706 ns/op |  226551 B/op |  2325 allocs/op |)

[//]: # (| BenchmarkPat_GithubAll         |       273 |   4381818 ns/op | 1483152 B/op | 26963 allocs/op |)

[//]: # (| BenchmarkPossum_GithubAll      |     10000 |    164367 ns/op |   84448 B/op |   609 allocs/op |)

[//]: # (| BenchmarkR2router_GithubAll    |     10000 |    160220 ns/op |   77328 B/op |   979 allocs/op |)

[//]: # (| BenchmarkRivet_GithubAll       |     14625 |     82453 ns/op |   16272 B/op |   167 allocs/op |)

[//]: # (| BenchmarkTango_GithubAll       |      6255 |    279611 ns/op |   63826 B/op |  1618 allocs/op |)

[//]: # (| BenchmarkTigerTonic_GithubAll  |      2008 |    687874 ns/op |  193856 B/op |  4474 allocs/op |)

[//]: # (| BenchmarkTraffic_GithubAll     |       355 |   3478508 ns/op |  820744 B/op | 14114 allocs/op |)

[//]: # (| BenchmarkVulcan_GithubAll      |      6885 |    193333 ns/op |   19894 B/op |   609 allocs/op |)

[//]: # ()
[//]: # (- &#40;1&#41;: Total Repetitions achieved in constant time, higher means more confident result)

[//]: # (- &#40;2&#41;: Single Repetition Duration &#40;ns/op&#41;, lower is better)

[//]: # (- &#40;3&#41;: Heap Memory &#40;B/op&#41;, lower is better)

[//]: # (- &#40;4&#41;: Average Allocations per Repetition &#40;allocs/op&#41;, lower is better)

[//]: # ()
[//]: # (## Gin v1. stable)

[//]: # ()
[//]: # (- [x] Zero allocation router.)

[//]: # (- [x] Still the fastest http router and framework. From routing to writing.)

[//]: # (- [x] Complete suite of unit tests.)

[//]: # (- [x] Battle tested.)

[//]: # (- [x] API frozen, new releases will not break your code.)

[//]: # ()
[//]: # (## Build with [jsoniter]&#40;https://github.com/json-iterator/go&#41;)

[//]: # ()
[//]: # (Gin uses `encoding/json` as default json package but you can change to [jsoniter]&#40;https://github.com/json-iterator/go&#41; by build from other tags.)

[//]: # ()
[//]: # (```sh)

[//]: # ($ go build -tags=jsoniter .)

[//]: # (```)

[//]: # ()
[//]: # (## API Examples)

[//]: # ()
[//]: # (You can find a number of ready-to-run examples at [Gin examples repository]&#40;https://github.com/gin-gonic/examples&#41;.)

[//]: # ()
[//]: # (### Using GET, POST, PUT, PATCH, DELETE and OPTIONS)

[//]: # ()
[//]: # (```go)

[//]: # (func main&#40;&#41; {)

[//]: # (	// Creates a gin router with default middleware:)

[//]: # (	// logger and recovery &#40;crash-free&#41; middleware)

[//]: # (	router := gin.Default&#40;&#41;)

[//]: # ()
[//]: # (	router.GET&#40;"/someGet", getting&#41;)

[//]: # (	router.POST&#40;"/somePost", posting&#41;)

[//]: # (	router.PUT&#40;"/somePut", putting&#41;)

[//]: # (	router.DELETE&#40;"/someDelete", deleting&#41;)

[//]: # (	router.PATCH&#40;"/somePatch", patching&#41;)

[//]: # (	router.HEAD&#40;"/someHead", head&#41;)

[//]: # (	router.OPTIONS&#40;"/someOptions", options&#41;)

[//]: # ()
[//]: # (	// By default it serves on :8080 unless a)

[//]: # (	// PORT environment variable was defined.)

[//]: # (	router.Run&#40;&#41;)

[//]: # (	// router.Run&#40;":3000"&#41; for a hard coded port)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (### Parameters in path)

[//]: # ()
[//]: # (```go)

[//]: # (func main&#40;&#41; {)

[//]: # (	router := gin.Default&#40;&#41;)

[//]: # ()
[//]: # (	// This handler will match /user/john but will not match /user/ or /user)

[//]: # (	router.GET&#40;"/user/:name", func&#40;c *gin.Context&#41; {)

[//]: # (		name := c.Param&#40;"name"&#41;)

[//]: # (		c.String&#40;http.StatusOK, "Hello %s", name&#41;)

[//]: # (	}&#41;)

[//]: # ()
[//]: # (	// However, this one will match /user/john/ and also /user/john/send)

[//]: # (	// If no other routers match /user/john, it will redirect to /user/john/)

[//]: # (	router.GET&#40;"/user/:name/*action", func&#40;c *gin.Context&#41; {)

[//]: # (		name := c.Param&#40;"name"&#41;)

[//]: # (		action := c.Param&#40;"action"&#41;)

[//]: # (		message := name + " is " + action)

[//]: # (		c.String&#40;http.StatusOK, message&#41;)

[//]: # (	}&#41;)

[//]: # ()
[//]: # (	// For each matched request Context will hold the route definition)

[//]: # (	router.POST&#40;"/user/:name/*action", func&#40;c *gin.Context&#41; {)

[//]: # (		c.FullPath&#40;&#41; == "/user/:name/*action" // true)

[//]: # (	}&#41;)

[//]: # ()
[//]: # (	// This handler will add a new router for /user/groups.)

[//]: # (	// Exact routes are resolved before param routes, regardless of the order they were defined.)

[//]: # (	// Routes starting with /user/groups are never interpreted as /user/:name/... routes)

[//]: # (	router.GET&#40;"/user/groups", func&#40;c *gin.Context&#41; {)

[//]: # (		c.String&#40;http.StatusOK, "The available groups are [...]", name&#41;)

[//]: # (	}&#41;)

[//]: # ()
[//]: # (	router.Run&#40;":8080"&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (### Querystring parameters)

[//]: # ()
[//]: # (```go)

[//]: # (func main&#40;&#41; {)

[//]: # (	router := gin.Default&#40;&#41;)

[//]: # ()
[//]: # (	// Query string parameters are parsed using the existing underlying request object.)

[//]: # (	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe)

[//]: # (	router.GET&#40;"/welcome", func&#40;c *gin.Context&#41; {)

[//]: # (		firstname := c.DefaultQuery&#40;"firstname", "Guest"&#41;)

[//]: # (		lastname := c.Query&#40;"lastname"&#41; // shortcut for c.Request.URL.Query&#40;&#41;.Get&#40;"lastname"&#41;)

[//]: # ()
[//]: # (		c.String&#40;http.StatusOK, "Hello %s %s", firstname, lastname&#41;)

[//]: # (	}&#41;)

[//]: # (	router.Run&#40;":8080"&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (### Multipart/Urlencoded Form)

[//]: # ()
[//]: # (```go)

[//]: # (func main&#40;&#41; {)

[//]: # (	router := gin.Default&#40;&#41;)

[//]: # ()
[//]: # (	router.POST&#40;"/form_post", func&#40;c *gin.Context&#41; {)

[//]: # (		message := c.PostForm&#40;"message"&#41;)

[//]: # (		nick := c.DefaultPostForm&#40;"nick", "anonymous"&#41;)

[//]: # ()
[//]: # (		c.JSON&#40;200, gin.H{)

[//]: # (			"status":  "posted",)

[//]: # (			"message": message,)

[//]: # (			"nick":    nick,)

[//]: # (		}&#41;)

[//]: # (	}&#41;)

[//]: # (	router.Run&#40;":8080"&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (### Another example: query + post form)

[//]: # ()
[//]: # (```)

[//]: # (POST /post?id=1234&page=1 HTTP/1.1)

[//]: # (Content-Type: application/x-www-form-urlencoded)

[//]: # ()
[//]: # (name=manu&message=this_is_great)

[//]: # (```)

[//]: # ()
[//]: # (```go)

[//]: # (func main&#40;&#41; {)

[//]: # (	router := gin.Default&#40;&#41;)

[//]: # ()
[//]: # (	router.POST&#40;"/post", func&#40;c *gin.Context&#41; {)

[//]: # ()
[//]: # (		id := c.Query&#40;"id"&#41;)

[//]: # (		page := c.DefaultQuery&#40;"page", "0"&#41;)

[//]: # (		name := c.PostForm&#40;"name"&#41;)

[//]: # (		message := c.PostForm&#40;"message"&#41;)

[//]: # ()
[//]: # (		fmt.Printf&#40;"id: %s; page: %s; name: %s; message: %s", id, page, name, message&#41;)

[//]: # (	}&#41;)

[//]: # (	router.Run&#40;":8080"&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (```)

[//]: # (id: 1234; page: 1; name: manu; message: this_is_great)

[//]: # (```)

[//]: # ()
[//]: # (### Map as querystring or postform parameters)

[//]: # ()
[//]: # (```)

[//]: # (POST /post?ids[a]=1234&ids[b]=hello HTTP/1.1)

[//]: # (Content-Type: application/x-www-form-urlencoded)

[//]: # ()
[//]: # (names[first]=thinkerou&names[second]=tianou)

[//]: # (```)

[//]: # ()
[//]: # (```go)

[//]: # (func main&#40;&#41; {)

[//]: # (	router := gin.Default&#40;&#41;)

[//]: # ()
[//]: # (	router.POST&#40;"/post", func&#40;c *gin.Context&#41; {)

[//]: # ()
[//]: # (		ids := c.QueryMap&#40;"ids"&#41;)

[//]: # (		names := c.PostFormMap&#40;"names"&#41;)

[//]: # ()
[//]: # (		fmt.Printf&#40;"ids: %v; names: %v", ids, names&#41;)

[//]: # (	}&#41;)

[//]: # (	router.Run&#40;":8080"&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (```)

[//]: # (ids: map[b:hello a:1234]; names: map[second:tianou first:thinkerou])

[//]: # (```)

[//]: # ()
[//]: # (### Upload files)

[//]: # ()
[//]: # (#### Single file)

[//]: # ()
[//]: # (References issue [#774]&#40;https://github.com/gin-gonic/gin/issues/774&#41; and detail [example code]&#40;https://github.com/gin-gonic/examples/tree/master/upload-file/single&#41;.)

[//]: # ()
[//]: # (`file.Filename` **SHOULD NOT** be trusted. See [`Content-Disposition` on MDN]&#40;https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Disposition#Directives&#41; and [#1693]&#40;https://github.com/gin-gonic/gin/issues/1693&#41;)

[//]: # ()
[//]: # (> The filename is always optional and must not be used blindly by the application: path information should be stripped, and conversion to the server file system rules should be done.)

[//]: # ()
[//]: # (```go)

[//]: # (func main&#40;&#41; {)

[//]: # (	router := gin.Default&#40;&#41;)

[//]: # (	// Set a lower memory limit for multipart forms &#40;default is 32 MiB&#41;)

[//]: # (	router.MaxMultipartMemory = 8 << 20  // 8 MiB)

[//]: # (	router.POST&#40;"/upload", func&#40;c *gin.Context&#41; {)

[//]: # (		// single file)

[//]: # (		file, _ := c.FormFile&#40;"file"&#41;)

[//]: # (		log.Println&#40;file.Filename&#41;)

[//]: # ()
[//]: # (		// Upload the file to specific dst.)

[//]: # (		c.SaveUploadedFile&#40;file, dst&#41;)

[//]: # ()
[//]: # (		c.String&#40;http.StatusOK, fmt.Sprintf&#40;"'%s' uploaded!", file.Filename&#41;&#41;)

[//]: # (	}&#41;)

[//]: # (	router.Run&#40;":8080"&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (How to `curl`:)

[//]: # ()
[//]: # (```bash)

[//]: # (curl -X POST http://localhost:8080/upload \)

[//]: # (  -F "file=@/Users/appleboy/test.zip" \)

[//]: # (  -H "Content-Type: multipart/form-data")

[//]: # (```)

[//]: # ()
[//]: # (#### Multiple files)

[//]: # ()
[//]: # (See the detail [example code]&#40;https://github.com/gin-gonic/examples/tree/master/upload-file/multiple&#41;.)

[//]: # ()
[//]: # (```go)

[//]: # (func main&#40;&#41; {)

[//]: # (	router := gin.Default&#40;&#41;)

[//]: # (	// Set a lower memory limit for multipart forms &#40;default is 32 MiB&#41;)

[//]: # (	router.MaxMultipartMemory = 8 << 20  // 8 MiB)

[//]: # (	router.POST&#40;"/upload", func&#40;c *gin.Context&#41; {)

[//]: # (		// Multipart form)

[//]: # (		form, _ := c.MultipartForm&#40;&#41;)

[//]: # (		files := form.File["upload[]"])

[//]: # ()
[//]: # (		for _, file := range files {)

[//]: # (			log.Println&#40;file.Filename&#41;)

[//]: # ()
[//]: # (			// Upload the file to specific dst.)

[//]: # (			c.SaveUploadedFile&#40;file, dst&#41;)

[//]: # (		})

[//]: # (		c.String&#40;http.StatusOK, fmt.Sprintf&#40;"%d files uploaded!", len&#40;files&#41;&#41;&#41;)

[//]: # (	}&#41;)

[//]: # (	router.Run&#40;":8080"&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (How to `curl`:)

[//]: # ()
[//]: # (```bash)

[//]: # (curl -X POST http://localhost:8080/upload \)

[//]: # (  -F "upload[]=@/Users/appleboy/test1.zip" \)

[//]: # (  -F "upload[]=@/Users/appleboy/test2.zip" \)

[//]: # (  -H "Content-Type: multipart/form-data")

[//]: # (```)

[//]: # ()
[//]: # (### Grouping routes)

[//]: # ()
[//]: # (```go)

[//]: # (func main&#40;&#41; {)

[//]: # (	router := gin.Default&#40;&#41;)

[//]: # ()
[//]: # (	// Simple group: v1)

[//]: # (	v1 := router.Group&#40;"/v1"&#41;)

[//]: # (	{)

[//]: # (		v1.POST&#40;"/login", loginEndpoint&#41;)

[//]: # (		v1.POST&#40;"/submit", submitEndpoint&#41;)

[//]: # (		v1.POST&#40;"/read", readEndpoint&#41;)

[//]: # (	})

[//]: # ()
[//]: # (	// Simple group: v2)

[//]: # (	v2 := router.Group&#40;"/v2"&#41;)

[//]: # (	{)

[//]: # (		v2.POST&#40;"/login", loginEndpoint&#41;)

[//]: # (		v2.POST&#40;"/submit", submitEndpoint&#41;)

[//]: # (		v2.POST&#40;"/read", readEndpoint&#41;)

[//]: # (	})

[//]: # ()
[//]: # (	router.Run&#40;":8080"&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (### Blank Gin without middleware by default)

[//]: # ()
[//]: # (Use)

[//]: # ()
[//]: # (```go)

[//]: # (r := gin.New&#40;&#41;)

[//]: # (```)

[//]: # ()
[//]: # (instead of)

[//]: # ()
[//]: # (```go)

[//]: # (// Default With the Logger and Recovery middleware already attached)

[//]: # (r := gin.Default&#40;&#41;)

[//]: # (```)

[//]: # ()
[//]: # ()
[//]: # (### Using middleware)

[//]: # (```go)

[//]: # (func main&#40;&#41; {)

[//]: # (	// Creates a router without any middleware by default)

[//]: # (	r := gin.New&#40;&#41;)

[//]: # ()
[//]: # (	// Global middleware)

[//]: # (	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.)

[//]: # (	// By default gin.DefaultWriter = os.Stdout)

[//]: # (	r.Use&#40;gin.Logger&#40;&#41;&#41;)

[//]: # ()
[//]: # (	// Recovery middleware recovers from any panics and writes a 500 if there was one.)

[//]: # (	r.Use&#40;gin.Recovery&#40;&#41;&#41;)

[//]: # ()
[//]: # (	// Per route middleware, you can add as many as you desire.)

[//]: # (	r.GET&#40;"/benchmark", MyBenchLogger&#40;&#41;, benchEndpoint&#41;)

[//]: # ()
[//]: # (	// Authorization group)

[//]: # (	// authorized := r.Group&#40;"/", AuthRequired&#40;&#41;&#41;)

[//]: # (	// exactly the same as:)

[//]: # (	authorized := r.Group&#40;"/"&#41;)

[//]: # (	// per group middleware! in this case we use the custom created)

[//]: # (	// AuthRequired&#40;&#41; middleware just in the "authorized" group.)

[//]: # (	authorized.Use&#40;AuthRequired&#40;&#41;&#41;)

[//]: # (	{)

[//]: # (		authorized.POST&#40;"/login", loginEndpoint&#41;)

[//]: # (		authorized.POST&#40;"/submit", submitEndpoint&#41;)

[//]: # (		authorized.POST&#40;"/read", readEndpoint&#41;)

[//]: # ()
[//]: # (		// nested group)

[//]: # (		testing := authorized.Group&#40;"testing"&#41;)

[//]: # (		testing.GET&#40;"/analytics", analyticsEndpoint&#41;)

[//]: # (	})

[//]: # ()
[//]: # (	// Listen and serve on 0.0.0.0:8080)

[//]: # (	r.Run&#40;":8080"&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (### Custom Recovery behavior)

[//]: # (```go)

[//]: # (func main&#40;&#41; {)

[//]: # (	// Creates a router without any middleware by default)

[//]: # (	r := gin.New&#40;&#41;)

[//]: # ()
[//]: # (	// Global middleware)

[//]: # (	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.)

[//]: # (	// By default gin.DefaultWriter = os.Stdout)

[//]: # (	r.Use&#40;gin.Logger&#40;&#41;&#41;)

[//]: # ()
[//]: # (	// Recovery middleware recovers from any panics and writes a 500 if there was one.)

[//]: # (	r.Use&#40;gin.CustomRecovery&#40;func&#40;c *gin.Context, recovered interface{}&#41; {)

[//]: # (		if err, ok := recovered.&#40;string&#41;; ok {)

[//]: # (			c.String&#40;http.StatusInternalServerError, fmt.Sprintf&#40;"error: %s", err&#41;&#41;)

[//]: # (		})

[//]: # (		c.AbortWithStatus&#40;http.StatusInternalServerError&#41;)

[//]: # (	}&#41;&#41;)

[//]: # ()
[//]: # (	r.GET&#40;"/panic", func&#40;c *gin.Context&#41; {)

[//]: # (		// panic with a string -- the custom middleware could save this to a database or report it to the user)

[//]: # (		panic&#40;"foo"&#41;)

[//]: # (	}&#41;)

[//]: # ()
[//]: # (	r.GET&#40;"/", func&#40;c *gin.Context&#41; {)

[//]: # (		c.String&#40;http.StatusOK, "ohai"&#41;)

[//]: # (	}&#41;)

[//]: # ()
[//]: # (	// Listen and serve on 0.0.0.0:8080)

[//]: # (	r.Run&#40;":8080"&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (### How to write log file)

[//]: # (```go)

[//]: # (func main&#40;&#41; {)

[//]: # (    // Disable Console Color, you don't need console color when writing the logs to file.)

[//]: # (    gin.DisableConsoleColor&#40;&#41;)

[//]: # ()
[//]: # (    // Logging to a file.)

[//]: # (    f, _ := os.Create&#40;"gin.log"&#41;)

[//]: # (    gin.DefaultWriter = io.MultiWriter&#40;f&#41;)

[//]: # ()
[//]: # (    // Use the following code if you need to write the logs to file and console at the same time.)

[//]: # (    // gin.DefaultWriter = io.MultiWriter&#40;f, os.Stdout&#41;)

[//]: # ()
[//]: # (    router := gin.Default&#40;&#41;)

[//]: # (    router.GET&#40;"/ping", func&#40;c *gin.Context&#41; {)

[//]: # (        c.String&#40;200, "pong"&#41;)

[//]: # (    }&#41;)

[//]: # ()
[//]: # (    router.Run&#40;":8080"&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (### Custom Log Format)

[//]: # (```go)

[//]: # (func main&#40;&#41; {)

[//]: # (	router := gin.New&#40;&#41;)

[//]: # ()
[//]: # (	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter)

[//]: # (	// By default gin.DefaultWriter = os.Stdout)

[//]: # (	router.Use&#40;gin.LoggerWithFormatter&#40;func&#40;param gin.LogFormatterParams&#41; string {)

[//]: # ()
[//]: # (		// your custom format)

[//]: # (		return fmt.Sprintf&#40;"%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",)

[//]: # (				param.ClientIP,)

[//]: # (				param.TimeStamp.Format&#40;time.RFC1123&#41;,)

[//]: # (				param.Method,)

[//]: # (				param.Path,)

[//]: # (				param.Request.Proto,)

[//]: # (				param.StatusCode,)

[//]: # (				param.Latency,)

[//]: # (				param.Request.UserAgent&#40;&#41;,)

[//]: # (				param.ErrorMessage,)

[//]: # (		&#41;)

[//]: # (	}&#41;&#41;)

[//]: # (	router.Use&#40;gin.Recovery&#40;&#41;&#41;)

[//]: # ()
[//]: # (	router.GET&#40;"/ping", func&#40;c *gin.Context&#41; {)

[//]: # (		c.String&#40;200, "pong"&#41;)

[//]: # (	}&#41;)

[//]: # ()
[//]: # (	router.Run&#40;":8080"&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (**Sample Output**)

[//]: # (```)

[//]: # (::1 - [Fri, 07 Dec 2018 17:04:38 JST] "GET /ping HTTP/1.1 200 122.767µs "Mozilla/5.0 &#40;Macintosh; Intel Mac OS X 10_11_6&#41; AppleWebKit/537.36 &#40;KHTML, like Gecko&#41; Chrome/71.0.3578.80 Safari/537.36" ")

[//]: # (```)

[//]: # ()
[//]: # (### Controlling Log output coloring)

[//]: # ()
[//]: # (By default, logs output on console should be colorized depending on the detected TTY.)

[//]: # ()
[//]: # (Never colorize logs:)

[//]: # ()
[//]: # (```go)

[//]: # (func main&#40;&#41; {)

[//]: # (    // Disable log's color)

[//]: # (    gin.DisableConsoleColor&#40;&#41;)

[//]: # ()
[//]: # (    // Creates a gin router with default middleware:)

[//]: # (    // logger and recovery &#40;crash-free&#41; middleware)

[//]: # (    router := gin.Default&#40;&#41;)

[//]: # ()
[//]: # (    router.GET&#40;"/ping", func&#40;c *gin.Context&#41; {)

[//]: # (        c.String&#40;200, "pong"&#41;)

[//]: # (    }&#41;)

[//]: # ()
[//]: # (    router.Run&#40;":8080"&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (Always colorize logs:)

[//]: # ()
[//]: # (```go)

[//]: # (func main&#40;&#41; {)

[//]: # (    // Force log's color)

[//]: # (    gin.ForceConsoleColor&#40;&#41;)

[//]: # ()
[//]: # (    // Creates a gin router with default middleware:)

[//]: # (    // logger and recovery &#40;crash-free&#41; middleware)

[//]: # (    router := gin.Default&#40;&#41;)

[//]: # ()
[//]: # (    router.GET&#40;"/ping", func&#40;c *gin.Context&#41; {)

[//]: # (        c.String&#40;200, "pong"&#41;)

[//]: # (    }&#41;)

[//]: # ()
[//]: # (    router.Run&#40;":8080"&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (### Model binding and validation)

[//]: # ()
[//]: # (To bind a request body into a type, use model binding. We currently support binding of JSON, XML, YAML and standard form values &#40;foo=bar&boo=baz&#41;.)

[//]: # ()
[//]: # (Gin uses [**go-playground/validator/v10**]&#40;https://github.com/go-playground/validator&#41; for validation. Check the full docs on tags usage [here]&#40;https://godoc.org/github.com/go-playground/validator#hdr-Baked_In_Validators_and_Tags&#41;.)

[//]: # ()
[//]: # (Note that you need to set the corresponding binding tag on all fields you want to bind. For example, when binding from JSON, set `json:"fieldname"`.)

[//]: # ()
[//]: # (Also, Gin provides two sets of methods for binding:)

[//]: # (- **Type** - Must bind)

[//]: # (  - **Methods** - `Bind`, `BindJSON`, `BindXML`, `BindQuery`, `BindYAML`, `BindHeader`)

[//]: # (  - **Behavior** - These methods use `MustBindWith` under the hood. If there is a binding error, the request is aborted with `c.AbortWithError&#40;400, err&#41;.SetType&#40;ErrorTypeBind&#41;`. This sets the response status code to 400 and the `Content-Type` header is set to `text/plain; charset=utf-8`. Note that if you try to set the response code after this, it will result in a warning `[GIN-debug] [WARNING] Headers were already written. Wanted to override status code 400 with 422`. If you wish to have greater control over the behavior, consider using the `ShouldBind` equivalent method.)

[//]: # (- **Type** - Should bind)

[//]: # (  - **Methods** - `ShouldBind`, `ShouldBindJSON`, `ShouldBindXML`, `ShouldBindQuery`, `ShouldBindYAML`, `ShouldBindHeader`)

[//]: # (  - **Behavior** - These methods use `ShouldBindWith` under the hood. If there is a binding error, the error is returned and it is the developer's responsibility to handle the request and error appropriately.)

[//]: # ()
[//]: # (When using the Bind-method, Gin tries to infer the binder depending on the Content-Type header. If you are sure what you are binding, you can use `MustBindWith` or `ShouldBindWith`.)

[//]: # ()
[//]: # (You can also specify that specific fields are required. If a field is decorated with `binding:"required"` and has a empty value when binding, an error will be returned.)

[//]: # ()
[//]: # (```go)

[//]: # (// Binding from JSON)

[//]: # (type Login struct {)

[//]: # (	User     string `form:"user" json:"user" xml:"user"  binding:"required"`)

[//]: # (	Password string `form:"password" json:"password" xml:"password" binding:"required"`)

[//]: # (})

[//]: # ()
[//]: # (func main&#40;&#41; {)

[//]: # (	router := gin.Default&#40;&#41;)

[//]: # ()
[//]: # (	// Example for binding JSON &#40;{"user": "manu", "password": "123"}&#41;)

[//]: # (	router.POST&#40;"/loginJSON", func&#40;c *gin.Context&#41; {)

[//]: # (		var json Login)

[//]: # (		if err := c.ShouldBindJSON&#40;&json&#41;; err != nil {)

[//]: # (			c.JSON&#40;http.StatusBadRequest, gin.H{"error": err.Error&#40;&#41;}&#41;)

[//]: # (			return)

[//]: # (		})

[//]: # ()
[//]: # (		if json.User != "manu" || json.Password != "123" {)

[//]: # (			c.JSON&#40;http.StatusUnauthorized, gin.H{"status": "unauthorized"}&#41;)

[//]: # (			return)

[//]: # (		})

[//]: # ()
[//]: # (		c.JSON&#40;http.StatusOK, gin.H{"status": "you are logged in"}&#41;)

[//]: # (	}&#41;)

[//]: # ()
[//]: # (	// Example for binding XML &#40;)

[//]: # (	//	<?xml version="1.0" encoding="UTF-8"?>)

[//]: # (	//	<root>)

[//]: # (	//		<user>user</user>)

[//]: # (	//		<password>123</password>)

[//]: # (	//	</root>&#41;)

[//]: # (	router.POST&#40;"/loginXML", func&#40;c *gin.Context&#41; {)

[//]: # (		var xml Login)

[//]: # (		if err := c.ShouldBindXML&#40;&xml&#41;; err != nil {)

[//]: # (			c.JSON&#40;http.StatusBadRequest, gin.H{"error": err.Error&#40;&#41;}&#41;)

[//]: # (			return)

[//]: # (		})

[//]: # ()
[//]: # (		if xml.User != "manu" || xml.Password != "123" {)

[//]: # (			c.JSON&#40;http.StatusUnauthorized, gin.H{"status": "unauthorized"}&#41;)

[//]: # (			return)

[//]: # (		})

[//]: # ()
[//]: # (		c.JSON&#40;http.StatusOK, gin.H{"status": "you are logged in"}&#41;)

[//]: # (	}&#41;)

[//]: # ()
[//]: # (	// Example for binding a HTML form &#40;user=manu&password=123&#41;)

[//]: # (	router.POST&#40;"/loginForm", func&#40;c *gin.Context&#41; {)

[//]: # (		var form Login)

[//]: # (		// This will infer what binder to use depending on the content-type header.)

[//]: # (		if err := c.ShouldBind&#40;&form&#41;; err != nil {)

[//]: # (			c.JSON&#40;http.StatusBadRequest, gin.H{"error": err.Error&#40;&#41;}&#41;)

[//]: # (			return)

[//]: # (		})

[//]: # ()
[//]: # (		if form.User != "manu" || form.Password != "123" {)

[//]: # (			c.JSON&#40;http.StatusUnauthorized, gin.H{"status": "unauthorized"}&#41;)

[//]: # (			return)

[//]: # (		})

[//]: # ()
[//]: # (		c.JSON&#40;http.StatusOK, gin.H{"status": "you are logged in"}&#41;)

[//]: # (	}&#41;)

[//]: # ()
[//]: # (	// Listen and serve on 0.0.0.0:8080)

[//]: # (	router.Run&#40;":8080"&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (**Sample request**)

[//]: # (```shell)

[//]: # ($ curl -v -X POST \)

[//]: # (  http://localhost:8080/loginJSON \)

[//]: # (  -H 'content-type: application/json' \)

[//]: # (  -d '{ "user": "manu" }')

[//]: # (> POST /loginJSON HTTP/1.1)

[//]: # (> Host: localhost:8080)

[//]: # (> User-Agent: curl/7.51.0)

[//]: # (> Accept: */*)

[//]: # (> content-type: application/json)

[//]: # (> Content-Length: 18)

[//]: # (>)

[//]: # (* upload completely sent off: 18 out of 18 bytes)

[//]: # (< HTTP/1.1 400 Bad Request)

[//]: # (< Content-Type: application/json; charset=utf-8)

[//]: # (< Date: Fri, 04 Aug 2017 03:51:31 GMT)

[//]: # (< Content-Length: 100)

[//]: # (<)

[//]: # ({"error":"Key: 'Login.Password' Error:Field validation for 'Password' failed on the 'required' tag"})

[//]: # (```)

[//]: # ()
[//]: # (**Skip validate**)

[//]: # ()
[//]: # (When running the above example using the above the `curl` command, it returns error. Because the example use `binding:"required"` for `Password`. If use `binding:"-"` for `Password`, then it will not return error when running the above example again.)

[//]: # ()
[//]: # (### Custom Validators)

[//]: # ()
[//]: # (It is also possible to register custom validators. See the [example code]&#40;https://github.com/gin-gonic/examples/tree/master/custom-validation/server.go&#41;.)

[//]: # ()
[//]: # (```go)

[//]: # (package main)

[//]: # ()
[//]: # (import &#40;)

[//]: # (	"net/http")

[//]: # (	"time")

[//]: # ()
[//]: # (	"github.com/gin-gonic/gin")

[//]: # (	"github.com/gin-gonic/gin/binding")

[//]: # (	"github.com/go-playground/validator/v10")

[//]: # (&#41;)

[//]: # ()
[//]: # (// Booking contains binded and validated data.)

[//]: # (type Booking struct {)

[//]: # (	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`)

[//]: # (	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`)

[//]: # (})

[//]: # ()
[//]: # (var bookableDate validator.Func = func&#40;fl validator.FieldLevel&#41; bool {)

[//]: # (	date, ok := fl.Field&#40;&#41;.Interface&#40;&#41;.&#40;time.Time&#41;)

[//]: # (	if ok {)

[//]: # (		today := time.Now&#40;&#41;)

[//]: # (		if today.After&#40;date&#41; {)

[//]: # (			return false)

[//]: # (		})

[//]: # (	})

[//]: # (	return true)

[//]: # (})

[//]: # ()
[//]: # (func main&#40;&#41; {)

[//]: # (	route := gin.Default&#40;&#41;)

[//]: # ()
[//]: # (	if v, ok := binding.Validator.Engine&#40;&#41;.&#40;*validator.Validate&#41;; ok {)

[//]: # (		v.RegisterValidation&#40;"bookabledate", bookableDate&#41;)

[//]: # (	})

[//]: # ()
[//]: # (	route.GET&#40;"/bookable", getBookable&#41;)

[//]: # (	route.Run&#40;":8085"&#41;)

[//]: # (})

[//]: # ()
[//]: # (func getBookable&#40;c *gin.Context&#41; {)

[//]: # (	var b Booking)

[//]: # (	if err := c.ShouldBindWith&#40;&b, binding.Query&#41;; err == nil {)

[//]: # (		c.JSON&#40;http.StatusOK, gin.H{"message": "Booking dates are valid!"}&#41;)

[//]: # (	} else {)

[//]: # (		c.JSON&#40;http.StatusBadRequest, gin.H{"error": err.Error&#40;&#41;}&#41;)

[//]: # (	})

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (```console)

[//]: # ($ curl "localhost:8085/bookable?check_in=2030-04-16&check_out=2030-04-17")

[//]: # ({"message":"Booking dates are valid!"})

[//]: # ()
[//]: # ($ curl "localhost:8085/bookable?check_in=2030-03-10&check_out=2030-03-09")

[//]: # ({"error":"Key: 'Booking.CheckOut' Error:Field validation for 'CheckOut' failed on the 'gtfield' tag"})

[//]: # ()
[//]: # ($ curl "localhost:8085/bookable?check_in=2000-03-09&check_out=2000-03-10")

[//]: # ({"error":"Key: 'Booking.CheckIn' Error:Field validation for 'CheckIn' failed on the 'bookabledate' tag"}%)

[//]: # (```)

[//]: # ()
[//]: # ([Struct level validations]&#40;https://github.com/go-playground/validator/releases/tag/v8.7&#41; can also be registered this way.)

[//]: # (See the [struct-lvl-validation example]&#40;https://github.com/gin-gonic/examples/tree/master/struct-lvl-validations&#41; to learn more.)

[//]: # ()
[//]: # (### Only Bind Query String)

[//]: # ()
[//]: # (`ShouldBindQuery` function only binds the query params and not the post data. See the [detail information]&#40;https://github.com/gin-gonic/gin/issues/742#issuecomment-315953017&#41;.)

[//]: # ()
[//]: # (```go)

[//]: # (package main)

[//]: # ()
[//]: # (import &#40;)

[//]: # (	"log")

[//]: # ()
[//]: # (	"github.com/gin-gonic/gin")

[//]: # (&#41;)

[//]: # ()
[//]: # (type Person struct {)

[//]: # (	Name    string `form:"name"`)

[//]: # (	Address string `form:"address"`)

[//]: # (})

[//]: # ()
[//]: # (func main&#40;&#41; {)

[//]: # (	route := gin.Default&#40;&#41;)

[//]: # (	route.Any&#40;"/testing", startPage&#41;)

[//]: # (	route.Run&#40;":8085"&#41;)

[//]: # (})

[//]: # ()
[//]: # (func startPage&#40;c *gin.Context&#41; {)

[//]: # (	var person Person)

[//]: # (	if c.ShouldBindQuery&#40;&person&#41; == nil {)

[//]: # (		log.Println&#40;"====== Only Bind By Query String ======"&#41;)

[//]: # (		log.Println&#40;person.Name&#41;)

[//]: # (		log.Println&#40;person.Address&#41;)

[//]: # (	})

[//]: # (	c.String&#40;200, "Success"&#41;)

[//]: # (})

[//]: # ()
[//]: # (```)

[//]: # ()
[//]: # (### Bind Query String or Post Data)

[//]: # ()
[//]: # (See the [detail information]&#40;https://github.com/gin-gonic/gin/issues/742#issuecomment-264681292&#41;.)

[//]: # ()
[//]: # (```go)

[//]: # (package main)

[//]: # ()
[//]: # (import &#40;)

[//]: # (	"log")

[//]: # (	"time")

[//]: # ()
[//]: # (	"github.com/gin-gonic/gin")

[//]: # (&#41;)

[//]: # ()
[//]: # (type Person struct {)

[//]: # (        Name       string    `form:"name"`)

[//]: # (        Address    string    `form:"address"`)

[//]: # (        Birthday   time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`)

[//]: # (        CreateTime time.Time `form:"createTime" time_format:"unixNano"`)

[//]: # (        UnixTime   time.Time `form:"unixTime" time_format:"unix"`)

[//]: # (})

[//]: # ()
[//]: # (func main&#40;&#41; {)

[//]: # (	route := gin.Default&#40;&#41;)

[//]: # (	route.GET&#40;"/testing", startPage&#41;)

[//]: # (	route.Run&#40;":8085"&#41;)

[//]: # (})

[//]: # ()
[//]: # (func startPage&#40;c *gin.Context&#41; {)

[//]: # (	var person Person)

[//]: # (	// If `GET`, only `Form` binding engine &#40;`query`&#41; used.)

[//]: # (	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` &#40;`form-data`&#41;.)

[//]: # (	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48)

[//]: # (        if c.ShouldBind&#40;&person&#41; == nil {)

[//]: # (                log.Println&#40;person.Name&#41;)

[//]: # (                log.Println&#40;person.Address&#41;)

[//]: # (                log.Println&#40;person.Birthday&#41;)

[//]: # (                log.Println&#40;person.CreateTime&#41;)

[//]: # (                log.Println&#40;person.UnixTime&#41;)

[//]: # (        })

[//]: # ()
[//]: # (	c.String&#40;200, "Success"&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (Test it with:)

[//]: # (```sh)

[//]: # ($ curl -X GET "localhost:8085/testing?name=appleboy&address=xyz&birthday=1992-03-15&createTime=1562400033000000123&unixTime=1562400033")

[//]: # (```)

[//]: # ()
[//]: # (### Bind Uri)

[//]: # ()
[//]: # (See the [detail information]&#40;https://github.com/gin-gonic/gin/issues/846&#41;.)

[//]: # ()
[//]: # (```go)

[//]: # (package main)

[//]: # ()
[//]: # (import "github.com/gin-gonic/gin")

[//]: # ()
[//]: # (type Person struct {)

[//]: # (	ID string `uri:"id" binding:"required,uuid"`)

[//]: # (	Name string `uri:"name" binding:"required"`)

[//]: # (})

[//]: # ()
[//]: # (func main&#40;&#41; {)

[//]: # (	route := gin.Default&#40;&#41;)

[//]: # (	route.GET&#40;"/:name/:id", func&#40;c *gin.Context&#41; {)

[//]: # (		var person Person)

[//]: # (		if err := c.ShouldBindUri&#40;&person&#41;; err != nil {)

[//]: # (			c.JSON&#40;400, gin.H{"msg": err}&#41;)

[//]: # (			return)

[//]: # (		})

[//]: # (		c.JSON&#40;200, gin.H{"name": person.Name, "uuid": person.ID}&#41;)

[//]: # (	}&#41;)

[//]: # (	route.Run&#40;":8088"&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (Test it with:)

[//]: # (```sh)

[//]: # ($ curl -v localhost:8088/thinkerou/987fbc97-4bed-5078-9f07-9141ba07c9f3)

[//]: # ($ curl -v localhost:8088/thinkerou/not-uuid)

[//]: # (```)

[//]: # ()
[//]: # (### Bind Header)

[//]: # ()
[//]: # (```go)

[//]: # (package main)

[//]: # ()
[//]: # (import &#40;)

[//]: # (	"fmt")

[//]: # (	"github.com/gin-gonic/gin")

[//]: # (&#41;)

[//]: # ()
[//]: # (type testHeader struct {)

[//]: # (	Rate   int    `header:"Rate"`)

[//]: # (	Domain string `header:"Domain"`)

[//]: # (})

[//]: # ()
[//]: # (func main&#40;&#41; {)

[//]: # (	r := gin.Default&#40;&#41;)

[//]: # (	r.GET&#40;"/", func&#40;c *gin.Context&#41; {)

[//]: # (		h := testHeader{})

[//]: # ()
[//]: # (		if err := c.ShouldBindHeader&#40;&h&#41;; err != nil {)

[//]: # (			c.JSON&#40;200, err&#41;)

[//]: # (		})

[//]: # ()
[//]: # (		fmt.Printf&#40;"%#v\n", h&#41;)

[//]: # (		c.JSON&#40;200, gin.H{"Rate": h.Rate, "Domain": h.Domain}&#41;)

[//]: # (	}&#41;)

[//]: # ()
[//]: # (	r.Run&#40;&#41;)

[//]: # ()
[//]: # (// client)

[//]: # (// curl -H "rate:300" -H "domain:music" 127.0.0.1:8080/)

[//]: # (// output)

[//]: # (// {"Domain":"music","Rate":300})

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (### Bind HTML checkboxes)

[//]: # ()
[//]: # (See the [detail information]&#40;https://github.com/gin-gonic/gin/issues/129#issuecomment-124260092&#41;)

[//]: # ()
[//]: # (main.go)

[//]: # ()
[//]: # (```go)

[//]: # (...)

[//]: # ()
[//]: # (type myForm struct {)

[//]: # (    Colors []string `form:"colors[]"`)

[//]: # (})

[//]: # ()
[//]: # (...)

[//]: # ()
[//]: # (func formHandler&#40;c *gin.Context&#41; {)

[//]: # (    var fakeForm myForm)

[//]: # (    c.ShouldBind&#40;&fakeForm&#41;)

[//]: # (    c.JSON&#40;200, gin.H{"color": fakeForm.Colors}&#41;)

[//]: # (})

[//]: # ()
[//]: # (...)

[//]: # ()
[//]: # (```)

[//]: # ()
[//]: # (form.html)

[//]: # ()
[//]: # (```html)

[//]: # (<form action="/" method="POST">)

[//]: # (    <p>Check some colors</p>)

[//]: # (    <label for="red">Red</label>)

[//]: # (    <input type="checkbox" name="colors[]" value="red" id="red">)

[//]: # (    <label for="green">Green</label>)

[//]: # (    <input type="checkbox" name="colors[]" value="green" id="green">)

[//]: # (    <label for="blue">Blue</label>)

[//]: # (    <input type="checkbox" name="colors[]" value="blue" id="blue">)

[//]: # (    <input type="submit">)

[//]: # (</form>)

[//]: # (```)

[//]: # ()
[//]: # (result:)

[//]: # ()
[//]: # (```)

[//]: # ({"color":["red","green","blue"]})

[//]: # (```)

[//]: # ()
[//]: # (### Multipart/Urlencoded binding)

[//]: # ()
[//]: # (```go)

[//]: # (type ProfileForm struct {)

[//]: # (	Name   string                `form:"name" binding:"required"`)

[//]: # (	Avatar *multipart.FileHeader `form:"avatar" binding:"required"`)

[//]: # ()
[//]: # (	// or for multiple files)

[//]: # (	// Avatars []*multipart.FileHeader `form:"avatar" binding:"required"`)

[//]: # (})

[//]: # ()
[//]: # (func main&#40;&#41; {)

[//]: # (	router := gin.Default&#40;&#41;)

[//]: # (	router.POST&#40;"/profile", func&#40;c *gin.Context&#41; {)

[//]: # (		// you can bind multipart form with explicit binding declaration:)

[//]: # (		// c.ShouldBindWith&#40;&form, binding.Form&#41;)

[//]: # (		// or you can simply use autobinding with ShouldBind method:)

[//]: # (		var form ProfileForm)

[//]: # (		// in this case proper binding will be automatically selected)

[//]: # (		if err := c.ShouldBind&#40;&form&#41;; err != nil {)

[//]: # (			c.String&#40;http.StatusBadRequest, "bad request"&#41;)

[//]: # (			return)

[//]: # (		})

[//]: # ()
[//]: # (		err := c.SaveUploadedFile&#40;form.Avatar, form.Avatar.Filename&#41;)

[//]: # (		if err != nil {)

[//]: # (			c.String&#40;http.StatusInternalServerError, "unknown error"&#41;)

[//]: # (			return)

[//]: # (		})

[//]: # ()
[//]: # (		// db.Save&#40;&form&#41;)

[//]: # ()
[//]: # (		c.String&#40;http.StatusOK, "ok"&#41;)

[//]: # (	}&#41;)

[//]: # (	router.Run&#40;":8080"&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (Test it with:)

[//]: # (```sh)

[//]: # ($ curl -X POST -v --form name=user --form "avatar=@./avatar.png" http://localhost:8080/profile)

[//]: # (```)

[//]: # ()
[//]: # (### XML, JSON, YAML and ProtoBuf rendering)

[//]: # ()
[//]: # (```go)

[//]: # (func main&#40;&#41; {)

[//]: # (	r := gin.Default&#40;&#41;)

[//]: # ()
[//]: # (	// gin.H is a shortcut for map[string]interface{})

[//]: # (	r.GET&#40;"/someJSON", func&#40;c *gin.Context&#41; {)

[//]: # (		c.JSON&#40;http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK}&#41;)

[//]: # (	}&#41;)

[//]: # ()
[//]: # (	r.GET&#40;"/moreJSON", func&#40;c *gin.Context&#41; {)

[//]: # (		// You also can use a struct)

[//]: # (		var msg struct {)

[//]: # (			Name    string `json:"user"`)

[//]: # (			Message string)

[//]: # (			Number  int)

[//]: # (		})

[//]: # (		msg.Name = "Lena")

[//]: # (		msg.Message = "hey")

[//]: # (		msg.Number = 123)

[//]: # (		// Note that msg.Name becomes "user" in the JSON)

[//]: # (		// Will output  :   {"user": "Lena", "Message": "hey", "Number": 123})

[//]: # (		c.JSON&#40;http.StatusOK, msg&#41;)

[//]: # (	}&#41;)

[//]: # ()
[//]: # (	r.GET&#40;"/someXML", func&#40;c *gin.Context&#41; {)

[//]: # (		c.XML&#40;http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK}&#41;)

[//]: # (	}&#41;)

[//]: # ()
[//]: # (	r.GET&#40;"/someYAML", func&#40;c *gin.Context&#41; {)

[//]: # (		c.YAML&#40;http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK}&#41;)

[//]: # (	}&#41;)

[//]: # ()
[//]: # (	r.GET&#40;"/someProtoBuf", func&#40;c *gin.Context&#41; {)

[//]: # (		reps := []int64{int64&#40;1&#41;, int64&#40;2&#41;})

[//]: # (		label := "test")

[//]: # (		// The specific definition of protobuf is written in the testdata/protoexample file.)

[//]: # (		data := &protoexample.Test{)

[//]: # (			Label: &label,)

[//]: # (			Reps:  reps,)

[//]: # (		})

[//]: # (		// Note that data becomes binary data in the response)

[//]: # (		// Will output protoexample.Test protobuf serialized data)

[//]: # (		c.ProtoBuf&#40;http.StatusOK, data&#41;)

[//]: # (	}&#41;)

[//]: # ()
[//]: # (	// Listen and serve on 0.0.0.0:8080)

[//]: # (	r.Run&#40;":8080"&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (#### SecureJSON)

[//]: # ()
[//]: # (Using SecureJSON to prevent json hijacking. Default prepends `"while&#40;1&#41;,"` to response body if the given struct is array values.)

[//]: # ()
[//]: # (```go)

[//]: # (func main&#40;&#41; {)

[//]: # (	r := gin.Default&#40;&#41;)

[//]: # ()
[//]: # (	// You can also use your own secure json prefix)

[//]: # (	// r.SecureJsonPrefix&#40;"&#41;]}',\n"&#41;)

[//]: # ()
[//]: # (	r.GET&#40;"/someJSON", func&#40;c *gin.Context&#41; {)

[//]: # (		names := []string{"lena", "austin", "foo"})

[//]: # ()
[//]: # (		// Will output  :   while&#40;1&#41;;["lena","austin","foo"])

[//]: # (		c.SecureJSON&#40;http.StatusOK, names&#41;)

[//]: # (	}&#41;)

[//]: # ()
[//]: # (	// Listen and serve on 0.0.0.0:8080)

[//]: # (	r.Run&#40;":8080"&#41;)

[//]: # (})

[//]: # (```)

[//]: # (#### JSONP)

[//]: # ()
[//]: # (Using JSONP to request data from a server  in a different domain. Add callback to response body if the query parameter callback exists.)

[//]: # ()
[//]: # (```go)

[//]: # (func main&#40;&#41; {)

[//]: # (	r := gin.Default&#40;&#41;)

[//]: # ()
[//]: # (	r.GET&#40;"/JSONP", func&#40;c *gin.Context&#41; {)

[//]: # (		data := gin.H{)

[//]: # (			"foo": "bar",)

[//]: # (		})

[//]: # ()
[//]: # (		//callback is x)

[//]: # (		// Will output  :   x&#40;{\"foo\":\"bar\"}&#41;)

[//]: # (		c.JSONP&#40;http.StatusOK, data&#41;)

[//]: # (	}&#41;)

[//]: # ()
[//]: # (	// Listen and serve on 0.0.0.0:8080)

[//]: # (	r.Run&#40;":8080"&#41;)

[//]: # ()
[//]: # (        // client)

[//]: # (        // curl http://127.0.0.1:8080/JSONP?callback=x)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (#### AsciiJSON)

[//]: # ()
[//]: # (Using AsciiJSON to Generates ASCII-only JSON with escaped non-ASCII characters.)

[//]: # ()
[//]: # (```go)

[//]: # (func main&#40;&#41; {)

[//]: # (	r := gin.Default&#40;&#41;)

[//]: # ()
[//]: # (	r.GET&#40;"/someJSON", func&#40;c *gin.Context&#41; {)

[//]: # (		data := gin.H{)

[//]: # (			"lang": "GO语言",)

[//]: # (			"tag":  "<br>",)

[//]: # (		})

[//]: # ()
[//]: # (		// will output : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"})

[//]: # (		c.AsciiJSON&#40;http.StatusOK, data&#41;)

[//]: # (	}&#41;)

[//]: # ()
[//]: # (	// Listen and serve on 0.0.0.0:8080)

[//]: # (	r.Run&#40;":8080"&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (#### PureJSON)

[//]: # ()
[//]: # (Normally, JSON replaces special HTML characters with their unicode entities, e.g. `<` becomes  `\u003c`. If you want to encode such characters literally, you can use PureJSON instead.)

[//]: # (This feature is unavailable in Go 1.6 and lower.)

[//]: # ()
[//]: # (```go)

[//]: # (func main&#40;&#41; {)

[//]: # (	r := gin.Default&#40;&#41;)

[//]: # ()
[//]: # (	// Serves unicode entities)

[//]: # (	r.GET&#40;"/json", func&#40;c *gin.Context&#41; {)

[//]: # (		c.JSON&#40;200, gin.H{)

[//]: # (			"html": "<b>Hello, world!</b>",)

[//]: # (		}&#41;)

[//]: # (	}&#41;)

[//]: # ()
[//]: # (	// Serves literal characters)

[//]: # (	r.GET&#40;"/purejson", func&#40;c *gin.Context&#41; {)

[//]: # (		c.PureJSON&#40;200, gin.H{)

[//]: # (			"html": "<b>Hello, world!</b>",)

[//]: # (		}&#41;)

[//]: # (	}&#41;)

[//]: # ()
[//]: # (	// listen and serve on 0.0.0.0:8080)

[//]: # (	r.Run&#40;":8080"&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (### Serving static files)

[//]: # ()
[//]: # (```go)

[//]: # (func main&#40;&#41; {)

[//]: # (	router := gin.Default&#40;&#41;)

[//]: # (	router.Static&#40;"/assets", "./assets"&#41;)

[//]: # (	router.StaticFS&#40;"/more_static", http.Dir&#40;"my_file_system"&#41;&#41;)

[//]: # (	router.StaticFile&#40;"/favicon.ico", "./resources/favicon.ico"&#41;)

[//]: # ()
[//]: # (	// Listen and serve on 0.0.0.0:8080)

[//]: # (	router.Run&#40;":8080"&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (### Serving data from file)

[//]: # ()
[//]: # (```go)

[//]: # (func main&#40;&#41; {)

[//]: # (	router := gin.Default&#40;&#41;)

[//]: # ()
[//]: # (	router.GET&#40;"/local/file", func&#40;c *gin.Context&#41; {)

[//]: # (		c.File&#40;"local/file.go"&#41;)

[//]: # (	}&#41;)

[//]: # ()
[//]: # (	var fs http.FileSystem = // ...)

[//]: # (	router.GET&#40;"/fs/file", func&#40;c *gin.Context&#41; {)

[//]: # (		c.FileFromFS&#40;"fs/file.go", fs&#41;)

[//]: # (	}&#41;)

[//]: # (})

[//]: # ()
[//]: # (```)

[//]: # ()
[//]: # (### Serving data from reader)

[//]: # ()
[//]: # (```go)

[//]: # (func main&#40;&#41; {)

[//]: # (	router := gin.Default&#40;&#41;)

[//]: # (	router.GET&#40;"/someDataFromReader", func&#40;c *gin.Context&#41; {)

[//]: # (		response, err := http.Get&#40;"https://raw.githubusercontent.com/gin-gonic/logo/master/color.png"&#41;)

[//]: # (		if err != nil || response.StatusCode != http.StatusOK {)

[//]: # (			c.Status&#40;http.StatusServiceUnavailable&#41;)

[//]: # (			return)

[//]: # (		})

[//]: # ()
[//]: # (		reader := response.Body)

[//]: # ( 		defer reader.Close&#40;&#41;)

[//]: # (		contentLength := response.ContentLength)

[//]: # (		contentType := response.Header.Get&#40;"Content-Type"&#41;)

[//]: # ()
[//]: # (		extraHeaders := map[string]string{)

[//]: # (			"Content-Disposition": `attachment; filename="gopher.png"`,)

[//]: # (		})

[//]: # ()
[//]: # (		c.DataFromReader&#40;http.StatusOK, contentLength, contentType, reader, extraHeaders&#41;)

[//]: # (	}&#41;)

[//]: # (	router.Run&#40;":8080"&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (### HTML rendering)

[//]: # ()
[//]: # (Using LoadHTMLGlob&#40;&#41; or LoadHTMLFiles&#40;&#41;)

[//]: # ()
[//]: # (```go)

[//]: # (func main&#40;&#41; {)

[//]: # (	router := gin.Default&#40;&#41;)

[//]: # (	router.LoadHTMLGlob&#40;"templates/*"&#41;)

[//]: # (	//router.LoadHTMLFiles&#40;"templates/template1.html", "templates/template2.html"&#41;)

[//]: # (	router.GET&#40;"/index", func&#40;c *gin.Context&#41; {)

[//]: # (		c.HTML&#40;http.StatusOK, "index.tmpl", gin.H{)

[//]: # (			"title": "Main website",)

[//]: # (		}&#41;)

[//]: # (	}&#41;)

[//]: # (	router.Run&#40;":8080"&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (templates/index.tmpl)

[//]: # ()
[//]: # (```html)

[//]: # (<html>)

[//]: # (	<h1>)

[//]: # (		{{ .title }})

[//]: # (	</h1>)

[//]: # (</html>)

[//]: # (```)

[//]: # ()
[//]: # (Using templates with same name in different directories)

[//]: # ()
[//]: # (```go)

[//]: # (func main&#40;&#41; {)

[//]: # (	router := gin.Default&#40;&#41;)

[//]: # (	router.LoadHTMLGlob&#40;"templates/**/*"&#41;)

[//]: # (	router.GET&#40;"/posts/index", func&#40;c *gin.Context&#41; {)

[//]: # (		c.HTML&#40;http.StatusOK, "posts/index.tmpl", gin.H{)

[//]: # (			"title": "Posts",)

[//]: # (		}&#41;)

[//]: # (	}&#41;)

[//]: # (	router.GET&#40;"/users/index", func&#40;c *gin.Context&#41; {)

[//]: # (		c.HTML&#40;http.StatusOK, "users/index.tmpl", gin.H{)

[//]: # (			"title": "Users",)

[//]: # (		}&#41;)

[//]: # (	}&#41;)

[//]: # (	router.Run&#40;":8080"&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (templates/posts/index.tmpl)

[//]: # ()
[//]: # (```html)

[//]: # ({{ define "posts/index.tmpl" }})

[//]: # (<html><h1>)

[//]: # (	{{ .title }})

[//]: # (</h1>)

[//]: # (<p>Using posts/index.tmpl</p>)

[//]: # (</html>)

[//]: # ({{ end }})

[//]: # (```)

[//]: # ()
[//]: # (templates/users/index.tmpl)

[//]: # ()
[//]: # (```html)

[//]: # ({{ define "users/index.tmpl" }})

[//]: # (<html><h1>)

[//]: # (	{{ .title }})

[//]: # (</h1>)

[//]: # (<p>Using users/index.tmpl</p>)

[//]: # (</html>)

[//]: # ({{ end }})

[//]: # (```)

[//]: # ()
[//]: # (#### Custom Template renderer)

[//]: # ()
[//]: # (You can also use your own html template render)

[//]: # ()
[//]: # (```go)

[//]: # (import "html/template")

[//]: # ()
[//]: # (func main&#40;&#41; {)

[//]: # (	router := gin.Default&#40;&#41;)

[//]: # (	html := template.Must&#40;template.ParseFiles&#40;"file1", "file2"&#41;&#41;)

[//]: # (	router.SetHTMLTemplate&#40;html&#41;)

[//]: # (	router.Run&#40;":8080"&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (#### Custom Delimiters)

[//]: # ()
[//]: # (You may use custom delims)

[//]: # ()
[//]: # (```go)

[//]: # (	r := gin.Default&#40;&#41;)

[//]: # (	r.Delims&#40;"{[{", "}]}"&#41;)

[//]: # (	r.LoadHTMLGlob&#40;"/path/to/templates"&#41;)

[//]: # (```)

[//]: # ()
[//]: # (#### Custom Template Funcs)

[//]: # ()
[//]: # (See the detail [example code]&#40;https://github.com/gin-gonic/examples/tree/master/template&#41;.)

[//]: # ()
[//]: # (main.go)

[//]: # ()
[//]: # (```go)

[//]: # (import &#40;)

[//]: # (    "fmt")

[//]: # (    "html/template")

[//]: # (    "net/http")

[//]: # (    "time")

[//]: # ()
[//]: # (    "github.com/gin-gonic/gin")

[//]: # (&#41;)

[//]: # ()
[//]: # (func formatAsDate&#40;t time.Time&#41; string {)

[//]: # (    year, month, day := t.Date&#40;&#41;)

[//]: # (    return fmt.Sprintf&#40;"%d%02d/%02d", year, month, day&#41;)

[//]: # (})

[//]: # ()
[//]: # (func main&#40;&#41; {)

[//]: # (    router := gin.Default&#40;&#41;)

[//]: # (    router.Delims&#40;"{[{", "}]}"&#41;)

[//]: # (    router.SetFuncMap&#40;template.FuncMap{)

[//]: # (        "formatAsDate": formatAsDate,)

[//]: # (    }&#41;)

[//]: # (    router.LoadHTMLFiles&#40;"./testdata/template/raw.tmpl"&#41;)

[//]: # ()
[//]: # (    router.GET&#40;"/raw", func&#40;c *gin.Context&#41; {)

[//]: # (        c.HTML&#40;http.StatusOK, "raw.tmpl", gin.H{)

[//]: # (            "now": time.Date&#40;2017, 07, 01, 0, 0, 0, 0, time.UTC&#41;,)

[//]: # (        }&#41;)

[//]: # (    }&#41;)

[//]: # ()
[//]: # (    router.Run&#40;":8080"&#41;)

[//]: # (})

[//]: # ()
[//]: # (```)

[//]: # ()
[//]: # (raw.tmpl)

[//]: # ()
[//]: # (```html)

[//]: # (Date: {[{.now | formatAsDate}]})

[//]: # (```)

[//]: # ()
[//]: # (Result:)

[//]: # (```)

[//]: # (Date: 2017/07/01)

[//]: # (```)

[//]: # ()
[//]: # (### Multitemplate)

[//]: # ()
[//]: # (Gin allow by default use only one html.Template. Check [a multitemplate render]&#40;https://github.com/gin-contrib/multitemplate&#41; for using features like go 1.6 `block template`.)

[//]: # ()
[//]: # (### Redirects)

[//]: # ()
[//]: # (Issuing a HTTP redirect is easy. Both internal and external locations are supported.)

[//]: # ()
[//]: # (```go)

[//]: # (r.GET&#40;"/test", func&#40;c *gin.Context&#41; {)

[//]: # (	c.Redirect&#40;http.StatusMovedPermanently, "http://www.google.com/"&#41;)

[//]: # (}&#41;)

[//]: # (```)

[//]: # ()
[//]: # (Issuing a HTTP redirect from POST. Refer to issue: [#444]&#40;https://github.com/gin-gonic/gin/issues/444&#41;)

[//]: # (```go)

[//]: # (r.POST&#40;"/test", func&#40;c *gin.Context&#41; {)

[//]: # (	c.Redirect&#40;http.StatusFound, "/foo"&#41;)

[//]: # (}&#41;)

[//]: # (```)

[//]: # ()
[//]: # (Issuing a Router redirect, use `HandleContext` like below.)

[//]: # ()
[//]: # (``` go)

[//]: # (r.GET&#40;"/test", func&#40;c *gin.Context&#41; {)

[//]: # (    c.Request.URL.Path = "/test2")

[//]: # (    r.HandleContext&#40;c&#41;)

[//]: # (}&#41;)

[//]: # (r.GET&#40;"/test2", func&#40;c *gin.Context&#41; {)

[//]: # (    c.JSON&#40;200, gin.H{"hello": "world"}&#41;)

[//]: # (}&#41;)

[//]: # (```)

[//]: # ()
[//]: # ()
[//]: # (### Custom Middleware)

[//]: # ()
[//]: # (```go)

[//]: # (func Logger&#40;&#41; gin.HandlerFunc {)

[//]: # (	return func&#40;c *gin.Context&#41; {)

[//]: # (		t := time.Now&#40;&#41;)

[//]: # ()
[//]: # (		// Set example variable)

[//]: # (		c.Set&#40;"example", "12345"&#41;)

[//]: # ()
[//]: # (		// before request)

[//]: # ()
[//]: # (		c.Next&#40;&#41;)

[//]: # ()
[//]: # (		// after request)

[//]: # (		latency := time.Since&#40;t&#41;)

[//]: # (		log.Print&#40;latency&#41;)

[//]: # ()
[//]: # (		// access the status we are sending)

[//]: # (		status := c.Writer.Status&#40;&#41;)

[//]: # (		log.Println&#40;status&#41;)

[//]: # (	})

[//]: # (})

[//]: # ()
[//]: # (func main&#40;&#41; {)

[//]: # (	r := gin.New&#40;&#41;)

[//]: # (	r.Use&#40;Logger&#40;&#41;&#41;)

[//]: # ()
[//]: # (	r.GET&#40;"/test", func&#40;c *gin.Context&#41; {)

[//]: # (		example := c.MustGet&#40;"example"&#41;.&#40;string&#41;)

[//]: # ()
[//]: # (		// it would print: "12345")

[//]: # (		log.Println&#40;example&#41;)

[//]: # (	}&#41;)

[//]: # ()
[//]: # (	// Listen and serve on 0.0.0.0:8080)

[//]: # (	r.Run&#40;":8080"&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (### Using BasicAuth&#40;&#41; middleware)

[//]: # ()
[//]: # (```go)

[//]: # (// simulate some private data)

[//]: # (var secrets = gin.H{)

[//]: # (	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},)

[//]: # (	"austin": gin.H{"email": "austin@example.com", "phone": "666"},)

[//]: # (	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},)

[//]: # (})

[//]: # ()
[//]: # (func main&#40;&#41; {)

[//]: # (	r := gin.Default&#40;&#41;)

[//]: # ()
[//]: # (	// Group using gin.BasicAuth&#40;&#41; middleware)

[//]: # (	// gin.Accounts is a shortcut for map[string]string)

[//]: # (	authorized := r.Group&#40;"/admin", gin.BasicAuth&#40;gin.Accounts{)

[//]: # (		"foo":    "bar",)

[//]: # (		"austin": "1234",)

[//]: # (		"lena":   "hello2",)

[//]: # (		"manu":   "4321",)

[//]: # (	}&#41;&#41;)

[//]: # ()
[//]: # (	// /admin/secrets endpoint)

[//]: # (	// hit "localhost:8080/admin/secrets)

[//]: # (	authorized.GET&#40;"/secrets", func&#40;c *gin.Context&#41; {)

[//]: # (		// get user, it was set by the BasicAuth middleware)

[//]: # (		user := c.MustGet&#40;gin.AuthUserKey&#41;.&#40;string&#41;)

[//]: # (		if secret, ok := secrets[user]; ok {)

[//]: # (			c.JSON&#40;http.StatusOK, gin.H{"user": user, "secret": secret}&#41;)

[//]: # (		} else {)

[//]: # (			c.JSON&#40;http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :&#40;"}&#41;)

[//]: # (		})

[//]: # (	}&#41;)

[//]: # ()
[//]: # (	// Listen and serve on 0.0.0.0:8080)

[//]: # (	r.Run&#40;":8080"&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (### Goroutines inside a middleware)

[//]: # ()
[//]: # (When starting new Goroutines inside a middleware or handler, you **SHOULD NOT** use the original context inside it, you have to use a read-only copy.)

[//]: # ()
[//]: # (```go)

[//]: # (func main&#40;&#41; {)

[//]: # (	r := gin.Default&#40;&#41;)

[//]: # ()
[//]: # (	r.GET&#40;"/long_async", func&#40;c *gin.Context&#41; {)

[//]: # (		// create copy to be used inside the goroutine)

[//]: # (		cCp := c.Copy&#40;&#41;)

[//]: # (		go func&#40;&#41; {)

[//]: # (			// simulate a long task with time.Sleep&#40;&#41;. 5 seconds)

[//]: # (			time.Sleep&#40;5 * time.Second&#41;)

[//]: # ()
[//]: # (			// note that you are using the copied context "cCp", IMPORTANT)

[//]: # (			log.Println&#40;"Done! in path " + cCp.Request.URL.Path&#41;)

[//]: # (		}&#40;&#41;)

[//]: # (	}&#41;)

[//]: # ()
[//]: # (	r.GET&#40;"/long_sync", func&#40;c *gin.Context&#41; {)

[//]: # (		// simulate a long task with time.Sleep&#40;&#41;. 5 seconds)

[//]: # (		time.Sleep&#40;5 * time.Second&#41;)

[//]: # ()
[//]: # (		// since we are NOT using a goroutine, we do not have to copy the context)

[//]: # (		log.Println&#40;"Done! in path " + c.Request.URL.Path&#41;)

[//]: # (	}&#41;)

[//]: # ()
[//]: # (	// Listen and serve on 0.0.0.0:8080)

[//]: # (	r.Run&#40;":8080"&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (### Custom HTTP configuration)

[//]: # ()
[//]: # (Use `http.ListenAndServe&#40;&#41;` directly, like this:)

[//]: # ()
[//]: # (```go)

[//]: # (func main&#40;&#41; {)

[//]: # (	router := gin.Default&#40;&#41;)

[//]: # (	http.ListenAndServe&#40;":8080", router&#41;)

[//]: # (})

[//]: # (```)

[//]: # (or)

[//]: # ()
[//]: # (```go)

[//]: # (func main&#40;&#41; {)

[//]: # (	router := gin.Default&#40;&#41;)

[//]: # ()
[//]: # (	s := &http.Server{)

[//]: # (		Addr:           ":8080",)

[//]: # (		Handler:        router,)

[//]: # (		ReadTimeout:    10 * time.Second,)

[//]: # (		WriteTimeout:   10 * time.Second,)

[//]: # (		MaxHeaderBytes: 1 << 20,)

[//]: # (	})

[//]: # (	s.ListenAndServe&#40;&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (### Support Let's Encrypt)

[//]: # ()
[//]: # (example for 1-line LetsEncrypt HTTPS servers.)

[//]: # ()
[//]: # (```go)

[//]: # (package main)

[//]: # ()
[//]: # (import &#40;)

[//]: # (	"log")

[//]: # ()
[//]: # (	"github.com/gin-gonic/autotls")

[//]: # (	"github.com/gin-gonic/gin")

[//]: # (&#41;)

[//]: # ()
[//]: # (func main&#40;&#41; {)

[//]: # (	r := gin.Default&#40;&#41;)

[//]: # ()
[//]: # (	// Ping handler)

[//]: # (	r.GET&#40;"/ping", func&#40;c *gin.Context&#41; {)

[//]: # (		c.String&#40;200, "pong"&#41;)

[//]: # (	}&#41;)

[//]: # ()
[//]: # (	log.Fatal&#40;autotls.Run&#40;r, "example1.com", "example2.com"&#41;&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (example for custom autocert manager.)

[//]: # ()
[//]: # (```go)

[//]: # (package main)

[//]: # ()
[//]: # (import &#40;)

[//]: # (	"log")

[//]: # ()
[//]: # (	"github.com/gin-gonic/autotls")

[//]: # (	"github.com/gin-gonic/gin")

[//]: # (	"golang.org/x/crypto/acme/autocert")

[//]: # (&#41;)

[//]: # ()
[//]: # (func main&#40;&#41; {)

[//]: # (	r := gin.Default&#40;&#41;)

[//]: # ()
[//]: # (	// Ping handler)

[//]: # (	r.GET&#40;"/ping", func&#40;c *gin.Context&#41; {)

[//]: # (		c.String&#40;200, "pong"&#41;)

[//]: # (	}&#41;)

[//]: # ()
[//]: # (	m := autocert.Manager{)

[//]: # (		Prompt:     autocert.AcceptTOS,)

[//]: # (		HostPolicy: autocert.HostWhitelist&#40;"example1.com", "example2.com"&#41;,)

[//]: # (		Cache:      autocert.DirCache&#40;"/var/www/.cache"&#41;,)

[//]: # (	})

[//]: # ()
[//]: # (	log.Fatal&#40;autotls.RunWithManager&#40;r, &m&#41;&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (### Run multiple service using Gin)

[//]: # ()
[//]: # (See the [question]&#40;https://github.com/gin-gonic/gin/issues/346&#41; and try the following example:)

[//]: # ()
[//]: # (```go)

[//]: # (package main)

[//]: # ()
[//]: # (import &#40;)

[//]: # (	"log")

[//]: # (	"net/http")

[//]: # (	"time")

[//]: # ()
[//]: # (	"github.com/gin-gonic/gin")

[//]: # (	"golang.org/x/sync/errgroup")

[//]: # (&#41;)

[//]: # ()
[//]: # (var &#40;)

[//]: # (	g errgroup.Group)

[//]: # (&#41;)

[//]: # ()
[//]: # (func router01&#40;&#41; http.Handler {)

[//]: # (	e := gin.New&#40;&#41;)

[//]: # (	e.Use&#40;gin.Recovery&#40;&#41;&#41;)

[//]: # (	e.GET&#40;"/", func&#40;c *gin.Context&#41; {)

[//]: # (		c.JSON&#40;)

[//]: # (			http.StatusOK,)

[//]: # (			gin.H{)

[//]: # (				"code":  http.StatusOK,)

[//]: # (				"error": "Welcome server 01",)

[//]: # (			},)

[//]: # (		&#41;)

[//]: # (	}&#41;)

[//]: # ()
[//]: # (	return e)

[//]: # (})

[//]: # ()
[//]: # (func router02&#40;&#41; http.Handler {)

[//]: # (	e := gin.New&#40;&#41;)

[//]: # (	e.Use&#40;gin.Recovery&#40;&#41;&#41;)

[//]: # (	e.GET&#40;"/", func&#40;c *gin.Context&#41; {)

[//]: # (		c.JSON&#40;)

[//]: # (			http.StatusOK,)

[//]: # (			gin.H{)

[//]: # (				"code":  http.StatusOK,)

[//]: # (				"error": "Welcome server 02",)

[//]: # (			},)

[//]: # (		&#41;)

[//]: # (	}&#41;)

[//]: # ()
[//]: # (	return e)

[//]: # (})

[//]: # ()
[//]: # (func main&#40;&#41; {)

[//]: # (	server01 := &http.Server{)

[//]: # (		Addr:         ":8080",)

[//]: # (		Handler:      router01&#40;&#41;,)

[//]: # (		ReadTimeout:  5 * time.Second,)

[//]: # (		WriteTimeout: 10 * time.Second,)

[//]: # (	})

[//]: # ()
[//]: # (	server02 := &http.Server{)

[//]: # (		Addr:         ":8081",)

[//]: # (		Handler:      router02&#40;&#41;,)

[//]: # (		ReadTimeout:  5 * time.Second,)

[//]: # (		WriteTimeout: 10 * time.Second,)

[//]: # (	})

[//]: # ()
[//]: # (	g.Go&#40;func&#40;&#41; error {)

[//]: # (		err := server01.ListenAndServe&#40;&#41;)

[//]: # (		if err != nil && err != http.ErrServerClosed {)

[//]: # (			log.Fatal&#40;err&#41;)

[//]: # (		})

[//]: # (		return err)

[//]: # (	}&#41;)

[//]: # ()
[//]: # (	g.Go&#40;func&#40;&#41; error {)

[//]: # (		err := server02.ListenAndServe&#40;&#41;)

[//]: # (		if err != nil && err != http.ErrServerClosed {)

[//]: # (			log.Fatal&#40;err&#41;)

[//]: # (		})

[//]: # (		return err)

[//]: # (	}&#41;)

[//]: # ()
[//]: # (	if err := g.Wait&#40;&#41;; err != nil {)

[//]: # (		log.Fatal&#40;err&#41;)

[//]: # (	})

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (### Graceful shutdown or restart)

[//]: # ()
[//]: # (There are a few approaches you can use to perform a graceful shutdown or restart. You can make use of third-party packages specifically built for that, or you can manually do the same with the functions and methods from the built-in packages.)

[//]: # ()
[//]: # (#### Third-party packages)

[//]: # ()
[//]: # (We can use [fvbock/endless]&#40;https://github.com/fvbock/endless&#41; to replace the default `ListenAndServe`. Refer to issue [#296]&#40;https://github.com/gin-gonic/gin/issues/296&#41; for more details.)

[//]: # ()
[//]: # (```go)

[//]: # (router := gin.Default&#40;&#41;)

[//]: # (router.GET&#40;"/", handler&#41;)

[//]: # (// [...])

[//]: # (endless.ListenAndServe&#40;":4242", router&#41;)

[//]: # (```)

[//]: # ()
[//]: # (Alternatives:)

[//]: # ()
[//]: # (* [manners]&#40;https://github.com/braintree/manners&#41;: A polite Go HTTP server that shuts down gracefully.)

[//]: # (* [graceful]&#40;https://github.com/tylerb/graceful&#41;: Graceful is a Go package enabling graceful shutdown of an http.Handler server.)

[//]: # (* [grace]&#40;https://github.com/facebookgo/grace&#41;: Graceful restart & zero downtime deploy for Go servers.)

[//]: # ()
[//]: # (#### Manually)

[//]: # ()
[//]: # (In case you are using Go 1.8 or a later version, you may not need to use those libraries. Consider using `http.Server`'s built-in [Shutdown&#40;&#41;]&#40;https://golang.org/pkg/net/http/#Server.Shutdown&#41; method for graceful shutdowns. The example below describes its usage, and we've got more examples using gin [here]&#40;https://github.com/gin-gonic/examples/tree/master/graceful-shutdown&#41;.)

[//]: # ()
[//]: # (```go)

[//]: # (// +build go1.8)

[//]: # ()
[//]: # (package main)

[//]: # ()
[//]: # (import &#40;)

[//]: # (	"context")

[//]: # (	"log")

[//]: # (	"net/http")

[//]: # (	"os")

[//]: # (	"os/signal")

[//]: # (	"syscall")

[//]: # (	"time")

[//]: # ()
[//]: # (	"github.com/gin-gonic/gin")

[//]: # (&#41;)

[//]: # ()
[//]: # (func main&#40;&#41; {)

[//]: # (	router := gin.Default&#40;&#41;)

[//]: # (	router.GET&#40;"/", func&#40;c *gin.Context&#41; {)

[//]: # (		time.Sleep&#40;5 * time.Second&#41;)

[//]: # (		c.String&#40;http.StatusOK, "Welcome Gin Server"&#41;)

[//]: # (	}&#41;)

[//]: # ()
[//]: # (	srv := &http.Server{)

[//]: # (		Addr:    ":8080",)

[//]: # (		Handler: router,)

[//]: # (	})

[//]: # ()
[//]: # (	// Initializing the server in a goroutine so that)

[//]: # (	// it won't block the graceful shutdown handling below)

[//]: # (	go func&#40;&#41; {)

[//]: # (		if err := srv.ListenAndServe&#40;&#41;; err != nil && errors.Is&#40;err, http.ErrServerClosed&#41; {)

[//]: # (			log.Printf&#40;"listen: %s\n", err&#41;)

[//]: # (		})

[//]: # (	}&#40;&#41;)

[//]: # ()
[//]: # (	// Wait for interrupt signal to gracefully shutdown the server with)

[//]: # (	// a timeout of 5 seconds.)

[//]: # (	quit := make&#40;chan os.Signal&#41;)

[//]: # (	// kill &#40;no param&#41; default send syscall.SIGTERM)

[//]: # (	// kill -2 is syscall.SIGINT)

[//]: # (	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it)

[//]: # (	signal.Notify&#40;quit, syscall.SIGINT, syscall.SIGTERM&#41;)

[//]: # (	<-quit)

[//]: # (	log.Println&#40;"Shutting down server..."&#41;)

[//]: # ()
[//]: # (	// The context is used to inform the server it has 5 seconds to finish)

[//]: # (	// the request it is currently handling)

[//]: # (	ctx, cancel := context.WithTimeout&#40;context.Background&#40;&#41;, 5*time.Second&#41;)

[//]: # (	defer cancel&#40;&#41;)

[//]: # ()
[//]: # (	if err := srv.Shutdown&#40;ctx&#41;; err != nil {)

[//]: # (		log.Fatal&#40;"Server forced to shutdown:", err&#41;)

[//]: # (	})

[//]: # ()
[//]: # (	log.Println&#40;"Server exiting"&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (### Build a single binary with templates)

[//]: # ()
[//]: # (You can build a server into a single binary containing templates by using [go-assets][].)

[//]: # ()
[//]: # ([go-assets]: https://github.com/jessevdk/go-assets)

[//]: # ()
[//]: # (```go)

[//]: # (func main&#40;&#41; {)

[//]: # (	r := gin.New&#40;&#41;)

[//]: # ()
[//]: # (	t, err := loadTemplate&#40;&#41;)

[//]: # (	if err != nil {)

[//]: # (		panic&#40;err&#41;)

[//]: # (	})

[//]: # (	r.SetHTMLTemplate&#40;t&#41;)

[//]: # ()
[//]: # (	r.GET&#40;"/", func&#40;c *gin.Context&#41; {)

[//]: # (		c.HTML&#40;http.StatusOK, "/html/index.tmpl",nil&#41;)

[//]: # (	}&#41;)

[//]: # (	r.Run&#40;":8080"&#41;)

[//]: # (})

[//]: # ()
[//]: # (// loadTemplate loads templates embedded by go-assets-builder)

[//]: # (func loadTemplate&#40;&#41; &#40;*template.Template, error&#41; {)

[//]: # (	t := template.New&#40;""&#41;)

[//]: # (	for name, file := range Assets.Files {)

[//]: # (		defer file.Close&#40;&#41;)

[//]: # (		if file.IsDir&#40;&#41; || !strings.HasSuffix&#40;name, ".tmpl"&#41; {)

[//]: # (			continue)

[//]: # (		})

[//]: # (		h, err := ioutil.ReadAll&#40;file&#41;)

[//]: # (		if err != nil {)

[//]: # (			return nil, err)

[//]: # (		})

[//]: # (		t, err = t.New&#40;name&#41;.Parse&#40;string&#40;h&#41;&#41;)

[//]: # (		if err != nil {)

[//]: # (			return nil, err)

[//]: # (		})

[//]: # (	})

[//]: # (	return t, nil)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (See a complete example in the `https://github.com/gin-gonic/examples/tree/master/assets-in-binary` directory.)

[//]: # ()
[//]: # (### Bind form-data request with custom struct)

[//]: # ()
[//]: # (The follow example using custom struct:)

[//]: # ()
[//]: # (```go)

[//]: # (type StructA struct {)

[//]: # (    FieldA string `form:"field_a"`)

[//]: # (})

[//]: # ()
[//]: # (type StructB struct {)

[//]: # (    NestedStruct StructA)

[//]: # (    FieldB string `form:"field_b"`)

[//]: # (})

[//]: # ()
[//]: # (type StructC struct {)

[//]: # (    NestedStructPointer *StructA)

[//]: # (    FieldC string `form:"field_c"`)

[//]: # (})

[//]: # ()
[//]: # (type StructD struct {)

[//]: # (    NestedAnonyStruct struct {)

[//]: # (        FieldX string `form:"field_x"`)

[//]: # (    })

[//]: # (    FieldD string `form:"field_d"`)

[//]: # (})

[//]: # ()
[//]: # (func GetDataB&#40;c *gin.Context&#41; {)

[//]: # (    var b StructB)

[//]: # (    c.Bind&#40;&b&#41;)

[//]: # (    c.JSON&#40;200, gin.H{)

[//]: # (        "a": b.NestedStruct,)

[//]: # (        "b": b.FieldB,)

[//]: # (    }&#41;)

[//]: # (})

[//]: # ()
[//]: # (func GetDataC&#40;c *gin.Context&#41; {)

[//]: # (    var b StructC)

[//]: # (    c.Bind&#40;&b&#41;)

[//]: # (    c.JSON&#40;200, gin.H{)

[//]: # (        "a": b.NestedStructPointer,)

[//]: # (        "c": b.FieldC,)

[//]: # (    }&#41;)

[//]: # (})

[//]: # ()
[//]: # (func GetDataD&#40;c *gin.Context&#41; {)

[//]: # (    var b StructD)

[//]: # (    c.Bind&#40;&b&#41;)

[//]: # (    c.JSON&#40;200, gin.H{)

[//]: # (        "x": b.NestedAnonyStruct,)

[//]: # (        "d": b.FieldD,)

[//]: # (    }&#41;)

[//]: # (})

[//]: # ()
[//]: # (func main&#40;&#41; {)

[//]: # (    r := gin.Default&#40;&#41;)

[//]: # (    r.GET&#40;"/getb", GetDataB&#41;)

[//]: # (    r.GET&#40;"/getc", GetDataC&#41;)

[//]: # (    r.GET&#40;"/getd", GetDataD&#41;)

[//]: # ()
[//]: # (    r.Run&#40;&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (Using the command `curl` command result:)

[//]: # ()
[//]: # (```)

[//]: # ($ curl "http://localhost:8080/getb?field_a=hello&field_b=world")

[//]: # ({"a":{"FieldA":"hello"},"b":"world"})

[//]: # ($ curl "http://localhost:8080/getc?field_a=hello&field_c=world")

[//]: # ({"a":{"FieldA":"hello"},"c":"world"})

[//]: # ($ curl "http://localhost:8080/getd?field_x=hello&field_d=world")

[//]: # ({"d":"world","x":{"FieldX":"hello"}})

[//]: # (```)

[//]: # ()
[//]: # (### Try to bind body into different structs)

[//]: # ()
[//]: # (The normal methods for binding request body consumes `c.Request.Body` and they)

[//]: # (cannot be called multiple times.)

[//]: # ()
[//]: # (```go)

[//]: # (type formA struct {)

[//]: # (  Foo string `json:"foo" xml:"foo" binding:"required"`)

[//]: # (})

[//]: # ()
[//]: # (type formB struct {)

[//]: # (  Bar string `json:"bar" xml:"bar" binding:"required"`)

[//]: # (})

[//]: # ()
[//]: # (func SomeHandler&#40;c *gin.Context&#41; {)

[//]: # (  objA := formA{})

[//]: # (  objB := formB{})

[//]: # (  // This c.ShouldBind consumes c.Request.Body and it cannot be reused.)

[//]: # (  if errA := c.ShouldBind&#40;&objA&#41;; errA == nil {)

[//]: # (    c.String&#40;http.StatusOK, `the body should be formA`&#41;)

[//]: # (  // Always an error is occurred by this because c.Request.Body is EOF now.)

[//]: # (  } else if errB := c.ShouldBind&#40;&objB&#41;; errB == nil {)

[//]: # (    c.String&#40;http.StatusOK, `the body should be formB`&#41;)

[//]: # (  } else {)

[//]: # (    ...)

[//]: # (  })

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (For this, you can use `c.ShouldBindBodyWith`.)

[//]: # ()
[//]: # (```go)

[//]: # (func SomeHandler&#40;c *gin.Context&#41; {)

[//]: # (  objA := formA{})

[//]: # (  objB := formB{})

[//]: # (  // This reads c.Request.Body and stores the result into the context.)

[//]: # (  if errA := c.ShouldBindBodyWith&#40;&objA, binding.JSON&#41;; errA == nil {)

[//]: # (    c.String&#40;http.StatusOK, `the body should be formA`&#41;)

[//]: # (  // At this time, it reuses body stored in the context.)

[//]: # (  } else if errB := c.ShouldBindBodyWith&#40;&objB, binding.JSON&#41;; errB == nil {)

[//]: # (    c.String&#40;http.StatusOK, `the body should be formB JSON`&#41;)

[//]: # (  // And it can accepts other formats)

[//]: # (  } else if errB2 := c.ShouldBindBodyWith&#40;&objB, binding.XML&#41;; errB2 == nil {)

[//]: # (    c.String&#40;http.StatusOK, `the body should be formB XML`&#41;)

[//]: # (  } else {)

[//]: # (    ...)

[//]: # (  })

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (* `c.ShouldBindBodyWith` stores body into the context before binding. This has)

[//]: # (a slight impact to performance, so you should not use this method if you are)

[//]: # (enough to call binding at once.)

[//]: # (* This feature is only needed for some formats -- `JSON`, `XML`, `MsgPack`,)

[//]: # (`ProtoBuf`. For other formats, `Query`, `Form`, `FormPost`, `FormMultipart`,)

[//]: # (can be called by `c.ShouldBind&#40;&#41;` multiple times without any damage to)

[//]: # (performance &#40;See [#1341]&#40;https://github.com/gin-gonic/gin/pull/1341&#41;&#41;.)

[//]: # ()
[//]: # (### http2 server push)

[//]: # ()
[//]: # (http.Pusher is supported only **go1.8+**. See the [golang blog]&#40;https://blog.golang.org/h2push&#41; for detail information.)

[//]: # ()
[//]: # (```go)

[//]: # (package main)

[//]: # ()
[//]: # (import &#40;)

[//]: # (	"html/template")

[//]: # (	"log")

[//]: # ()
[//]: # (	"github.com/gin-gonic/gin")

[//]: # (&#41;)

[//]: # ()
[//]: # (var html = template.Must&#40;template.New&#40;"https"&#41;.Parse&#40;`)

[//]: # (<html>)

[//]: # (<head>)

[//]: # (  <title>Https Test</title>)

[//]: # (  <script src="/assets/app.js"></script>)

[//]: # (</head>)

[//]: # (<body>)

[//]: # (  <h1 style="color:red;">Welcome, Ginner!</h1>)

[//]: # (</body>)

[//]: # (</html>)

[//]: # (`&#41;&#41;)

[//]: # ()
[//]: # (func main&#40;&#41; {)

[//]: # (	r := gin.Default&#40;&#41;)

[//]: # (	r.Static&#40;"/assets", "./assets"&#41;)

[//]: # (	r.SetHTMLTemplate&#40;html&#41;)

[//]: # ()
[//]: # (	r.GET&#40;"/", func&#40;c *gin.Context&#41; {)

[//]: # (		if pusher := c.Writer.Pusher&#40;&#41;; pusher != nil {)

[//]: # (			// use pusher.Push&#40;&#41; to do server push)

[//]: # (			if err := pusher.Push&#40;"/assets/app.js", nil&#41;; err != nil {)

[//]: # (				log.Printf&#40;"Failed to push: %v", err&#41;)

[//]: # (			})

[//]: # (		})

[//]: # (		c.HTML&#40;200, "https", gin.H{)

[//]: # (			"status": "success",)

[//]: # (		}&#41;)

[//]: # (	}&#41;)

[//]: # ()
[//]: # (	// Listen and Server in https://127.0.0.1:8080)

[//]: # (	r.RunTLS&#40;":8080", "./testdata/server.pem", "./testdata/server.key"&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (### Define format for the log of routes)

[//]: # ()
[//]: # (The default log of routes is:)

[//]: # (```)

[//]: # ([GIN-debug] POST   /foo                      --> main.main.func1 &#40;3 handlers&#41;)

[//]: # ([GIN-debug] GET    /bar                      --> main.main.func2 &#40;3 handlers&#41;)

[//]: # ([GIN-debug] GET    /status                   --> main.main.func3 &#40;3 handlers&#41;)

[//]: # (```)

[//]: # ()
[//]: # (If you want to log this information in given format &#40;e.g. JSON, key values or something else&#41;, then you can define this format with `gin.DebugPrintRouteFunc`.)

[//]: # (In the example below, we log all routes with standard log package but you can use another log tools that suits of your needs.)

[//]: # (```go)

[//]: # (import &#40;)

[//]: # (	"log")

[//]: # (	"net/http")

[//]: # ()
[//]: # (	"github.com/gin-gonic/gin")

[//]: # (&#41;)

[//]: # ()
[//]: # (func main&#40;&#41; {)

[//]: # (	r := gin.Default&#40;&#41;)

[//]: # (	gin.DebugPrintRouteFunc = func&#40;httpMethod, absolutePath, handlerName string, nuHandlers int&#41; {)

[//]: # (		log.Printf&#40;"endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers&#41;)

[//]: # (	})

[//]: # ()
[//]: # (	r.POST&#40;"/foo", func&#40;c *gin.Context&#41; {)

[//]: # (		c.JSON&#40;http.StatusOK, "foo"&#41;)

[//]: # (	}&#41;)

[//]: # ()
[//]: # (	r.GET&#40;"/bar", func&#40;c *gin.Context&#41; {)

[//]: # (		c.JSON&#40;http.StatusOK, "bar"&#41;)

[//]: # (	}&#41;)

[//]: # ()
[//]: # (	r.GET&#40;"/status", func&#40;c *gin.Context&#41; {)

[//]: # (		c.JSON&#40;http.StatusOK, "ok"&#41;)

[//]: # (	}&#41;)

[//]: # ()
[//]: # (	// Listen and Server in http://0.0.0.0:8080)

[//]: # (	r.Run&#40;&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (### Set and get a cookie)

[//]: # ()
[//]: # (```go)

[//]: # (import &#40;)

[//]: # (    "fmt")

[//]: # ()
[//]: # (    "github.com/gin-gonic/gin")

[//]: # (&#41;)

[//]: # ()
[//]: # (func main&#40;&#41; {)

[//]: # ()
[//]: # (    router := gin.Default&#40;&#41;)

[//]: # ()
[//]: # (    router.GET&#40;"/cookie", func&#40;c *gin.Context&#41; {)

[//]: # ()
[//]: # (        cookie, err := c.Cookie&#40;"gin_cookie"&#41;)

[//]: # ()
[//]: # (        if err != nil {)

[//]: # (            cookie = "NotSet")

[//]: # (            c.SetCookie&#40;"gin_cookie", "test", 3600, "/", "localhost", false, true&#41;)

[//]: # (        })

[//]: # ()
[//]: # (        fmt.Printf&#40;"Cookie value: %s \n", cookie&#41;)

[//]: # (    }&#41;)

[//]: # ()
[//]: # (    router.Run&#40;&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (## Don't trust all proxies)

[//]: # ()
[//]: # (Gin lets you specify which headers to hold the real client IP &#40;if any&#41;,)

[//]: # (as well as specifying which proxies &#40;or direct clients&#41; you trust to)

[//]: # (specify one of these headers.)

[//]: # ()
[//]: # (The `TrustedProxies` slice on your `gin.Engine` specifes network addresses or)

[//]: # (network CIDRs from where clients which their request headers related to client)

[//]: # (IP can be trusted. They can be IPv4 addresses, IPv4 CIDRs, IPv6 addresses or)

[//]: # (IPv6 CIDRs.)

[//]: # ()
[//]: # (```go)

[//]: # (import &#40;)

[//]: # (	"fmt")

[//]: # ()
[//]: # (	"github.com/gin-gonic/gin")

[//]: # (&#41;)

[//]: # ()
[//]: # (func main&#40;&#41; {)

[//]: # ()
[//]: # (	router := gin.Default&#40;&#41;)

[//]: # (	router.TrustedProxies = []string{"192.168.1.2"})

[//]: # ()
[//]: # (	router.GET&#40;"/", func&#40;c *gin.Context&#41; {)

[//]: # (		// If the client is 192.168.1.2, use the X-Forwarded-For)

[//]: # (		// header to deduce the original client IP from the trust-)

[//]: # (		// worthy parts of that header.)

[//]: # (		// Otherwise, simply return the direct client IP)

[//]: # (		fmt.Printf&#40;"ClientIP: %s\n", c.ClientIP&#40;&#41;&#41;)

[//]: # (	}&#41;)

[//]: # (	router.Run&#40;&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (## Testing)

[//]: # ()
[//]: # (The `net/http/httptest` package is preferable way for HTTP testing.)

[//]: # ()
[//]: # (```go)

[//]: # (package main)

[//]: # ()
[//]: # (func setupRouter&#40;&#41; *gin.Engine {)

[//]: # (	r := gin.Default&#40;&#41;)

[//]: # (	r.GET&#40;"/ping", func&#40;c *gin.Context&#41; {)

[//]: # (		c.String&#40;200, "pong"&#41;)

[//]: # (	}&#41;)

[//]: # (	return r)

[//]: # (})

[//]: # ()
[//]: # (func main&#40;&#41; {)

[//]: # (	r := setupRouter&#40;&#41;)

[//]: # (	r.Run&#40;":8080"&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (Test for code example above:)

[//]: # ()
[//]: # (```go)

[//]: # (package main)

[//]: # ()
[//]: # (import &#40;)

[//]: # (	"net/http")

[//]: # (	"net/http/httptest")

[//]: # (	"testing")

[//]: # ()
[//]: # (	"github.com/stretchr/testify/assert")

[//]: # (&#41;)

[//]: # ()
[//]: # (func TestPingRoute&#40;t *testing.T&#41; {)

[//]: # (	router := setupRouter&#40;&#41;)

[//]: # ()
[//]: # (	w := httptest.NewRecorder&#40;&#41;)

[//]: # (	req, _ := http.NewRequest&#40;"GET", "/ping", nil&#41;)

[//]: # (	router.ServeHTTP&#40;w, req&#41;)

[//]: # ()
[//]: # (	assert.Equal&#40;t, 200, w.Code&#41;)

[//]: # (	assert.Equal&#40;t, "pong", w.Body.String&#40;&#41;&#41;)

[//]: # (})

[//]: # (```)

[//]: # ()
[//]: # (## Users)

[//]: # ()
[//]: # (Awesome project lists using [Gin]&#40;https://github.com/gin-gonic/gin&#41; web framework.)

[//]: # ()
[//]: # (* [gorush]&#40;https://github.com/appleboy/gorush&#41;: A push notification server written in Go.)

[//]: # (* [fnproject]&#40;https://github.com/fnproject/fn&#41;: The container native, cloud agnostic serverless platform.)

[//]: # (* [photoprism]&#40;https://github.com/photoprism/photoprism&#41;: Personal photo management powered by Go and Google TensorFlow.)

[//]: # (* [krakend]&#40;https://github.com/devopsfaith/krakend&#41;: Ultra performant API Gateway with middlewares.)

[//]: # (* [picfit]&#40;https://github.com/thoas/picfit&#41;: An image resizing server written in Go.)

[//]: # (* [brigade]&#40;https://github.com/brigadecore/brigade&#41;: Event-based Scripting for Kubernetes.)

[//]: # (* [dkron]&#40;https://github.com/distribworks/dkron&#41;: Distributed, fault tolerant job scheduling system.)

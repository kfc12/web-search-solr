package main

import (
    "github.com/kataras/iris"

    "github.com/kataras/iris/middleware/logger"
    "github.com/kataras/iris/middleware/recover"
    //"github.com/wirelessregistry/gora"
    //"github.com/vanng822/go-solr/solr"
    //"fmt"
)

func main() {
    app := iris.New()
    // Optionally, add two built'n handlers
    // that can recover from any http-relative panics
    // and log the requests to the terminal.
    app.Use(recover.New())
    app.Use(logger.New())

    // Method:   GET
    // Resource: http://localhost:8080/
    app.Handle("GET", "/", func(ctx iris.Context) {
        ctx.HTML("<b>Welcome!</b>")
    })

    app.StaticWeb("/static", "./web")

    // same as app.Handle("GET", "/ping", [...])
    // Method:   GET
    // Resource: http://localhost:8080/ping
    app.Get("/ping", func(ctx iris.Context) {
        ctx.WriteString("pong")
    })

    // Method:   GET
    // Resource: http://localhost:8080/hello
    app.Get("/hello", func(ctx iris.Context) {
        ctx.JSON(iris.Map{"message": "Hello iris web framework."})
    })

    // http://localhost:8080
    // http://localhost:8080/ping
    // http://localhost:8080/hello
    app.Run(iris.Addr(":6060"), iris.WithoutServerError(iris.ErrServerClosed), iris.WithoutVersionChecker)
    print("running")
}
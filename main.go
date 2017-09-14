package main

import (
	"github.com/kataras/iris"

	"github.com/kataras/iris/websocket"
	"wandering-server/server"
)

func main() {
	app := iris.New()

	app.Use(iris.Gzip)
	setupStaticServer(app)
	setupWebsocket(app)

	app.Run(iris.Addr(":8080"))
}

func setupStaticServer(app *iris.Application) {
	//app.Get("/", func(ctx handler.Context) {
	//	ctx.ServeFile("./static/index.html", false) // second parameter: enable gzip?
	//})
	app.StaticWeb("/", "./static")
	assetHandler := app.StaticHandler("./static", false, true)
	app.SPA(assetHandler)
}

func setupWebsocket(app *iris.Application) {
	// create our echo websocket server
	ws := websocket.New(websocket.Config{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	})
	ws.OnConnection(server.HandleConnection)

	// register the server on an endpoint.
	// see the inline javascript code in the websockets.html, this endpoint is used to connect to the server.
	app.Get("/echo", ws.Handler())
}

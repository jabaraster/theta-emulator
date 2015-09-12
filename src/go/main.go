package main

import (
	"./env"
	"./web/handler"
	"github.com/jabaraster/webtool"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
)

const (
	ASSET_ROOT = "./assets"
)


func main() {
	env.Dump()

	// modelMuxやhtmlMuxには認証ミドルウェアを仕込む必要がある.

	resourceMux := web.New()
//	resourceMux.Use(middleware.AjaxAuthenticator)
	resourceMux.Get ("/resource/property/", handler.GetAllPropertiesHandler)

	staticMux := web.New()
	staticMux.Get("/css/*", webtool.GetAssetsHandlerWithContentType("text/css", ASSET_ROOT))
	staticMux.Get("/js/*", webtool.GetAssetsHandlerWithContentType("text/javascript", ASSET_ROOT))
	staticMux.Get("/*", http.FileServer(http.Dir(ASSET_ROOT)))

	htmlMux := web.New()
//	htmlMux.Use(middleware.PageAuthenticator)
	htmlMux.Get("/page/:page/", webtool.GetHtmlHandler(ASSET_ROOT+"/html", ASSET_ROOT))
	htmlMux.Get("/", webtool.GetHtmlPathHandler(ASSET_ROOT+"/html/index.html", ASSET_ROOT))

	// 各MuxをURLに割り当てる
	// MuxでもURLが登場するので、冗長と言えば冗長.
	// 指定の順番は要注意.
	// 先に割り当てた方が先にマッチするので
	// 競合する指定は優先させたいMuxを先に記述する必要がある.
	defaultMux := goji.DefaultMux
	defaultMux.Handle("/", htmlMux)
	defaultMux.Handle("/resource/*", resourceMux)

	defaultMux.Handle("/css/*", staticMux)
	defaultMux.Handle("/js/*", staticMux)

	defaultMux.Handle("/page/*", htmlMux)

	defaultMux.Handle("/*", staticMux)

	goji.Serve()
}

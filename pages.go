package main

import (
	"log"

	"github.com/Hatch1fy/httpserve"
	pages "github.com/ShibuiCMS/pages/lib"
	"github.com/vroomy/plugins"
)

var p *pages.Pages

func init() {
	var err error
	if p, err = pages.New("./data"); err != nil {
		log.Fatal(err)
	}
}

// Backend returns the underlying backend to the plugin
func Backend() interface{} {
	return p
}

// OnInit will be called by Vroomie on initialization
func OnInit(p *plugins.Plugins, flags, env map[string]string) (err error) {
	return
}

// Close will close the underlying back-end
func Close() (err error) {
	return p.Close()
}

// New will create a new page
func New(ctx *httpserve.Context) httpserve.Response {
	var (
		e   pages.Entry
		err error
	)

	if err = ctx.BindJSON(&e); err != nil {
		return httpserve.NewJSONResponse(400, err)
	}

	if e.Key, err = p.New(e.Name, e.Template, e.Data); err != nil {
		return httpserve.NewJSONResponse(400, err)
	}

	return httpserve.NewJSONResponse(200, e.Key)
}

// EditData will update the data for an existing page
func EditData(ctx *httpserve.Context) httpserve.Response {
	var (
		d   pages.Data
		err error
	)

	key := ctx.Param("key")
	if err = ctx.BindJSON(&d); err != nil {
		return httpserve.NewJSONResponse(400, err)
	}

	if err = p.EditData(key, d); err != nil {
		return httpserve.NewJSONResponse(400, err)
	}

	return httpserve.NewNoContentResponse()
}

// EditTemplate will update the template for an existing page
func EditTemplate(ctx *httpserve.Context) httpserve.Response {
	var (
		e   pages.Entry
		err error
	)

	key := ctx.Param("key")
	if err = ctx.BindJSON(&e); err != nil {
		return httpserve.NewJSONResponse(400, err)
	}

	if err = p.EditTemplate(key, e.Template); err != nil {
		return httpserve.NewJSONResponse(400, err)
	}

	return httpserve.NewNoContentResponse()
}

// Get will retrieve an existing page
func Get(ctx *httpserve.Context) httpserve.Response {
	var (
		e   *pages.Entry
		err error
	)

	key := ctx.Param("key")

	if e, err = p.Get(key); err != nil {
		return httpserve.NewJSONResponse(400, err)
	}

	return httpserve.NewJSONResponse(200, e)
}

// Remove will delete an existing page (and it's history)
func Remove(ctx *httpserve.Context) httpserve.Response {
	var err error
	key := ctx.Param("key")

	if err = p.Remove(key); err != nil {
		return httpserve.NewJSONResponse(400, err)
	}

	return httpserve.NewNoContentResponse()
}

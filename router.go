package nego

import (
	"net/http"
	"sync"
)

type Handle func(http.ResponseWriter, *http.Request, Params)

type Router struct {
	trees map[string]*node
	paramsPool sync.Pool
	maxParams uint16
}

func (me *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	method := req.Method
	root := me.trees[method]

	if nil == root {
		http.NotFound(res, req)
		return
	}

	path := req.URL.Path

	handle, ps, _ := root.getValue(path, me.getParams)

	if nil == handle {
		http.NotFound(res, req)
		return
	}

	if nil == ps {
		handle(res, req, nil)
	} else {
		handle(res, req, *ps)
		me.putParams(ps)
	}
}

func (me *Router) getParams() *Params {
	ps, _ := me.paramsPool.Get().(*Params)
	*ps = (*ps)[0:0]
	return ps
}

func (me *Router) putParams(ps *Params) {
	if nil != ps {
		me.paramsPool.Put(ps)
	}
}

func (me *Router) GET(path string, mw Handle) {
	me.handle(http.MethodGet, path, mw)
}

func (me *Router) handle(method string, path string, mw Handle) {
	if nil == mw {
		panic("handler of path[" + path + "] should not be empty")
	}

	if nil == me.trees {
		me.trees = make(map[string]*node)
	}

	root := me.trees[method]

	if nil == root {
		root = new(node)
		me.trees[method] = root
	}

	root.addRoute(path, mw)
}

func New() *Router {
	return &Router{
	}
}

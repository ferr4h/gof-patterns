package structural

type Server interface {
	HandleRequest(url, method string) (int, string)
}

type Application struct {
}

func (app *Application) HandleRequest(url, method string) (int, string) {
	if url != "/user" {
		return 404, "Not found"
	}
	if method != "GET" {
		return 405, "Method not allowed"
	}
	return 200, "OK"
}

type Proxy struct {
	app              *Application
	maxRequestPerUrl int
	requestPerUrl    map[string]int
}

func NewProxy(app *Application, maxRequestPerUrl int) *Proxy {
	return &Proxy{app: app, maxRequestPerUrl: maxRequestPerUrl, requestPerUrl: make(map[string]int)}
}

func (proxy *Proxy) HandleRequest(url, method string) (int, string) {
	allowed := proxy.checkRequestLimit(url)
	if !allowed {
		return 403, "Forbidden"
	}
	return proxy.app.HandleRequest(url, method)
}

func (proxy *Proxy) checkRequestLimit(url string) bool {
	if proxy.requestPerUrl[url] >= proxy.maxRequestPerUrl {
		return false
	} else {
		proxy.requestPerUrl[url]++
		return true
	}
}

type RouteEntry struct {
    Path    string
    Method  string
    Handler http.HandlerFunc
}
type Router struct {
    routes []RouteEntry
}

func (rtr *Router) Route(method, path string, handlerFunc http.HandlerFunc) {
    e := RouteEntry{
        Method:      method,
        Path:        path,
        HandlerFunc: handlerFunc,
    }
    rtr.routes = append(rtr.routes, e)
}
func (re *RouteEntry) Match(r *http.Request) bool {
    if r.Method != re.Method {
        return false // Method mismatch
    }

    if r.URL.Path != re.Path {
        return false // Path mismatch
    }

    return true
}
func (rtr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    for _, e := range rtr.routes {
        match := e.Match(r)
        if !match {
            continue
        }

        // We have a match! Call the handler, and return
        e.HandlerFunc.ServeHTTP(w, r)
        return
    }

    // No matches, so it's a 404
    http.NotFound(w, r)
}
r := &Router{}
r.Route("GET", "/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("The Best Router!"))
})


//------------------------------------------------------>

//accessing the parameters
// r.Route("GET", `/hello/(?P<Message>\w+)`, func(w http.ResponseWriter, r *http.Request) {
//     message := URLParam(r, "Message")
//     w.Write([]byte("Hello " + message))
// })

// // URLParam extracts a parameter from the URL by name
// func URLParam(r *http.Request, name string) string {
//     ctx := r.Context()

//     // ctx.Value returns an `interface{}` type, so we
//     // also have to cast it to a map, which is the 
//     // type we'll be using to store our parameters.
//     params := ctx.Value("params").(map[string]string)
//     return params[name]
// }
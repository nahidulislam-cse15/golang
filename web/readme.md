Date:24 March 2022
## Handler
Handler is a function which holds chunks of code which is suppose to run for every request made at that route. A handler function has to fulfill a certain signature
## Route
Route is a path of the URL. 
## http handle
Handle registers the handler for the given pattern in the DefaultServeMux
```
func Handle(pattern string, handler Handler)
```
handler- interface
## func HandleFunc 
HandleFunc registers the handler function for the given pattern in the DefaultServeMux. The documentation for ServeMux explains how patterns are matched.
```
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))

```
The pointer to <b>Request<b> above is an struct which holds data from the client. <b>ResponseWriter</b> is an interface which allows us to respond to the request elegantly.
http.ResponseWriter has a method signature of Write([]byte) (int, error) which means any type that implements io.Writer can also write to the w; io.WriteString being one of them.
## The Server Multiplexer
http.NewServeMux returns a ServeMux, which looks like this:
```
type ServeMux struct {
	mu    sync.RWMutex
	m     map[string]muxEntry
	es    []muxEntry // slice of entries sorted from longest to shortest.
	hosts bool       // whether any patterns contain hostnames
}

```
ServerMux is the canonical abstraction of all routes and handlers.
```
func (mux *ServeMux) Handle(pattern string, handler http.Handler)
func (mux *ServeMux) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string)
func (mux *ServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request)

```
ServerMux has 4 public methods, namely:

- mux.Handle(pattern string, handler Handler) - This takes a URL pattern and a type which implements a Handler. What is a Handler once again? An interface with method signature of ServeHTTP(w http.ResponseWriter, r *http.Request).
```
type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "dog doggy doggggy")
}

func main() {
    var d hotdog

    mux := http.NewServeMux()
    mux.Handle("/dog/", d)

    http.ListenAndServe(":8080", mux)
}

```
Above we can use d variable, which is a hotdog type which implements the Handler interface. The underlaying data type could be anything. In this case, it’s it. But it could have been a struct without any side-effect.

- http.HandlerFunc - HandlerFunc is a kind of a helper function that converts a standalone function (more on this next) to what mux.Handle takes. Let’s add into example above.

```
func madDog(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "I am a mad dog.")
}

func main() {
    var d hotdog

    mux := http.NewServeMux()
    mux.Handle("/dog/", d)
    mux.Handle("/maddog/", madDog)

    http.ListenAndServe(":8080", mux)
}
```
As you can see in the following screenshot, I can’t simply use the madDog function as argument to http.ServeMux.Handle. This is because mux.Handle is looking for a type which implements a ServerHTTP method.To overcome this, we can wrap our function with http.HandlerFunc which makes the function mux.Handle compatible.
- http.HandleFunc - http.HandleFunc takes a standalone function instead of taking a type which implements Handler inteface.
http.ListenAndServe takes a server address and any object which implements http.Handler to start a server. Normally we put a ServeMux, but it will also take any custom type which implements ServeHTTP(w http.ResponseWriter, r *http.Response).
# What is http.Handle vs mux.Handle?
They both are same. When you use http.Handle, program will automatically create a default server multiplexer. But in most cases developers create a new mux. mux := http.NewServeMux().

# In above example you once passed Handler and then a Mux to the ListenAndServe. How is that?
ServeMux implements ServeHTTP(w http.ResponseWriter, r *http.Request), so it’s also a handler.


## func ListenAndServe 
ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections. Accepted connections are configured to enable TCP keep-alives.The handler is typically nil, in which case the DefaultServeMux is used.


ListenAndServe always returns a non-nil error.
```
func ListenAndServe(addr string, handler Handler) error
```
## http.Handler
interface with the ServeHTTP method

## HandlerFunc
http.HandlerFunc-a function type that accepts same args as servehttp method also implements http.handler
```
type HandlerFunc func(ResponseWriter, *Request)
// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
```

sources;:
- [santosh](https://santoshk.dev/posts/2020/difference-between-handler-handle-and-handlerfunc/)
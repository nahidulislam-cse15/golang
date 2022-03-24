Date:24 March 2022
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
## func ListenAndServe 
ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections. Accepted connections are configured to enable TCP keep-alives.The handler is typically nil, in which case the DefaultServeMux is used.
ListenAndServe always returns a non-nil error.
```
func ListenAndServe(addr string, handler Handler) error
```
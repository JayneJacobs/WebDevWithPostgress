### Templates in Go

# Tol install local assets:

1. install NodeJs(https://nodejs.org)
2. install `grunt-cli` globally via `npm install -g grunt-cli`
3. run `npm install` from root directory of the project
4. run `grunt` from the root of the project to compile CSS

## to run the application
1. install `http0server` via `npm install -g http-server`
2.[ run `http-server` in applications root. ]* navigate to `http://localhost:8080/html/home.html`
''
## HandleFunc
type HandlerFunc func(ResponseWriter, * Request)
#### has method
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)

### notFoundHandler

func NOtFoundHandler Handler - returns 404

### RedirectHandler (similar to middle ware)
func RedirectHJandler(url string, code int) Handler

### StripPrefix
func StripPrefix(prefix string, h handler) Handler

### FileServer
func FileServer(root FileSystem) Handler
type FileSystem interface {
    Open(name string) (File, error)
}

type Dir string
func (d Dir) Open(name string) (File, error)


# Templates

### Example
Dear {{.FirstName}}

We would like to personally invite you to {{.Event}}!

Sincerely
{{.AuthorName}}




type Data struct {
    FirstName string
    Event string
    AuthorName string
}


### Https (Automatically upgrades to HTTP2)

1. Generate certs
   1. go run $GOROOT/src/crypto/tls/generate_cert.go -host localhost (run in the application root)
   2. Add http.ListenAndServeTLS(":8070", "cert.pem", "key.pem", &middleware.TimeoutMiddleware{new(middleware.GzipMiddleware)})

Header and Body are sent together in HTTP1

Header cannot be encrypted. 


  *Http/2  Streaming* - **Independant pipelines within an application**
Information is sent in frames. optimized independantly

Headers exchjanged 
continuation frames. 
DATA sent independantly
Each data type optimized independantly

Compressend and independant of the data. 
Advantages: 
1. Request multiplexing. 
2. Header compression. 
3. Reduced data
4. Secure by default. 
   * server push sends related stylesheets ahead. preemptively
   * reduces requests

## Caveat: 

    Can make app less efficient.  Resources are pushed every time so that client side cache has not benefit. 
    

## Testing:
Prefix with Test                  Aways take a pointer to a test object
func *Test*UpdatesLastLoginTime(t *testing.T) {
    ....
}

Controller Layer

func (h home) handleHome (w http.ResponseWriter, r *http.request){}


net/http/httptest

func NewRequest(method, target string, body io.Reader) *http.Request

type ResponseRecorder {
    Code int
    HeaderMap http.Header
    Body *byptes.Buffer
    Flushed bool
}


Profiling

import _ "net/http/pprof"

to tool pprof http://localhost:8070/degbug/pprof/heap //memory
to tool pprof http://localhost:8070/degbug/pprof/profile //CPU
to tool pprof http://localhost:8070/degbug/pprof/block //goroutines
to tool pprof http://localhost:8070/degbug/pprof/trace?seconds=5 //trace

http://localhost:8070/debug/pprof


# Commandline tool
 
 go tool pprof http://localhost:8080/debug/pprof/heap

 top

 top
Showing nodes accounting for 2454.85kB, 100% of 2454.85kB total
Showing top 10 nodes out of 39
      flat  flat%   sum%        cum   cum%
  902.59kB 36.77% 36.77%   902.59kB 36.77%  compress/flate.NewWriter
  528.17kB 21.52% 58.28%   528.17kB 21.52%  io.copyBuffer
  512.09kB 20.86% 79.14%   512.09kB 20.86%  compress/flate.newHuffmanEncoder (inline)
  512.01kB 20.86%   100%   512.01kB 20.86%  vendor/golang_org/x/net/http2/hpack.addDecoderNode
         0     0%   100%   512.09kB 20.86%  compress/flate.init
         0     0%   100%   512.09kB 20.86%  compress/flate.init.0
         0     0%   100%   902.59kB 36.77%  compress/gzip.(*Writer).Write
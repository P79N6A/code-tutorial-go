package main
/*
使用
1.zipkin 收集端
下载包 wget -O zipkin.jar 'https://search.maven.org/remote_content?g=io.zipkin.java&a=zipkin-server&v=LATEST&c=exec'
运行 /Users/gaolichuang/workspace/opentracing >java -jar zipkin.jar

2.server
运行下边的main 指定actor 为server
./trace --actor=server

3.client
运行下边的main 指定actor 为client
./trace --actor=client


启动server， 启动zipkin， 调用client后，client和server会将结果发送到zipkin
这个例子中：
client 调用server地址  http://localhost:8000/
server 的 / 会rewrite到 /gettime 然后返回结果


http request 实现原理， 在request的header里面添加uber-trace-id 头
并且client和server做http request的时候，都会讲相关信息发送给zipkin。  通过wireshark，显示过滤器指定http即可。
Frame 7: 240 bytes on wire (1920 bits), 240 bytes captured (1920 bits) on interface 0
Null/Loopback
Internet Protocol Version 6, Src: localhost (::1), Dst: localhost (::1)
Transmission Control Protocol, Src Port: 63818 (63818), Dst Port: irdmi (8000), Seq: 1, Ack: 1, Len: 164
Hypertext Transfer Protocol
    GET / HTTP/1.1\r\n
        [Expert Info (Chat/Sequence): GET / HTTP/1.1\r\n]
        Request Method: GET
        Request URI: /
        Request Version: HTTP/1.1
    Host: localhost:8000\r\n
    User-Agent: Go-http-client/1.1\r\n
    Uber-Trace-Id: 194f7f08e3042392:53fdfe4d35236dd9:48b1de0f73aca7df:1\r\n
    Accept-Encoding: gzip\r\n
    \r\n
    [Full request URI: http://localhost:8000/]
    [HTTP request 1/2]
    [Response in frame: 9]
    [Next request in frame: 13]



这个是zipkin收到的json:
[
  {
    "traceId": "0bf09e880a9a71f1",
    "parentId": "232a351255633896",
    "id": "2d9964a78619e2c0",
    "kind": "SERVER",
    "name": "http get",
    "timestamp": 1539766777058145,
    "duration": 45,
    "localEndpoint": {
      "serviceName": "server",
      "ipv4": "192.168.39.11"
    },
    "tags": {
      "component": "net/http",
      "hostname": "gaolichuangdeMacBook-Pro.local",
      "http.method": "GET",
      "http.status_code": "301",
      "http.url": "/",
      "jaeger.version": "Go-2.15.1dev"
    }
  },
  {
    "traceId": "0bf09e880a9a71f1",
    "parentId": "1fa40247df4a488f",
    "id": "02b9cbfbfc25fafa",
    "kind": "SERVER",
    "name": "http get",
    "timestamp": 1539766777058742,
    "duration": 41,
    "localEndpoint": {
      "serviceName": "server",
      "ipv4": "192.168.39.11"
    },
    "tags": {
      "component": "net/http",
      "hostname": "gaolichuangdeMacBook-Pro.local",
      "http.method": "GET",
      "http.status_code": "200",
      "http.url": "/gettime",
      "jaeger.version": "Go-2.15.1dev"
    }
  },
  {
    "traceId": "0bf09e880a9a71f1",
    "parentId": "1875079373dccdde",
    "id": "232a351255633896",
    "kind": "CLIENT",
    "name": "http get",
    "timestamp": 1539766777048953,
    "duration": 9569,
    "localEndpoint": {
      "serviceName": "client",
      "ipv4": "192.168.39.11"
    },
    "annotations": [
      {
        "timestamp": 1539766777049894,
        "value": "GetConn"
      },
      {
        "timestamp": 1539766777052219,
        "value": "{\"event\":\"DNSStart\",\"host\":\"localhost\"}"
      },
      {
        "timestamp": 1539766777056244,
        "value": "{\"addr\":\"127.0.0.1\",\"event\":\"DNSDone\"}"
      },
      {
        "timestamp": 1539766777056578,
        "value": "{\"addr\":\"[::1]:8000\",\"event\":\"ConnectStart\",\"network\":\"tcp\"}"
      },
      {
        "timestamp": 1539766777057429,
        "value": "{\"addr\":\"[::1]:8000\",\"event\":\"ConnectDone\",\"network\":\"tcp\"}"
      },
      {
        "timestamp": 1539766777057472,
        "value": "GotConn"
      },
      {
        "timestamp": 1539766777057971,
        "value": "WroteHeaders"
      },
      {
        "timestamp": 1539766777057979,
        "value": "WroteRequest"
      },
      {
        "timestamp": 1539766777058271,
        "value": "GotFirstResponseByte"
      },
      {
        "timestamp": 1539766777058512,
        "value": "PutIdleConn"
      },
      {
        "timestamp": 1539766777058522,
        "value": "ClosedBody"
      }
    ],
    "tags": {
      "component": "net/http",
      "http.method": "GET",
      "http.status_code": "301",
      "http.url": "localhost:8000"
    }
  },
  {
    "traceId": "0bf09e880a9a71f1",
    "parentId": "1875079373dccdde",
    "id": "1fa40247df4a488f",
    "kind": "CLIENT",
    "name": "http get",
    "timestamp": 1539766777058533,
    "duration": 434,
    "localEndpoint": {
      "serviceName": "client",
      "ipv4": "192.168.39.11"
    },
    "annotations": [
      {
        "timestamp": 1539766777058582,
        "value": "GetConn"
      },
      {
        "timestamp": 1539766777058587,
        "value": "GotConn"
      },
      {
        "timestamp": 1539766777058619,
        "value": "WroteHeaders"
      },
      {
        "timestamp": 1539766777058621,
        "value": "WroteRequest"
      },
      {
        "timestamp": 1539766777058848,
        "value": "GotFirstResponseByte"
      },
      {
        "timestamp": 1539766777058937,
        "value": "PutIdleConn"
      },
      {
        "timestamp": 1539766777058967,
        "value": "ClosedBody"
      }
    ],
    "tags": {
      "component": "net/http",
      "http.method": "GET",
      "http.status_code": "200",
      "http.url": "localhost:8000"
    }
  },
  {
    "traceId": "0bf09e880a9a71f1",
    "parentId": "0bf09e880a9a71f1",
    "id": "1875079373dccdde",
    "name": "http client",
    "timestamp": 1539766777048946,
    "duration": 10023,
    "localEndpoint": {
      "serviceName": "client",
      "ipv4": "192.168.39.11"
    },
    "tags": {
      "lc": "client"
    }
  },
  {
    "traceId": "0bf09e880a9a71f1",
    "id": "0bf09e880a9a71f1",
    "name": "client",
    "timestamp": 1539766777048751,
    "duration": 10218,
    "localEndpoint": {
      "serviceName": "client",
      "ipv4": "192.168.39.11"
    },
    "tags": {
      "component": "client",
      "hostname": "gaolichuangdeMacBook-Pro.local",
      "jaeger.version": "Go-2.15.1dev",
      "lc": "client",
      "sampler.type": "const"
    }
  }
]
*/

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"

    "github.com/opentracing-contrib/go-stdlib/nethttp"
    "github.com/opentracing/opentracing-go"
    "github.com/opentracing/opentracing-go/ext"
    otlog "github.com/opentracing/opentracing-go/log"
    "golang.org/x/net/context"
    "time"
    "io"
    "flag"
    "github.com/uber/jaeger-client-go/transport/zipkin"
    "github.com/uber/jaeger-client-go"
    "net/http/httputil"
)

func runClient(tracer opentracing.Tracer) {
    // nethttp.Transport from go-stdlib will do the tracing
    c := &http.Client{Transport: &nethttp.Transport{}}

    // create a top-level span to represent full work of the client
    span := tracer.StartSpan(client)
    span.SetTag(string(ext.Component), client)
    defer span.Finish()
    ctx := opentracing.ContextWithSpan(context.Background(), span)

    req, err := http.NewRequest(
        "GET",
        fmt.Sprintf("http://localhost:%s/", *serverPort),
        nil,
    )
    if err != nil {
        onError(span, err)
        return
    }

    req = req.WithContext(ctx)
    // wrap the request in nethttp.TraceRequest
    req, ht := nethttp.TraceRequest(tracer, req)
    defer ht.Finish()

    reqstr,_ := httputil.DumpRequest(req,true)
    fmt.Println(string(reqstr))
    res, err := c.Do(req)
    if err != nil {
        onError(span, err)
        return
    }
    resstr,_ := httputil.DumpResponse(res,true)
    fmt.Println(string(resstr))
    defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        onError(span, err)
        return
    }
    fmt.Printf("Received result: %s\n", string(body))
}

func onError(span opentracing.Span, err error) {
    // handle errors by recording them in the span
    span.SetTag(string(ext.Error), true)
    span.LogKV(otlog.Error(err))
    log.Print(err)
}
///////////////////////
func getTime(w http.ResponseWriter, r *http.Request) {
    log.Print("Received getTime request")
    t := time.Now()
    ts := t.Format("Mon Jan _2 15:04:05 2006")
    io.WriteString(w, fmt.Sprintf("The time is %s", ts))
}

func redirect(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r,
        fmt.Sprintf("http://localhost:%s/gettime", *serverPort), 301)
}

func runServer(tracer opentracing.Tracer) {
    http.HandleFunc("/gettime", getTime)
    http.HandleFunc("/", redirect)
    log.Printf("Starting server on port %s", *serverPort)
    http.ListenAndServe(
        fmt.Sprintf(":%s", *serverPort),
        // use nethttp.Middleware to enable OpenTracing for server
        nethttp.Middleware(tracer, http.DefaultServeMux))
}
///////////////////////////


var (
    zipkinURL = flag.String("url",
        "http://localhost:9411/api/v1/spans", "Zipkin server URL")
    serverPort = flag.String("port", "8000", "server port")
    actorKind  = flag.String("actor", "server", "server or client")
)

const (
    server = "server"
    client = "client"
)

func main() {
    flag.Parse()

    if *actorKind != server && *actorKind != client {
        log.Fatal("Please specify '-actor server' or '-actor client'")
    }

    // Jaeger tracer can be initialized with a transport that will
    // report tracing Spans to a Zipkin backend
    transport, err := zipkin.NewHTTPTransport(
        *zipkinURL,
        zipkin.HTTPBatchSize(1),
        zipkin.HTTPLogger(jaeger.StdLogger),
    )
    if err != nil {
        log.Fatalf("Cannot initialize HTTP transport: %v", err)
    }
    // create Jaeger tracer
    tracer, closer := jaeger.NewTracer(
        *actorKind,
        jaeger.NewConstSampler(true), // sample all traces
        jaeger.NewRemoteReporter(transport),
    )

    if *actorKind == server {
        runServer(tracer)
        return
    }

    runClient(tracer)

    // Close the tracer to guarantee that all spans that could
    // be still buffered in memory are sent to the tracing backend
    closer.Close()
}

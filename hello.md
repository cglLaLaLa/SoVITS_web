# My-GPT-SoVITS-Web-v1

## （一）项目介绍

### 1.背景

GPT-SoVITS是一个语音大模型，能根据提供的一分钟左右的语音素材，训练出音色，语气都一致的音频文件（wav,mp3等格式）。

但是由于GPT-SoVITS是一个学术项目，因此其WEB端完全是单体服务，不能商用，同时项目结构也比较混乱，难以维护和拓展。

因此，本项目用golang重写了GPT-SoVITS的服务端，采用DDD（领域驱动设计）的思想，对外提供http服务，通过调用MY-GPT-SoVITS（推理引擎）提供的rpc接口，实现了模拟语音的服务

### 2.项目结构

#### 2.1internal模块

##### 2.1.1 main.go

```go
func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    //（1）service.NewApplication,见2.1.1.1
    application, cleanup := service.NewApplication(ctx)
    defer cleanup()
    //（2）server.RunHttpServer,见2.1.1.2
    server.RunHttpServer("sovits", func(router *gin.Engine) {
       //router.StaticFile("/success", "../../public/success.html")
       sliceoapi.RegisterHandlersWithOptions(router, HttpServer{
          app: application,
       }, sliceoapi.GinServerOptions{
          BaseURL:      "/api",
          Middlewares:  nil,
          ErrorHandler: nil,
       })
    })
}
```

###### 2.1.1.1其中，service.NewApplication注册了一个grpc服务，该方法位于preprocess领域的service包下，是提供预处理服务的

```go
func NewApplication(ctx context.Context) (app.Application, func()) {
    //（1） client.NewSliceGRPCClient，通用
    sliceClient, closeSliceClient, err := client.NewSliceGRPCClient(ctx)
    if err != nil {
       panic(err)
    }
    //（2）grpc.NewSliceGRPC，不同领域有不同的适配层
    sliceGRPc := grpc.NewSliceGRPC(sliceClient)
    //（3）newApplicationClient
    return newApplicationClient(ctx, sliceGRPc), func() {
       _ = closeSliceClient()
    }
}
```

（1） client.NewSliceGRPCClient注册了一个指定地址（带端口号）的gprc客户端，由于我们认为提供grpc的客户端是一项通用能力，我们将其放在internal/common/slice下，定义其为slice领域提供的一项通用功能

```go
func NewSliceGRPCClient(ctx context.Context) (client slice.SliceServiceClient, close func() error, err error) {
    //（1）等待一段时间，超时则连接失败
    if !waitFor(addr, timeout) {
       logrus.Error("addr:%s dial failed,timeout:%d", addr, timeout)
       return nil, nil, errors.New("grpc not available")
    }
    //（2）golang的标准grpc库里的方法，无需赘述，注意，这个http连接是无证书验证的，后期可以修改，具体参考官网demo
    opts := grpcDialOpts(addr)
    conn, err := grpc.NewClient(addr, opts...)
    if err != nil {
       return nil, func() error { return nil }, err
    }
    //（3）调用common/slice/genproto/start_sclie_grpc.pb.go是根据proto文件生成的
    return slice.NewSliceServiceClient(conn), conn.Close, nil
}

func grpcDialOpts(_ string) []grpc.DialOption {
    return []grpc.DialOption{
       grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithStatsHandler(otelgrpc.NewClientHandler())}
}
```

（2）grpc.NewSliceGRPC，是一个适配器层，将grpc客户端和gprc服务端都封装进去，该段代码，由于不同的领域实现的具体grpc服务不同，请求和响应体也不同，所以我们将其定义为不同领域层的适配层，该段代码在/internal/preprocess/adapter/grpc包下

```go
type SliceGRPC struct {
    client slice.SliceServiceClient
}

func NewSliceGRPC(client slice.SliceServiceClient) *SliceGRPC {
    return &SliceGRPC{client: client}
}

func (s SliceGRPC) StartSlice(ctx context.Context, slice *slice.SliceRequest) (*slice.StartResponse, error) {
    resp, err := s.client.StartSlice(ctx, slice)
    return resp, err
}
```

（3）newApplicationClient，该方法组装了一个application

```go
func newApplicationClient(ctx context.Context, serviceClient command.CommandService) app.Application {
    //（1）日志
    logger := logrus.NewEntry(logrus.StandardLogger())
    //（2）链路追踪
    metric := common.MyMetric{}
    //（3）applicaiton结构体
    return app.Application{
       Commands: app.Commands{
          StartSliceCommand: command.NewStartSliceHandler(ctx, serviceClient, logger, metric),
       },
    }
}
```

该application结构如下，在具体领域的app目录下，例如该结构体位于preprocess/app下，封装了command和query，由于具体的command和query与领域耦合，所以在领域层下

```go
type Application struct {
    Commands Commands
}

type Commands struct {
    StartSliceCommand command.StartSliceHandler
}
```

###### 2.1.1.2 该函数提供了http服务，不管你是openapi还是其他rpc框架生成的http接口，我们统一接受一个wrapper函数，由这个wrapeer函数来进行http服务的注册，因此该段代码应该是通用的，在internal/common/server下

```go
func RunHttpServer(serviceName string, wrapper func(router *gin.Engine)) {
    //todo 后期配置化，不要写死
    addr := "127.0.0.1:14888"
    if addr == "" {
       panic("empty http address")
    }
    RunHttpServerOnAddr(addr, wrapper)
}

func RunHttpServerOnAddr(addr string, wrapper func(router *gin.Engine)) {
    apiRouter := gin.New()
    wrapper(apiRouter)
    apiRouter.Group("/api")
    apiRouter.GET("/ping", func(c *gin.Context) {
       c.JSON(200, "pong")
    })
    if err := apiRouter.Run(addr); err != nil {
       panic(err)
    }
}
```

我们可以看到,openapi的wrapper函数，我们做了如下定义

```go
func(router *gin.Engine) {
    //router.StaticFile("/success", "../../public/success.html")
    sliceoapi.RegisterHandlersWithOptions(router, HttpServer{
       app: application,
    }, sliceoapi.GinServerOptions{
       BaseURL:      "/api",
       Middlewares:  nil,
       ErrorHandler: nil,
    })
}
```

其中sliceoapi，是我们用openapi框架生成的gen.go文件，其函数签名如下：

```go
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
    errorHandler := options.ErrorHandler
    if errorHandler == nil {
       errorHandler = func(c *gin.Context, err error, statusCode int) {
          c.JSON(statusCode, gin.H{"msg": err.Error()})
       }
    }

    wrapper := ServerInterfaceWrapper{
       Handler:            si,
       HandlerMiddlewares: options.Middlewares,
       ErrorHandler:       errorHandler,
    }
	//关键看这一行代码就行，其实就是，在gin的router中注册了一个post方法（router实际上是一个压缩字段树结构）
    router.POST(options.BaseURL+"/preprocess/:customer_id/slice", wrapper.PostPreprocessCustomerIdSlice)
}
```

ServerInterface的接口签名如下,已经定义好了调用的方法和路由：

```go
type ServerInterface interface {

    // (POST /preprocess/{customer_id}/slice)
    PostPreprocessCustomerIdSlice(c *gin.Context, customerId string)
}
```

所以我们只需要实现具体领域的PostPreprocessCustomerIdSlice方法即可，如下,这个http服务是与领域相关的，所以放在preprocess下：

```go
type HttpServer struct {
    app app.Application
}

func (s HttpServer) PostPreprocessCustomerIdSlice(c *gin.Context, customerId string) {
    var (
       req  sliceoapi.SliceRequest
       err  error
       resp struct {
          CustomerID  string `json:"customer_id"`
          RedirectURL string `json:"redirect_url"`
       }
    )
    if err = c.ShouldBindJSON(&req); err != nil {
       return
    }
    r, err := s.app.Commands.StartSliceCommand.Handle(c.Request.Context(), command.StartSlice{
       CustomerID: customerId,
       Request:    convertor.NewSliceStartConverter().ClientToEntity(&req),
    })
    if err != nil {
       return
    }
    resp.CustomerID = customerId
    //todo
    resp.RedirectURL = "todo" + r.SliceId
}
```

看这一段代码,其实就是我们在2.1.1.1中写好的代码，调用具体application的handle方法，其中coomand.StartSlice是对请求体的封装，在ddd的领域层中，我们并不希望出现与底层技术相关的一些命名，例如xxxrequest等，我们希望领域层完全与业务相关：

```go
r, err := s.app.Commands.StartSliceCommand.Handle(c.Request.Context(), command.StartSlice{
    CustomerID: customerId,
    Request:    convertor.NewSliceStartConverter().ClientToEntity(&req),
})
```

点进这个handle方法里看一下,实际上就是我们又把grpc的客户端封装了一层，这段代码在具体的领域层的app/command目录下：

```go
type startSliceHandler struct {
    logger *logrus.Entry
    client CommandService
}

func (s startSliceHandler) Handle(ctx context.Context, q StartSlice) (*SliceResult, error) {
    resp, err := s.client.StartSlice(ctx, convertor.NewSliceStartConverter().EntityToProto(q.Request))
    if err == nil {
       s.logger.WithFields(logrus.Fields{
          "slice error": err,
       })
    }
    //todo 规定一下常量值
    if resp.Status == "1" {
       return &SliceResult{SliceId: q.CustomerID}, nil
    }
    return nil, errors.New("返回错误")
}
```

那么我们为啥还要封装一下呢，看一下这个handler的new方法,实际上我们才考了go的一些web项目的中间件的写法，采用装饰器模式：

```go
type StartSliceHandler decorator.CommandHandler[StartSlice, *SliceResult]

func NewStartSliceHandler(ctx context.Context, client CommandService, logger *logrus.Entry, metric decorator.MetricClient) StartSliceHandler {
    if client == nil {
       panic("[NewStartSliceHandler] client is nil")
    }
    return decorator.ApplyCommandDecorators[StartSlice, *SliceResult](
       startSliceHandler{client: client},
       logger,
       metric,
    )
}
```

其中ApplyCommandDecorators，是一种通用的能力，我们写在common/decorator下,所有的command的handler，都要加一个loggingDecorator装饰器

```go
type CommandHandler[C any, Q any] interface {
    Handle(ctx context.Context, q C) (Q, error)
}

func ApplyCommandDecorators[H, R any](handler CommandHandler[H, R], logger *logrus.Entry, client MetricClient) CommandHandler[H, R] {
    return loggingDecorator[H, R]{
       logger: logger,
       base: metricDecorator[H, R]{
          client: client,
          base:   handler,
       },
    }
}
```

loggingDecorator，可以看到，这个装饰器的handle方法，会先打一下日志，然后再调用CommandHandler的handle方法：

```go
type loggingDecorator[Q, R any] struct {
    logger *logrus.Entry
    base   CommandHandler[Q, R]
}

func (l loggingDecorator[Q, R]) Handle(ctx context.Context, cmd Q) (result R, err error) {
    logger := l.logger.WithFields(logrus.Fields{
       "query":      generateActionName(cmd),
       "query_body": fmt.Sprintf("%#v", cmd),
    })

    logger.Debug("Executing")
    return l.base.Handle(ctx, cmd)
}
```

关键来了，那么handle方法在哪里被调用呢？正是在PostPreprocessCustomerIdSlice这个方法中被调用了，那么一切就串起来来了：

在main函数中，我们创建application的时候，创建了这个带有装饰器的handler，并将其作为一个command封装进一个application中。

然后我们把这个application作为一个httpServer封装了一下，注册成为一个openapi接口！这样一旦这个command的handle方法被调用了，那么就会先走一下日志的装饰器，然后再调用具体的业务逻辑！具体接口如下所示

!!!请见openapi/slice.yml


以后我们新增接口的时候，只需要在CommandService中新写一个函数签名：

```go
type CommandService interface {
    StartSlice(ctx context.Context, slice *slice.SliceRequest) (*slice.StartResponse, error)
    //新增在下面
}
```

然后把具体的实现写在app/command下就行了，对之前的代码毫无侵入！
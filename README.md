# TarsGo  Document
[点我查看中文版](README.zh.md)
## About
- Tarsgo is high performance RPC framework in Golang programing language using the tars protocol.
- Go has become popular for programming with the rise of containerization technology such as docker, k8s, and etcd.
- Go's goroutine concurrency mechanism means Go is very suitable for large-scale high-concurrency back-end server program development. The Go language has nearly C/C++ performance and near Python productivity.
- In Tencent, part of the existing C++ development team has gradually turned into Go developers. Tars, a widely used RPC framework, supports C++, Java, NodeJS, and PHP, and now Go. The combination with Go language has become a general trend. Therefore, in the voice of users, we launched Tarsgo, and we have applied to Tencent map application, YingYongbao application, Internet plus and other projects.
- Learn more about the whole Tars architecture and design at [Introduction](https://tarscloud.github.io/TarsDocs_en/SUMMARY.html).


## Function & features
- Tars2go tool: tars file is automatically generated and converted into Go language, contains RPC server/client code; 
- Serialization and deserialization of tars protocol in Go.
- Auto service discovery.
- TCP/UDP/Http server & Client.
- The support of local logging and remote logging.
- The support of statistical reporting,  property statistics, and anomaly reporting.
- The support of set division.
- The support of protocol buffers. See more in [pb2tarsgo](tars/tools/pb2tarsgo/README.md).
- The support of filter.
- The support of Zipkin OpenTracing.


## Install
- For installing OSS and other basic servers, see the [Installation](https://tarscloud.github.io/TarsDocs_en/installation/) document.
- Install go 1.13.x or above (for example go install path: `/usr/local/go`), set `GOROOT`, `GOPATH`, for example: in Linux:

```bash
export GOROOT=/usr/local/go  
export GOPATH=/root/gocode   
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

If you in china, you can set go proxy:

```bash
go env -w GOPROXY=https://goproxy.cn   
```

Please set go mod:

```bash
go env -w GO111MODULE=auto
```

install `tars2go`:

```bash
# < go 1.17 
go get -u github.com/jslyzt/tarsgo/tars/tools/tars2go
# >= go 1.17
go install github.com/jslyzt/tarsgo/tars/tools/tars2go@latest
```

## Quickstart
- For quickstart, see [tars_go_quickstart_en.md](docs/tars_go_quickstart_en.md)

## Performance
- For performance, see [tars_go_performance.md](docs/tars_go_performance.md)

## Usage
### 1 Server
 - Below is a full example illustrating how to use tarsgo to build a server.

#### 1.1 Interface Definition

Create a tars file, like hello.tars, under `$GOPATH/src` (for example, `$GOPATH/src/TestApp/TestServer/hello.tars`).

For more detail about tars protocol, see [tars_protocol](https://tarscloud.github.io/TarsDocs_en/base/tars-protocol.html)
Tars protocol is a binary, IDL-based protocol similar to protocol buffers.
	
```go
module TestApp
{
	
	interface Hello
	{
	    int test();
	    int testHello(string sReq, out string sRsp);
	};
	
};	
```


#### 1.2 Compile Interface Definition File

##### 1.2.1 Build `tars2go`
If not install `tars2go`, compile and install the `tars2go` tools.

```bash
# < go 1.17 
go get -u github.com/jslyzt/tarsgo/tars/tools/tars2go
# >= go 1.17
go install github.com/jslyzt/tarsgo/tars/tools/tars2go@latest
```

##### 1.2.2 Compile the Tars File and Translate into Go File

```bash
tars2go --outdir=./vendor hello.tars
```

#### 1.3 Implement the Interface

```go
package main

import (
    "github.com/jslyzt/tarsgo/tars"

    "TestApp"
)

type HelloImp struct {
}

//implete the Test interface
func (imp *HelloImp) Test() (int32, error) {
    return 0, nil 
}

//implete the testHello interface

func (imp *HelloImp) TestHello(in string, out *string) (int32, error) {
    *out = in
    return 0, nil 
}


func main() { //Init servant
    imp := new(HelloImp)                                    //New Imp
    app := new(TestApp.Hello)                               //New init the A Tars
    cfg := tars.GetServerConfig()                           //Get Config File Object
    app.AddServant(imp, cfg.App+"."+cfg.Server+".HelloObj") //Register Servant
    tars.Run()
}
```

Illustration:
- HelloImp is the struct where you implement the Hello & Test interface, notice that Test & Hello must start with an upper letter to be exported, which is the only place distinguished from tars file definition. 
- TestApp.Hello is generated by `tar2go` tools, which could be found in `./vendor/TestApp/Hello_IF.go`, with a package named TestApp which is the same as module TestApp in the tars file.
-  `tars.GetServerConfig()` is used to get server-side configuration.
-  `cfg.App+"."+cfg.Server+".HelloObj"` is the object name bind to the Servant, which the client will use this name to access the server.



#### 1.4 ServerConfig

`tars.GetServerConfig()` return a server config, which is defined as below:

```go
type serverConfig struct {
	Node      string
	App       string
	Server    string
	LogPath   string
	LogSize   string
	LogLevel  string
	Version   string
	LocalIP   string
	BasePath  string
	DataPath  string
	config    string
	notify    string
	log       string
	netThread int
	Adapters  map[string]adapterConfig

	Container   string
	Isdocker    bool
	Enableset   bool
	Setdivision string
}
```

- Node: local tarsnode address, only if you use tars platform to deploy will use this parameter.
- APP: The application name.
- Server: The server name.
- LogPath: The  directory to save logs.
- LogSize: The size when rotate logs.
- LogLevel: The rotate log level.
- Version: Tarsgo version.
- LocalIP: Local IP address.
- BasePath: Base path for the binary.
- DataPath: Path for store some cache files.
- config: The configuration center for getting configuration, like `tars.tarsconfig.ConfigObj`.
- notify: The notify center for report notify report, like `tars.tarsnotify.NotifyObj`.
- log: The remote log centre, like `tars.tarslog.LogObj`.
- netThread: Reserved  for controlling the go routine that receives and sends packages.
- Adapters:  The specified configuration for each adapter.
- Container: Reserved for later use, to store the container name.
- Isdocker: Reserved for later use, to specify if the server running inside a container.
- Enableset: True if used set division.
- Setdivision: To specify  which set division, like `gray.sz.*`.

A server-side configuration looks like:

```xml
<tars>
  <application>
      enableset=Y
      setdivision=gray.sz.*
    <server>
       node=tars.tarsnode.ServerObj@tcp -h 10.120.129.226 -p 19386 -t 60000
       app=TestApp
       server=HelloServer
       localip=10.120.129.226
       local=tcp -h 127.0.0.1 -p 20001 -t 3000
       basepath=/usr/local/app/tars/tarsnode/data/TestApp.HelloServer/bin/
       datapath=/usr/local/app/tars/tarsnode/data/TestApp.HelloServer/data/
       logpath=/usr/local/app/tars/app_log/
       logsize=10M
       config=tars.tarsconfig.ConfigObj
       notify=tars.tarsnotify.NotifyObj
       log=tars.tarslog.LogObj
       #timeout for deactiving,  ms.
       deactivating-timeout=2000
       logLevel=DEBUG
    </server>
  </application>
</tars>
```

#### 1.5 Adapter
An adapter represents a bind IP port for a certain object. `
app.AddServant(imp, cfg.App+"."+cfg.Server+".HelloObj")` in the server implement code example, fulfill the binding of the adapter configuration and implement for the `HelloObj`.

A full example for an adapter, see below:

```xml
<tars>
  <application>
    <server>
       #each adapter configuration 
       <TestApp.HelloServer.HelloObjAdapter>
            #allow Ip for white list.
            allow
            # ip and port to listen on  
            endpoint=tcp -h 10.120.129.226 -p 20001 -t 60000
            #handlegroup
            handlegroup=TestApp.HelloServer.HelloObjAdapter
            #max connection 
            maxconns=200000
            #portocol, only tars for now.
            protocol=tars
            #max capbility in handle queue.
            queuecap=10000
            #timeout in ms for the request in the queue.
            queuetimeout=60000
            #servant 
            servant=TestApp.HelloServer.HelloObj
            #threads in handle server side implement code. goroutine for golang.
            threads=5
       </TestApp.HelloServer.HelloObjAdapter>
    </server>
  </application>
</tars>
```

#### 1.6 Start the Server 

The command to start the server:

```bash
./HelloServer --config=config.conf
```

See below for a full example of `config.conf`, we will explain the client-side configuration later.

```xml
<tars>
  <application>
    enableset=n
    setdivision=NULL
    <server>
       node=tars.tarsnode.ServerObj@tcp -h 10.120.129.226 -p 19386 -t 60000
       app=TestApp
       server=HelloServer
       localip=10.120.129.226
       local=tcp -h 127.0.0.1 -p 20001 -t 3000
       basepath=/usr/local/app/tars/tarsnode/data/TestApp.HelloServer/bin/
       datapath=/usr/local/app/tars/tarsnode/data/TestApp.HelloServer/data/
       logpath=/usr/local/app/tars/app_log/
       logsize=10M
       config=tars.tarsconfig.ConfigObj
       notify=tars.tarsnotify.NotifyObj
       log=tars.tarslog.LogObj
       deactivating-timeout=2000
       logLevel=DEBUG
       <TestApp.HelloServer.HelloObjAdapter>
            allow
            endpoint=tcp -h 10.120.129.226 -p 20001 -t 60000
            handlegroup=TestApp.HelloServer.HelloObjAdapter
            maxconns=200000
            protocol=tars
            queuecap=10000
            queuetimeout=60000
            servant=TestApp.HelloServer.HelloObj
            threads=5
       </TestApp.HelloServer.HelloObjAdapter>
    </server>
    <client>
       locator=tars.tarsregistry.QueryObj@tcp -h 10.120.129.226 -p 17890
       sync-invoke-timeout=3000
       async-invoke-timeout=5000
       refresh-endpoint-interval=60000
       report-interval=60000
       sample-rate=100000
       max-sample-count=50
       asyncthread=3
       modulename=TestApp.HelloServer
    </client>
  </application>
</tars>
```

### 2 Client
Users can write a client-side code easily without writing any protocol-specified communicating code.
#### 2.1 Client Example
A client side example:

```go
package main

import (
    "fmt"
    "github.com/jslyzt/tarsgo/tars"
    "TestApp"
)
//tars.Communicator should only init once and be global
var comm *tars.Communicator

func main() {
    comm = tars.NewCommunicator()
    obj := "TestApp.TestServer.HelloObj@tcp -h 127.0.0.1 -p 10015 -t 60000"
    app := new(TestApp.Hello)
    comm.StringToProxy(obj, app)
	var req string="Hello World"
    var out string
    ret, err := app.TestHello(req, &out)
    if err != nil {
        fmt.Println(err)
        return
    }   
    fmt.Println(ret, out)
}
```

Illustration:
- package TestApp was generated by `tars2go` tool using the tars protocol file.
- comm: A Communicator is used for communicating with the server side which should only init once and be global.
- obj: object name which to specify the server IP and port. Usually, we just need the object name before the "@" character.
- app: Application that associated with the interface in the tars file. In this case, it's `TestApp.Hello`.
- `StringToProxy`: `StringToProxy` method is used for binding the object name and the application, if skip this, the communicator won't know who to communicate with for the application.
- `req, res`: In and out parameter which define in the tars file for the TestHello method.
- `app.TestHello` is used to call the method defined in the tars file, and ret and err will be returned.

#### 2.2 Communicator
A communicator represents a group of resources for sending and receiving packages for the client-side, which in the end manages the socket communicating for each object.

You will only need one communicator in a program.

```go
var comm *tars.Communicato
comm = tars.NewCommunicator()
comm.SetProperty("property", "tars.tarsproperty.PropertyObj")
comm.SetProperty("locator", "tars.tarsregistry.QueryObj@tcp -h ... -p ...")
```

Description:
> * The communicator's configuration file format will be described later.
> * Communicators can be configured without a configuration file, and all parameters have default values.
> * The communicator can also be initialized directly through the `SetProperty` method.
> * If you don't want a configuration file, you must set the locator parameter yourself.

Communicator attribute description:
> * locator: The address of the registry service must be in the format `ip port`. If you do not need the registry to locate the service, you do not need to configure this item.
> * important  async-invoke-timeout: The maximum timeout (in milliseconds) for client calls. The default value for this configuration is `3000`.
> * sync-invoke-timeout: Unused for tarsgo right now.
> * refresh-endpoint-interval: The interval (in milliseconds) for periodically accessing the registry to obtain information. The default value for this configuration is one minute.
> * stat: The address of the service is called between modules. If this item is not configured, it means that the reported data will be directly discarded.
> * property: The address that the service reports its attribute. If it is not configured, this means that the reported data is directly discarded.
> * report-interval: Unused for tarsgo for now.
> * asyncthread: Discarded for tarsgo.
> * modulename: The module name, the default value is the name of the executable program.

The format of the communicator's configuration file is as follows:

```xml
<tars>
  <application>
    #The configuration required by the proxy
    <client>
        #address
        locator                     = tars.tarsregistry.QueryObj@tcp -h 127.0.0.1 -p 17890
        #The maximum timeout (in milliseconds) for synchronous calls.
        sync-invoke-timeout         = 3000
        #The maximum timeout (in milliseconds) for asynchronous calls.
        async-invoke-timeout        = 5000
        #The maximum timeout (in milliseconds) for synchronous calls.
        refresh-endpoint-interval   = 60000
        #Used for inter-module calls
        stat                        = tars.tarsstat.StatObj
        #Address used for attribute reporting
        property                    = tars.tarsproperty.PropertyObj
        #report time interval
        report-interval             = 60000
        #The number of threads that process asynchronous responses
        asyncthread                 = 3
        #The module name
        modulename                  = Test.HelloServer
    </client>
  </application>
</tars>
```

#### 2.3 Timeout Control
if you want to use timeout control on the client-side, use `TarsSetTimeout` which in ms.

```go
app := new(TestApp.Hello)
comm.StringToProxy(obj, app)
app.TarsSetTimeout(3000)
```

#### 2.4 Call Interface

This section details how the Tars client remotely invokes the server.

First, briefly describe the addressing mode of the Tars client. Secondly, it will introduce the calling method of the client, including but not limited to one-way calling, synchronous calling, asynchronous calling, hash calling, and so on.

##### 2.4.1. Introduction to Addressing Mode

The addressing mode of the Tars service can usually be divided into two ways: the service name is registered in the master and the service name is not registered in the master. A master is a name server (routing server) dedicated to registering service node information.

The service name added in the name server is implemented through the operation management platform.

For services that are not registered with the master, it can be classified as direct addressing, that is, the IP address of the service provider needs to be specified before calling the service. The client needs to specify the specific address of the `HelloObj` object when calling the service:

that is `Test.HelloServer.HelloObj@tcp -h 127.0.0.1 -p 9985`

`Test.HelloServer.HelloObj`: Object name

`tcp`: Tcp protocol

`-h`: Specify the host address, here is `127.0.0.1`

`-p`: Port, here is `9985`

If `HelloServer` is running on two servers, the app is initialized as follows:

```go
obj:= "Test.HelloServer.HelloObj@tcp -h 127.0.0.1 -p 9985:tcp -h 192.168.1.1 -p 9983"
app := new(TestApp.Hello)
comm.StringToProxy(obj, app)
```

The address of `HelloObj` is set to the address of the two servers. At this point, the request will be distributed to two servers (distribution method can be specified, not introduced here). If one server is down, the request will be automatically assigned to another one, and the server will be restarted periodically.

For services registered in the master, the service is addressed based on the service name. When the client requests the service, it does not need to specify the specific address of the `HelloServer`, but it needs to specify the address of the `registry` when generating the communicator or initializing the communicator.

The following shows the address of the registry by setting the parameters of the communicator:

```go
var *tars.Communicator
comm = tars.NewCommunicator()
comm.SetProperty("locator", "tars.tarsregistry.QueryObj@tcp -h ... -p ...")
```

Since the client needs to rely on the registry's address, the registry must also be fault-tolerant. The registry's fault-tolerant method is the same as above, specifying the address of the two registries.

##### 2.4.2. One-way Call
TODO. Unsupported yet in tarsgo.

##### 2.4.3. Synchronous Call

```go
package main

import (
    "fmt"
    "github.com/jslyzt/tarsgo/tars"
    "TestApp"
)

var *tars.Communicator
func main() {
    comm = tars.NewCommunicator()
    obj := "TestApp.TestServer.HelloObj@tcp -h 127.0.0.1 -p 10015 -t 60000"
    app := new(TestApp.Hello)
    comm.StringToProxy(obj, app)
	var req string="Hello World"
    var out string
    ret, err := app.TestHello(req, &out)
    if err != nil {
        fmt.Println(err)
        return
    }   
    fmt.Println(ret, out)
}
```

##### 2.4.4 Asynchronous Call
tarsgo can use Asynchronous calls easily using goroutine. Unlike C++, we don't need to implement a callback function.

```go
package main

import (
    "fmt"
    "github.com/jslyzt/tarsgo/tars"
    "time"
    "TestApp"
)
var *tars.Communicator
func main() {
    comm = tars.NewCommunicator()
    obj := "TestApp.TestServer.HelloObj@tcp -h 127.0.0.1 -p 10015 -t 60000"
    app := new(TestApp.Hello)
    comm.StringToProxy(obj, app)
	go func(){
		var req string="Hello World"
    	var out string
    	ret, err := app.TestHello(req, &out)
    	if err != nil {
        	fmt.Println(err)
        	return
    	} 
		fmt.Println(ret, out)
	}()
    time.Sleep(1)  
}
```

##### 2.4.5 Call by Set
The client can call server by set through configuration file mentioned. Which enableset will be y and setdivision will set `like gray.sz.*`. See [IDC Set](https://tarscloud.github.io/TarsDocs_en/dev/tars-idc-set.html) for more detail.

```go
package main

import (
    "fmt"
    "github.com/jslyzt/tarsgo/tars"
    "TestApp"
)

var *tars.Communicator
func main() {
    comm = tars.NewCommunicator()
    app := new(TestApp.Hello)
    obj := "TestApp.HelloGo.SayHelloObj"
    comm.SetProperty("locator", "tars.tarsregistry.QueryObj@tcp -h ... -p ...")
    comm.SetProperty("enableset", true)
    comm.SetProperty("setdivision", "gray.sz.*")
    
    var req string="Hello Wold"
    var res string
    ret, err := app.TestHello(req, &out)
    if err != nil {
        fmt.Println(err)
        return
    }   
    fmt.Println(ret, res)
}
```



##### 2.4.6. Hash Call
Since multiple servers can be deployed, client requests are randomly distributed to the server, but in some cases, certain requests should be always sent to a particular server. In this case, Tars provides a simple way to achieve which is called hash-call. Tarsgo Has supported this feature In version v1.1.5.
```go
package main

import (
    "fmt"
    "github.com/jslyzt/tarsgo/tars"
    "github.com/jslyzt/tarsgo/tars/util/current"
    "context"
    "time"
    "TestApp"
)
func main() {
    var comm *tars.Communicator
    comm = tars.NewCommunicator()
    obj := "TestApp.TestServer.HelloObj@tcp -h 127.0.0.1 -p 10015 -t 60000"
    app := new(TestApp.Hello)
    comm.StringToProxy(obj, app)
	go func(){
        var req string="Hello Wold"
    	var res string
        ctx := context.Background()
        ctx = current.ContextWithClientCurrent(ctx)
        // the request parameter hashtype, ModHash is 0, ConsistentHash is 1
        hashType := 0
        hashCode := uint32(123)
        current.SetClientHash(ctx, hashType, hashCode)
    	ret, err := app.TestHelloWithContext(ctx, req, &res)
    	if err != nil {
        	fmt.Println(err)
        	return
    	} 
		fmt.Println(ret, res)
	}()
    time.Sleep(1)  
}

```

### 3 Return Code Defined by Tars.
```go
//Define the return code given by the TARS service
const int TARSSERVERSUCCESS       = 0;    //Server-side processing succeeded
const int TARSSERVERDECODEERR     = -1;   //Server-side decoding exception
const int TARSSERVERENCODEERR     = -2;   //Server-side encoding exception
const int TARSSERVERNOFUNCERR     = -3;   //There is no such function on the server side
const int TARSSERVERNOSERVANTERR  = -4;   //The server does not have the Servant object
const int TARSSERVERRESETGRID     = -5;   // server grayscale state is inconsistent
const int TARSSERVERQUEUETIMEOUT  = -6;   //server queue exceeds limit
const int TARSASYNCCALLTIMEOUT    = -7;   // Asynchronous call timeout
const int TARSINVOKETIMEOUT       = -7;   //call timeout
const int TARSPROXYCONNECTERR     = -8;   //proxy link exception
const int TARSSERVEROVERLOAD      = -9;   //Server overload, exceeding queue length
const int TARSADAPTERNULL         = -10;  //The client routing is empty, the service does not exist or all services are down.
const int TARSINVOKEBYINVALIDESET = -11;  //The client calls the set rule illegally
const int TARSCLIENTDECODEERR     = -12;  //Client decoding exception
const int TARSSERVERUNKNOWNERR    = -99;  //The server is in an abnormal position
```

### 4 Log

A quick example for using tarsgo rotating log:

```go
TLOG := tars.GetLogger("TLOG")
TLOG.Debug("Debug logging")
```

This is will create a `*Rogger.Logger`, which was defined in `tars/util/rogger`, and after `GetLogger` was called, and a logfile is created under Logpath defined in the `config.conf` which name is `cfg.App + "." + cfg.Server + "_" + name`, and will be rotated after `100MB`(default),  and max rotated file is `10`(default). 

If you don't want to rotate log by file size. For example, you want to rotate by day, then use:

```go
TLOG := tars.GetDayLogger("TLOG",1)
TLOG.Debug("Debug logging")
```

For rotating by hour, use `GetHourLogger("TLOG",1)`.
If you want to log to a remote server, which is defined in `config.conf` named `tars.tarslog.LogObj`. A full tars file definition can be found in `tars/protocol/res/LogF.tars`. You have to setup a log server before doing this. A log server can be found under `Tencent/Tars/cpp/framework/LogServer`. A quick example:

```go
TLOG := GetRemoteLogger("TLOG")
TLOG.Debug("Debug logging")
```

If you want to set the log level, you can set it from the OSS platform provided by tars project under `Tencent/Tars/web`.
If you want to customize your logger, see more detail in `tars/util/rogger`,  `tars/logger.go` and `tars/remotelogger.go`.

### 5 Service Management

The Tars server framework supports dynamic receiving commands to handle related business logic, such as dynamic update configuration.

tarsgo currently has `tars.viewversion` / `tars.setloglevel` administration commands for now. Users can send admin command from OSS to see what version is or setting log level mentioned about.

If you want to defined ur own admin commands, see this example:

```go
func helloAdmin(who string ) (string, error) {
	return who, nil
}
tars.RegisterAdmin("tars.helloAdmin", helloAdmin)
```

Then you can send self-defined admin command `tars.helloAdmin` tarsgo and tarsgo will be shown in the browser.

Illustration:
```go
// A function  should be in this format
type adminFn func(string) (string, error)

//then u should registry this function using

func RegisterAdmin(name string, fn adminFn)
```

### 6 Statistical Reporting

Reporting statistics information is the logic of reporting the time-consuming information and other information to `tarsstat` inside the Tars framework. No user development is required. After the relevant information is correctly set during program initialization, it can be automatically reported inside the framework (including the client and the server).

After the client called the reporting interface, it is temporarily stored in memory. When it reaches a certain time point, it is reported to the `tarsstat` service (the default is once reporting 1 minute). We call the time gap between the two reporting time points as a statistical interval and perform the operations such as accumulating and comparing the same key in a statistical interval.
The sample code is as follows:

```go
//for error
ReportStat(msg, 0, 1, 0)

//for success
ReportStat(msg, 1, 0, 0)

//func ReportStat(msg *Message, succ int32, timeout int32, exec int32)
//see more detail in tars/statf.go
```

Description:
> * Normally, we don't have to concern about  the Statistical reporting, the tarsgo framework will do this report after every client call server, no matter success or failure. And the success rate, fail rate, average  cost time, and so on will be shown in the web management system if you set up properly.
> *  If the main service is deployed on the web management system, you do not need to define Communicator set the configurations of `tarsregistry`, `tarsstat`, etc., the service will be automatically reported.
> * If the main service or program is not deployed on the web management system, you need to define the Communicator, set the `tarsregistry`, `tarsstat`, etc., so that you can view the service monitoring of the called service on the web management system.
> * The reported data is reported regularly and can be set in the configuration of the communicator.


### 7 Anomaly Reporting
For better monitoring, the TARS framework supports reporting abnormal situation directly to `tarsnotify` in the program and can be viewed on the Web management page.

The framework provides three macros to report different kinds of exceptions:

```go
tars.reportNotifyInfo("Get data from mysql error!")
```

`Info` is a string, which can directly report the string to `tarsnotify`. The reported string can be seen on the page, subsequently, we can alarm according to the reported information.

### 8 Attribute (Property) Statistics
To facilitate business statistics, the TARS framework also supports the display of information on the web management platform. 


The types of statistics currently supported include the following:
> * `Sum(sum)`: calculate the sum of every report value.
> * `Average(avg)`: calculate the average of every report value.
> * `Distribution(distr)`: calculate the distribution of every report, which parameter is a list, and calculate the probability distribution of each interval.
> * `Maximum(max)`: calculate the maximum of every report value.
> * `Minimum(min)`: calculate the minimum of every report value.
> * `Count(count)`: calculate the count of  report times.

The sample code is the following:

```go
    sum := tars.NewSum()
    count := tars.NewCount()
    max := tars.NewMax()
    min := tars.NewMin()
    d := []int{10, 20, 30, 50} 
    distr := tars.NewDistr(d)
    p := tars.CreatePropertyReport("testproperty", sum, count, max, min, distr)
    for i := 0; i < 5; i++ {
        v := rand.Intn(100)
        p.Report(v)

    }
```

Description:
> * Data is reported regularly, and can be set in the configuration of the communicator, currently once per minute.
> * Create a `PropertyReportPtr` function: The parameter `createPropertyReport` can be any collection of statistical methods, the example uses six statistical methods, usually only need to use one or two.
> * Note that when you call `createPropertyReport`, you must create and save the created object after the service is enabled, and then just take the object to report, do not create it each time you use.

### 9 Remote Configuration
Users can set up remote configuration from OSS. See more detail in [tars-config](https://tarscloud.github.io/TarsDocs_en/dev/tarsphp/Framework/tars-config.html)
That is an example to illustrate how to use this API to get a configuration file from remote.

```go
import "github.com/jslyzt/tarsgo/tars"
...
cfg := tars.GetServerConfig()
remoteConf := tars.NewRConf(cfg.App, cfg.Server, cfg.BasePath)
config, _ := remoteConf.GetConfig("test.conf")
...
```

### 10 setting.go
`setting.go` in package tars is used to control tarsgo performance and characteristics. Some options should be updated from `Getserverconfig()`.

```go
//number of worker routines to handle client request
//zero means no control, just one goroutine for a client request.
//runtime.NumCPU() usually best performance in the benchmark.
var MaxInvoke int = 0

const (
	//for now, some option should update from remote config

	//version
	TarsVersion string = "1.0.0"

	//server

	AcceptTimeout time.Duration = 500 * time.Millisecond
	//zero for not set read deadline for Conn (better  performance)
	ReadTimeout time.Duration = 0 * time.Millisecond
	//zero for not set write deadline for Conn (better performance)
	WriteTimeout time.Duration = 0 * time.Millisecond
	//zero for not set deadline for invoke user interface (better performance)
	HandleTimeout  time.Duration = 0 * time.Millisecond
	IdleTimeout    time.Duration = 600000 * time.Millisecond
	ZombileTimeout time.Duration = time.Second * 10
	QueueCap       int           = 10000000

	//client
	ClientQueueLen     int           = 10000
	ClientIdleTimeout  time.Duration = time.Second * 600
	ClientReadTimeout  time.Duration = time.Millisecond * 100
	ClientWriteTimeout time.Duration = time.Millisecond * 3000
	ReqDefaultTimeout  int32         = 3000
	ObjQueueMax        int32         = 10000

	//report
	PropertyReportInterval time.Duration = 10 * time.Second
	StatReportInterval     time.Duration = 10 * time.Second

	//mainloop
	MainLoopTicker time.Duration = 10 * time.Second

	//adapter
	AdapterProxyTicker     time.Duration = 10 * time.Second
	AdapterProxyResetCount int           = 5

	//communicator default, update from remote config
	refreshEndpointInterval int = 60000
	reportInterval          int = 10000
	AsyncInvokeTimeout      int = 3000

	//tcp network config
	TCPReadBuffer  = 128 * 1024 * 1024
	TCPWriteBuffer = 128 * 1024 * 1024
	TCPNoDelay     = false
)
```

### 11 HTTP Support

`tars.TarsHttpMux` is multiplexer like [http.ServeMux](https://golang.org/pkg/net/http/#ServeMux),the `pattern` parameter is used as the interface name in monitoring report. 

Here is a sample of HTTP server:

```go
package main

import (
	"net/http"
	"github.com/jslyzt/tarsgo/tars"
)

func main() {
	mux := &tars.TarsHttpMux{}
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello tafgo"))
	})

    cfg := tars.GetServerConfig()
	tars.AddHttpServant(mux, cfg.App+"."+cfg.Server+".HttpObj") //Register http server
	tars.Run()
}
```

###  12 Using Context
### Context
In the past, TarsGo did not use context in the generated client code, or the implementation code passed in by the user. This makes us want to pass some framework information, such as client IP, port, etc., or the user passes some information about the call chain to the framework, which is difficult to implement. Through a refactoring of the interface, the context is supported, and these information will be implemented through the context. This refactoring is designed to be compatible with older user behavior and is fully compatible.


Server-Side Context:

```go
type ContextTestImp struct {
}
//only need to add  ctx context.Context parameter
func (imp *ContextTestImp) Add(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	//We can use context to get some usefull infomation we need, such Client, ip, port and tracing infomation
	//read more detail under tars/util/current
	ip, ok := current.GetClientIPFromContext(ctx)
    if !ok {
        logger.Error("Error getting ip from context")
    }  
	return 0, nil
}
//just change AddServant into AddServantWithContext
app.AddServantWithContext(imp, cfg.App+"."+cfg.Server+".ContextTestObj")
```

Client-Side Context:

```go

    ctx := context.Background()
    c := make(map[string]string)
    c["a"] = "b" 
//juse change app.Add into app.AddWithContext, now you can pass context to framework, 
//if you want to setting request package's context, you can pass a optional parameter, just like c, which is ...[string]string
    ret, err := app.AddWithContext(ctx, i, i*2, &out, c)
```

Read full demo client and server under `_examples/ContextTestServer`


### 13 Filter & Zipkin Plugin 
For supporting the writing plugin, we provide the filter concept to the framework. We have the client-side filter and the server-side filter. 

```go
//ServerFilter, dispatch and f is passed as parameter,  for dispatching user's implement. 
//req and resp is  
type ServerFilter func(ctx context.Context, d Dispatch, f interface{}, req *requestf.RequestPacket, resp *requestf.ResponsePacket, withContext bool) (err error)
//
type ClientFilter func(ctx context.Context, msg *Message, invoke Invoke, timeout time.Duration) (err error)
//RegisterServerFilter registers the server side filter
//func RegisterServerFilter(f ServerFilter)
//RegisterClientFilter registers the client side filter
//func RegisterClientFilter(f ClientFilter)
```

Having these filters, now we can add OpenTracing for every request.
Let's take a look at the client-side filter for OpenTracing.

```go
//ZipkinClientFilter returns a client side tars filter, for hooking zipking opentracing.
func ZipkinClientFilter() tars.ClientFilter {
	return func(ctx context.Context, msg *tars.Message, invoke tars.Invoke, timeout time.Duration) (err error) {
		var pCtx opentracing.SpanContext
		req := msg.Req
		//If span context is passed in the context, we use this context as parent span, else start a new span.
		//The method name of the rpc request,  is used as span's name.
		if parent := opentracing.SpanFromContext(ctx); parent != nil {
			pCtx = parent.Context()
		}
		cSpan := opentracing.GlobalTracer().StartSpan(
			req.SFuncName,
			opentracing.ChildOf(pCtx),
			ext.SpanKindRPCClient,
		)
		defer cSpan.Finish()
		cfg := tars.GetServerConfig()

		//set additional information for the span, like method, interface, protocol, vesion, ip and port etc.
		cSpan.SetTag("client.ipv4", cfg.LocalIP)
		cSpan.SetTag("tars.interface", req.SServantName)
		cSpan.SetTag("tars.method", req.SFuncName)
		cSpan.SetTag("tars.protocol", "tars")
		cSpan.SetTag("tars.client.version", tars.TarsVersion)

		//inject the span context into the request package's status, which is map[string]string
		if req.Status != nil {
			err = opentracing.GlobalTracer().Inject(cSpan.Context(), opentracing.TextMap, opentracing.TextMapCarrier(req.Status))
			if err != nil {
				logger.Error("inject span to status error:", err)
			}
		} else {
			s := make(map[string]string)
			err = opentracing.GlobalTracer().Inject(cSpan.Context(), opentracing.TextMap, opentracing.TextMapCarrier(s))
			if err != nil {
				logger.Error("inject span to status error:", err)
			} else {
				req.Status = s
			}
		}
		//Nothing tho change,  just invoke the request.
		err = invoke(ctx, msg, timeout)
		if err != nil {
			//invoke error, logging the error information to the span.
			ext.Error.Set(cSpan, true)
			cSpan.LogFields(oplog.String("event", "error"), oplog.String("message", err.Error()))
		}

		return err
	}
```

The server will add filters, which exact the span context from the request package's status and start a new span.

Read more under `TarsGo/tars/plugin/zipkintracing`. For client-side and server-side example code, read `ZipkinTraceClient` & `ZipkinTraceServer` under the examples.

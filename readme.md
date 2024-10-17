### vmess link util 

provides vmess link `encode/decode` methods

#### install 

```shell
go get github.com/clearcodecn/vmess
```

### usage 

```go

// decode vmess link 
var vmessString = "vmess://xxxx"
vmessLink, err := vmess.FromVmessLink(vmessString)


// encode vmess link to "vmess://"

vmessString = vmessLink.ToVmessLink()
```


### 参考

```go
v: 配置文件版本号,主要用来识别当前配置
ps: 备注或别名
add: 地址IP或域名
port: 端口号
id: UUID
aid: alterId
scy: 加密方式(security),没有时值默认auto
net: 传输协议(tcp\kcp\ws\h2\quic)
type: 伪装类型(none\http\srtp\utp\wechat-video) *tcp or kcp or QUIC
host: 伪装的域名
1)http(tcp)->host中间逗号(,)隔开
2)ws->host
3)h2->host
4)QUIC->securty
path: path
1)ws->path
2)h2->path
3)QUIC->key/Kcp->seed
4)grpc->serviceName
tls: 传输层安全(tls)
sni: serverName
alpn: h2,http/1.1
fp: fingerprint
```
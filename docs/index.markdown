---
# Feel free to add content and custom Front Matter to this file.
# To modify the layout, see https://jekyllrb.com/docs/themes/#overriding-theme-defaults

layout: home
---
<img align="center" width="1250" src="./images/tunnel.png">
<iframe src="https://ghbtns.com/github-btn.html?user=narendranathreddythota&repo=podtnl&type=star&count=true&size=large" frameborder="0" scrolling="0" width="170" height="30" title="GitHub"></iframe>
<iframe src="https://ghbtns.com/github-btn.html?user=narendranathreddythota&repo=podtnl&type=fork&count=true&size=large" frameborder="0" scrolling="0" width="170" height="30" title="GitHub"></iframe>
<iframe src="https://ghbtns.com/github-btn.html?user=narendranathreddythota&type=follow&count=true&size=large" frameborder="0" scrolling="0" width="350" height="30" title="GitHub"></iframe>
<hr>
<br>

Access Pod Online using Podtnl
==========
[![wercker status](https://app.wercker.com/status/11cd0df4d8d696f68146c8014eb042c3/s/master "wercker status")](https://app.wercker.com/project/byKey/11cd0df4d8d696f68146c8014eb042c3)
[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go Report Card](https://goreportcard.com/badge/github.com/narendranathreddythota/podtnl)](https://goreportcard.com/report/github.com/narendranathreddythota/podtnl)
[![Release](https://img.shields.io/badge/release-1.0-brightgreen.svg)](https://github.com/narendranathreddythota/podtnl/releases/tag/1.0)

 - Expose your pod to Online easily from any kubernetes clusters without creating a kubernetes service.

 - Clusters including minikube, kind, PKS, AKS, GKE, DK, etc

 - No need to worry about accessing your application during development, forgot about the following buzz words 
   [ingress, controller, loadbalancer, Public IP, etc]

**Podtnl** uses two concepts: 
> - Port Forward
> - Tunnel

## Installation
Podtnl is available in [homebrew](http://brew.sh/)
```shell
$ brew tap narendranathreddythota/podtnl
$ brew install podtnl
```
Download Binary
```shell
https://github.com/narendranathreddythota/podtnl/releases/download/1.0/podtnl
```

**Build from Source**
```shell
$ git clone https://github.com/narendranathreddythota/podtnl
$ cd podtnl
$ go build 
$ cp podtnl /var/local/bin
```

## Tunnel Providers
**Podtnl** is built in a way that it can support any tunnel provider. 
Currently **Podtnl** support only Ngrok as tunnel provider
### Ngrok
  - Install Ngrok from their [website](https://dashboard.ngrok.com/get-started/setup) 

## Usage:
### Available Flags
```shell
  version        : Output Podtnl Version
  provider       : Input Tunnel Provider
  providerPath   : Input Tunnel Provider Path
  podname        : Input Pod Name
  protocol       : Input Type of Protocol
  namespace      : Input Namespace
  podport        : Input Pod Port
  auth           : Need Authentication ? Applicable for HTTP
```
```shell
➜  ~ podtnl --help

 Usage: podtnl := Please Provide necessary Arguments..
  -v    prints current podtnl version
  -auth
        Need to secure the exposed pod with Basic Auth? (default true)
  -namespace string
         Please Provide Namespace where pod is running.. (default "default")
  -podname string
        Please Provide Pod Name
  -podport int
        Please Provide Pod Port
  -protocol string
        Please Provide Type of Protocol HTTP or TCP (default "http")
  -provider string
        Input Tunnel Provider (default "ngrok")
  -providerPath string
        Please Provide Tunnel Provider Path (default "/usr/local/bin/ngrok")
```
### HTTP
```shell
$ podtnl -provider ngrok -podname couchdb0-64d95cccc5-5phqz -podport 5984

Expected Output:
[INFO] ...Tunnel provider ngrok
[INFO] NGROK is Ready
[INFO] Username: QoElOkoMmFv45f8kNKOCSVzyJz9zfakb
[INFO] Password: jApdHSdVLdpXYCkVgYmzYzSD70j9tipdYvgWLhrZ1mdHGvcngdHiHpdvJfQjAins
Forwarding from 127.0.0.1:5984 -> 5984
Forwarding from [::1]:5984 -> 5984
[INFO] mytunnel is created and Live: -> https://fa8df289.ngrok.io
Handling connection for 5984
Handling connection for 5984
Handling connection for 5984
Handling connection for 5984

^C[WARN] Shutting down all open tunnels..
[DBUG] Closing tunnel in https://fa8df289.ngrok.io
```
### TCP
```shell
$ podtnl -provider ngrok -podname orderer1-7cb4b7565-nv95k -podport 7050 -protocol tcp

Expected Output:
[INFO] ...Tunnel provider ngrok
[INFO] NGROK is Ready
[INFO] mytunnel is created and Live: -> tcp://0.tcp.ngrok.io:10467

^C[WARN] Shutting down all open tunnels..
[DBUG] Closing tunnel in tcp://0.tcp.ngrok.io:10467
```
## Note:

- Please switch context to the target cluster which is running in Minikube or Kind or AKS or PKS, etc
- Do not forgot to hit <<<< ^C [CTL + C] >>>> in order to close the open tunnels
  
```shell
$ kubectl config use-context {cluster_name} 
```
`Kubectl get all` should give the following output
```shell
NAME                 TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
service/kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   179d
```
If the following works perfectly 

```shell
$ kubectl port-forward pod/couchdb0-64d95cccc5-5phqz 5984:5984
```
Then the following also works
```shell
$ podtnl -provider ngrok -providerPath /usr/local/bin/ngrok -podname couchdb0-64d95cccc5-5phqz -podport 5984 -auth=false
```

## Demo:
<img position="absolute" width="850" src="./images/recorder.gif">
<hr>
<br>
## Error Scenarios:
### Wrong Pod 
```shell
➜  ~ podtnl -provider ngrok -podname somedummypodnameorwrongname -podport 5984
[INFO] ...Tunnel provider ngrok
[FTAL] error upgrading connection: pods "somedummypodnameorwrongname" not found
```
### Wrong Provider Path 
```shell
➜  ~ podtnl -provider ngrok -providerPath /usr/local/bin/ng -podname $POD -podport 5984
[INFO] ...Tunnel provider ngrok
2020/05/24 00:39:59 fork/exec /usr/local/bin/ng: no such file or directory
```
### Socat Error
#### For wrong Port or unable to port forward situations
```shell
E0524 00:42:00.609685   24854 portforward.go:400] an error occurred forwarding 5989 -> 5989: error forwarding port 5989 to pod 2f82fec0fb08efb22d2efcf02beb02b89a73398dd82c9a3e82103346c3f074f3, uid : exit status 1: 2020/05/23 20:42:00 socat[30542] E connect(5, AF=2 127.0.0.1:5989, 16): Connection refused
```
[Thank you gonnel](https://github.com/afdalwahyu/gonnel)

Licensing
=========
**Podtnl** is licensed under the Apache License, Version 2.0. See
[LICENSE](https://github.com/narendranathreddythota/podtnl/blob/master/LISCENSE) for the full
license text.
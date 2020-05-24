/*
Copyright [2020] [Narendranath Reddy]

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
package main
*/

package flags

import (
	"flag"

	"github.com/narendranathreddythota/podtnl/tunnel/types"
)

var provider string
var protocol string
var providerPath string
var podname string
var auth bool
var ns string
var podport int

func init() {

	flag.StringVar(&provider, "provider", "ngrok", "Provides Tunnel provider")
	flag.StringVar(&providerPath, "providerPath", "/usr/local/bin/ngrok", "Tunnel provider Path")
	flag.StringVar(&podname, "podname", "", "Pod Name")
	flag.StringVar(&protocol, "protocol", types.Protocol(types.HTTP).ToString(), "Type of Protocol HTTP or TCP")
	flag.StringVar(&ns, "namespace", "default", "Namespace where pod is running..")
	flag.IntVar(&podport, "podport", 0, "Pod Port")
	flag.BoolVar(&auth, "auth", true, "Need to secure the exposed pod with Basic Auth?")
	flag.Parse()
}

func GetTunnelProvider() string {

	return provider
}

func GetSelectedProtocol() string {

	if protocol == types.Protocol(types.TCP).ToString() {
		return types.Protocol(types.TCP).ToString()
	}
	return protocol
}

func GetTunnelProviderPath() string {

	return providerPath
}
func GetPodName() string {

	return podname
}
func GetAuth() bool {

	return auth
}
func GetPodPort() int {

	return podport
}

func GetNamespace() string {

	return ns
}

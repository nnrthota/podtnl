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
	"fmt"
	"os"

	"github.com/narendranathreddythota/podtnl/tunnel/types"
)

var provider string
var protocol string
var providerPath string
var podname string
var auth bool
var ns string
var podport int

// AppVersion 1.0
const AppVersion = "1.0"

// MyUsage Special instructions
func MyUsage() {
	fmt.Printf(" Expose your pod to Online easily from any kubernetes clusters without creating a kubernetes service.\n")
	fmt.Printf(`
  Available Flage:

  provider       : Input Tunnel Provider
  providerPath   : Input Tunnel Provider Path
  podname        : Input Pod Name
  protocol       : Input Type of Protocol
  namespace      : Input Namespace
  podport        : Input Pod Port
  auth           : Need Authentication ? Applicable for HTTP
	`)
	fmt.Printf("\n Usage: %s Please Provide necessary Arguments..\n", os.Args[0])
	flag.PrintDefaults()
}

//InitFlags init all flags
func InitFlags() {
	initEnvFlag()
}

func initEnvFlag() {
	version := flag.Bool("v", false, "prints current podtnl version")
	flag.StringVar(&provider, "provider", "ngrok", "Input Tunnel Provider")
	flag.StringVar(&providerPath, "providerPath", "/usr/local/bin/ngrok", "Please Provide Tunnel Provider Path")
	flag.StringVar(&podname, "podname", "", "Please Provide Pod Name")
	flag.StringVar(&protocol, "protocol", types.Protocol(types.HTTP).ToString(), "Please Provide Type of Protocol HTTP or TCP")
	flag.StringVar(&ns, "namespace", "default", " Please Provide Namespace where pod is running..")
	flag.IntVar(&podport, "podport", 0, "Please Provide Pod Port")
	flag.BoolVar(&auth, "auth", true, "Need to secure the exposed pod with Basic Auth?")
	flag.Parse()
	if *version {
		fmt.Println(AppVersion)
		os.Exit(0)
	}
}

func veryFlagInput() {
	if len(os.Args) < 2 {
		flag.Usage = MyUsage
		flag.Usage()
		os.Exit(1)
	}
}

//GetTunnelProvider return tunnel provider
func GetTunnelProvider() string {
	veryFlagInput()
	return provider
}

//GetSelectedProtocol return protocol
func GetSelectedProtocol() string {
	veryFlagInput()
	if protocol == types.Protocol(types.TCP).ToString() {
		return types.Protocol(types.TCP).ToString()
	}
	return protocol
}

// GetTunnelProviderPath return selected provider path
func GetTunnelProviderPath() string {
	veryFlagInput()
	return providerPath
}

// GetPodName return provided podname
func GetPodName() string {
	veryFlagInput()
	return podname
}

// GetAuth return provided auth status
func GetAuth() bool {
	veryFlagInput()
	return auth
}

// GetPodPort return provided port number
func GetPodPort() int {
	veryFlagInput()
	return podport
}

// GetNamespace return provided namespace
func GetNamespace() string {
	veryFlagInput()
	return ns
}

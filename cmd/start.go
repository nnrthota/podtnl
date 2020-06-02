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

package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/kataras/golog"
	sysFlags "github.com/narendranathreddythota/podtnl/flags"
	"github.com/narendranathreddythota/podtnl/kubeconfig"
	"github.com/narendranathreddythota/podtnl/portforward"
	"github.com/narendranathreddythota/podtnl/tunnel"
	providers "github.com/narendranathreddythota/podtnl/tunnel/providers"
)

func runPortForward(podName string, exposedPort int) {
	config, _ := kubeconfig.GetKubeConfig()
	stopChan := make(chan struct{})
	defer close(stopChan)
	if err := portforward.PortForward(config, stopChan, podName, exposedPort); err != nil {
		golog.Fatal(err)
	}
}

func init() {
	golog.SetTimeFormat("")
	golog.SetLevel("debug")
}

func openTunnels(provider providers.ITunnelProvider) []*providers.Tunnel {
	auth := sysFlags.GetAuth()
	proto := sysFlags.GetSelectedProtocol()
	tunnels := []*providers.Tunnel{
		{
			Proto:     proto,
			Name:      "mytunnel",
			LocalIP:   "127.0.0.1",
			LocalPort: strconv.Itoa(int(sysFlags.GetPodPort())),
		},
	}
	if auth {
		username := GetCode(false)
		password := GetCode(true)
		tunnels[0].Auth = fmt.Sprintf("%s:%s", username, password)
		golog.Info("Username: ", username)
		golog.Info("Password: ", password)
	}
	provider.OpenManyTunnels(tunnels)
	return tunnels
}

//Start the podtnl
func Start() error {
	sysFlags.InitFlags()
	podname := sysFlags.GetPodName()
	podport := sysFlags.GetPodPort()
	go runPortForward(podname, podport)

	provider := tunnel.GetTunnelProvider()

	defer provider.End()
	provider.Start()

	tunnels := openTunnels(provider)

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
	<-signalChannel

	golog.Warn("Shutting down all open tunnels..")
	provider.CloseManyTunnels(tunnels)
	return nil
}

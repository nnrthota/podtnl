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

package tunnel

import (
	"strings"

	"github.com/kataras/golog"
	sysFlags "github.com/narendranathreddythota/podtnl/flags"
	providers "github.com/narendranathreddythota/podtnl/tunnel/providers"
	ngrok "github.com/narendranathreddythota/podtnl/tunnel/providers/ngrok"
)

const (
	//NGROK ngrok
	NGROK = "ngrok"
)

//GetTunnelProvider factory where different tunnels are supported
func GetTunnelProvider() providers.ITunnelProvider {

	provider := sysFlags.GetTunnelProvider()
	providerPath := sysFlags.GetTunnelProviderPath()
	golog.Info("...Tunnel provider ", provider)
	if strings.ToLower(provider) == NGROK {
		return ngrok.NewNgrokProvider(providerPath)
	}
	return nil
}

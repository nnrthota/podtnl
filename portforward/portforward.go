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

package portforward

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/kataras/golog"
	sysFlags "github.com/narendranathreddythota/podtnl/flags"
	"k8s.io/apimachinery/pkg/util/httpstream"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/portforward"
	"k8s.io/client-go/transport/spdy"
)

func getDialer(podName string, config *rest.Config) (httpstream.Dialer, error) {
	ns := sysFlags.GetNamespace()
	path := fmt.Sprintf("/api/v1/namespaces/%s/pods/%s/portforward", ns, podName)
	hostIP := strings.TrimLeft(config.Host, "https:/")
	serverURL := url.URL{
		Scheme: "https",
		Path:   path,
		Host:   hostIP,
	}
	roundTripper, upgrader, err := spdy.RoundTripperFor(config)
	if err != nil {
		golog.Error(err)
		return nil, err
	}
	dialer := spdy.NewDialer(upgrader,
		&http.Client{Transport: roundTripper},
		http.MethodPost, &serverURL,
	)
	return dialer, nil
}

func getForwarder(ports []string, stopChan chan struct{}, dialer httpstream.Dialer) (*portforward.PortForwarder, error) {
	readyChan := make(chan struct{})
	forwarder, err := portforward.New(dialer,
		ports,
		stopChan,
		readyChan,
		os.Stdout,
		os.Stderr,
	)
	if err != nil {
		golog.Error("error while preparing dialer: ", err)
		return nil, err
	}
	return forwarder, nil
}

// PortForward This piece of function makes your pod available to localhost which is running in any kube cluster
func PortForward(config *rest.Config, stopChan chan struct{}, podName string, selectedPort int) error {

	errChan := make(chan error)
	ports := []string{
		strconv.Itoa(selectedPort) + ":" + strconv.Itoa(selectedPort),
	}

	dialer, err := getDialer(podName, config)
	if err != nil {
		return err
	}

	forwarder, err := getForwarder(ports, stopChan, dialer)
	if err != nil {
		return err
	}

	if err := forwarder.ForwardPorts(); err != nil {
		return err
	}

	go func() {
		errChan <- forwarder.ForwardPorts()
		close(errChan)
	}()

	<-forwarder.Ready
	forwardedPorts, err := forwarder.GetPorts()

	if err != nil {
		return err
	}

	if len(forwardedPorts) != 1 {
		golog.Error("expected 1 port, got ", len(ports))
		return errors.New("Error in the port forward")
	}

	port := forwardedPorts[0]
	if port.Local == 0 {
		return errors.New("Error in the port forward")
	}

	return nil
}

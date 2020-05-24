package ngrok

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/kataras/golog"
	tunnelProvider "github.com/narendranathreddythota/podtnl/tunnel/providers"
)

// NgrokProvider object
type NgrokProvider struct {
	Client NGROKClient
}

//NewNgrokProvider return provider
func NewNgrokProvider(binaryPath string) tunnelProvider.ITunnelProvider {
	client, _ := NewClient(Options{
		BinaryPath: binaryPath,
	})
	return &NgrokProvider{
		Client: *client,
	}
}

//Start NGROK server
func (np *NgrokProvider) Start() error {
	done := make(chan bool)
	go np.Client.StartServer(done)
	<-done
	return nil
}

//CreateTunnel open tunnel using NGROK Server
func (np *NgrokProvider) CreateTunnel(t *tunnelProvider.Tunnel) error {

	for attempt := uint(0); attempt <= maxRetries; attempt++ {
		err := func() error {
			//log.Printf("Creating tunnel %d attempt \n", attempt)
			time.Sleep(1 * time.Second)
			var record TunnelResponse
			jsonData := map[string]interface{}{
				"addr":    fmt.Sprintf("%s:%s", t.LocalIP, t.LocalPort),
				"proto":   t.Proto,
				"name":    t.Name,
				"inspect": t.Inspect,
				"auth":    t.Auth,
			}

			if string(t.Proto) == "http" {
				jsonData["bind_tls"] = true
			}
			url := fmt.Sprintf("http://%s/api/tunnels", np.Client.API)
			jsonValue, err := json.Marshal(jsonData)
			if err != nil {
				golog.Error(err)
				return err
			}
			res, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
			if err != nil {
				golog.Error(err)
				return err
			}
			defer res.Body.Close()

			if res.StatusCode < 200 || res.StatusCode > 299 {
				res, _ := ioutil.ReadAll(res.Body)
				return errors.New("error api: " + string(res))
			}

			if err := json.NewDecoder(res.Body).Decode(&record); err != nil {
				return err
			}

			t.RemoteAddress = record.PublicURL
			t.IsCreated = true
			golog.Info(t.Name + " is created and Live: -> " + t.RemoteAddress)
			return nil
		}()
		if np.Client.LogAPI && err != nil {
			golog.Error(err)
		}
		if err == nil {
			break
		}
	}
	return nil
}

//CloseTunnel it closes open tunnels
func (np *NgrokProvider) CloseTunnel(t *tunnelProvider.Tunnel) error {
	for attempt := uint(0); attempt <= maxRetries; attempt++ {
		err := func() error {
			golog.Debug("Closing tunnel in " + t.RemoteAddress)
			url := fmt.Sprintf("http://%s/api/tunnels/%s", np.Client.API, t.Name)
			req, err := http.NewRequest("DELETE", url, nil)
			if err != nil {
				log.Println(err)
				return err
			}
			client := &http.Client{}
			res, err := client.Do(req)
			if err != nil {
				log.Println(err)
				return err
			}
			defer res.Body.Close()

			if res.StatusCode < 200 || res.StatusCode > 299 {
				res, _ := ioutil.ReadAll(res.Body)
				return errors.New("error api: " + string(res))
			}

			t.RemoteAddress = ""
			t.IsCreated = false
			golog.Info("Tunnel " + t.Name + " successfully closed")
			return nil
		}()
		if np.Client.LogAPI && err != nil {
			golog.Error(err)
		}
		if err == nil {
			break
		}
	}
	return nil
}

//End close the running NGROK Server
func (np *NgrokProvider) End() error {
	return np.Client.runningCmd.Process.Kill()
}

// OpenManyTunnels open multiple tunnels simultaneously
func (np *NgrokProvider) OpenManyTunnels(tunnels []*tunnelProvider.Tunnel) error {
	wg := &sync.WaitGroup{}
	// api request post to /api/tunnels
	if len(tunnels) < 1 {
		return errors.New("need at least 1 tunnel to connect")
	}

	for _, t := range tunnels {
		if !t.IsCreated {
			wg.Add(1)
			go func(x *tunnelProvider.Tunnel) {
				np.CreateTunnel(x)
				wg.Done()
			}(t)
		}
	}

	wg.Wait()
	return nil
}

// CloseManyTunnels many tunnels
func (np *NgrokProvider) CloseManyTunnels(tunnels []*tunnelProvider.Tunnel) error {
	wg := &sync.WaitGroup{}
	//	api request delete to /api/tunnels/:Name
	if len(tunnels) < 1 {
		return errors.New("need at least 1 tunnel to disconnect")
	}

	for _, t := range tunnels {
		if t.IsCreated {
			wg.Add(1)
			go func() {
				np.CloseTunnel(t)
				wg.Done()
			}()
		}
	}

	wg.Wait()
	return nil
}

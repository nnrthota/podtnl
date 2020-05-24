package ngrok

import (
	"bytes"
	"errors"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"regexp"
	"syscall"

	"github.com/kataras/golog"
	tunnelProv "github.com/narendranathreddythota/podtnl/tunnel/providers"
)

type Options struct {
	SubDomain  string // Sub domain config if you're using premium plan
	AuthToken  string // Auth token to authenticate client
	Region     string // Region that will tunneling from
	ConfigPath string // Path config to store auth token or specific WebUI port
	BinaryPath string // Binary file that will be running
	LogBinary  bool   // You can watch binary log or not
}
type NGROKClient struct {
	Options    *Options             // Options that will be used for command
	Tunnel     []*tunnelProv.Tunnel // List of all tunnel
	API        string               // Client server for API communication
	LogApi     bool                 // Log response from API or not
	commands   []string             // result of commands that will be used to run binary
	runningCmd *exec.Cmd            // Pointer of command that running
}

func NewClient(opt Options) (*NGROKClient, error) {
	if opt.BinaryPath == "" {
		return nil, errors.New("binary path required")
	}

	if opt.Region == "" {
		opt.Region = "us"
	}

	if opt.AuthToken != "" {
		err := opt.AuthTokenCommand()
		if err != nil {
			return nil, err
		}
	}

	c := NGROKClient{Options: &opt}
	return &c, nil
}

// AuthTokenCommand that will be authenticate api token
func (o *Options) AuthTokenCommand() error {
	if o.AuthToken == "" {
		return errors.New("token missing")
	}

	if o.BinaryPath == "" {
		return errors.New("binary path file is missing")
	}

	commands := make([]string, 0)
	commands = append(commands, []string{"authtoken", o.AuthToken}...)

	if o.ConfigPath != "" {
		commands = append(commands, "--config"+o.ConfigPath)
	}

	cmd := exec.Command(o.BinaryPath, commands...)
	var outBuffer, errBuffer bytes.Buffer
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer

	if err := cmd.Start(); err != nil {
		return err
	}

	if errBuffer.String() != "" {
		return errors.New(errBuffer.String())
	}

	log.Println(outBuffer.String())
	return nil
}

// StartServer will be run command from previous options
//
// Channel needed to send information about WebUI started or not.
// stdout will be pipe and check using regex.
func (c *NGROKClient) StartServer(isReady chan bool) {
	commands := c.Options.generateCommands()
	cmd := exec.Command(c.Options.BinaryPath, commands...)
	c.runningCmd = cmd
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(
		signalChan, syscall.SIGHUP,
		syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGQUIT)
	go c.handleSignalInput(signalChan)

	checkNGReady, err := regexp.Compile(ngReady)
	if err != nil {
		log.Fatalln(err)
	}

	checkNGInUse, err := regexp.Compile(ngInUse)
	if err != nil {
		log.Fatalln(err)
	}

	checkSessionLimit, err := regexp.Compile(ngSessionLimited)
	if err != nil {
		log.Fatalln(err)
	}

	checkWebURI, err := regexp.Compile(webURI)
	if err != nil {
		log.Fatalln(err)
	}

	chunk := make([]byte, 256)
	for {
		n, err := stdout.Read(chunk)
		if err != nil {
			log.Fatalln(err)
			os.Exit(1)
		}

		if n < 1 {
			continue
		}

		if c.Options.LogBinary {
			log.Print("Client-Bin-Log: ", string(chunk[:n]))
		}
		// handle regex (output) that search local ip and port for web ui
		if checkNGReady.Match(chunk[:n]) {
			host := checkWebURI.FindStringSubmatch(string(chunk[:n]))
			if len(host) >= 1 {
				golog.Info("NGROK is Ready")
				c.API = host[0]
				isReady <- true
			}
		}
		if checkNGInUse.Match(chunk[:n]) {
			golog.Fatal("Address already in use")
		}
		if checkSessionLimit.Match(chunk[:n]) {
			golog.Error("Limit session reached for this account")
		}
	}
}

// generateCommands return array of commands
// that will be run on binary
func (o *Options) generateCommands() []string {
	commands := make([]string, 0)
	commands = append(commands, []string{"start", "--none", "--log=stdout"}...)
	commands = append(commands, "--region="+o.Region)

	if o.ConfigPath != "" {
		commands = append(commands, "--config="+o.ConfigPath)
	}

	if o.SubDomain != "" {
		commands = append(commands, "-subdomain="+o.SubDomain)
	}

	return commands
}

// handleSignalInput to handle signal form command,
// is program received signal or not
func (c *NGROKClient) handleSignalInput(signalChan chan os.Signal) {
	for {
		s := <-signalChan
		switch s {
		default:
			log.Println(s)
			c.Signal(s)
			os.Exit(1)
		}
	}
}

// Close running command and send kill signal to ngrok binary
func (c *NGROKClient) Close() error {
	return c.runningCmd.Process.Kill()
}

// Signal handle signal input and proceed to command
func (c *NGROKClient) Signal(signal os.Signal) error {
	return c.runningCmd.Process.Signal(signal)
}

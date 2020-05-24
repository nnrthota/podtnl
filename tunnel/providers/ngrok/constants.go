package ngrok

// Constant regex that will be used for handling stdout command
const (
	maxRetries       = 100
	ngReady          = `starting web service.*addr=(\d+\.\d+\.\d+\.\d+:\d+)`   // check if ngrok ready
	ngInUse          = `address already in use`                                // check if port already in use
	ngSessionLimited = `is limited to (\d+) simultaneous ngrok client session` // Check limit ngrok
	webURI           = `\d+\.\d+\.\d+\.\d+:\d+`                                // Find client server
)

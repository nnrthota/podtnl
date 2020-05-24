package ngrok

// TunnelResponse NGROK resp after creating tunnel
type TunnelResponse struct {
	Name      string `json:"Name"`
	URI       string `json:"uri"`
	PublicURL string `json:"public_url"`
	Proto     string `json:"Proto"`
	Config    struct {
		Addr    string `json:"addr"`
		Inspect bool   `json:"Inspect"`
	} `json:"config"`
	Metrics struct {
		Conns info `json:"conns"`
		HTTP  info `json:"http"`
	} `json:"metrics"`
}

type info struct {
	Count  int `json:"count"`
	Gauge  int `json:"gauge"`
	Rate1  int `json:"rate1"`
	Rate5  int `json:"rate5"`
	Rate15 int `json:"rate15"`
	P50    int `json:"p50"`
	P90    int `json:"p90"`
	P95    int `json:"p95"`
	P99    int `json:"p99"`
}

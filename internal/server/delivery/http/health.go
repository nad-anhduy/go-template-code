package httpdelivery

type healthResponse struct {
	Status     string            `json:"status"`
	Components map[string]string `json:"components,omitempty"`
}

package aksk

type Config struct {
	Name       string       `json:"name"`
	Driver     string       `json:"driver"`
	akskConfig []AKSKConfig `json:"user"`
}

type AKSKConfig struct {
	AK             string            `json:"ak"`
	SK             string            `json:"sk"`
	HideCredential bool              `json:"hide_credential"`
	Labels         map[string]string `json:"labels"`
	Expire         int64             `json:"expire"`
}

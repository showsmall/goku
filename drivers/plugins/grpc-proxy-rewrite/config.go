package grpc_proxy_rewrite

type Config struct {
	Service         string            `json:"service" label:"服务名称"`
	Method          string            `json:"method" label:"方法名称"`
	Authority       string            `json:"authority" label:"虚拟主机域名(Authority)"`
	Headers         map[string]string `json:"headers" label:"请求头部"`
	Tls             bool              `json:"tls" label:"TLS传输"`
	SkipCertificate bool              `json:"skip_certificate" label:"跳过证书检查"`
}

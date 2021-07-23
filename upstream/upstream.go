package upstream

import (
	http_context "github.com/eolinker/goku-eosc/node/http-context"
	"github.com/eolinker/goku-eosc/node/http-proxy/backend"
	"github.com/eolinker/goku-eosc/service"
)

//CheckSkill 检测目标技能是否符合
func CheckSkill(skill string) bool {
	return skill == "github.com/eolinker/goku-eosc/upstream.upstream.IUpstream"
}

//IUpstream 实现了负载发送请求方法
type IUpstream interface {
	Send(ctx *http_context.Context, serviceDetail service.IServiceDetail) (backend.IResponse, error)
}

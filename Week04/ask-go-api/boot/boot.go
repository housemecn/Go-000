package boot

import (
	"github.com/gogf/gf/frame/g"

	_ "git.code.tencent.com/houseme/ask-go-api/packed"
)

func init() {
	env := g.Cfg().GetString("env.app_env")
	g.Cfg().SetFileName("config." + env + ".toml")
}

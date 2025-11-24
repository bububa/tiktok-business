package tool

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tool"
)

// Language 获取受众语言列表
// 您可以使用本接口获取受众语言代码枚举值列表。
// 获取语言代码后，可将语言代码传入 /adgroup/create/ 中的 languages 字段用于受众语言定向。
func Language(ctx context.Context, clt *core.SDKClient, req *tool.LanguageRequest, accessToken string) ([]tool.Language, error) {
	var ret tool.LanguageResponse
	if err := clt.Get(ctx, "v1.3/tool/language/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data.Languages, nil
}

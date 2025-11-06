package business

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business"
)

// Benchmark 获取企业类别的基准数据
// 您可以使用本节课获取某个企业类别的基准数据。
// 获取企业类别的基准数据后，您可使用 /business/get/ 拉取您的企业号的主页数据，并依据行业标准评估该企业号的表现。
func Benchmark(ctx context.Context, clt *core.SDKClient, req *business.BenchmarkRequest, accessToken string) (*business.Benchmark, error) {
	var resp business.BenchmarkResponse
	if err := clt.Get(ctx, "v1.3/business/benchmark/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

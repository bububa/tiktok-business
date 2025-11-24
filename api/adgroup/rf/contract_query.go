package rf

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/adgroup/rf"
)

// ContractQuery 查询合同有效期
// 您可以使用本接口接口来获取您的品牌商务合同有效期。商务合同是具有时效性的，您可以使用本接口通过合同有效期内的一个日期，来查询商务合同的起止日期。
func ContractQuery(ctx context.Context, clt *core.SDKClient, req *rf.ContractQueryRequest, accessToken string) (*rf.ContractQueryResult, error) {
	var resp rf.ContractQueryResponse
	if err := clt.Get(ctx, "v1.3/rf/contract/query/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

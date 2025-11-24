package tool

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tool"
)

// DeviceModel 获取设备机型列表
func DeviceModel(ctx context.Context, clt *core.SDKClient, req *tool.DeviceModelRequest, accessToken string) ([]tool.DeviceModel, error) {
	var ret tool.DeviceModelResponse
	if err := clt.Get(ctx, "v1.3/tool/device_model/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data.DeviceModels, nil
}

package music

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/file/music"
)

// Get 获取音乐列表
// 您可以使用本接口获取可用于创建视频或轮播广告的音乐列表。列表中包括广告主上传的私有音乐素材（简称用户音乐）或者音频库（Audio Library，又称商用音乐库）中提供的公用音乐素材（简称系统音乐）。音频库更新频率比较低，不需要频繁获取。建议本地缓存，一个月更新一次即可。
func Get(ctx context.Context, clt *core.SDKClient, req *music.GetRequest, accessToken string) (*music.GetResult, error) {
	var ret music.GetResponse
	if err := clt.Get(ctx, "v1.3/file/music/get/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}

package music

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
)

type Music struct {
	// MusicID 音乐 ID
	MusicID string `json:"music_id,omitempty"`
	// MaterialID 素材 ID
	MaterialID string `json:"material_id,omitempty"`
	// Sources 音乐来源。USER 表示用户上传的音乐，SYSTEM 表示系统音乐
	Sources []enum.MusicSource `json:"sources,omitempty"`
	// Author 仅为可用于轮播广告的音乐返回本字段。
	// 音乐作者
	Author string `json:"author,omitempty"`
	// Liked 仅为可用于轮播广告的音乐返回本字段。
	// 此音乐是否已添加至收藏列表中
	Liked bool `json:"liked,omitempty"`
	// CoverURL 仅为可用于轮播广告的音乐返回本字段。
	// 音乐的封面图片的URL
	CoverURL string `json:"cover_url,omitempty"`
	// URL 音乐试听链接。有效时长：12h
	URL string `json:"url,omitempty"`
	// Duration 音乐时长，单位秒。仅系统音乐返回该字段
	Duration float64 `json:"duration,omitempty"`
	// Style 音乐风格。仅系统音乐返回该字段
	Style string `json:"style,omitempty"`
	// Signature 音乐文件MD5
	Signature string `json:"signature,omitempty"`
	// FileName 文件名称
	FileName string `json:"file_name,omitempty"`
	// Copyright 版权信息。枚举值：MUSIC_FORBID_VIDEO_ALLOW (音乐不允许下载，但是生成的视频允许下载) ，MUSIC_FORBID_VIDEO_FORBID（音乐不允许下载，且生成的视频不允许下载）
	Copyright string `json:"copyright,omitempty"`
	// CreateTime 创建时间。自定义音乐会返回该字段，系统音乐不返回该字段。UTC 时间，格式：2020-06-10T07:39:14Z
	CreateTime model.DateTime `json:"create_time,omitzero"`
	// ModifyTime 修改时间。自定义音乐会返回该字段，系统音乐不返回该字段。UTC 时间，格式：2020-06-10T07:39:14Z
	ModifyTime model.DateTime `json:"modify_time,omitzero"`
}

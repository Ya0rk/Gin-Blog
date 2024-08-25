package model

import (
	"Gin-Blog/utils"
	"Gin-Blog/utils/errmsg"
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
)

var AccessKey = utils.AccessKey
var SecretKey = utils.SecretKey
var Bucket = utils.Bucket
var ImgUrl = utils.QiniuServer

/*
使用七牛云的api 将图片上传到七牛存储 具体方法参考七牛文档
用第三放暂存图片 不会占用本地带宽和减少磁盘读写 加载更快
*/

func UpLoadFile(file multipart.File, fileSize int64) (string, int) {
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}

	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	putExtra := storage.PutExtra{}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		return "", errmsg.ERROR
	}

	url := ImgUrl + ret.Key
	return url, errmsg.SUCCESS
}

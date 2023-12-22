package util

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
	"strconv"
	"time"
)

func getClient() *gclient.Client {
	client := g.Client()
	client.SetTimeout(time.Second * 10)
	return client
}

// DownloadFile :下载文件到指定路径下
func DownloadFile(ctx context.Context, url string, filePath string) error {
	data := g.Client().GetBytes(ctx, url)
	if data == nil {
		return gerror.New("下载文件失败")
	}

	// 将字节数组写入到文件中
	err := gfile.PutBytes(filePath, data)
	if err != nil {
		return gerror.Wrap(err, "下载文件，文件写入失败")
	}
	return nil
}

func UploadFile(ctx context.Context, url string, filePath string) error {
	cli := g.Client()
	cli.SetTimeout(time.Second * 60)
	data := gfile.GetBytes(filePath)
	if data == nil {
		return gerror.New("读取文件失败")
	}
	cli.SetHeader("Content-Type", "application/octet-stream")
	cli.SetHeader("x-oss-meta-author", "aliy")
	cli.SetHeader("Content-Length", strconv.Itoa(len(data)))

	resp, err := cli.Put(ctx, url, data)
	if err != nil {
		return gerror.Wrap(err, "上传文件失败")
	}
	defer func(resp *gclient.Response) {
		_ = resp.Close()
	}(resp)
	return nil
}

func Get(ctx context.Context, url string) (string, error) {
	glog.Info(ctx, "get url:", url)
	resp, err := getClient().DoRequest(ctx, "GET", url)

	if err != nil {
		return "", err
	}
	defer func(resp *gclient.Response) {
		_ = resp.Close()
	}(resp)
	return resp.ReadAllString(), nil
}

func GetObject(ctx context.Context, url string, v interface{}) error {
	json, err := Get(ctx, url)
	if err != nil {
		return err
	}
	return gjson.DecodeTo(json, v)
}

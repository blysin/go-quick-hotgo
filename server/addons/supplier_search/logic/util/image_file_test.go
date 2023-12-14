package util

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"testing"
)

func TestImageFileModify(t *testing.T) {
	fmt.Println("TestImageFileModify")

	imageFile := gfile.GetContents("C:\\Users\\daishaokun1\\Desktop\\6.18\\images.txt")
	imageFileMap := toMap(imageFile)
	//g.Dump(imageFileMap)

	//读取docker-compose文件
	dockerCompose := gfile.GetContents("C:\\Users\\daishaokun1\\Desktop\\6.18\\changtai-docker-compose-app.yml")

	//逐行分析docker-compose文件，如果是‘image:’开头的，就根据imageFileMap进行替换
	for _, v := range gstr.SplitAndTrim(dockerCompose, "\n") {
		v = gstr.Trim(v)
		if gstr.HasPrefix(v, "image:") {
			imageUrl := gstr.TrimLeft(v, "image:")

			imageName := gstr.SplitAndTrim(imageUrl, ":")[0]

			//fmt.Println("解析到镜像名称", imageName)

			newImageUrl := imageFileMap[imageName]
			if newImageUrl != "" {
				fmt.Println("替换镜像地址", imageUrl, "为", newImageUrl)
				dockerCompose = gstr.Replace(dockerCompose, imageUrl, newImageUrl)
			}
		}
	}

	err := gfile.PutContents("C:\\Users\\daishaokun1\\Desktop\\6.18\\changtai-docker-compose-app-new.yml", dockerCompose)
	if err != nil {
		return
	}
}

func toMap(imageFile string) map[string]string {
	m := make(map[string]string)
	for _, v := range gstr.SplitAndTrim(imageFile, "\n") {
		v = gstr.Trim(v)
		if len(v) == 0 {
			continue
		}
		// 将镜像地址和版本号分开，key为镜像地址，value为版本号
		key := gstr.SplitAndTrim(v, ":")[0]
		if m[key] == "" {
			m[key] = v
		} else {
			fmt.Println("镜像地址重复", key)
			//字符串比较哪个比较大
			if gstr.Compare(m[key], v) == 1 {
				fmt.Println("替换镜像地址", m[key], "为", v)
				m[key] = v
			}
		}

	}
	return m
}

func TestCheckMulti(t *testing.T) {
	imageFile := gfile.GetContents("C:\\Users\\daishaokun1\\Desktop\\6.18\\images.txt")

	m := make(map[string]string)
	for _, v := range gstr.SplitAndTrim(imageFile, "\n") {
		v = gstr.Trim(v)
		if len(v) == 0 {
			continue
		}
		// 将镜像地址和版本号分开，key为镜像地址，value为版本号
		key := gstr.SplitAndTrim(v, ":")[0]
		if m[key] == "" {
			m[key] = v
		} else {
			t.Error("镜像地址重复", key)
			//字符串比较哪个比较大
			if gstr.Compare(m[key], v) == 1 {
				fmt.Println("替换镜像地址", m[key], "为", v)
				m[key] = v
			}
		}

	}
}

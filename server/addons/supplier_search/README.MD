## 供应商检索

### 简介

数据上传、供应商检索


### 使用说明

1、xxx
2、xxx
3、xxx


### 迁移或安装

1、安装 HotGo (2.11.5及以上)

项目介绍：https://github.com/bufanyun/hotgo

2、将当前插件项目拷贝进 HotGo 根目录的 server/addons 目录下

3、在 HotGo 根目录的 server/addons/modules 目录下创建go文件:supplier_search.go，内容如下：
```go
package modules

import _ "hotgo/addons/supplier_search"
```

4、HotGo 后台进入 开发工具->插件管理->找到 供应商检索 (supplier_search) 进行安装

5、重启服务即可生效


### 常用命令行

```shell
# 接口维护-gen service
gf gen service -s=addons/supplier_search/logic -d=addons/supplier_search/service

```

### 语法记录

切片参数的不同传参方式

```go
func UpdateIndex(ctx context.Context, detail []*entity.VendorDetail, index int) (err error){}
func UpdateIndex(ctx context.Context, detail *[]entity.VendorDetail, index int) (err error){}
```

- 指针类型
    - 写法1： `detail []*entity.VendorDetail` 表示 detail 是一个指向 entity.VendorDetail 类型的指针的**切片**
      。这意味着你可以修改切片中的元素。
    - 写法2： `detail *[]entity.VendorDetail` 表示 detail 是一个指向 entity.VendorDetail 类型的切片的**指针**
      。这意味着你可以修改切片本身，但不能修改切片中的元素。
- 修改权限：
    - 写法1：可以修改切片中的元素，不能修改切片大小
    - 写法2：可以修改切片大小，但不能修改切片中的元素
- 内存效率：
    - 写法1：比较高效，go的切片作为参数传递是，传递的是切片的引用，不会复制数据，而且方法内部修改切片元素，会影响到原始切片
    - 写法2：比较低效，但是能调整切片大小
- 使用建议：尽量使用 `detail []*entity.VendorDetail` 的写法，因为这种写法更高效，而且更符合go的语法习惯
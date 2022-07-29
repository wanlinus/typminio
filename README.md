# Typora minio client

[toc]

Typora的minio 图床客户端, 提供图片上传功能和显示功能

## 安装

[下载](https://github.com/wanlinus/typminio/releases)二进制文件`typminio`

将config.json配置文件放到typminio 同级, 具体配置如下:

```json
{
  "endpoint": "minio地址 exp: xxx.xom:9000",
  "username": "username",
  "password": "userpassword",
  "buckets": "tuchuang",
  "preUrl": "nginx 代理前缀 https://xxx.xxx:88"
}
```

## 食用方法

将程序下载到指定位置如`/Users/wanli/usr/typminio/typminio`, 配置文件和二进制文件保持同一目录, 然后配置typora如下

<img src="https://wanlinus.site:88/tuchuang-2022/image-20220729155815075-dqXqZ.png" alt="image-20220729155815075" style="zoom:50%;" />

点击验证图片上传选项, 等待成功

<img src="https://wanlinus.site:88/tuchuang-2022/image-20220729155835498-REsma.png" alt="image-20220729155835498" style="zoom:50%;" />

### 开发命令

```bash
go build -o typminio ./src/main.go
```



### License

Apache License 2.0 @ Wanli


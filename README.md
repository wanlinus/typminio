# Typora minio client

[toc]

Typora的minio 图床客户端, 提供图片上传功能和显示功能

## 安装

下载二进制文件`typminio`

配置对应的配置

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

<img src="https://wanlinus.site:88/tuchuang-2022/fcb5a35af7834f928f2e6d9e12a4dd18.png" alt="image-20220725144744740" style="zoom:50%;" />

点击验证图片上传选项, 等待成功

<img src="https://wanlinus.site:88/tuchuang-2022/06a97946538a450c961b807085722123.png" alt="image-20220725144823256" style="zoom:50%;" />

### License

Apache License 2.0 @ Wanli


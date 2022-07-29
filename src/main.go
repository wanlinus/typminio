package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
	"typminio/src/utils"
)

type Config struct {
	Endpoint string `json:"endpoint,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Buckets  string `json:"buckets,omitempty"`
	PreUrl   string `json:"preUrl,omitempty"`
}

func main() {
	// 读取命令行参数
	args := os.Args
	if len(args) == 1 {
		log.Fatalln("请传入文件路径")
		return
	}

	// 读取配置文件
	binPath, _ := os.Executable()
	dir := filepath.Dir(binPath)
	jsonFile, err2 := os.Open(dir + "/config.json")
	if err2 != nil {
		log.Fatalln("请配置config.json文件, 字段有endpoint,username,password,buckets")
		return
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			log.Fatalln("文件关闭失败")
		}
	}(jsonFile)

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var config Config
	err := json.Unmarshal(byteValue, &config)
	if err != nil {
		log.Fatalln("json解析失败")
		return
	}

	if config.PreUrl == "" {
		log.Fatalln("图片访问地址错误")
	}

	// Initialize minio client object.
	minioClient, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.Username, config.Password, ""),
		Secure: true,
	})
	if err != nil {
		log.Fatalln(err)
	}

	bucketName := config.Buckets + "-" + time.Now().Format("2006")
	location := "zh-east-1"

	ctx := context.Background()
	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			//log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	}

	var finalStrin []string
	for _, f := range args[1:] {
		objectName := buildFileName(f)
		finalStrin = append(finalStrin, config.PreUrl+"/"+bucketName+"/"+objectName)
		_, err := minioClient.FPutObject(ctx, bucketName, objectName, f, minio.PutObjectOptions{ContentType: ""})
		if err != nil {
			log.Fatalln(err)
		}
	}
	fmt.Println("Upload Success:")
	fmt.Print(strings.Join(finalStrin, "\n"))
}

// 生成文件名
func buildFileName(f string) string {
	if !fileExists(f) {
		fmt.Println("文件不存在")
	}

	fileNameAll := path.Base(f)
	fileExt := path.Ext(f)
	// 限制上传的文件名长度
	s := fileNameAll[0 : len(fileNameAll)-len(fileExt)]
	fileLen := len([]rune(s))
	min := int(math.Min(10, float64(fileLen)))
	filePrefix := string([]rune(s)[:min])
	return filePrefix + "-" + utils.RandStr(5) + fileExt

}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

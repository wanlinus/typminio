package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
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

	ctx := context.Background()

	// Initialize minio client object.
	minioClient, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.Username, config.Password, ""),
		Secure: true,
	})
	if err != nil {
		log.Fatalln(err)
	}

	if config.PreUrl == "" {
		log.Fatalln("图片访问地址错误")
	}

	bucketName := config.Buckets + "-" + time.Now().Format("2006")
	location := "us-east-1"

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			//log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		//log.Printf("Successfully created %s\n", bucketName)
	}

	var finalStrin []string
	for _, f := range args[1:] {
		ext := f[strings.LastIndex(f, "."):]
		name := strings.Replace(uuid.Must(uuid.NewV4()).String(), "-", "", -1)

		// Upload the zip file with FPutObject
		objectName := name + ext
		finalStrin = append(finalStrin, config.PreUrl+"/"+bucketName+"/"+objectName)
		_, err := minioClient.FPutObject(ctx, bucketName, objectName, f, minio.PutObjectOptions{ContentType: ""})
		if err != nil {
			log.Fatalln(err)
		}
	}
	fmt.Println("Upload Success:")
	fmt.Print(strings.Join(finalStrin, "\n"))
}

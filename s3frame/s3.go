package s3frame

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gogf/gf/frame/g"
	"image"
	"io/ioutil"
	"os"
)

type config struct {
	Endpoint string
	Key      string
	Secret   string
	Bucket   string
	Region   string
}

var (
	cfg      *config
	s3client *s3.S3
)

func init() {
	if cfg == nil {
		cfg = &config{
			Endpoint: g.Cfg().GetString("s3.S3_ENDPOINT"),
			Key:      g.Cfg().GetString("s3.S3_KEY"),
			Secret:   g.Cfg().GetString("s3.S3_SECRET"),
			Bucket:   g.Cfg().GetString("s3.S3_BUCKET"),
			Region:   g.Cfg().GetString("s3.S3_REGION"),
		}
	}
	if s3client == nil {
		s3client = s3.New(session.New(&aws.Config{
			Credentials:      credentials.NewStaticCredentials(cfg.Key, cfg.Secret, ""),
			Endpoint:         aws.String(cfg.Endpoint),
			Region:           aws.String(cfg.Region),
			DisableSSL:       aws.Bool(true),
			S3ForcePathStyle: aws.Bool(true),
		}))
	}
}

// AddBase64File 添加base64编码文件
func AddBase64File(name, base64str string) error {
	buf, err := base64.StdEncoding.DecodeString(base64str)
	if err != nil {
		return err
	}
	_, err = s3client.PutObject(&s3.PutObjectInput{
		Body:   bytes.NewReader(buf),
		Bucket: aws.String(cfg.Bucket),
		Key:    aws.String(name),
	})
	if err != nil {
		return err
	}
	return nil
}

// GetFileURL 获取文件URL
func GetFileURL(name string) string {
	if name != "" {
		return cfg.Endpoint + "/" + cfg.Bucket + "/" + name
	}
	return ""
}

// GetFileBytes 获取文件
func GetFileBytes(name string) ([]byte, error) {
	out, err := s3client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(cfg.Bucket),
		Key:    aws.String(name),
	})
	if err != nil {
		return nil, err
	}
	buf, err := ioutil.ReadAll(out.Body)
	defer out.Body.Close()
	return buf, err
}

// DeleteFile 删除文件
func DeleteFile(nameList ...string) error {
	objs := make([]*s3.ObjectIdentifier, 0, len(nameList))
	for _, name := range nameList {
		if name != "" {
			objs = append(objs, &s3.ObjectIdentifier{
				Key: aws.String(name),
			})
		}
	}
	_, err := s3client.DeleteObjects(&s3.DeleteObjectsInput{
		Bucket: aws.String(cfg.Bucket),
		Delete: &s3.Delete{
			Objects: objs,
		},
	})
	return err
}

// GetFileConfig 获取图片信息
func GetFileConfig(name string) (*image.Config, error) {
	url := GetFileURL(name)
	reader, err := os.Open(url)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	im, _, err := image.DecodeConfig(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", url, err)
		return nil, err
	}
	return &im, err
}

// GetFileBase64 获取图片base64字符串
func GetFileBase64(name string) (string, error) {
	b, err := GetFileBytes(name)
	if err != nil {
		return "", err
	}
	s := base64.StdEncoding.EncodeToString(b)
	return s, nil
}

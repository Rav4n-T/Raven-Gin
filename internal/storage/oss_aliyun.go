package storage

import (
	"io"
	"strconv"
	"sync"

	alioss "github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type AliyunCof struct {
	AccessKeyId     string `mapstructure:"access_key_id" json:"access_key_id" yaml:"access_key_id"`
	AccessKeySecret string `mapstructure:"access_key_secret" json:"access_key_secret" yaml:"access_key_secret"`
	Bucket          string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	Endpoint        string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	IsSsl           bool   `mapstructure:"is_ssl" json:"is_ssl" yaml:"is_ssl"`
	IsPrivate       bool   `mapstructure:"is_private" json:"is_private" yaml:"is_private"`
}

type aliyunoss struct {
	config *AliyunCof
	client *alioss.Client
	bucket *alioss.Bucket
}

var (
	o       *aliyunoss
	initErr error
)

func InitOss(config AliyunCof) (Storage, error) {
	once = &sync.Once{}
	once.Do(func() {
		o = &aliyunoss{}
		o.config = &config

		o.client, initErr = alioss.New(config.Endpoint, config.AccessKeyId, config.AccessKeySecret)
		if initErr != nil {
			return
		}

		o.bucket, initErr = o.client.Bucket(config.Bucket)
		if initErr != nil {
			return
		}

		Register(Oss, o)
	})
	if initErr != nil {
		return nil, initErr
	}
	return o, nil
}

func (o *aliyunoss) Put(key string, r io.Reader, dataLength int64) error {
	key = NormalizeKey(key)

	err := o.bucket.PutObject(key, r)
	if err != nil {
		return err
	}

	return nil
}

func (o *aliyunoss) PutFile(key string, localFile string) error {
	key = NormalizeKey(key)

	err := o.bucket.PutObjectFromFile(key, localFile)
	if err != nil {
		return err
	}

	return nil
}

func (o *aliyunoss) Get(key string) (io.ReadCloser, error) {
	key = NormalizeKey(key)

	body, err := o.bucket.GetObject(key)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (o *aliyunoss) Rename(srcKey string, destKey string) error {
	srcKey = NormalizeKey(srcKey)
	destKey = NormalizeKey(destKey)

	_, err := o.bucket.CopyObject(srcKey, destKey)
	if err != nil {
		return err
	}

	err = o.Delete(srcKey)
	if err != nil {
		return err
	}

	return nil
}

func (o *aliyunoss) Copy(srcKey string, destKey string) error {
	srcKey = NormalizeKey(srcKey)
	destKey = NormalizeKey(destKey)

	_, err := o.bucket.CopyObject(srcKey, destKey)
	if err != nil {
		return err
	}

	return nil
}

func (o *aliyunoss) Exists(key string) (bool, error) {
	key = NormalizeKey(key)

	return o.bucket.IsObjectExist(key)
}

func (o *aliyunoss) Size(key string) (int64, error) {
	key = NormalizeKey(key)

	props, err := o.bucket.GetObjectDetailedMeta(key)
	if err != nil {
		return 0, err
	}

	size, err := strconv.ParseInt(props.Get("Content-Length"), 10, 64)
	if err != nil {
		return 0, err
	}

	return size, nil
}

func (o *aliyunoss) Delete(key string) error {
	key = NormalizeKey(key)

	err := o.bucket.DeleteObject(key)
	if err != nil {
		return err
	}

	return nil
}

func (o *aliyunoss) Url(key string) string {
	var prefix string
	key = NormalizeKey(key)

	if o.config.IsSsl {
		prefix = "https://"
	} else {
		prefix = "http://"
	}

	if o.config.IsPrivate {
		url, err := o.bucket.SignURL(key, alioss.HTTPGet, 3600)
		if err == nil {
			return url
		}
	}

	return prefix + o.config.Bucket + "." + o.config.Endpoint + "/" + key
}

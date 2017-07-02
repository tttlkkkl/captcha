package image

import "image"

//Store 图片存储接口
type Store interface {
	Open(src interface{}) (image.Image, error)
	Save(src interface{}, img image.Image) error
}

//FileStore 文件存储
type FileStore struct {
}

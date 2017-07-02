package image

//Image 一个图像实例
type Image struct {
}

//New 返回一个图像实例
func New() *Image {
	return new(Image)
}

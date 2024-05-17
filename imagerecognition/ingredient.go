package imagerecognition

const (
	ingredient = "https://aip.baidubce.com/rest/2.0/image-classify/v1/classify/ingredient"
)

type IngredientRequest struct {
	Image string `url:"image,omitempty" form:"image,omitempty" json:"image"`     // 图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式。注意：图片需要base64编码、去掉编码头（data:image/jpg;base64,）后，再进行urlencode。
	URL   string `url:"url,omitempty" form:"url,omitempty" json:"url,omitempty"` // 图片完整URL，URL长度不超过1024字节，URL对应的图片base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式，当image字段存在时url字段失效。
}

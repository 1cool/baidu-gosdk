package imagerecognition

import (
	"encoding/json"
	"errors"
	query2 "github.com/google/go-querystring/query"
	"net/http"
	"net/url"
	"strings"
)

const (
	Dish = "https://aip.baidubce.com/rest/2.0/image-classify/v2/dish"
)

type DishResponse struct {
	Result []struct {
		Probability string `json:"probability"`
		HasCalorie  bool   `json:"has_calorie"`
		Calorie     string `json:"calorie"`
		Name        string `json:"name"`
	} `json:"result"`
	ResultNum int    `json:"result_num"`
	LogID     int64  `json:"log_id"`
	ErrorMsg  string `json:"error_msg,omitempty"`
	ErrorCode int    `json:"error_code,omitempty"`
}

type DishRequest struct {
	Image           string  `url:"image,omitempty" form:"image,omitempty" json:"image,omitempty"`                                  // 图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式。注意：图片需要base64编码、去掉编码头（data:image/jpg;base64,）后，再进行urlencode。
	URL             string  `url:"url,omitempty" form:"url,omitempty" json:"url,omitempty"`                                        // 图片完整URL，URL长度不超过1024字节，URL对应的图片base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式，当image字段存在时url字段失效。
	TopNum          uint32  `url:"top_num,omitempty" form:"top_num,omitempty" json:"top_num,omitempty"`                            // 返回结果top n,默认5.
	FilterThreshold float32 `url:"filter_threshold,omitempty" form:"filter_threshold,omitempty" json:"filter_threshold,omitempty"` // 默认0.95，可以通过该参数调节识别效果，降低非菜识别率.
	BaikeNum        int     `url:"baike_num,omitempty" form:"baike_num,omitempty" json:"baike_num,omitempty"`                      // 用于控制返回结果是否带有百科信息，若不输入此参数，则默认不返回百科结果；若输入此参数，会根据输入的整数返回相应个数的百科信息
}

func (srv *service) Dish(in DishRequest) (DishResponse, error) {
	parse, err := url.Parse(Dish)
	if err != nil {
		return DishResponse{}, err
	}

	parse.RawQuery = srv.accessToken.Encode()
	values, err := query2.Values(in)
	if err != nil {
		return DishResponse{}, err
	}

	req, err := http.NewRequest(http.MethodPost, parse.String(), strings.NewReader(values.Encode()))

	if err != nil {
		return DishResponse{}, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := srv.client.Do(req)
	if err != nil {
		return DishResponse{}, err
	}

	defer response.Body.Close()

	var d DishResponse

	err = json.NewDecoder(response.Body).Decode(&d)
	if err != nil {
		return DishResponse{}, err
	}

	if d.ErrorMsg != "" {
		return DishResponse{}, errors.New(d.ErrorMsg)
	}

	return d, nil
}

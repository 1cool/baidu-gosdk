package internal

import (
	"github.com/1cool/baidu-gosdk/internal/imagerecognition"
	"net/http"
	"net/url"
	"time"
)

type BaiduBce interface {
	NewImageRecognition() imagerecognition.Service
	setAccessToken(clientID, clientSecret string) error
}

type baiduBce struct {
	accessToken url.Values
	client      *http.Client
}

func NewBaiduBce(clientID, clientSecret string) (BaiduBce, error) {
	b := &baiduBce{
		accessToken: url.Values{},
		client: &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				MaxIdleConnsPerHost: 30,
			},
		},
	}

	err := b.setAccessToken(clientID, clientSecret)
	if err != nil {
		return nil, err
	}
	return b, nil
}

var _ BaiduBce = (*baiduBce)(nil)

func (b *baiduBce) NewImageRecognition() imagerecognition.Service {
	return imagerecognition.NewImageRecognition(b.accessToken, b.client)
}

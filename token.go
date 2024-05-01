package baidu_gosdk

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

const (
	Oauth = "https://aip.baidubce.com/oauth/2.0/token"
)

type OAuthErrorResponse struct {
	ErrorDescription string `json:"error_description"`
	Error            string `json:"error"`
}

type OAuthSuccessResponse struct {
	RefreshToken  string `json:"refresh_token"`
	ExpiresIn     int    `json:"expires_in"`
	SessionKey    string `json:"session_key"`
	AccessToken   string `json:"access_token"`
	Scope         string `json:"scope"`
	SessionSecret string `json:"session_secret"`
}

func (b *baiduBce) setAccessToken(clientID, clientSecret string) error {
	query := url.Values{}
	query.Set("client_id", clientID)
	query.Set("client_secret", clientSecret)
	query.Set("grant_type", "client_credentials")
	parse, err := url.Parse(Oauth)
	if err != nil {
		return err
	}

	parse.RawQuery = query.Encode()
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, parse.String(), nil)

	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if http.StatusOK == res.StatusCode {
		var s OAuthSuccessResponse
		err = json.NewDecoder(res.Body).Decode(&s)
		if err != nil {
			return err
		}

		b.accessToken.Set("access_token", s.AccessToken)
		return nil
	}

	var e OAuthErrorResponse
	err = json.NewDecoder(res.Body).Decode(&e)
	if err != nil {
		return err
	}
	return errors.New(e.ErrorDescription)
}

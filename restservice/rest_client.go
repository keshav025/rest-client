package restservice

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"

	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"
)

// var restClient *RestClient

type RestClient struct {
	Client          *http.Client
	URL             string
	Username        string
	Password        string
	SessionKey      string
	ProxySessionKey string
	Header          *http.Header
}

func GetRestClient(urlstr, user, password string) *RestClient {

	skipTls := &http.Transport{

		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	httpClient := &http.Client{Transport: skipTls}
	client := &RestClient{
		Client:   httpClient,
		URL:      urlstr,
		Username: user,
		Password: password,
	}
	return client
}

func (r *RestClient) SetSessionKey(sessionKey string) *RestClient {
	r.SessionKey = sessionKey
	return r

}
func (r *RestClient) SetproxySessionKey(sessionKey string) *RestClient {
	r.ProxySessionKey = sessionKey
	return r

}

func (r *RestClient) SetHeader(header *http.Header) *RestClient {
	r.Header = header
	return r

}

func (r RestClient) Get(data *url.Values, result interface{}) error {
	response, err := r.sendHttpRequest("GET", data, nil)
	res := ""
	if response != nil {
		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			logrus.Errorf("error in get call: %v", res)
			return err
		}
		err = json.Unmarshal(body, result)
		if err != nil {
			logrus.Errorf("error in parsing get call: %v", res)
			return err
		}

		res = string(body)
	}
	if err != nil {
		logrus.Errorf("error in get call: %v", res)
		return err
	}

	return nil

}

func (r RestClient) Post(data *url.Values, jsonBody interface{}, result interface{}) error {
	response, err := r.sendHttpRequest("POST", data, jsonBody)
	if err != nil {
		logrus.Errorf("error in post call: %v", err)
		return err
	}

	if response != nil {
		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			logrus.Errorf("error in get call: %v", err)
			return err
		}
		err = json.Unmarshal(body, result)
		if err != nil {
			logrus.Errorf("error in parsing get call: %v", err)
			return err
		}
	}

	return nil

}

func (r RestClient) Put(data *url.Values, jsonBody interface{}, result interface{}) error {
	response, err := r.sendHttpRequest("PUT", data, jsonBody)
	if err != nil {
		logrus.Errorf("error in post call: %v", err)
		return err
	}

	if response != nil {
		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			logrus.Errorf("error in get call: %v", err)
			return err
		}
		err = json.Unmarshal(body, result)
		if err != nil {
			logrus.Errorf("error in parsing get call: %v", err)
			return err
		}
	}

	return nil

}

func (r RestClient) Patch(data *url.Values, jsonBody interface{}, result interface{}) error {
	response, err := r.sendHttpRequest("PATCH", data, jsonBody)
	if err != nil {
		logrus.Errorf("error in post call: %v", err)
		return err
	}

	if response != nil {
		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			logrus.Errorf("error in get call: %v", err)
			return err
		}
		err = json.Unmarshal(body, result)
		if err != nil {
			logrus.Errorf("error in parsing get call: %v", err)
			return err
		}
	}

	return nil

}

func (r RestClient) Delete(result interface{}) error {
	response, err := r.sendHttpRequest("DELETE", nil, nil)
	if err != nil {
		logrus.Errorf("error in post call: %v", err)
		return err
	}

	if response != nil {
		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			logrus.Errorf("error in get call: %v", err)
			return err
		}
		err = json.Unmarshal(body, result)
		if err != nil {
			logrus.Errorf("error in parsing get call: %v", err)
			return err
		}
	}

	return nil

}

func (r RestClient) sendHttpRequest(method string, data *url.Values, jsonBody interface{}) (*http.Response, error) {

	var payload io.Reader
	if data != nil {
		payload = bytes.NewBufferString(data.Encode())
	}
	if jsonBody != nil {
		js, err := json.Marshal(jsonBody)
		if err != nil {
			return nil, err
		}
		payload = bytes.NewBuffer(js)
		logrus.Infof("rest payload: %v", string(js))
	}

	request, err := http.NewRequest(method, r.URL, payload)
	if err != nil {
		return nil, err
	}
	if r.Header != nil {
		for key, values := range *r.Header {
			if len(values) > 0 {
				request.Header.Set(key, values[0])
			}
		}
	}

	if r.ProxySessionKey != "" {
		request.Header.Add("Proxy-Authorization", r.ProxySessionKey)
		//request.Header.Add("Proxy-Authorization", fmt.Sprintf("Bearer %v", "eyJhbGciOiJSUzI1NiIsImprdSI6Imh0dHBzOi8vaHhtLWNmLXRvb2xzLWFuZC1pbnRlZ3JhdGlvbi00YWV3NjYyeS5hdXRoZW50aWNhdGlvbi51czMwLmhhbmEub25kZW1hbmQuY29tL3Rva2VuX2tleXMiLCJraWQiOiJkZWZhdWx0LWp3dC1rZXktLTIxMTUwMzMwNTIiLCJ0eXAiOiJKV1QifQ.eyJqdGkiOiJlYzEzNTFkNWNhMTA0MWNkYmMxNmI5YmVkNGZkOTUzYSIsImV4dF9hdHRyIjp7ImVuaGFuY2VyIjoiWFNVQUEiLCJzdWJhY2NvdW50aWQiOiI2Njk2NTBiZC02MzBlLTQwODgtYmQxZC05YTU2NTkwYTE0NmQiLCJ6ZG4iOiJoeG0tY2YtdG9vbHMtYW5kLWludGVncmF0aW9uLTRhZXc2NjJ5Iiwic2VydmljZWluc3RhbmNlaWQiOiI0ZTU4YzNlZi1mZTYwLTQwZTktYjQ4Mi03NjBlMmNlY2VlNTgifSwic3ViIjoic2ItY2xvbmU0ZTU4YzNlZmZlNjA0MGU5YjQ4Mjc2MGUyY2VjZWU1OCFiMjc0OXxjb25uZWN0aXZpdHkhYjMzIiwiYXV0aG9yaXRpZXMiOlsidWFhLnJlc291cmNlIiwiY29ubmVjdGl2aXR5IWIzMy5wcm94eSJdLCJzY29wZSI6WyJ1YWEucmVzb3VyY2UiLCJjb25uZWN0aXZpdHkhYjMzLnByb3h5Il0sImNsaWVudF9pZCI6InNiLWNsb25lNGU1OGMzZWZmZTYwNDBlOWI0ODI3NjBlMmNlY2VlNTghYjI3NDl8Y29ubmVjdGl2aXR5IWIzMyIsImNpZCI6InNiLWNsb25lNGU1OGMzZWZmZTYwNDBlOWI0ODI3NjBlMmNlY2VlNTghYjI3NDl8Y29ubmVjdGl2aXR5IWIzMyIsImF6cCI6InNiLWNsb25lNGU1OGMzZWZmZTYwNDBlOWI0ODI3NjBlMmNlY2VlNTghYjI3NDl8Y29ubmVjdGl2aXR5IWIzMyIsImdyYW50X3R5cGUiOiJjbGllbnRfY3JlZGVudGlhbHMiLCJyZXZfc2lnIjoiMjhmMDBiYjciLCJpYXQiOjE2NzM4Njc5NjEsImV4cCI6MTY3MzkxMTE2MSwiaXNzIjoiaHR0cHM6Ly9oeG0tY2YtdG9vbHMtYW5kLWludGVncmF0aW9uLTRhZXc2NjJ5LmF1dGhlbnRpY2F0aW9uLnVzMzAuaGFuYS5vbmRlbWFuZC5jb20vb2F1dGgvdG9rZW4iLCJ6aWQiOiI2Njk2NTBiZC02MzBlLTQwODgtYmQxZC05YTU2NTkwYTE0NmQiLCJhdWQiOlsiY29ubmVjdGl2aXR5IWIzMyIsInVhYSIsInNiLWNsb25lNGU1OGMzZWZmZTYwNDBlOWI0ODI3NjBlMmNlY2VlNTghYjI3NDl8Y29ubmVjdGl2aXR5IWIzMyJdfQ.YCnKn3Pbroeuc0se1baYW2YUUxpylUb2a-u1wsMaTwWgzJi4JbgvRYJhRiL-RvBdltDqjRHFaSH4XjHzayJzVEqLPpqsTYtwNwPcwLDF8rlg2egA13Rlt8S4jy7K4JNmGFyTJeof328ZhqyJ2-i-X5HgE6joUqHE6LYxIfq82VTF-jzZDAY48NN96VucvAmvhqjhyCirdpuUkzngNXKMg1aqX_eoS5tWg034j09Wuy9Kru0TXaOGEzhgcQDg8FG3NeCbHu-51M4AgNml4DJWf5jVWuv1SkJXcVwEGL6ZJf4vbIasfb8KFyYdoshBx-WFPnriumxK2nftDPQSyjSixw"))
	}
	if r.SessionKey != "" {
		request.Header.Add("Authorization", r.SessionKey)
	} else {
		if r.Username != "" && r.Password != "" {
			request.SetBasicAuth(r.Username, r.Password)
		}
	}
	// for k, v := range request.Header {
	// 	logrus.Infof("request Headers: %v:%v", k, v)
	// }
	// logrus.Infof("request headers: %v", request.Header)
	// request.Header = request.Header
	response, err := r.Client.Do(request)
	statusOK := false
	statusCode := 500
	status := "api call failed with no response."
	if response != nil {
		statusOK = response.StatusCode >= 200 && response.StatusCode < 300
		statusCode = response.StatusCode
		status = response.Status
		// bodyBytes, _ := io.ReadAll(response.Body)

		// bodyString := string(bodyBytes)
		// logrus.Info(bodyString)

		// logrus.Infof("response Headers: %v", response.Header)
	}
	if err != nil || !statusOK {
		if err != nil {
			logrus.Error("Error in rest call: ", err.Error())
		}
		return response, fmt.Errorf("api call failed: %d : %v", statusCode, status)
	}
	return response, err
}

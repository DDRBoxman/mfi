package mfi

import (
	"net/http"
	"net/http/cookiejar"
	"math/rand"
	"fmt"
	"net/url"
	"time"
	)

type MFIClient struct {
	client http.Client
	hostname string
}

type SensorData struct {
	Port int `json:"port"`
	Output int `json:"output"`
    // Power
    // Enabled
    // Current
    Voltage float64 `json:"voltage"`

	// powerfactor

	// relay

	// lock

	// prevmonth

	// thismonth

}

func GenerateCookie() string {
	bytes := make([]byte, 32)
	for i := 0; i < 32; i++ {
		bytes[i] = byte(48 + rand.Intn(10))  //0=48 and Z = 48+10
	}
	return string(bytes)
}

func (client *MFIClient) generateURL(path string) string {
	return fmt.Sprintf("http://%s/%s", client.hostname, path)
}

func MakeMFIClient(hostname string) (*MFIClient, error) {
	cookieJar, _ := cookiejar.New(nil)

	httpClient := http.Client{
		Jar: cookieJar,
	}

	cookie := &http.Cookie{Name: "AIROS_SESSIONID", Value: GenerateCookie(), Expires: time.Now().Add(876000 * time.Hour)} // 100 year expiration

	url, err := url.Parse(fmt.Sprintf("http://%s", hostname))
	if err != nil {
		return nil, err
	}

	cookieJar.SetCookies(url, []*http.Cookie{cookie})

	return &MFIClient{
		hostname: hostname,
		client: httpClient,
	}, nil
}

func (client *MFIClient) Auth(username, password string) error {
	values := url.Values{}
	values.Add("username", username)
	values.Add("password", password)
	resp, err := client.client.PostForm(client.generateURL("login.cgi"), values)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to login with status: %d", resp.StatusCode)
	}

	return nil
}

/*func (client *MFIClient) GetSensorData() []SensorData {
 // curl -b "AIROS_SESSIONID=01234567890123456789012345678901" 10.0.0.1/sensors
}*/

func (client *MFIClient) SetOutputEnabled(number int, enabled bool) error {
	values := url.Values{}
	if enabled {
		values.Add("output", "1")
	} else {
		values.Add("output", "0")
	}
	_, err := client.client.PostForm(client.generateURL(fmt.Sprintf("sensors/%d", number)), values)

	if err != nil {
		return err
	}

	return nil
}
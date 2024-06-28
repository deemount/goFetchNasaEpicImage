package epic

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var apiKey string

func SetAPIKey(key string) {
	apiKey = key
}

// ApiURL is the base URL for the NASA EPIC API
var ApiURL = "https://epic.gsfc.nasa.gov"

func GetLatestDate() (time.Time, error) {
	req, err := http.NewRequest("GET", ApiURL+"/api/natural", nil)
	if err != nil {
		return time.Time{}, err
	}
	req.Header.Set("api-key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return time.Time{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return time.Time{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var data []struct {
		Date string `json:"date"`
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return time.Time{}, err
	}

	if len(data) == 0 {
		return time.Time{}, fmt.Errorf("no data received from API")
	}

	latestDate, err := time.Parse("20060102", data[0].Date)
	if err != nil {
		return time.Time{}, err
	}

	return latestDate, nil
}

func DownloadImages(date time.Time, targetFolder string) error {
	urls, err := GetImageURLs(date)
	if err != nil {
		return err
	}

	for _, url := range urls {
		filename := filepath.Base(url)
		targetPath := filepath.Join(targetFolder, filename)

		err = downloadFile(url, targetPath)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetImageURLs(date time.Time) ([]string, error) {
	apiURL := fmt.Sprintf(ApiURL+"/api/natural/date/%s", date.Format("20060102"))
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("api-key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var data []struct {
		Image string `json:"image"`
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	urls := make([]string, len(data))
	for i, item := range data {
		urls[i] = fmt.Sprintf(ApiURL+"/archive/natural/%s/png/%s.png", date.Format("20060102"), item.Image)
	}

	return urls, nil
}

func downloadFile(url, targetPath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

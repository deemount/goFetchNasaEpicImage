package epic_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/deemount/goFetchNasaEpicImage/internal/epic"
)

func TestGetLatestDate(t *testing.T) {
	// Set up a mock server to simulate the NASA EPIC API response
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"date":"20230415"}]`))
	}))
	defer mockServer.Close()

	// Override the NASA EPIC API URL with the mock server URL
	originalURL := epic.ApiURL
	defer func() { epic.ApiURL = originalURL }() // Restore the original URL after the test
	epic.ApiURL = mockServer.URL + "/api/natural"

	// Call the GetLatestDate function
	latestDate, err := epic.GetLatestDate()
	if err != nil {
		t.Errorf("GetLatestDate() returned an error: %v", err)
	}

	// Check if the returned date is correct
	expectedDate := time.Date(2023, time.April, 15, 0, 0, 0, 0, time.UTC)
	if !latestDate.Equal(expectedDate) {
		t.Errorf("GetLatestDate() returned %v, expected %v", latestDate, expectedDate)
	}
}

func TestGetImageURLs(t *testing.T) {
	// Set up a mock server to simulate the NASA EPIC API response
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"image":"20230415000000"},{"image":"20230415000001"}]`))
	}))
	defer mockServer.Close()

	// Override the NASA EPIC API URL with the mock server URL
	originalURL := epic.ApiURL
	defer func() { epic.ApiURL = originalURL }() // Restore the original URL after the test
	epic.ApiURL = mockServer.URL

	// Set the test date
	testDate := time.Date(2023, time.April, 15, 0, 0, 0, 0, time.UTC)

	// Call the GetImageURLs function
	imageURLs, err := epic.GetImageURLs(testDate)
	if err != nil {
		t.Errorf("GetImageURLs() returned an error: %v", err)
	}

	// Check if the returned image URLs are correct
	expectedURLs := []string{
		"https://epic.gsfc.nasa.gov/archive/natural/20230415/png/20230415000000.png",
		"https://epic.gsfc.nasa.gov/archive/natural/20230415/png/20230415000001.png",
	}

	if len(imageURLs) != len(expectedURLs) {
		t.Errorf("GetImageURLs() returned %d URLs, expected %d", len(imageURLs), len(expectedURLs))
	}

	for i, url := range imageURLs {
		if url != expectedURLs[i] {
			t.Errorf("GetImageURLs() returned %s, expected %s", url, expectedURLs[i])
		}
	}
}

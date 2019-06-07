package gcpmisc

import (
	"io/ioutil"
	"net/http"
)

func GetProjectId() (string, error) {
	// Creates http client object
	var client http.Client

	// Creates the request
	req, err := http.NewRequest(
		"GET",
		"http://metadata.google.internal/computeMetadata/v1/project/projectid",
		nil,
	)

	// Adds this HTTP header otherwise google will give you an error.
	req.Header.Add("MetadataFlavor", "Google")

	// Makes the HTTP request
	resp, err := client.Do(req)

	// Checks for errors when doing the http request
	if err != nil {
		return "", err
	}

	// Closes the TCP connection when function returns
	defer resp.Body.Close()

	// Converts the io.Reader object to string
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	// Returns the projectID
	return bodyString, nil
}

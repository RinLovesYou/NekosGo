package NekosGo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const apiUrl = "https://api.nekos.dev/api/v3/"

func Image(endpoint interface{}, contentType ContentType) (string, error) {
	if v, ok := endpoint.(NSFW); ok {
		return getImage(string(v), "sfw", contentType)
	} else if v, ok := endpoint.(SFW); ok {
		return getImage(string(v), "sfw", contentType)
	}
	return "", errors.New("invalid endpoint provided")
}

func getImage(endpoint string, sfw string, ct ContentType) (string, error) {
	url := fmt.Sprintf("%s%s/%s/%s/", apiUrl, string(ct), sfw, endpoint)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	response, err := UnmarshalNekos(body)
	if err != nil {
		return "", err
	}

	if !response.Data.Status.Success {
		return "", fmt.Errorf("ERROR %v: %s", response.Data.Status.Code, response.Data.Status.Message)
	}

	return response.Data.Response.URL, nil
}

package NekosGo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const apiUrl = "https://api.nekos.dev/api/v3/images"

func Image(endpoint interface{}) (string, error) {
	if v, ok := endpoint.(NSFWImage); ok {
		return getImage(string(v), "nsfw", "img")
	} else if v, ok := endpoint.(SFWImage); ok {
		return getImage(string(v), "sfw", "img")
	} else if v, ok := endpoint.(NSFWGif); ok {
		return getImage(string(v), "nsfw", "gif")
	} else if v, ok := endpoint.(SFWGif); ok {
		return getImage(string(v), "sfw", "gif")
	}
	return "", errors.New("invalid endpoint provided")
}

func getImage(endpoint string, sfw string, ct string) (string, error) {
	url := fmt.Sprintf("%s/%s/%s/%s/", apiUrl, sfw, ct, endpoint)

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

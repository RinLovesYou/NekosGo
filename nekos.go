package NekosGo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const apiUrl = "https://api.nekos.dev/api/v3/images/"

func Image(ep interface{}) (string, error) {
	if v, ok := ep.(NSFW); ok {
		return getImage(string(v), "nsfw")
	} else if v, ok := ep.(SFW); ok {
		return getImage(string(v), "sfw")
	}
	return "", errors.New("invalid endpoint provided")
}

func getImage(ep string, sfw string) (string, error) {
	url := fmt.Sprintf("%s%s/%s", apiUrl, sfw, ep)

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

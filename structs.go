package NekosGo

import "encoding/json"

type NekosResponse struct {
	Data struct {
		Response struct {
			URL string `json:"url"`
		} `json:"response"`
		Status struct {
			Code       int         `json:"code"`
			Message    interface{} `json:"message"`
			RenderedIn string      `json:"rendered_in"`
			Success    bool        `json:"success"`
		} `json:"status"`
	} `json:"data"`
}

func UnmarshalNekos(data []byte) (*NekosResponse, error) {
	response := &NekosResponse{}
	err := json.Unmarshal(data, &response)
	return response, err
}

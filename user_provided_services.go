package cfclient

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type UserProvidedServicesResponse struct {
	Count     int                            `json:"total_results"`
	Pages     int                            `json:"total_pages"`
	NextUrl   string                         `json:"next_url"`
	Resources []UserProvidedServicesResource `json:"resources"`
}

type UserProvidedServicesResource struct {
	Meta   Meta                 `json:"metadata"`
	Entity UserProvidedServices `json:"entity"`
}

type UserProvidedServices struct {
	Guid        string            `json:"guid"`
	Name        string            `json:"name"`
	Credentials map[string]string `json:"credentials"`
	c           *Client
}

func (c *Client) ListUserProvidedServices() []UserProvidedServices {
	var services []UserProvidedServices

	requestUrl := "/v2/user_provided_service_instances"

	for {

		var serviceResp UserProvidedServicesResponse
		r := c.newRequest("GET", requestUrl)

		resp, err := c.doRequest(r)
		if err != nil {
			log.Printf("Error requesting services %v", err)
		}
		resBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error reading services request %v", resBody)
		}

		err = json.Unmarshal(resBody, &serviceResp)
		if err != nil {
			log.Printf("Error unmarshaling services %v", err)
		}
		for _, service := range serviceResp.Resources {
			service.Entity.Guid = service.Meta.Guid
			service.Entity.c = c
			services = append(services, service.Entity)
		}
		requestUrl = serviceResp.NextUrl
		if requestUrl == "" {
			break
		}

	}

	return services
}

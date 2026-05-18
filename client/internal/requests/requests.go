package requests

import (
	"fmt"
	"net/http"
	"os"
)

func Get(objectID string, authToken string) {

	// Create GET request
	req, err := http.NewRequest("GET", os.Getenv("SERVER_URL")+"/"+objectID, nil)

	// Want error to be nil, indicating success in building the request
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}

	// Set UAuth and log request
	req.Header.Set("Authorization", "Bearer "+authToken)
	fmt.Println(req)

	// Send the request to the server. Later, the server will send the relevant node information
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("Error fetching object:", err)
		return
	}

	fmt.Println(res)

}

func Send(objectID string, authToken string) {
	// Pass
}

func Main() {
	Get("", "")
}

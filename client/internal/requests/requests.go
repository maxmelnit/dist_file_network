package requests

import (
	"fmt"
	"net/http"
	"os"
)

func Get(objectID string, authToken string) {
	// Create GET request
	req, err := http.NewRequest("GET", os.Getenv("SERVER_URL")+"/"+objectID, nil)

	// Want error to be nil, indicating success
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
	}

	// Set UAuth and log request
	req.Header.Set("Authorization", "Bearer: "+authToken)
	fmt.Println(req)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("Error fetching object:", err)
	}

	fmt.Println(res)

}

func Main() {
	Get("", "")
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/Pallinder/go-randomdata"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "register-dummy-user",
				Usage: "register a dummy user",
				Action: func(*cli.Context) error {
					registerDummyUser()
					return nil
				},
			},
			{
				Name:  "get",
				Usage: "get a user",
				Action: func(c *cli.Context) error {
					username := c.Args().First()
					if username == "" {
						fmt.Println("Please provide a username")
						return nil
					}
					getUserByUsername(username)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}

func registerDummyUser() {
	// dummy user
	profile := randomdata.GenerateProfile(randomdata.Male | randomdata.Female | randomdata.RandomGender)

	apiURL := "http://localhost:8000/users"

	// Prepare form data
	data := url.Values{}
	data.Set("username", profile.Login.Username)
	data.Set("name", profile.Name.First+" "+profile.Name.Last)
	data.Set("gender", profile.Gender)
	data.Set("phone", profile.Phone)
	data.Set("address", randomdata.Address())
	data.Set("consented", strconv.FormatBool(randomdata.Boolean()))

	// Create the request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error executing request:", err)
		return
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusCreated {
		fmt.Println("Failed to register user, status code:", resp.StatusCode)
		return
	}

	// Parse the response
	var response map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	// Print the response
	fmt.Println(response["message"], ":", response["username"])
}

func getUserByUsername(username string) {
	apiURL := fmt.Sprintf("http://localhost:8000/users/%s", username)

	// Create the request
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error executing request:", err)
		return
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Failed to fetch user:", resp.StatusCode)
		return
	}

	// Parse the response
	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	// Print the response
	fmt.Printf(
		"Username: %s\nName: %s\nGender: %s\nPhone: %s\nAddress: %s\nConsented: %t\nCreated At: %s\n",
		response["username"], response["name"], response["gender"], response["phone"], response["address"], response["consented"], response["created_at"],
	)
}

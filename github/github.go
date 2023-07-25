package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/lami-health/run-after-approvals/models"
)

// GetReviews perform an GET request to github API and return the reviews in an array objects, the return is unmarshalled into the target struct provided by the user.
func GetReviews(client *http.Client, url, token string, target interface{}) error {
	var err error
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token))
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Println(url)
	fmt.Println(string(body))

	return json.Unmarshal(body, target)
}

type githubPullRequest struct {
	Number int `json:"number"`
}

type githubEventPath struct {
	PullRequest githubPullRequest `json:"pull_request"`
}

// GetPullRequestNumber read github json file and return pull_request number from it.
func GetPullRequestNumber(path string) string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Error while opening github file on path: %s -> %v", path, err)
	}

	var payload githubEventPath
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatalf("Error while parsing github file on path: %s -> %v", path, err)
	}

	return strconv.Itoa(payload.PullRequest.Number)
}

func CalculateValidApprovals(reviews []models.Review) int {
	acc := 0
	for _, review := range reviews {
		if review.State != "APPROVED" {
			acc = 0
			continue
		}

		if review.State == "APPROVED" {
			acc++
		}
	}

	if acc >= 0 {
		return acc
	}

	return acc
}

package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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

	return json.Unmarshal(body, target)
}

func CalculateValidApprovals(reviews []models.Review) (int, string) {
	acc := 0
	for _, review := range reviews {
		if review.State == "CHANGES_REQUESTED" {
			acc = 0
			continue
		}

		if review.State == "APPROVED" {
			acc++
		}
	}

	if acc >= 0 {
		return acc, "APPROVED"
	}

	return acc, "DENIED"
}

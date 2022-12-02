package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/lami-health/run-after-approvals/env"
	"github.com/lami-health/run-after-approvals/github"
	"github.com/lami-health/run-after-approvals/models"
)

var GH_REPO string = env.Getenv("GITHUB_REPOSITORY", "")
var GH_PR_NUMBER string = env.Getenv("GITHUB_PULL_REQUEST", "")
var APPROVALS string = env.Getenv("APPROVALS", "2")
var TOKEN string = env.Getenv("GITHUB_TOKEN", "")

var client = &http.Client{Timeout: 10 * time.Second}

func main() {
	var reviews []models.Review
	user := strings.Split(GH_REPO, "/")[0]
	repo := strings.Split(GH_REPO, "/")[1]

	fmt.Printf("TESTE -> %s", GH_PR_NUMBER)

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/pulls/%s/reviews", user, repo, GH_PR_NUMBER)

	if err := github.GetReviews(client, url, TOKEN, &reviews); err != nil {
		log.Fatalf("Could not get the github reviews for specified repo: %v", err)
	}

	acc := github.CalculateValidApprovals(reviews)

	approvals, err := strconv.Atoi(APPROVALS)

	if err != nil {
		log.Fatalf("Could not convert APPROVALS from string to int: %v", err)
	}

	if acc >= approvals {
		fmt.Printf("status=approved")
		setTag("approved")
	} else {
		fmt.Printf("status=denied")
		setTag("denied")
	}
}

func setTag(value string) {
	fmt.Printf(`::set-output name=status::%s`, value)
	fmt.Print("\n")
	// TODO: Not sure it's needed twice
	// Ref: https://stackoverflow.com/questions/71357973/github-actions-set-two-output-names-from-custom-action-in-golang-code
	fmt.Printf(`::set-output name=status::%s`, value)
	fmt.Print("\n")
}

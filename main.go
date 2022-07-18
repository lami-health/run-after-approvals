package main

import (
	"fmt"
	"github.com/lami-health/run-after-approvals/env"
	"github.com/lami-health/run-after-approvals/github"
	"github.com/lami-health/run-after-approvals/models"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var GH_USER string = env.Getenv("GITHUB_USER", "lami-health")
var GH_REPO string = env.Getenv("GITHUB_REPOSITORY", "website")
var GH_PR_NUMBER string = env.Getenv("GITHUB_PULL_REQUEST", "720")
var APPROVALS string = env.Getenv("APPROVALS", "2")

var client = &http.Client{Timeout: 10 * time.Second}

func main() {
	var reviews []models.Review
	token := ""
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/pulls/%s/reviews", GH_USER, GH_REPO, GH_PR_NUMBER)

	if err := github.GetReviews(client, url, token, &reviews); err != nil {
		log.Fatalf("Could not get the github reviews for specified repo: %v", err)
	}

	acc, label := github.CalculateValidApprovals(reviews)

	fmt.Printf("%d/%s Approvals - %s", acc, APPROVALS, label)

	approvals, err := strconv.Atoi(APPROVALS)

	if err != nil {
		log.Fatalf("Could not convert APPROVALS from string to int: %v", err)
	}

	if acc < approvals {
		os.Exit(1)
	}
}

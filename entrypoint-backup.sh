#!/bin/bash
set -e

if [[ -z "$APPROVALS" ]]; then
  echo "Set the APPROVALS env variable."
  exit 1
fi

if [[ -z "$GITHUB_TOKEN" ]]; then
  echo "Set the GITHUB_TOKEN env variable."
  exit 1
fi

if [[ -z "$GITHUB_TOKEN" ]]; then
  echo "Set the GITHUB_TOKEN env variable."
  exit 1
fi

if [[ -z "$GITHUB_REPOSITORY" ]]; then
  echo "Set the GITHUB_REPOSITORY env variable."
  exit 1
fi

if [[ -z "$GITHUB_EVENT_PATH" ]]; then
  echo "Set the GITHUB_EVENT_PATH env variable."
  exit 1
fi

URI="https://api.github.com"
API_HEADER="Accept: application/vnd.github.v3+json"
AUTH_HEADER="Authorization: token ${GITHUB_TOKEN}"

action=$(jq --raw-output .action "$GITHUB_EVENT_PATH")
state=$(jq --raw-output .review.state "$GITHUB_EVENT_PATH")
number=$(jq --raw-output .pull_request.number "$GITHUB_EVENT_PATH")

break_if_not_approved() {
  body=$(curl -sSL -H "${AUTH_HEADER}" -H "${API_HEADER}" "${URI}/repos/${GITHUB_REPOSITORY}/pulls/${number}/reviews?per_page=100")
  reviews=$(echo "$body" | jq --raw-output '.[] | {state: .state} | @base64')

  approvals=0

  for r in $reviews; do
    review="$(echo "$r" | base64 -d)"
    rState=$(echo "$review" | jq --raw-output '.state')

    if [[ "$rState" == "APPROVED" ]]; then
      approvals=$((approvals+1))
    fi

    echo "${approvals}/${APPROVALS} approvals"

    if [[ "$approvals" -lt "$APPROVALS" ]]; then
      exit 1
    fi
  done
}

if [[ "$action" == "submitted" ]] && [[ "$state" == "approved" ]]; then
  break_if_not_approved
else
  echo "Ignoring event ${action}/${state}"
fi

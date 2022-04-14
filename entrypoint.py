#!/usr/bin/env python3

import os
import json
import requests

approvals = os.getenv("APPROVALS") or 2
github_token = os.getenv("GITHUB_TOKEN")
github_repository = os.getenv("GITHUB_REPOSITORY")
github_event_path = os.getenv("GITHUB_EVENT_PATH")

if not github_token:
    print("Set the GITHUB_TOKEN env variable")
    exit(1)

if not github_repository:
    print("Set the GITHUB_REPOSITORY env variable")
    exit(1)

if not github_event_path:
    print("Set the GITHUB_EVENT_PATH env variable")
    exit(1)

pull_request_number = github_event_path["pull_request"]["number"]
api_url = f"https://api.github.com/repos/{github_repository}/pulls/{pull_request_number}/reviews?per_page=100"
headers = {
    "Authorization": f"token {github_token}",
    "Accept": "application/vnd.github.v3+json",
}


def make_request():
    response = requests.get(api_url, headers=headers).json()
    acc = 0
    for res in response:
        if res["state"] != "APPROVED":
            acc = 0
            continue

        acc = acc + 1

    label = "APPROVED" if acc >= approvals else "DENIED"
    print(f"{acc}/{approvals} Approvals - {label}")

    if acc < approvals:
        exit(1)
    return


make_request()

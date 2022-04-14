#!/usr/bin/env python3

import os
import json
import requests

api_url = "https://api.github.com/repos/cherryramatisdev/website-ci/pulls/16/reviews?per_page=100"
headers = {
    "Authorization": "token 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",
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

    label = 'APPROVED' if acc >= 2 else 'DENIED'
    print(f"{acc}/2 Approvals - {label}")

    if acc < 2:
        exit(1)
    return


make_request()

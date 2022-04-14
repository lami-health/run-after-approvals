FROM alpine:3.10.3

LABEL "com.github.actions.name"="Run after approvals"
LABEL "com.github.actions.description"="Only permit the PR running if has a certain number of approvals"
LABEL "com.github.actions.icon"="tag"
LABEL "com.github.actions.color"="gray-dark"

LABEL version="0.0.1"
LABEL repository="http://github.com/lami-health/run-after-approvals"
LABEL homepage="http://github.com/lami-health/run-after-approvals"
LABEL maintainer="Lami Team <tech@lamimed.it>"

RUN apk add --no-cache bash curl jq

ADD entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

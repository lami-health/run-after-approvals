FROM alpine:latest

LABEL "com.github.actions.name"="Run after approvals"
LABEL "com.github.actions.description"="Only permit the PR running if has a certain number of approvals"
LABEL "com.github.actions.icon"="tag"
LABEL "com.github.actions.color"="gray-dark"

LABEL version="0.0.2"
LABEL repository="http://github.com/lami-health/run-after-approvals"
LABEL homepage="http://github.com/lami-health/run-after-approvals"
LABEL maintainer="Lami Team <tech@lamimed.it>"

COPY ./run-after-approvals /run-after-approvals
ENTRYPOINT ["/run-after-approvals"]

.PHONY=linux

linux:
	echo "Building for linux GOARCH=amd64 GOOS=linux"
	GOOS=linux GOARCH=amd64 go build -o run-after-approvals
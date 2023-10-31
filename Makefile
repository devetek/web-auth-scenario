run-nodejs:
	node nodejs/index.js

run-golang:
	go run golang/main.go


run-hotreload:
	nodemon --signal SIGTERM
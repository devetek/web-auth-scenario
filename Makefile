.init:
	@npm install -g nodemon

run-nodejs: .init
	@nodemon --signal SIGTERM --config nodemon-nodejs.json

run-golang: .init
	@nodemon --signal SIGTERM --config nodemon-golang.json
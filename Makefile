swagger:
	rd /s /q api && mkdir api
	swagger generate server -t api --exclude-main
	go mod tidy
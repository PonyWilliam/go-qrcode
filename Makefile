docker:
	export GOOS=linux
	GOOS=linux
	go build -o go-qrcode main.go
	docker build -t ponywilliam/go-qrcode .
	docker tag ponywilliam/go-qrcode ponywilliam/go-qrcode
	docker push ponywilliam/go-qrcode
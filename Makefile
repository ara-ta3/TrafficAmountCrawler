
setup: .env

run: 
	go run main.go

.env: sample.env
	cp -f $< $@

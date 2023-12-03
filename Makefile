goos_opt=GOOS=$(GOOS)
goarch_opt=GOARCH=$(GOARCH)
out=trafic_amount_crawler
out_opt=-o $(out)

setup: .env

run: 
	go run main.go

build:
	$(goos_opt) $(goarch_opt) go build $(out_opt)

build_for_linux:
	$(MAKE) build GOOS=linux GOARCH=amd64 

.env: sample.env
	cp -f $< $@


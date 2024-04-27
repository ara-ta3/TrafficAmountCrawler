goos_opt=GOOS=$(GOOS)
goarch_opt=GOARCH=$(GOARCH)
goarm_opt=GOARM=$(GOARM)
out=traffic_amount_crawler
out_opt=-o $(out)

setup: .env

run: 
	go run cmd/traffic_amount_crawler.go

build:
	$(goos_opt) $(goarch_opt) $(goarm_opt) go build $(out_opt) cmd/traffic_amount_crawler.go

build_for_rapsberry_pi:
	$(MAKE) build GOOS=linux GOARCH=arm GOARM=7

.env: sample.env
	cp -f $< $@

renovate:
	npx --yes --package renovate -- renovate-config-validator --strict ./renovate.json 

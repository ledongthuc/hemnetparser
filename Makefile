GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run 
GOTEST=$(GOCMD) test
BINARY_NAME=hemnet-parser
DOCKER_IMG_NAME=$(BINARY_NAME)

# TEST_URL=https://www.hemnet.se/bostad/lagenhet-2rum-stockholms-kommun-hasselby-torg-2-16554758?utm_campaign=bevaka&utm_content=rad1&utm_medium=html&utm_source=mail&utm_term=basic
# TEST_URL=https://www.hemnet.se/salda/1132777?utm_campaign=slutprismail&utm_content=objekt&utm_medium=html&utm_source=mail
TEST_URL=https://www.hemnet.se/bostad/lagenhet-3rum-hasselby-lovsta-alle-stockholms-kommun-astrakangatan-101-16699595?utm_campaign=bevaka&utm_content=rad1&utm_medium=html&utm_source=mail&utm_term=basic
SHEET_ID=1S0Xzpl_a3SFGnkFWCw7M1t_NXa2VZV42Lxh70G9PTlo

all: test build
test:
	$(GOTEST) ./...
build: 
	$(GOBUILD) -o ./output/$(BINARY_NAME) -v ./cmd/hemnet-parser/main.go
run:
	$(GORUN) ./cmd/hemnet-parser/main.go --format=json $(TEST_URL)
run-sheet:
	$(GORUN) ./cmd/hemnet-parser/main.go \
		--format=googlesheet \
		--credential=credentials.json \
		--sheet-id=$(SHEET_ID) \
		--skipheader=1

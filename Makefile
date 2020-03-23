GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run 
GOTEST=$(GOCMD) test
BINARY_NAME=hemnet-parser
DOCKER_IMG_NAME=$(BINARY_NAME)

TEST_URL=https://www.hemnet.se/bostad/lagenhet-2rum-bromma-beckomberga-stockholms-kommun-agatvagen-26-16065891
SHEET_ID=1S0Xzpl_a3SFGnkFWCw7M1t_NXa2VZV42Lxh70G9PTlo

all: test build
build: 
	$(GOBUILD) -o ./output/$(BINARY_NAME) -v
run:
	$(GORUN) ./cmd/hemnet-parser/main.go --format=json $(TEST_URL)
run-sheet:
	$(GORUN) ./cmd/hemnet-parser/main.go \
		--format=googlesheet \
		--credential=credentials.json \
		--sheet-id=$(SHEET_ID) \
		--skipheader=1

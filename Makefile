
GENERATED_FILES = int.go string.go
#complex128.go float32.go int16.go int64.go uint16.go uint64.go byte.go complex64.go float64.go int32.go int8.go rune.go uint.go uint32.go uint8.go

.PHONY: generate genset

test:
	@go test

build: generate
	@go build

vet:
	gofmt -s -w $(GENERATED_FILES) set.go
	golint $(GENERATED_FILES) set.go
	govet $(GENERATED_FILES) set.go
	errcheck $(GENERATED_FILES) set.go
	staticcheck $(GENERATED_FILES) set.go
	unused $(GENERATED_FILES) set.go
	gosimple $(GENERATED_FILES) set.go

genset: genset/genset.go
	@go install kkn.fi/set/genset

generate: genset
	@go generate

clean:
	@rm -rf $(GENERATED_FILES)

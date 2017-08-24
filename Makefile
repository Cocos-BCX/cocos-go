
all: generate

clean_ffjson_base: 
	@rm -rf objects/ffjson-inception* ||:
	@rm -rf operations/ffjson-inception* ||:
	@rm objects/*_ffjson_expose.go ||:
	@rm operations/*_ffjson_expose.go ||:

clean_ffjson_gen:
	@rm objects/*_ffjson.go ||:
	@rm operations/*_ffjson.go ||:

generate: clean_ffjson_base	
	-go generate ./...

generate_new: clean_ffjson_base clean_ffjson_gen		
	-go generate ./...

test:
	go test -v ./tests/
#go test -v ./tests/common_test.go
#go test -v ./tests/subscribe_test.go
# go test -v ./tests/walletapi_test.go
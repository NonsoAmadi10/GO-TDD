test:
	go test 

test_verbose:
	go test -v 

coverage:
	go test -cover 

errcheck:
	errcheck ./...
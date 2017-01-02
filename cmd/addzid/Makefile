.phony: build, clean, testbuild

build:
	# version gets its data here:
	rm -f gitcommit.go
	/bin/echo "package main" > gitcommit.go
	/bin/echo "var LASTGITCOMMITHASH string" >> gitcommit.go
	/bin/echo "func init() { LASTGITCOMMITHASH = \"$(shell git rev-parse HEAD)\" }" >> gitcommit.go
	go build
	go install

clean:
	rm -rf testdir_* ; rm -f *~; rm -rf diffdir_*

testbuild:
	go test -c -gcflags "-N -l" -v

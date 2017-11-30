
# NOTE: This Makefile is only necessary if you 
# plan on developing the msgp tool and library.
# Installation can still be performed with a
# normal `go install`.


# generated integration test files
GGEN = ./_generated/generated.go ./_generated/generated_test.go
# generated unit test files

MGEN = ./msgp/defgen_test.go ./msgp/nestedgen_test.go

# generated zebra layer above msgp
ZGEN = ./zebra/zebra_gen.go

SHELL := /bin/bash

BIN = $(GOBIN)/zebrapack

.PHONY: clean wipe install get-deps bench all dev

dev: clean install test

$(BIN): */*.go *.go
	@go install && (cd ./cmd/addzid && make)

install:
	go install  && (cd ./cmd/addzid && make)

$(GGEN): ./_generated/def.go
	go generate ./_generated

$(MGEN): ./msgp/defs_test.go
	go generate ./msgp

$(ZGEN): ./zebra/zebra.go
	go install
	go generate ./zebra
	zebrapack -method-prefix Zebra -o zebra/zebra_zebra.go -file zebra/zebra.go

test: all
	go test -v ./parse
	go test -v ./msgp
	go test -v ./_generated
	go test -v ./zebra
	# and test addzid
	go test -v ./cmd/addzid
	# build and run on testdata/
	go build -o ./zebrapack && ./zebrapack -write-schema testdata/my.z -file testdata/my.go && go test -v ./testdata
	cd testdata && go generate && go test -v
	./zebrapack -file testdata/my.go && go test -v ./testdata/my_gen_test.go ./testdata/my.go ./testdata/my_gen.go
	./zebrapack -file testdata/my.go -msgp -o testdata/my_msgp_gen.go -method-prefix=MSGP -tests=false -io=false # test the -method-prefix flag
	./zebrapack -schema-to-go testdata/my.z > /tmp/remy.go && echo "func main() {}" >> /tmp/remy.go && go run /tmp/remy.go && rm /tmp/remy.go

bench: all
	go test -bench . ./msgp
	go test -bench . ./_generated

clean:
	$(RM) $(GGEN) $(MGEN) 

wipe: clean
	$(RM) $(BIN)

get-deps:
	go get -d -t ./...

all: install $(GGEN) $(MGEN) $(ZGEN)

# travis CI enters here
travis:
	go get -d -t ./...
	go build -o "$${GOPATH%%:*}/bin/msgp" .
	go generate ./msgp
	go generate ./_generated
	go test ./msgp
	go test ./_generated

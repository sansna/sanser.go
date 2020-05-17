PB_TARGETS=$(patsubst proto/%.proto, proto/%.pb.go, $(wildcard proto/*.proto))
all: $(PB_TARGETS)
	go run src/main.go

$(PB_TARGETS):proto/%.pb.go:proto/%.proto
	/opt/soft/protoc/bin/protoc --gogofast_out=. proto/$*.proto

clean:
	rm proto/*.pb.go

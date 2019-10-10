all: server


run_server: server
	output/bin/xcross.server

.PHONY: server
server: output
	env GO111MODULE=on go build -o output/bin/xcross.server ./server

output:
	mkdir -p output/bin

clean:
	rm -r ./output/*


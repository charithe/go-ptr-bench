N:=15
COUNT:=5
RESULTS_DIR:=results
# Valid value are: Primitive, Reference
BENCHMARK:=Primitive

.PHONY: run draw stat deps clean

run:
	@go run main.go -n=$(N)
	@gofmt -w ./bench
	@mkdir -p $(RESULTS_DIR)
	@go test -count=$(COUNT) -bench='Benchmark$(BENCHMARK)Fields' -benchmem ./bench/... | tee $(RESULTS_DIR)/$(BENCHMARK).txt

draw: run
	@benchdraw --filter='Benchmark$(BENCHMARK)Fields' --group=size --x=method --input=$(RESULTS_DIR)/$(BENCHMARK).txt --output=$(RESULTS_DIR)/$(BENCHMARK).svg
	@xdg-open $(RESULTS_DIR)/$(BENCHMARK).svg

stat: run
	@benchstat $(RESULTS_DIR)/$(BENCHMARK).txt

deps:
	@GO111MODULE=off go get -u github.com/cep21/benchdraw golang.org/x/perf/cmd/benchstat 
	@go mod download

clean:
	-rm -rf $(RESULTS_DIR)

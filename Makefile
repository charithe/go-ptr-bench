N:=15
COUNT:=5
RESULTS_DIR:=results

.PHONY: run
run:
	@go run main.go -n=$(N)
	@gofmt -w ./bench/...
	@mkdir -p $(RESULTS_DIR)
	@go test -count=$(COUNT) -bench=. -benchmem ./bench/... | tee $(RESULTS_DIR)/results.txt
	@benchdraw --group=fields --x=method --input=$(RESULTS_DIR)/results.txt --output=$(RESULTS_DIR)/results.svg
	@xdg-open $(RESULTS_DIR)/results.svg

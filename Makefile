ALL := \
	bench \
	json_bench \
	parallele_bench \
	json_parallele_bench

all: $(ALL)

.PHONY: $(ALL)

$(ALL):
	@echo TEST: $@
	@go run ./$@.go 2>/dev/null
	@echo --------------------
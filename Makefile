ALL := \
	bench \
	json_bench \
	parallele_bench \
	json_parallele_bench

all: $(ALL)

all-plot: $(ALL:%=%_plot)

.PHONY: all all-plot $(ALL) $(ALL:%=%_plot)

$(ALL):
	@echo TEST: $@ $$MESSAGES_NUM
	@# print a sample of the log to check what it is
	@echo sample: >&2
	@MESSAGES_NUM=1 go run ./$@.go >/dev/null
	@echo --------------------
	@echo results: >&2
	@go run ./$@.go 2>/dev/null
	@echo _____________________
	@echo ---------------------

$(ALL:%=%_plot):
	@for i in 10 100 1000 10000 100000 1000000 ; do \
		MESSAGES_NUM=$$i $(MAKE) $(@:%_plot=%) 2>&1 ; \
	done

README.md: _README.md
	@mv _README.md README.md
	@echo Generated README.md

README_TARGET = all-plot

_README.md:
	@echo \# Golang logger benchmark > _README.md
	@echo $$ >> _README.md
	@echo Generated README.md >> _README.md
	@echo "> Makefile : $(ALL) $(ALL:%=%_plot)" >> _README.md
	@echo \`\`\`sh >> _README.md
	@echo make $(README_TARGET) >> _README.md
	@echo \`\`\` >> _README.md
	@echo \`\`\`sh >> _README.md
	@$(MAKE) $(README_TARGET) | tee -a _README.md

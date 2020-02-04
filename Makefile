ALL := \
	bench \
	burst_bench \
	json_bench \
	burst_json_bench \
	parallele_bench \
	burst_parallele_bench \
	json_parallele_bench \
	burst_json_parallele_bench

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

define DISCLAMER
This benchmark introduces a conscious bias. A considered fair bias considering UX.
This should be interpreted as modulo the "most user-friendly call".

Each library also have the benefit/penality of logging its own name. It's considered fair in this benchmark.

This benchmark does not intend to cover the advanced usages of the different loggers, just basics:

* `logger.Info(msg, loggerlog.Ctx("Logger", dummy))`
* `logrus.WithField("Logrus", dummy).Info(msg)`
* `zap.Info(msg, zap.Stringer("Zap", dummy))`
* `zap.Sugar().Infow(msg, "ZapSugar", dummy)`
* `zerolog.Info().Interface("Zerolog", dummy).Msg (msg)`

This is considered fair as some libraries doesn't support disabling the
addition of a call stack inside there logs while using the standard jsonEncoder by example.
This is a fair UX criteria not to forget in this benchmark.
endef

export DISCLAMER

_README.md:
	@echo \# Golang logger benchmark > _README.md
	@echo "" >> _README.md
	@echo Generated README.md >> _README.md
	@echo "> Makefile : $(ALL) $(ALL:%=%_plot)" >> _README.md
	@echo "" >> _README.md
	@echo "$$DISCLAMER" | tee -a _README.md
	@echo "" >> _README.md
	@echo \`\`\`sh >> _README.md
	@echo make $(README_TARGET) >> _README.md
	@echo \`\`\` >> _README.md
	@echo "" >> _README.md
	@echo \`\`\`sh >> _README.md
	@$(MAKE) $(README_TARGET) | tee -a _README.md
.PHONY: _README.md

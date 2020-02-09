all: app

RM := rm -rf
CP := cp
MKDIR := mkdir -p
STATICDIR := static
OUTDIR := public

app: public go js static

public:
	$(MKDIR) $(OUTDIR)

go: $(wildcard go/*.go)
	make -C go
	$(CP) go/main.wasm $(OUTDIR)

js: vendor

vendor:
	$(CP) js/vendor/wasm_exec.js $(OUTDIR)

static: $(wildcard $(STATICDIR)/*.*)
	$(CP) $^ $(OUTDIR)

.PHONY: clean
clean:
	$(RM) $(OUTDIR)
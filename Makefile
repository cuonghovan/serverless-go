TOPTARGETS := build deploy
SUBDIRS := resthandler

$(TOPTARGETS): $(SUBDIRS)

$(SUBDIRS):
	@$(MAKE) -C $@ $(MAKECMDGOALS)

clean:
	rm -rf ./bin

.PHONY: deps clean $(TOPTARGETS) $(SUBDIRS)

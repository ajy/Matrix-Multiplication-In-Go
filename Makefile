include $(GOROOT)/src/Make.inc

TARG=mypackage
GOFILES=\
	beta1.go\
	structure_defs.go\
	ChannelRowColMultiplier.go\
	read_csv.go\

include $(GOROOT)/src/Make.pkg

include ~/go/src/Make.inc

TARG=matmul

GOFILES=\
	structure_defs.go\
	ChannelRowColMultiplier.go\
	read_csv.go\

include ~/go/src/Make.cmd

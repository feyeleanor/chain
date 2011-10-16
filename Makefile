include $(GOROOT)/src/Make.inc

TARG=feyeleanor/chain

GOFILES=\
	node.go\
	cell.go\
	association_list.go\
	chain.go

include $(GOROOT)/src/Make.pkg
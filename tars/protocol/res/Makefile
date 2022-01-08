all:
	tars2go -add-servant=false -tarsPath github.com/jslyzt/tarsgo/tars *.tars
	#override file
	sed -i 's|"endpointf"|"github.com/jslyzt/tarsgo/tars/protocol/res/endpointf"|g' queryf/QueryF_IF.go

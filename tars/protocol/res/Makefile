all:
	tars2go -add-servant=false -tarsPath tarsgo/tars *.tars
	#override file
	sed -i 's|"endpointf"|"tarsgo/tars/protocol/res/endpointf"|g' queryf/QueryF_IF.go

APP       := ZipkinTraceApp
TARGET    := ZipkinTraceClient
MFLAGS    :=
DFLAGS    :=
CONFIG    := client
STRIP_FLAG:= N
J2GO_FLAG:= 

libpath=${subst :, ,$(GOPATH)}
$(foreach path,$(libpath),$(eval -include $(path)/src/github.com/jslyzt/tarsgo/tars/makefile.tars))

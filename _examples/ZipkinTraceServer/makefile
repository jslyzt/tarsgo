APP       := ZipkinTraceApp
TARGET    := ZipkinTraceServer
MFLAGS    :=
DFLAGS    :=
CONFIG    := client
STRIP_FLAG:= N
J2GO_FLAG:= 

libpath=${subst :, ,$(GOPATH)}
$(foreach path,$(libpath),$(eval -include $(path)/src/tarsgo/tars/makefile.tars))

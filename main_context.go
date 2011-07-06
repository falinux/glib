package glib

/*
#include <glib-object.h>
*/
import "C"

type MainContext struct {
	Object
}

func (k MainContext) GMainContext() *C.GMainContext {
	return (*C.GMainContext)(k.GPointer())
}

func (k MainContext) Iteration(may_block bool) bool {
	return C.g_main_context_iteration(k.GMainContext(), GBoolean(may_block))!=0
}

func (k MainContext) Pending() bool {
	return C.g_main_context_pending(k.GMainContext()) != 0
}

func NewMainContext() *MainContext {
	k := new(MainContext)
	k.Set(Pointer(C.g_main_context_new()))
	return k
}

func DefaultMainContext() *MainContext {
	k := new(MainContext)
	k.Set(Pointer(C.g_main_context_default()))
	return k
}

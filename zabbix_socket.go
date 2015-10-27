package main

/*
#include <sysinc.h>
#include <module.h>

typedef int (*agent_item_handler)(AGENT_REQUEST*, AGENT_RESULT*);

int socket_send(AGENT_REQUEST*, AGENT_RESULT*);
*/
import "C"
import "unsafe"

func main() {}

//export zbx_module_api_version
func zbx_module_api_version() C.int {
	return C.ZBX_MODULE_API_VERSION_ONE
}

//export zbx_module_init
func zbx_module_init() C.int {
	return C.ZBX_MODULE_OK
}

//export zbx_module_item_list
func zbx_module_item_list() *C.ZBX_METRIC {
	metrics := make([]C.ZBX_METRIC, 3)

	metrics[0] = C.ZBX_METRIC{
		key:        C.CString("socket.send"),
		function:   C.agent_item_handler(unsafe.Pointer(C.socket_send)),
		flags:      C.CF_HAVEPARAMS,
		test_param: nil,
	}

	return &metrics[0]
}

package main

/*
#include <sysinc.h>
#include <module.h>

static char* rparam(AGENT_REQUEST* request, int i) {
	return get_rparam(request, i);
}
*/
import "C"
import (
	"bytes"
	"fmt"
	"net"
	"strings"
	"time"
)

func return_error(result *C.AGENT_RESULT, format string, args ...interface{}) C.int {
	result._type = C.AR_MESSAGE
	result.msg = C.CString(fmt.Sprintf(format, args...))
	return C.SYSINFO_RET_FAIL
}

func return_string(result *C.AGENT_RESULT, data string) C.int {
	result._type = C.AR_STRING
	result.str = C.CString(data)
	return C.SYSINFO_RET_OK
}

func send(mode, address *C.char, command string) ([]byte, error) {
	conn, err := net.Dial(C.GoString(mode), C.GoString(address))
	if err != nil {
		return nil, err
	}

	conn.SetDeadline(time.Now().Add(time.Second * 5))
	defer conn.Close()

	fmt.Fprintln(conn, command)
	var reply bytes.Buffer
	if _, err := reply.ReadFrom(conn); err != nil {
		return nil, err
	}

	return reply.Bytes(), nil
}

//export socket_send
func socket_send(request *C.AGENT_REQUEST, result *C.AGENT_RESULT) C.int {
	if request.nparam < 3 {
		return return_error(result, "usage: socket.send <mode> <address> <args...>")
	}

	mode := C.rparam(request, 0)
	address := C.rparam(request, 1)
	var args []string
	for i := C.int(2); i < request.nparam; i++ {
		args = append(args, C.GoString(C.rparam(request, i)))
	}

	fmt.Println(mode, address, args)

	if reply, err := send(mode, address, strings.Join(args, " ")); err != nil {
		return return_error(result, err.Error())
	} else {
		return return_string(result, string(reply))
	}
}

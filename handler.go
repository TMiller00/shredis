package main

var Handlers = map[string]func([]Value) Value{
	"PING": ping,
}

func ping(args []Value) Value {
	if len(args) > 1 {
		return Value{typ: "error", str: "ERR wrong number of arguments for 'ping' command"}
	}

	if len(args) == 0 {
		return Value{typ: "string", str: "PONG"}
	}

	return Value{typ: "string", str: args[0].bulk}
}

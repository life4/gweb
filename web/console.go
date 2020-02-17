package web

type Console struct {
	Value
}

// LOGGING

func (console Console) log(fname, format string, args []interface{}) {
	if format == "" {
		console.Call(fname, args...)
	} else {
		console.Call(fname, append([]interface{}{format}, args...)...)
	}
}

func (console Console) Log(format string, args ...interface{}) {
	console.log("log", format, args)
}

func (console Console) Debug(format string, args ...interface{}) {
	console.log("debug", format, args)
}

func (console Console) Info(format string, args ...interface{}) {
	console.log("info", format, args)
}

func (console Console) Warning(format string, args ...interface{}) {
	console.log("warn", format, args)
}

func (console Console) Error(format string, args ...interface{}) {
	console.log("error", format, args)
}

// OTHER METHODS

func (console Console) callWithLabel(fname, label string) {
	if label == "" {
		console.Call(fname)
	} else {
		console.Call(fname, label)
	}
}

func (console Console) Clear() {
	console.Call("clear")
}

func (console Console) Count(label string) {
	console.callWithLabel("count", label)
}

func (console Console) CountReset(label string) {
	console.callWithLabel("countReset", label)
}

func (console Console) Group(label string) {
	console.callWithLabel("group", label)
}

func (console Console) GroupCollapsed(label string) {
	console.callWithLabel("groupCollapsed", label)
}

func (console Console) GroupEnd() {
	console.Call("groupEnd")
}

func (console Console) Profile(label string) {
	console.callWithLabel("profile", label)
}

func (console Console) ProfileEnd(label string) {
	console.callWithLabel("profileEnd", label)
}

func (console Console) Time(label string) {
	console.callWithLabel("time", label)
}

func (console Console) TimeEnd(label string) {
	console.callWithLabel("timeEnd", label)
}

func (console Console) TimeLog(label string) {
	console.callWithLabel("timeLog", label)
}

func (console Console) Trace(args ...interface{}) {
	console.Call("trace", args...)
}

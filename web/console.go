package web

// https://developer.mozilla.org/en-US/docs/Web/API/Console
type Console struct {
	Value
}

// LOGGING

func (console Console) log(fname, format string, args []any) {
	if format == "" {
		console.Call(fname, args...)
	} else {
		console.Call(fname, append([]any{format}, args...)...)
	}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Console/log
func (console Console) Log(format string, args ...any) {
	console.log("log", format, args)
}

// https://developer.mozilla.org/en-US/docs/Web/API/Console/debug
func (console Console) Debug(format string, args ...any) {
	console.log("debug", format, args)
}

// https://developer.mozilla.org/en-US/docs/Web/API/Console/info
func (console Console) Info(format string, args ...any) {
	console.log("info", format, args)
}

// https://developer.mozilla.org/en-US/docs/Web/API/Console/warn
func (console Console) Warning(format string, args ...any) {
	console.log("warn", format, args)
}

// https://developer.mozilla.org/en-US/docs/Web/API/Console/error
func (console Console) Error(format string, args ...any) {
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

// https://developer.mozilla.org/en-US/docs/Web/API/Console/clear
func (console Console) Clear() {
	console.Call("clear")
}

// https://developer.mozilla.org/en-US/docs/Web/API/Console/count
func (console Console) Count(label string) {
	console.callWithLabel("count", label)
}

// https://developer.mozilla.org/en-US/docs/Web/API/Console/countReset
func (console Console) CountReset(label string) {
	console.callWithLabel("countReset", label)
}

// https://developer.mozilla.org/en-US/docs/Web/API/Console/group
func (console Console) Group(label string) {
	console.callWithLabel("group", label)
}

// https://developer.mozilla.org/en-US/docs/Web/API/Console/groupCollapsed
func (console Console) GroupCollapsed(label string) {
	console.callWithLabel("groupCollapsed", label)
}

// https://developer.mozilla.org/en-US/docs/Web/API/Console/groupEnd
func (console Console) GroupEnd() {
	console.Call("groupEnd")
}

// https://developer.mozilla.org/en-US/docs/Web/API/Console/profile
func (console Console) Profile(label string) {
	console.callWithLabel("profile", label)
}

// https://developer.mozilla.org/en-US/docs/Web/API/Console/profileEnd
func (console Console) ProfileEnd(label string) {
	console.callWithLabel("profileEnd", label)
}

// https://developer.mozilla.org/en-US/docs/Web/API/Console/time
func (console Console) Time(label string) {
	console.callWithLabel("time", label)
}

// https://developer.mozilla.org/en-US/docs/Web/API/Console/timeEnd
func (console Console) TimeEnd(label string) {
	console.callWithLabel("timeEnd", label)
}

// https://developer.mozilla.org/en-US/docs/Web/API/Console/timeLog
func (console Console) TimeLog(label string) {
	console.callWithLabel("timeLog", label)
}

// https://developer.mozilla.org/en-US/docs/Web/API/Console/trace
func (console Console) Trace(args ...any) {
	console.Call("trace", args...)
}

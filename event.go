package gweb

// https://developer.mozilla.org/en-US/docs/Web/API/Event
type Event struct {
	Value
}

// GETTERS

// https://developer.mozilla.org/en-US/docs/Web/API/Event/bubbles
func (event *Event) Bubbles() bool {
	return event.Get("bubbles").Bool()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Event/cancelable
func (event *Event) Cancelable() bool {
	return event.Get("cancelable").Bool()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Event/composed
func (event *Event) Composed() bool {
	return event.Get("composed").Bool()
}

func (event *Event) CurrentTarget() Value {
	return event.Get("currentTarget")
}

// https://developer.mozilla.org/en-US/docs/Web/API/Event/isTrusted
func (event *Event) Trusted() bool {
	return event.Get("isTrusted").Bool()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Event/type
func (event *Event) EventType() string {
	return event.Get("type").String()
}

// METHODS

// https://developer.mozilla.org/en-US/docs/Web/API/Event/preventDefault
func (event *Event) PreventDefault() {
	event.Call("preventDefault")
}

// https://developer.mozilla.org/en-US/docs/Web/API/Event/stopImmediatePropagation
func (event *Event) StopImmediatePropagation() {
	event.Call("stopImmediatePropagation")
}

// https://developer.mozilla.org/en-US/docs/Web/API/Event/stopPropagation
func (event *Event) StopPropagation() {
	event.Call("stopPropagation")
}

// EVENT TYPES

type EventType string

// EVENT_TYPE_ERROR means that A resource failed to load.
const EVENT_TYPE_ERROR = EventType("error")

// EVENT_TYPE_ABORT means that The loading of a resource has been aborted.
const EVENT_TYPE_ABORT = EventType("abort")

// EVENT_TYPE_LOAD means that A resource and its dependent resources have finished loading.
const EVENT_TYPE_LOAD = EventType("load")

// EVENT_TYPE_BEFOREUNLOAD means that The window, the document and its resources are about to be unloaded.
const EVENT_TYPE_BEFOREUNLOAD = EventType("beforeunload")

// EVENT_TYPE_UNLOAD means that The document or a dependent resource is being unloaded.
const EVENT_TYPE_UNLOAD = EventType("unload")

// EVENT_TYPE_ONLINE means that The browser has gained access to the network.
const EVENT_TYPE_ONLINE = EventType("online")

// EVENT_TYPE_OFFLINE means that The browser has lost access to the network.
const EVENT_TYPE_OFFLINE = EventType("offline")

// EVENT_TYPE_FOCUS means that An element has received focus (does not bubble).
const EVENT_TYPE_FOCUS = EventType("focus")

// EVENT_TYPE_BLUR means that An element has lost focus (does not bubble).
const EVENT_TYPE_BLUR = EventType("blur")

// EVENT_TYPE_OPEN means that A WebSocket connection has been established.
const EVENT_TYPE_OPEN = EventType("open")

// EVENT_TYPE_MESSAGE means that A message is received through a WebSocket.
const EVENT_TYPE_MESSAGE = EventType("message")

// EVENT_TYPE_CLOSE means that A WebSocket connection has been closed.
const EVENT_TYPE_CLOSE = EventType("close")

// EVENT_TYPE_PAGEHIDE means that A session history entry is being traversed from.
const EVENT_TYPE_PAGEHIDE = EventType("pagehide")

// EVENT_TYPE_PAGESHOW means that A session history entry is being traversed to.
const EVENT_TYPE_PAGESHOW = EventType("pageshow")

// EVENT_TYPE_POPSTATE means that A session history entry is being navigated to (in certain cases).
const EVENT_TYPE_POPSTATE = EventType("popstate")

// EVENT_TYPE_ANIMATIONSTART means that A CSS animation has started.
const EVENT_TYPE_ANIMATIONSTART = EventType("animationstart")

// EVENT_TYPE_ANIMATIONCANCEL means that A CSS animation has aborted.
const EVENT_TYPE_ANIMATIONCANCEL = EventType("animationcancel")

// EVENT_TYPE_ANIMATIONEND means that A CSS animation has completed.
const EVENT_TYPE_ANIMATIONEND = EventType("animationend")

// EVENT_TYPE_ANIMATIONITERATION means that A CSS animation is repeated.
const EVENT_TYPE_ANIMATIONITERATION = EventType("animationiteration")

// EVENT_TYPE_TRANSITIONSTART means that A CSS transition has actually started (fired after any delay).
const EVENT_TYPE_TRANSITIONSTART = EventType("transitionstart")

// EVENT_TYPE_TRANSITIONCANCEL means that A CSS transition has been cancelled.
const EVENT_TYPE_TRANSITIONCANCEL = EventType("transitioncancel")

// EVENT_TYPE_TRANSITIONEND means that A CSS transition has completed.
const EVENT_TYPE_TRANSITIONEND = EventType("transitionend")

// EVENT_TYPE_TRANSITIONRUN means that A CSS transition has begun running (fired before any delay starts).
const EVENT_TYPE_TRANSITIONRUN = EventType("transitionrun")

// EVENT_TYPE_RESET means that The reset button is pressed
const EVENT_TYPE_RESET = EventType("reset")

// EVENT_TYPE_SUBMIT means that The submit button is pressed
const EVENT_TYPE_SUBMIT = EventType("submit")

// EVENT_TYPE_BEFOREPRINT means that The print dialog is opened
const EVENT_TYPE_BEFOREPRINT = EventType("beforeprint")

// EVENT_TYPE_AFTERPRINT means that The print dialog is closed
const EVENT_TYPE_AFTERPRINT = EventType("afterprint")

// EVENT_TYPE_COMPOSITIONSTART means that The composition of a passage of text is prepared (similar to keydown for a keyboard input, but works with other inputs such as speech recognition).
const EVENT_TYPE_COMPOSITIONSTART = EventType("compositionstart")

// EVENT_TYPE_COMPOSITIONUPDATE means that A character is added to a passage of text being composed.
const EVENT_TYPE_COMPOSITIONUPDATE = EventType("compositionupdate")

// EVENT_TYPE_COMPOSITIONEND means that The composition of a passage of text has been completed or canceled.
const EVENT_TYPE_COMPOSITIONEND = EventType("compositionend")

// EVENT_TYPE_FULLSCREENCHANGE means that An element was turned to fullscreen mode or back to normal mode.
const EVENT_TYPE_FULLSCREENCHANGE = EventType("fullscreenchange")

// EVENT_TYPE_FULLSCREENERROR means that It was impossible to switch to fullscreen mode for technical reasons or because the permission was denied.
const EVENT_TYPE_FULLSCREENERROR = EventType("fullscreenerror")

// EVENT_TYPE_RESIZE means that The document view has been resized.
const EVENT_TYPE_RESIZE = EventType("resize")

// EVENT_TYPE_SCROLL means that The document view or an element has been scrolled.
const EVENT_TYPE_SCROLL = EventType("scroll")

// EVENT_TYPE_CUT means that The selection has been cut and copied to the clipboard
const EVENT_TYPE_CUT = EventType("cut")

// EVENT_TYPE_COPY means that The selection has been copied to the clipboard
const EVENT_TYPE_COPY = EventType("copy")

// EVENT_TYPE_PASTE means that The item from the clipboard has been pasted
const EVENT_TYPE_PASTE = EventType("paste")

// EVENT_TYPE_KEYDOWN means that ANY key is pressed
const EVENT_TYPE_KEYDOWN = EventType("keydown")

// EVENT_TYPE_KEYPRESS means that ANY key except Shift, Fn, CapsLock is in pressed position. (Fired continously.)
const EVENT_TYPE_KEYPRESS = EventType("keypress")

// EVENT_TYPE_KEYUP means that ANY key is released
const EVENT_TYPE_KEYUP = EventType("keyup")

// EVENT_TYPE_AUXCLICK means that A pointing device button (ANY non-primary button) has been pressed and released on an element.
const EVENT_TYPE_AUXCLICK = EventType("auxclick")

// EVENT_TYPE_CLICK means that A pointing device button (ANY button; soon to be primary button only) has been pressed and released on an element.
const EVENT_TYPE_CLICK = EventType("click")

// EVENT_TYPE_CONTEXTMENU means that The right button of the mouse is clicked (before the context menu is displayed).
const EVENT_TYPE_CONTEXTMENU = EventType("contextmenu")

// EVENT_TYPE_DBLCLICK means that A pointing device button is clicked twice on an element.
const EVENT_TYPE_DBLCLICK = EventType("dblclick")

// EVENT_TYPE_MOUSEDOWN means that A pointing device button is pressed on an element.
const EVENT_TYPE_MOUSEDOWN = EventType("mousedown")

// EVENT_TYPE_MOUSEENTER means that A pointing device is moved onto the element that has the listener attached.
const EVENT_TYPE_MOUSEENTER = EventType("mouseenter")

// EVENT_TYPE_MOUSELEAVE means that A pointing device is moved off the element that has the listener attached.
const EVENT_TYPE_MOUSELEAVE = EventType("mouseleave")

// EVENT_TYPE_MOUSEMOVE means that A pointing device is moved over an element. (Fired continously as the mouse moves.)
const EVENT_TYPE_MOUSEMOVE = EventType("mousemove")

// EVENT_TYPE_MOUSEOVER means that A pointing device is moved onto the element that has the listener attached or onto one of its children.
const EVENT_TYPE_MOUSEOVER = EventType("mouseover")

// EVENT_TYPE_MOUSEOUT means that A pointing device is moved off the element that has the listener attached or off one of its children.
const EVENT_TYPE_MOUSEOUT = EventType("mouseout")

// EVENT_TYPE_MOUSEUP means that A pointing device button is released over an element.
const EVENT_TYPE_MOUSEUP = EventType("mouseup")

// EVENT_TYPE_POINTERLOCKCHANGE means that The pointer was locked or released.
const EVENT_TYPE_POINTERLOCKCHANGE = EventType("pointerlockchange")

// EVENT_TYPE_POINTERLOCKERROR means that It was impossible to lock the pointer for technical reasons or because the permission was denied.
const EVENT_TYPE_POINTERLOCKERROR = EventType("pointerlockerror")

// EVENT_TYPE_SELECT means that Some text is being selected.
const EVENT_TYPE_SELECT = EventType("select")

// EVENT_TYPE_WHEEL means that A wheel button of a pointing device is rotated in any direction.
const EVENT_TYPE_WHEEL = EventType("wheel")

// EVENT_TYPE_DRAG means that An element or text selection is being dragged (Fired continuously every 350ms).
const EVENT_TYPE_DRAG = EventType("drag")

// EVENT_TYPE_DRAGEND means that A drag operation is being ended (by releasing a mouse button or hitting the escape key).
const EVENT_TYPE_DRAGEND = EventType("dragend")

// EVENT_TYPE_DRAGENTER means that A dragged element or text selection enters a valid drop target.
const EVENT_TYPE_DRAGENTER = EventType("dragenter")

// EVENT_TYPE_DRAGSTART means that The user starts dragging an element or text selection.
const EVENT_TYPE_DRAGSTART = EventType("dragstart")

// EVENT_TYPE_DRAGLEAVE means that A dragged element or text selection leaves a valid drop target.
const EVENT_TYPE_DRAGLEAVE = EventType("dragleave")

// EVENT_TYPE_DRAGOVER means that An element or text selection is being dragged over a valid drop target. (Fired continuously every 350ms.)
const EVENT_TYPE_DRAGOVER = EventType("dragover")

// EVENT_TYPE_DROP means that An element is dropped on a valid drop target.
const EVENT_TYPE_DROP = EventType("drop")

// EVENT_TYPE_AUDIOPROCESS means that The input buffer of a ScriptProcessorNode is ready to be processed.
const EVENT_TYPE_AUDIOPROCESS = EventType("audioprocess")

// EVENT_TYPE_CANPLAY means that The browser can play the media, but estimates that not enough data has been loaded to play the media up to its end without having to stop for further buffering of content.
const EVENT_TYPE_CANPLAY = EventType("canplay")

// EVENT_TYPE_CANPLAYTHROUGH means that The browser estimates it can play the media up to its end without stopping for content buffering.
const EVENT_TYPE_CANPLAYTHROUGH = EventType("canplaythrough")

// EVENT_TYPE_COMPLETE means that The rendering of an OfflineAudioContext is terminated.
const EVENT_TYPE_COMPLETE = EventType("complete")

// EVENT_TYPE_DURATIONCHANGE means that The duration attribute has been updated.
const EVENT_TYPE_DURATIONCHANGE = EventType("durationchange")

// EVENT_TYPE_EMPTIED means that The media has become empty; for example, this event is sent if the media has already been loaded (or partially loaded), and the load() method is called to reload it.
const EVENT_TYPE_EMPTIED = EventType("emptied")

// EVENT_TYPE_ENDED means that Playback has stopped because the end of the media was reached.
const EVENT_TYPE_ENDED = EventType("ended")

// EVENT_TYPE_LOADEDDATA means that The first frame of the media has finished loading.
const EVENT_TYPE_LOADEDDATA = EventType("loadeddata")

// EVENT_TYPE_LOADEDMETADATA means that The metadata has been loaded.
const EVENT_TYPE_LOADEDMETADATA = EventType("loadedmetadata")

// EVENT_TYPE_PAUSE means that Playback has been paused.
const EVENT_TYPE_PAUSE = EventType("pause")

// EVENT_TYPE_PLAY means that Playback has begun.
const EVENT_TYPE_PLAY = EventType("play")

// EVENT_TYPE_PLAYING means that Playback is ready to start after having been paused or delayed due to lack of data.
const EVENT_TYPE_PLAYING = EventType("playing")

// EVENT_TYPE_RATECHANGE means that The playback rate has changed.
const EVENT_TYPE_RATECHANGE = EventType("ratechange")

// EVENT_TYPE_SEEKED means that A seek operation completed.
const EVENT_TYPE_SEEKED = EventType("seeked")

// EVENT_TYPE_SEEKING means that A seek operation began.
const EVENT_TYPE_SEEKING = EventType("seeking")

// EVENT_TYPE_STALLED means that The user agent is trying to fetch media data, but data is unexpectedly not forthcoming.
const EVENT_TYPE_STALLED = EventType("stalled")

// EVENT_TYPE_SUSPEND means that Media data loading has been suspended.
const EVENT_TYPE_SUSPEND = EventType("suspend")

// EVENT_TYPE_TIMEUPDATE means that The time indicated by the currentTime attribute has been updated.
const EVENT_TYPE_TIMEUPDATE = EventType("timeupdate")

// EVENT_TYPE_VOLUMECHANGE means that The volume has changed.
const EVENT_TYPE_VOLUMECHANGE = EventType("volumechange")

// EVENT_TYPE_WAITING means that Playback has stopped because of a temporary lack of data.
const EVENT_TYPE_WAITING = EventType("waiting")

// EVENT_TYPE_LOADEND means that Progress has stopped (after "error", "abort" or "load" have been dispatched).
const EVENT_TYPE_LOADEND = EventType("loadend")

// EVENT_TYPE_LOADSTART means that Progress has begun.
const EVENT_TYPE_LOADSTART = EventType("loadstart")

// EVENT_TYPE_PROGRESS means that In progress.
const EVENT_TYPE_PROGRESS = EventType("progress")

// EVENT_TYPE_TIMEOUT means that Progression is terminated due to preset time expiring.
const EVENT_TYPE_TIMEOUT = EventType("timeout")

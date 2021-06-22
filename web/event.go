package web

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
func (event *Event) EventType() EventType {
	return EventType(event.Get("type").String())
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
// https://developer.mozilla.org/en-US/docs/Web/Events
type EventType string

const (
	// EventTypeError means that a resource failed to load.
	EventTypeError = EventType("error")

	// EventTypeAbort means that the loading of a resource has been aborted.
	EventTypeAbort = EventType("abort")

	// EventTypeLoad means that a resource and its dependent resources have finished loading.
	EventTypeLoad = EventType("load")

	// EventTypeBeforeUnload means that the window, the document and its resources are about to be unloaded.
	EventTypeBeforeUnload = EventType("beforeunload")

	// EventTypeUnload means that the document or a dependent resource is being unloaded.
	EventTypeUnload = EventType("unload")

	// EventTypeOnline means that the browser has gained access to the network.
	EventTypeOnline = EventType("online")

	// EventTypeOffline means that the browser has lost access to the network.
	EventTypeOffline = EventType("offline")

	// EventTypeFocus means that an element has received focus (does not bubble).
	EventTypeFocus = EventType("focus")

	// EventTypeBlur means that an element has lost focus (does not bubble).
	EventTypeBlur = EventType("blur")

	// EventTypeOpen means that a WebSocket connection has been established.
	EventTypeOpen = EventType("open")

	// EventTypeMessage means that a message is received through a WebSocket.
	EventTypeMessage = EventType("message")

	// EventTypeClose means that a WebSocket connection has been closed.
	EventTypeClose = EventType("close")

	// EventTypePageHide means that a session history entry is being traversed from.
	EventTypePageHide = EventType("pagehide")

	// EventTypePageShow means that a session history entry is being traversed to.
	EventTypePageShow = EventType("pageshow")

	// EventTypePopState means that a session history entry is being navigated to (in certain cases).
	EventTypePopState = EventType("popstate")

	// EventTypeAnimationStart means that a CSS animation has started.
	EventTypeAnimationStart = EventType("animationstart")

	// EventTypeAnimationCancel means that a CSS animation has aborted.
	EventTypeAnimationCancel = EventType("animationcancel")

	// EventTypeAnimationEnd means that a CSS animation has completed.
	EventTypeAnimationEnd = EventType("animationend")

	// EventTypeAnimationiteration means that a CSS animation is repeated.
	EventTypeAnimationIteration = EventType("animationiteration")

	// EventTypeTransitionStart means that a CSS transition has actually started (fired after any delay).
	EventTypeTransitionStart = EventType("transitionstart")

	// EventTypeTransitionCancel means that a CSS transition has been cancelled.
	EventTypeTransitionCancel = EventType("transitioncancel")

	// EventTypeTransitionEnd means that a CSS transition has completed.
	EventTypeTransitionEnd = EventType("transitionend")

	// EventTypeTransitionRun means that a CSS transition has begun running (fired before any delay starts).
	EventTypeTransitionRun = EventType("transitionrun")

	// EventTypeReset means that the reset button is pressed.
	// https://developer.mozilla.org/en-US/docs/Web/API/HTMLFormElement/reset_event
	EventTypeReset = EventType("reset")

	// EventTypeSubmit means that the submit button is pressed.
	// https://developer.mozilla.org/en-US/docs/Web/API/HTMLFormElement/submit_event
	EventTypeSubmit = EventType("submit")

	// EventTypeBeforePrint means that the print dialog is opened.
	EventTypeBeforePrint = EventType("beforeprint")

	// EventTypeAfterPrint means that the print dialog is closed.
	EventTypeAfterPrint = EventType("afterprint")

	// EventTypeCompositionStart means that the composition of a passage of text is prepared (similar to keydown for a keyboard input, but works with other inputs such as speech recognition).
	EventTypeCompositionStart = EventType("compositionstart")

	// EventTypeCompositionUpdate means that a character is added to a passage of text being composed.
	EventTypeCompositionUpdate = EventType("compositionupdate")

	// EventTypeCompositionEnd means that the composition of a passage of text has been completed or canceled.
	EventTypeCompositionEnd = EventType("compositionend")

	// EventTypeFullscreenChange means that an element was turned to fullscreen mode or back to normal mode.
	EventTypeFullscreenChange = EventType("fullscreenchange")

	// EventTypeFullscreenError means that It was impossible to switch to fullscreen mode for technical reasons or because the permission was denied.
	EventTypeFullscreenError = EventType("fullscreenerror")

	// EventTypeResize means that the document view has been resized.
	EventTypeResize = EventType("resize")

	// EventTypeScroll means that the document view or an element has been scrolled.
	EventTypeScroll = EventType("scroll")

	// EventTypeCut means that the selection has been cut and copied to the clipboard
	EventTypeCut = EventType("cut")

	// EventTypeCopy means that the selection has been copied to the clipboard
	EventTypeCopy = EventType("copy")

	// EventTypePaste means that the item from the clipboard has been pasted
	EventTypePaste = EventType("paste")

	// EventTypeKeyDown means that aNY key is pressed
	EventTypeKeyDown = EventType("keydown")

	// EventTypeKeyPress means that aNY key except Shift, Fn, CapsLock is in pressed position. (Fired continously.)
	EventTypeKeyPress = EventType("keypress")

	// EventTypeKeyUp means that aNY key is released
	EventTypeKeyUp = EventType("keyup")

	// EventTypeAuxClick means that a pointing device button (ANY non-primary button) has been pressed and released on an element.
	EventTypeAuxClick = EventType("auxclick")

	// EventTypeClick means that a pointing device button (ANY button; soon to be primary button only) has been pressed and released on an element.
	EventTypeClick = EventType("click")

	// EventTypeContextMenu means that the right button of the mouse is clicked (before the context menu is displayed).
	EventTypeContextMenu = EventType("contextmenu")

	// EventTypeDoubleClick means that a pointing device button is clicked twice on an element.
	EventTypeDoubleClick = EventType("dblclick")

	// EventTypeMouseDown means that a pointing device button is pressed on an element.
	EventTypeMouseDown = EventType("mousedown")

	// EventTypeMouseEnter means that a pointing device is moved onto the element that has the listener attached.
	EventTypeMouseEnter = EventType("mouseenter")

	// EventTypeMouseLeave means that a pointing device is moved off the element that has the listener attached.
	EventTypeMouseLeave = EventType("mouseleave")

	// EventTypeMouseMove means that a pointing device is moved over an element. (Fired continously as the mouse moves.)
	EventTypeMouseMove = EventType("mousemove")

	// EventTypeMouseOver means that a pointing device is moved onto the element that has the listener attached or onto one of its children.
	EventTypeMouseOver = EventType("mouseover")

	// EventTypeMouseOut means that a pointing device is moved off the element that has the listener attached or off one of its children.
	EventTypeMouseOut = EventType("mouseout")

	// EventTypeMouseUp means that a pointing device button is released over an element.
	EventTypeMouseUp = EventType("mouseup")

	// EventTypePointerLockChange means that the pointer was locked or released.
	EventTypePointerLockChange = EventType("pointerlockchange")

	// EventTypePointerLockError means that It was impossible to lock the pointer for technical reasons or because the permission was denied.
	EventTypePointerLockError = EventType("pointerlockerror")

	// EventTypeSelect means that Some text is being selected.
	EventTypeSelect = EventType("select")

	// EventTypeWheel means that a wheel button of a pointing device is rotated in any direction.
	EventTypeWheel = EventType("wheel")

	// EventTypeDrag means that an element or text selection is being dragged (Fired continuously every 350ms).
	EventTypeDrag = EventType("drag")

	// EventTypeDragEnd means that a drag operation is being ended (by releasing a mouse button or hitting the escape key).
	EventTypeDragEnd = EventType("dragend")

	// EventTypeDragEnter means that a dragged element or text selection enters a valid drop target.
	EventTypeDragEnter = EventType("dragenter")

	// EventTypeDragStart means that the user starts dragging an element or text selection.
	EventTypeDragStart = EventType("dragstart")

	// EventTypeDragLeave means that a dragged element or text selection leaves a valid drop target.
	EventTypeDragLeave = EventType("dragleave")

	// EventTypeDragOver means that an element or text selection is being dragged over a valid drop target. (Fired continuously every 350ms.)
	EventTypeDragOver = EventType("dragover")

	// EventTypeDrop means that an element is dropped on a valid drop target.
	EventTypeDrop = EventType("drop")

	// EventTypeAudioProcess means that the input buffer of a ScriptProcessorNode is ready to be processed.
	EventTypeAudioProcess = EventType("audioprocess")

	// EventTypeCanPlay means that the browser can play the media, but estimates that not enough data has been loaded to play the media up to its end without having to stop for further buffering of content.
	EventTypeCanPlay = EventType("canplay")

	// EventTypeCanPlayThrough means that the browser estimates it can play the media up to its end without stopping for content buffering.
	EventTypeCanPlayThrough = EventType("canplaythrough")

	// EventTypeComplete means that the rendering of an OfflineAudioContext is terminated.
	EventTypeComplete = EventType("complete")

	// EventTypeDurationChange means that the duration attribute has been updated.
	EventTypeDurationChange = EventType("durationchange")

	// EventTypeEmptied means that the media has become empty; for example, this event is sent if the media has already been loaded (or partially loaded), and the load() method is called to reload it.
	EventTypeEmptied = EventType("emptied")

	// EventTypeEnded means that Playback has stopped because the end of the media was reached.
	EventTypeEnded = EventType("ended")

	// EventTypeLoadedData means that the first frame of the media has finished loading.
	EventTypeLoadedData = EventType("loadeddata")

	// EventTypeLoadedMetadata means that the metadata has been loaded.
	EventTypeLoadedMetadata = EventType("loadedmetadata")

	// EventTypePause means that Playback has been paused.
	EventTypePause = EventType("pause")

	// EventTypePlay means that Playback has begun.
	EventTypePlay = EventType("play")

	// EventTypePlaying means that Playback is ready to start after having been paused or delayed due to lack of data.
	EventTypePlaying = EventType("playing")

	// EventTypeRateChange means that the playback rate has changed.
	EventTypeRateChange = EventType("ratechange")

	// EventTypeSeeked means that a seek operation completed.
	EventTypeSeeked = EventType("seeked")

	// EventTypeSeeking means that a seek operation began.
	EventTypeSeeking = EventType("seeking")

	// EventTypeStalled means that the user agent is trying to fetch media data, but data is unexpectedly not forthcoming.
	EventTypeStalled = EventType("stalled")

	// EventTypeSuspend means that Media data loading has been suspended.
	EventTypeSuspend = EventType("suspend")

	// EventTypeTimeUpdate means that the time indicated by the currentTime attribute has been updated.
	EventTypeTimeUpdate = EventType("timeupdate")

	// EventTypeVolumeChange means that the volume has changed.
	EventTypeVolumeChange = EventType("volumechange")

	// EventTypeWaiting means that Playback has stopped because of a temporary lack of data.
	EventTypeWaiting = EventType("waiting")

	// EventTypeLoadEnd means that Progress has stopped (after "error", "abort" or "load" have been dispatched).
	EventTypeLoadEnd = EventType("loadend")

	// EventTypeLoadStart means that Progress has begun.
	EventTypeLoadStart = EventType("loadstart")

	// EventTypeProgress means that In progress.
	EventTypeProgress = EventType("progress")

	// EventTypeTimeout means that Progression is terminated due to preset time expiring.
	EventTypeTimeout = EventType("timeout")
)

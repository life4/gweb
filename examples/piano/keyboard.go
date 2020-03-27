package main

import (
	"fmt"

	"github.com/life4/gweb/audio"
	"github.com/life4/gweb/web"
)

type KeyBoard struct {
	notes   map[int]map[string]float64
	context audio.AudioContext
	doc     web.Document
	sounds  map[int]map[string]*Sound
	octave  int
}

func (kbd KeyBoard) Octaves() []int {
	max := 0
	for octave := range kbd.notes {
		if octave > max {
			max = octave
		}
	}
	result := make([]int, max+1)
	for n := 0; n <= max; n++ {
		result[n] = n
	}
	return result
}

func (kbd KeyBoard) Notes() []string {
	return []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
}

func (kbd KeyBoard) Render(doc web.Document) web.HTMLElement {
	root := doc.CreateElement("div")
	for _, octave := range kbd.Octaves() {
		row := doc.CreateElement("div")
		row.SetID(fmt.Sprintf("octave-%d", octave))

		number := doc.CreateElement("span")
		number.SetText(fmt.Sprintf("%d", octave))
		number = StyleBlock(number)
		row.Node().AppendChild(number.Node())

		for _, note := range kbd.Notes() {
			_, ok := kbd.notes[octave][note]
			if !ok {
				holder := doc.CreateElement("span")
				holder = StyleBlock(holder)
				row.Node().AppendChild(holder.Node())
				continue
			}

			key := Key{Octave: octave, Note: note}
			element := key.Render(doc)
			element.EventTarget().Listen(web.EventTypeMouseDown, kbd.handlePress)
			// element.EventTarget().Listen(web.EventTypeMouseOver, kbd.handlePress)
			element.EventTarget().Listen(web.EventTypeMouseUp, kbd.handleRelease)
			element.EventTarget().Listen(web.EventTypeMouseLeave, kbd.handleRelease)
			row.Node().AppendChild(element.Node())
		}

		root.Node().AppendChild(row.Node())
	}

	doc.EventTarget().Listen(web.EventTypeKeyDown, kbd.handleKeyDown)
	doc.EventTarget().Listen(web.EventTypeKeyUp, kbd.handleKeyUp)
	kbd.doc = doc
	return root
}

func (kbd KeyBoard) play(octave int, note string) Sound {
	freq := kbd.notes[octave][note]
	return Play(kbd.context, freq)
}

func (kbd *KeyBoard) Press(octave int, note string) {
	old, ok := kbd.sounds[octave][note]
	if ok && old != nil {
		return
	}

	sound := kbd.play(octave, note)
	sounds := kbd.sounds[octave]
	if sounds == nil {
		kbd.sounds[octave] = make(map[string]*Sound)
	}
	kbd.sounds[octave][note] = &sound
}

func (kbd *KeyBoard) Release(octave int, note string) {
	sound, ok := kbd.sounds[octave][note]
	if !ok || sound == nil {
		return
	}
	sound.Stop()
	kbd.sounds[octave][note] = nil
}

func (kbd *KeyBoard) SetOctave(octave int) {
	if octave == kbd.octave {
		return
	}
	mod := len(kbd.Octaves())
	kbd.octave = (mod + octave) % mod

	// if octave has been changed, release all pressed keys
	for octave, sounds := range kbd.sounds {
		for note := range sounds {
			key := KeyFromNote(kbd.doc, octave, note)
			key.Release()
			kbd.Release(octave, note)
		}
	}
}

// handlers

func (kbd *KeyBoard) handlePress(event web.Event) {
	element := event.CurrentTarget().HTMLElement()
	key := KeyFromElement(element)
	key.Press()
	kbd.Press(key.Octave, key.Note)
}

func (kbd *KeyBoard) handleRelease(event web.Event) {
	element := event.CurrentTarget().HTMLElement()
	key := KeyFromElement(element)
	key.Release()
	kbd.Release(key.Octave, key.Note)
}

func (kbd *KeyBoard) handleKeyDown(event web.Event) {
	keyCode := event.Get("keyCode").Int()

	// change octave if arrow up or down is pressed
	if keyCode == 38 {
		kbd.SetOctave(kbd.octave - 1)
		return
	}
	if keyCode == 40 {
		kbd.SetOctave(kbd.octave + 1)
		return
	}
	// change octave on numbers pressed
	mod := len(kbd.Octaves())
	if keyCode >= 48 && keyCode <= 48+mod {
		kbd.SetOctave(keyCode - 48)
	}

	note, offset := keyToNote(keyCode)
	if note == "" {
		return
	}
	octave := (mod + kbd.octave + offset) % mod
	key := KeyFromNote(kbd.doc, octave, note)

	// if no key for the given note
	if key.Note == "" {
		return
	}

	key.Press()
	kbd.Press(octave, note)
}

func (kbd *KeyBoard) handleKeyUp(event web.Event) {
	keyCode := event.Get("keyCode").Int()
	mod := len(kbd.Octaves())

	note, offset := keyToNote(keyCode)
	if note == "" {
		return
	}
	octave := (mod + kbd.octave + offset) % mod
	key := KeyFromNote(kbd.doc, octave, note)

	// if no key for the given note
	if key.Note == "" {
		return
	}
	key.Release()
	kbd.Release(octave, note)
}

// funcs

func getNotes() map[int]map[string]float64 {
	notes := make(map[int]map[string]float64)
	notes[0] = map[string]float64{
		"A":  27.500000000000000,
		"A#": 29.135235094880619,
		"B":  30.867706328507756,
	}
	notes[1] = map[string]float64{
		"C":  32.703195662574829,
		"C#": 34.647828872109012,
		"D":  36.708095989675945,
		"D#": 38.890872965260113,
		"E":  41.203444614108741,
		"F":  43.653528929125485,
		"F#": 46.249302838954299,
		"G":  48.999429497718661,
		"G#": 51.913087197493142,
		"A":  55.000000000000000,
		"A#": 58.270470189761239,
		"B":  61.735412657015513,
	}
	notes[2] = map[string]float64{
		"C":  65.406391325149658,
		"C#": 69.295657744218024,
		"D":  73.416191979351890,
		"D#": 77.781745930520227,
		"E":  82.406889228217482,
		"F":  87.307057858250971,
		"F#": 92.498605677908599,
		"G":  97.998858995437323,
		"G#": 103.826174394986284,
		"A":  110.000000000000000,
		"A#": 116.540940379522479,
		"B":  123.470825314031027,
	}

	notes[3] = map[string]float64{
		"C":  130.812782650299317,
		"C#": 138.591315488436048,
		"D":  146.832383958703780,
		"D#": 155.563491861040455,
		"E":  164.813778456434964,
		"F":  174.614115716501942,
		"F#": 184.997211355817199,
		"G":  195.997717990874647,
		"G#": 207.652348789972569,
		"A":  220.000000000000000,
		"A#": 233.081880759044958,
		"B":  246.941650628062055,
	}

	notes[4] = map[string]float64{
		"C":  261.625565300598634,
		"C#": 277.182630976872096,
		"D":  293.664767917407560,
		"D#": 311.126983722080910,
		"E":  329.627556912869929,
		"F":  349.228231433003884,
		"F#": 369.994422711634398,
		"G":  391.995435981749294,
		"G#": 415.304697579945138,
		"A":  440.000000000000000,
		"A#": 466.163761518089916,
		"B":  493.883301256124111,
	}

	notes[5] = map[string]float64{
		"C":  523.251130601197269,
		"C#": 554.365261953744192,
		"D":  587.329535834815120,
		"D#": 622.253967444161821,
		"E":  659.255113825739859,
		"F":  698.456462866007768,
		"F#": 739.988845423268797,
		"G":  783.990871963498588,
		"G#": 830.609395159890277,
		"A":  880.000000000000000,
		"A#": 932.327523036179832,
		"B":  987.766602512248223,
	}

	notes[6] = map[string]float64{
		"C":  1046.502261202394538,
		"C#": 1108.730523907488384,
		"D":  1174.659071669630241,
		"D#": 1244.507934888323642,
		"E":  1318.510227651479718,
		"F":  1396.912925732015537,
		"F#": 1479.977690846537595,
		"G":  1567.981743926997176,
		"G#": 1661.218790319780554,
		"A":  1760.000000000000000,
		"A#": 1864.655046072359665,
		"B":  1975.533205024496447,
	}

	notes[7] = map[string]float64{
		"C":  2093.004522404789077,
		"C#": 2217.461047814976769,
		"D":  2349.318143339260482,
		"D#": 2489.015869776647285,
		"E":  2637.020455302959437,
		"F":  2793.825851464031075,
		"F#": 2959.955381693075191,
		"G":  3135.963487853994352,
		"G#": 3322.437580639561108,
		"A":  3520.000000000000000,
		"A#": 3729.310092144719331,
		"B":  3951.066410048992894,
	}

	notes[8] = map[string]float64{
		"C": 4186.009044809578154,
	}
	return notes
}

func keyToNote(key int) (string, int) {
	switch key + 32 {
	case int('z'):
		return "A", 1
	case int('x'):
		return "B", 1
	case int('c'):
		return "C", 1
	case int('v'):
		return "D", 1
	case int('b'):
		return "E", 1
	case int('n'):
		return "F", 1
	case int('m'):
		return "G", 1
	}

	switch key + 32 {
	case int('a'):
		return "A", 0
	case int('s'):
		return "B", 0
	case int('d'):
		return "C", 0
	case int('f'):
		return "D", 0
	case int('g'):
		return "E", 0
	case int('h'):
		return "F", 0
	case int('j'):
		return "G", 0
	}

	switch key + 32 {
	case int('q'):
		return "A", -1
	case int('w'):
		return "B", -1
	case int('e'):
		return "C", -1
	case int('r'):
		return "D", -1
	case int('t'):
		return "E", -1
	case int('y'):
		return "F", -1
	case int('u'):
		return "G", -1
	}

	return "", 0
}

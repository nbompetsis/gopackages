package gadget

import "fmt"

type Player interface {
	Play(string)
	Stop()
}

// TapePlayer type
type TapePlayer struct {
	Batteries string
}

// Play method of type
func (t TapePlayer) Play(song string) {
	fmt.Println("Playing song", song)
}

// Stop method of type
func (t TapePlayer) Stop() {
	fmt.Println("Stopped!")
}

// TapeRecorder type
type TapeRecorder struct {
	Microphone int
}

// Play method of type
func (t TapeRecorder) Play(song string) {
	fmt.Println("Playing song", song)
}

// Record method of type
func (t TapeRecorder) Record() {
	fmt.Println("Recording")
}

// Stop method of type
func (t TapeRecorder) Stop() {
	fmt.Println("Stopped!")
}

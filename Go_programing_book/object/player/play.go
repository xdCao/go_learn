package player

import "fmt"

type Player interface {
	Play(source string)
}

func Play(source, fileType string) {
	var p Player
	switch fileType {
	case "MP3":
		p = &Mp3Player{}
	case "WAV":
		p = &WavPlayer{}
	}
	p.Play(source)
}

type Mp3Player struct {
}

func (this *Mp3Player) Play(source string) {
	fmt.Println("Mp3 playing")
}

type WavPlayer struct {
}

func (this *WavPlayer) Play(source string) {
	fmt.Println("Wav playing")
}

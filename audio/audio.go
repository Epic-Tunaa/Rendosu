package audio

import (
	"time"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/vorbis"
	"github.com/faiface/beep/speaker"
)

type AudioEngine struct {
	sampleRate beep.SampleRate
	format beep.Format
	streamer beep.StreamSeeker
	ctrl *beep.ctrl
	timeChan chan time.Duration
	syncOffset time.Duration
}

func Init(cfg config.AudioConfig) {
	sr := beep.SampleRate(cfg.SampleRate)
	speaker.Init(sr, sr.N(time.Duration(cfg.BufferSize)))
}

func (a*AudioEngine) Load(filename string) error{
	f, err := os.Open(filename)
	if err != nil{
		return err
	}


var streamer beep.StreamSeeker
switch filepath.Ext(filename) {
case ".mp3":
	stream, format, err = mp3.Decode(f)
case ".ogg":
	streamer, format, err = vorbis.Decode(f)
}

a.ctrl = &beep.Ctrl{Streamer: beep.Loop(-1, streamer)}
a.sampleRate = format.SampleRate
a.streamer = streamers

return nil
}

func(a*AudioEngine) Play() {
	speaker.Play(a.ctrl)
}

func(a*AudioEngine) GetTime() time.Duration {
	return a.syncOffset + a.sampleRate.D(a.stream.Position())
}
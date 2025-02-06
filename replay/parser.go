package replay

import (
	"compress/zlib"
	"encoding/binary"
	"time"
)

type Replay struct {
	Mode	byte
	Version	int32
	BeatmapHash	string
	PlayerName	string
	ReplayHash	string
	Timestamp	time.Time
	Actions		[]Action
}

type Action struct {
	Time	time.Duration
	Position	mgl32.Vec2
	Keys	byte
}

func ParseReplay(filename string) (*Replay, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
}

r := bytes.NewReader(data)
replay := &Replay{}

//Parse header
binary.Read(r, binary.LittleEndian, &replay.Mode)
binary.Read(r, binary.LittleEndian, &replay.Version)

//parse compressed data
var compressedSize int32
binary.Read(r, binary.LittleEndian, &compressedSize)

zr, err := zlib.NewReader(r)
if err != nil {
	return nil, err
}

//Parse actions
actionReader := bytes.NewReader(decompressed)
for {
	var timeDiff int32
	if err := binary.Read(actionReader, binary.LittleEndian, &timeDiff); err != nil {
		break
	}

	var x, y float32
	var keys byte
	binary.Read(actionReader, binary.LittleEndian, &x)
	binary.Read(actionReader, binary.LittleEndian, &y)
	binary.Read(actionReader, binary.LittleEndian, &keys)

	replay.Actions = append(replay.Actions, Actions{
		Time: time.Duration(timeDiff) * time.Millisecond,
		Position: mgl32.Vec2{x, y}
	})
}

return replay, nil
package ffprobe

import (
	"bytes"
	"encoding/json"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"go.senan.xyz/gonic/tags/tagcommon"
)

type FFProbe struct{}

func (FFProbe) CanRead(absPath string) bool {
	switch ext := strings.ToLower(filepath.Ext(absPath)); ext {
	case ".webm", ".mp4", ".mkv":
		return true
	}
	return false
}

func (FFProbe) Read(absPath string) (tagcommon.Info, error) {
	cmd := exec.Command(
		"ffprobe",
		"-v", "quiet",
		"-print_format", "json",
		"-show_format",
		"-show_streams",
	)
	stdout := bytes.NewBuffer(make([]byte, 0, 1024))
	stderr := bytes.NewBuffer(make([]byte, 0, 1024))
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	cmd.Args = append(cmd.Args, absPath)
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	var mi MediaInfo
	if err := json.NewDecoder(stdout).Decode(&mi); err != nil {
		return nil, err
	}

	if len(mi.Streams) == 0 {
		return nil, ErrNoMediaStreams
	}

	if mi.Format.ProbeScore < 100 {
		return nil, ErrFFProbeScroeNotEnough
	}

	if mi.Format.Duration == "" {
		return nil, ErrNoMediaDuration
	}

	ret := &info{absPath, mi}

	return ret, nil
}

type info struct {
	abspath   string
	mediaInfo MediaInfo
}

func (i *info) Title() string          { return "" }
func (i *info) BrainzID() string       { return "" }
func (i *info) Artist() string         { return tagcommon.FallbackArtist }
func (i *info) Artists() []string      { return []string{tagcommon.FallbackArtist} }
func (i *info) Album() string          { return "" }
func (i *info) AlbumArtist() string    { return "" }
func (i *info) AlbumArtists() []string { return []string{} }
func (i *info) AlbumBrainzID() string  { return "" }
func (i *info) Genre() string          { return tagcommon.FallbackGenre }
func (i *info) Genres() []string       { return []string{tagcommon.FallbackGenre} }
func (i *info) TrackNumber() int       { return 0 }
func (i *info) DiscNumber() int        { return 0 }
func (i *info) Year() int              { return 0 }

func (i *info) ReplayGainTrackGain() float32 { return 0 }
func (i *info) ReplayGainTrackPeak() float32 { return 0 }
func (i *info) ReplayGainAlbumGain() float32 { return 0 }
func (i *info) ReplayGainAlbumPeak() float32 { return 0 }

func (i *info) Length() int {
	ret, _ := strconv.ParseFloat(i.mediaInfo.Format.Duration, 64)
	return int(ret / 1024)
}

func (i *info) Bitrate() int {
	ret, _ := strconv.Atoi(i.mediaInfo.Format.BitRate)
	return ret
}

func (i *info) AbsPath() string { return i.abspath }

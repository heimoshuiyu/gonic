package ffprobe

type Stream struct {
	Index         int    `json:"index"`
	CodecName     string `json:"codec_name"`
	CodecLongName string `json:"codec_long_name"`
	CodecType     string `json:"codec_type"`
	SampleRate    string `json:"sample_rate"`
	Channels      int    `json:"channels"`
	ChannelLayout string `json:"channel_layout"`
}

type Format struct {
	Filename       string `json:"filename"`
	NbStreams      int    `json:"nb_streams"`
	FormatName     string `json:"format_name"`
	FormatLongName string `json:"format_long_name"`
	Duration       string `json:"duration"`
	Size           string `json:"size"`
	BitRate        string `json:"bit_rate"`
	ProbeScore     int    `json:"probe_score"`
}

type MediaInfo struct {
	Streams []Stream `json:"streams"`
	Format  Format   `json:"format"`
}

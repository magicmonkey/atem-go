// Code generated by "stringer -type=VideoMode -linecomment"; DO NOT EDIT.

package models

import "strconv"

const (
	_VideoMode_name_0 = "525i59.94 NTSC625i50 PAL525i59.94 NTSC 16:9625i50 PAL 16:9720p50720p59.941080i501080i59.941080p23.981080p241080p251080p29.971080p501080p59.942160p23.982160p242160p252160p29.97"
	_VideoMode_name_1 = "1080p301080p60"
)

var (
	_VideoMode_index_0 = [...]uint8{0, 14, 24, 43, 58, 64, 73, 80, 90, 100, 107, 114, 124, 131, 141, 151, 158, 165, 175}
	_VideoMode_index_1 = [...]uint8{0, 7, 14}
)

func (i VideoMode) String() string {
	switch {
	case 0 <= i && i <= 17:
		return _VideoMode_name_0[_VideoMode_index_0[i]:_VideoMode_index_0[i+1]]
	case 26 <= i && i <= 27:
		i -= 26
		return _VideoMode_name_1[_VideoMode_index_1[i]:_VideoMode_index_1[i+1]]
	default:
		return "VideoMode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}

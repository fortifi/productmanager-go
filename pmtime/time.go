package pmtime

import "time"

type Time struct {
	t time.Time
}

func (t Time) ForTransport() int64 { return t.t.UnixNano() / int64(time.Millisecond) }
func Now() Time                    { return Time{t: time.Now()} }
func FromTime(t time.Time) Time    { return Time{t: t} }
func FromTransport(t int64) Time   { return Time{t: time.Unix(0, t*int64(time.Millisecond))} }

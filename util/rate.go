package util

import (
	"io"
	"time"
)

// Copy all data from r to w.
// Param b is a buffer for copy data.
// Param rate is bytes per second, rate<1 means no limit.
// Param dur is timer duration.
func LimitRateCopy(w io.Writer, r io.Reader, b []byte, rate int, dur int) (n int64, err error) {
	// No limit.
	if rate < 1 {
		return io.Copy(w, r)
	}
	if len(b) < 1 {
		return 0, io.ErrShortBuffer
	}
	if dur < 1 {
		dur = 1
	}
	// Copy n bytes per dur millisecond.
	nPerDur := int(float64(rate)/float64(1000)) * dur
	// Timer.
	d := time.Duration(dur) * time.Millisecond
	timer := time.NewTimer(d)
	defer timer.Stop()
	// Loop.
	for {
		// Copy n bytes
		err = CopyN(w, r, b, nPerDur)
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return
		}
		n += int64(nPerDur)
		<-timer.C
		timer.Reset(d)
	}
}

// Copy n bytes from r to w.
// Param b is a buffer for copy data.
func CopyN(w io.Writer, r io.Reader, b []byte, n int) error {
	var m int
	var err error
	p := b
	for n > 0 {
		// Read data.
		if n > len(p) {
			m, err = r.Read(p)
		} else {
			m, err = r.Read(p[:n])
		}
		if err != nil {
			if err == io.EOF {
				p = p[m:]
				break
			}
			return err
		}
		n -= m
		p = p[m:]
		// f.Buff is full.
		if len(p) < 1 {
			_, err = w.Write(b)
			if err != nil {
				return err
			}
			p = b
		}
	}
	// Has data.
	if len(p) < len(b) {
		_, err = w.Write(b[:len(b)-len(p)])
	}
	return err
}

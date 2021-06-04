package util

import (
	"io"
	"time"
)

// Limit the read and write rate of IO.
type RateLimiter struct {
	// Buffer for copy data.
	Buff []byte
	// Bytes per second, Rate<1 means no limit.
	Rate int
	// Duration of tick, default is 1 millisecond.
	Dur int
}

// Copy all data from r to w.
func (l *RateLimiter) Copy(w io.Writer, r io.Reader) (n int64, err error) {
	// No limit.
	if l.Rate < 1 {
		return io.Copy(w, r)
	}
	if len(l.Buff) < 1 {
		l.Buff = make([]byte, 10240)
	}
	l.Dur = MaxInt(l.Dur, 1)
	// Copy n bytes per dur millisecond.
	nPerDur := int(float64(l.Rate)/float64(1000)) * l.Dur
	// Timer.
	dur := time.Duration(l.Dur) * time.Millisecond
	timer := time.NewTimer(dur)
	defer timer.Stop()
	// Loop.
	for {
		// Copy n bytes
		err = l.copyN(w, r, nPerDur)
		if err != nil {
			return
		}
		n += int64(nPerDur)
		<-timer.C
		timer.Reset(dur)
	}
}

// Copy n bytes from r to w.
func (l *RateLimiter) copyN(w io.Writer, r io.Reader, n int) error {
	var m int
	var err error
	p := l.Buff
	for n > 0 {
		// Read data.
		if n > len(p) {
			m, err = r.Read(p)
		} else {
			m, err = r.Read(p[:n])
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		n -= m
		p = p[m:]
		// f.Buff is full.
		if len(p) < 1 {
			_, err = w.Write(l.Buff)
			if err != nil {
				return err
			}
			p = l.Buff
		}
	}
	// Has data.
	if len(p) < len(l.Buff) {
		_, err = w.Write(l.Buff[:len(l.Buff)-len(p)])
	}
	return err
}

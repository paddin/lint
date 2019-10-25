package sample

import (
	// "fmt"
	"io"
)

type Sample struct {
	Item string
}

func (s *Sample) GetItem(l io.Writer) bool {
	if l == nil {
		// fmt.Printf("nil Writer")
		return false
	}
	l.Write([]byte(s.Item))
	return true
}

func (s *Sample) ItemIsEmpty(a, b, c, d int) bool {
	if len(s.Item) == 0 {
		return true
	}
	return false
}

func SampleCompare(s1, s2 *Sample) bool {
	if s1 != nil {
		if s2 != nil {
			if s1.Item == s2.Item {
				return true
			}
		} else {
			// fmt.Printf("s2 is nil")
		}
	} else {
		// fmt.Printf("s1 is nil")
	}
	return false
}

func InvalidSlices(slice1 []string, slice2 []int) (bool, int) {
	if slice1 == nil {
		return false, 0
	}
	if slice2 == nil {
		return false, 42
	}
	slice1[42] = "12" // crash!
	if slice2 == nil {
		sliceOfSlice := slice2[:12]
		_ = sliceOfSlice
	}
	return true, 12
}

func (s *Sample) SetItem(l io.Reader) {
	if l == nil {
		// 	// fmt.Printf("nil Reader")
	}
	buf := make([]byte, 256)
	n, _ := l.Read(buf)
	s.Item = string(buf[:n])
}


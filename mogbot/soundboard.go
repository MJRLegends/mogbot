package mogbot

import "sync"

type Soundboard struct {
	mu        *sync.Mutex
	isPlaying bool
}

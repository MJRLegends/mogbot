package mogbot

import (
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Soundboard struct {
	mu     *sync.Mutex
	buffer [][]byte
}

type Sound struct {
	ID   int64
	Name string
	Path string
}

type SoundService interface {
	AddSound(*Sound) error
	GetSoundByID(id string) (*Sound, error)
	GetSoundByName(name string) (*Sound, error)
	GetAllSounds() ([]*Sound, error)
	UpdateSound(id string, fields map[string]interface{}) error
	RemoveSound(id string)
}

func (s *Soundboard) FillBuffer(b [][]byte) {
	s.mu.Lock()
	s.buffer = b
	s.mu.Unlock()
}

func (s *Soundboard) ClearBuffer() {
	s.mu.Lock()
	s.buffer = nil
	s.mu.Unlock()
}

func (s *Soundboard) PlayBuffer(ses *discordgo.Session, guildID, channelID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Join the provided voice channel.
	vc, err := ses.ChannelVoiceJoin(guildID, channelID, false, true)
	if err != nil {
		return err
	}

	// Sleep for a specified amount of time before playing the sound
	time.Sleep(250 * time.Millisecond)

	// Start speaking.
	vc.Speaking(true)

	// Send the buffer data.
	for _, buff := range s.buffer {
		vc.OpusSend <- buff
	}

	// Stop speaking
	if err := vc.Speaking(false); err != nil {
		return err
	}

	// Sleep for a specificed amount of time before ending.
	time.Sleep(250 * time.Millisecond)

	// Disconnect from the provided voice channel.
	if err := vc.Disconnect(); err != nil {
		return err
	}

	return nil
}

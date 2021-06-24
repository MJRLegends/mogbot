package handlers

import "github.com/bwmarrin/discordgo"

func OnBibleVerse(s *discordgo.Session, m *discordgo.MessageCreate) {
	_ = `^\[(?P<book>(?:\d\s*)?[A-Z]?[a-z]+)\s*(?P<chapter>\d+):(?P<verses>(?P<start>\d+)(?:-(?P<end>\d+))?)(?:\s(?P<version>[A-Z]?[a-z]+))?\]$`
}

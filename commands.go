package main

import "github.com/bwmarrin/discordgo"

// Commands is a slice of all my commands
var Commands = []*discordgo.ApplicationCommand{
	{
		Name:        "ping",
		Description: "Pings the mogbot. Returns Pong with latency.",
	},
	{
		Name:        "ban",
		Description: "Bans a user",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "The user you want to ban",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "reason",
				Description: "The ban reason",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "days",
				Description: "The number of days of messages to delete",
				Required:    false,
			},
		},
	},
	{
		Name:        "join",
		Description: "Joins a voice channel and starts recording",
	},
	{
		Name:        "warn",
		Description: "Warns a user",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "The user you want to warn",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "reason",
				Description: "The warn reason",
				Required:    true,
			},
		},
	},
	{
		Name:        "close",
		Description: "Closes a modmail ticket",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "ticket",
				Description: "The ticket number you want to close",
				Required:    false,
			},
		},
	},
	{
		Name:        "reply",
		Description: "Replies to a modmail ticket",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "content",
				Description: "The reply content",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "ticket",
				Description: "The ticket number you want to close",
				Required:    false,
			},
		},
	},
	{
		Name:        "mute",
		Description: "Mutes a user",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "The user you want to mute",
			},
		},
	},
	{
		Name:        "move",
		Description: "Moves all users from one channel to another",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionChannel,
				Name:        "from",
				Description: "The channel you want to move people from",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionChannel,
				Name:        "to",
				Description: "The channel you want to move people to",
				Required:    true,
			},
		},
	},
	{
		Name:        "avatar",
		Description: "Returns the avatar of the user",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "The user whose avatar you want",
				Required:    false,
			},
		},
	},
	{
		Name:        "log",
		Description: "Creates, edits, or deletes a log setup",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        "default",
				Description: "Creates the default log setup",
				Required:    false,
			},
		},
	},
	{
		Name:        "embed",
		Description: "Creates, edits, or appends to an embed",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        "create",
				Required:    false,
				Description: "Creates an embed",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionChannel,
						Name:        "channel",
						Description: "The channel the embed will be posted in",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "title",
						Description: "The title of your embed",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "description",
						Description: "The description of your embed (One line only, use append subcommand to add more.)",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionInteger,
						Name:        "color",
						Description: "The color of your embed",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionUser,
						Name:        "author",
						Description: "The author of your embed",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionSubCommand,
						Name:        "footer",
						Description: "The footer of your embed",
						Required:    false,
						Options: []*discordgo.ApplicationCommandOption{
							{
								Type:        discordgo.ApplicationCommandOptionString,
								Name:        "text",
								Description: "The text of your embed footer",
								Required:    false,
							},
							{
								Type:        discordgo.ApplicationCommandOptionString,
								Name:        "icon",
								Description: "The icon URL of the footer",
								Required:    false,
							},
						},
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "thumbnail",
						Description: "The thumbnail of your embed",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionBoolean,
						Name:        "timestamp",
						Description: "The timestamp of your embed (if true, it will set the current time)",
						Required:    false,
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        "edit",
				Required:    false,
				Description: "Edits an embed",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionChannel,
						Name:        "channel",
						Description: "The channel the embed is in",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "messageid",
						Description: "The message ID of the embed",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "title",
						Description: "The new title of your embed",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "description",
						Description: "The new description of your embed",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionInteger,
						Name:        "color",
						Description: "The new color of your embed",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionUser,
						Name:        "author",
						Description: "The new author of your embed",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "footertext",
						Description: "The new footer text of your embed",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "footericon",
						Description: "The new icon URL of the footer",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "thumbnail",
						Description: "The new thumbnail of your embed",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionBoolean,
						Name:        "timestamp",
						Description: "Change the timestamp of your embed (if true, it will set the current time)",
						Required:    false,
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        "append",
				Required:    false,
				Description: "Appends a line to the embed description.",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionChannel,
						Name:        "channel",
						Description: "The channel the embed is in",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "messageid",
						Description: "The message ID of the embed",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "text",
						Description: "The text to be appended",
						Required:    true,
					},
				},
			},
		},
	},
}

package bot

import (
	"log"
	"strings"

	"github.com/ameijboom/chime/internal/bot/commands"
	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	session  *discordgo.Session
	guildID  string
	commands []*discordgo.ApplicationCommand
}

func (b *Bot) GetSession() *discordgo.Session {
	return b.session
}

func (b *Bot) GetGuildID() string {
	return b.guildID
}

func (b *Bot) Open() error {
	// Add Ready handler
	b.session.AddHandler(b.onReady)

	err := b.session.Open()
	if err != nil {
		return err
	}

	// Register commands after session is opened
	b.commands, err = commands.ProcessCommandRegistry(b.session, b.guildID)
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) onReady(s *discordgo.Session, r *discordgo.Ready) {
	log.Printf("Bot is ready! Logged in as: %s#%s", r.User.Username, r.User.Discriminator)
	log.Printf("Bot ID: %s", r.User.ID)
	log.Printf("Connected to %d guilds:", len(r.Guilds))

	for _, guild := range r.Guilds {
		// Fetch detailed guild info to get full name
		g, err := s.Guild(guild.ID)
		if err != nil {
			log.Printf("  - Guild ID: %s (Failed to fetch name: %v)", guild.ID, err)
			continue
		}

		log.Printf("  - %s (ID: %s)", g.Name, g.ID)

		// Get bot member info to check permissions
		member, err := s.GuildMember(guild.ID, s.State.User.ID)
		if err != nil {
			log.Printf("    Failed to fetch bot member: %v", err)
			continue
		}

		// Calculate permissions
		perms, err := s.State.UserChannelPermissions(s.State.User.ID, guild.ID)
		if err != nil {
			// Fallback: get permissions from roles
			perms = int64(0)
			for _, roleID := range member.Roles {
				for _, role := range g.Roles {
					if role.ID == roleID {
						perms |= role.Permissions
						break
					}
				}
			}
		}

		log.Printf("    Permissions: %d (0x%x)", perms, perms)
		logPermissions(perms)
	}
}

func logPermissions(perms int64) {
	permMap := map[string]int64{
		"Administrator":         0x0000000008,
		"Manage Guild":          0x0000000020,
		"Manage Channels":       0x0000000010,
		"Manage Events":         0x0000000002,
		"Create Events":         0x0100000000,
		"Send Messages":         0x0000000800,
		"Embed Links":           0x0000004000,
		"Attach Files":          0x0000008000,
		"Read Message History":  0x0000010000,
		"Mention Everyone":      0x0000020000,
		"Use External Emojis":   0x0000040000,
		"Add Reactions":         0x0000000040,
		"View Channel":          0x0000000400,
		"Manage Messages":       0x0000002000,
		"Manage Roles":          0x0000010000000,
		"Create Instant Invite": 0x0000000001,
	}

	var grantedPerms []string
	for name, perm := range permMap {
		if perms&perm != 0 {
			grantedPerms = append(grantedPerms, name)
		}
	}

	if len(grantedPerms) > 0 {
		log.Printf("    Granted permissions: %v", strings.Join(grantedPerms, ", "))
	}
}

func (b *Bot) Close() error {
	return b.session.Close()
}

func NewBot(token string, guildID string) (*Bot, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	return &Bot{
		session: session,
		guildID: guildID,
	}, nil
}

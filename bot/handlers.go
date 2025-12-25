package bot

import (
	"fmt"
	"log"
	"strings"

	"exitlag-bot/config"

	"github.com/bwmarrin/discordgo"
)

func MessageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "!setup") {
		// Check admin permissions here if needed
		_, err := s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
			Embeds:     []*discordgo.MessageEmbed{GetShopEmbed()},
			Components: GetShopButtons(),
		})
		if err != nil {
			log.Println("Error sending shop embed:", err)
		}
	}
}

func InteractionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	cfg := config.LoadConfig()

	switch i.Type {
	case discordgo.InteractionMessageComponent:
		customID := i.MessageComponentData().CustomID

		switch customID {
		case "create_ticket":
			handleCreateTicket(s, i, cfg)
		case "pay_1m":
			handlePaymentResponse(s, i, "1 Месяц", "398")
		case "pay_3m":
			handlePaymentResponse(s, i, "3 Месяца", "911")
		case "pay_12m":
			handlePaymentResponse(s, i, "12 Месяцев", "2622")
		case "close_ticket":
			s.ChannelDelete(i.ChannelID)
		}
	}
}

func handleCreateTicket(s *discordgo.Session, i *discordgo.InteractionCreate, cfg *config.Config) {
	guildID := i.GuildID
	user := i.Member.User

	// Create channel
	channelName := fmt.Sprintf("ticket-%s", user.Username)

	// Permission overwrites
	permissionOverwrites := []*discordgo.PermissionOverwrite{
		{
			ID:   guildID, // @everyone
			Type: discordgo.PermissionOverwriteTypeRole,
			Deny: discordgo.PermissionViewChannel,
		},
		{
			ID:    user.ID,
			Type:  discordgo.PermissionOverwriteTypeMember,
			Allow: discordgo.PermissionViewChannel | discordgo.PermissionSendMessages,
		},
		{
			ID:    s.State.User.ID, // Bot
			Type:  discordgo.PermissionOverwriteTypeMember,
			Allow: discordgo.PermissionViewChannel | discordgo.PermissionSendMessages | discordgo.PermissionManageChannels,
		},
	}

	// Create the channel
	ch, err := s.GuildChannelCreateComplex(guildID, discordgo.GuildChannelCreateData{
		Name:                 channelName,
		Type:                 discordgo.ChannelTypeGuildText,
		PermissionOverwrites: permissionOverwrites,
		ParentID:             "", // Add Category ID from config if needed
	})

	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Error creating ticket channel.",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		log.Println("Error creating channel:", err)
		return
	}

	// Respond to the interaction (Ephemeral)
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Тикет создан: <#%s>", ch.ID),
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})

	// Send message inside the new ticket
	_, err = s.ChannelMessageSendComplex(ch.ID, &discordgo.MessageSend{
		Content:    fmt.Sprintf("Привет, %s! Выберите срок подписки:", user.Mention()),
		Components: GetTicketButtons(),
	})
	if err != nil {
		log.Println("Error sending ticket welcome:", err)
	}
}

func handlePaymentResponse(s *discordgo.Session, i *discordgo.InteractionCreate, planName string, price string) {
	// Reply with the T-Bank link and amount
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("**Оплата тарифа: %s**\nСумма к оплате: **%s₽**\n\n1. Нажмите кнопку оплаты ниже.\n2. Введите сумму **%s**.\n", planName, price, price),
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label: "Оплатить",
							Style: discordgo.LinkButton,
							URL:   "https://funpay.com/go/ovldershop",
							Emoji: &discordgo.ComponentEmoji{Name: "💳"},
						},
					},
				},
			},
		},
	})
	if err != nil {
		log.Println("Error responding with payment info:", err)
	}
}

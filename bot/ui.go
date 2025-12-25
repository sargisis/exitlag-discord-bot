package bot

import "github.com/bwmarrin/discordgo"

func GetShopEmbed() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       "ExitLag Store",
		Description: "**🚀 Уменьши пинг и убери лаги с ExitLag!**\n\n🔹 **Игровой VPN №1**\n🔹 **Оптимизация маршрутов**\n🔹 **Повышение FPS**\n\n**🛒 Тарифы:**\n🕐 **1 Месяц** — 349₽\n🕒 **3 Месяца** — 799₽\n📅 **12 Месяцев** — 2299₽\n\nНажмите кнопку ниже, чтобы оформить подписку.",
		Color:       0xFF0000, // Red
		Footer: &discordgo.MessageEmbedFooter{
			Text: "ExitLag Store | Low Ping High FPS",
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://yt3.googleusercontent.com/ytc/AIdro_k4XX_Wv3u5v2gV3y2xYe3q2w8v8z8z8z8z8z8=s900-c-k-c0x00ffffff-no-rj", // ExitLag Logo or similar
		},
	}
}

func GetShopButtons() []discordgo.MessageComponent {
	return []discordgo.MessageComponent{
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Label:    "Купить / Создать тикет",
					Style:    discordgo.PrimaryButton,
					CustomID: "create_ticket",
					Emoji:    &discordgo.ComponentEmoji{Name: "🎫"},
				},
			},
		},
	}
}

// Buttons inside the ticket for choosing duration
func GetTicketButtons() []discordgo.MessageComponent {
	return []discordgo.MessageComponent{
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Label:    "1 Месяц (398₽)",
					Style:    discordgo.SecondaryButton,
					CustomID: "pay_1m",
				},
				discordgo.Button{
					Label:    "3 Месяца (911₽)",
					Style:    discordgo.SecondaryButton,
					CustomID: "pay_3m",
				},
				discordgo.Button{
					Label:    "12 Месяцев (2622₽)",
					Style:    discordgo.SecondaryButton,
					CustomID: "pay_12m",
				},
			},
		},
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Label:    "Закрыть тикет",
					Style:    discordgo.DangerButton,
					CustomID: "close_ticket",
				},
			},
		},
	}
}

package handlers

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/gorcon/rcon"
	"github.com/johnnylin-a/discord-wol-extras/pkg/data"
)

var PowerOffWithRCONSafety = CommandHandler{
	Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// Check if skip-players-check is set
		options := i.ApplicationCommandData().Options
		skip := false
		for _, v := range options {
			if v.Name == "skip-players-check" {
				skip = v.BoolValue()
			}
		}

		// RCON and see online players
		rconNoReply := false

		if !skip {
			conn, err := rcon.Dial(data.ConfigDatabase.RCONHost+":"+data.ConfigDatabase.RCONPort, data.ConfigDatabase.RCONPassword)
			if err == nil {
				defer conn.Close()
				response, err := conn.Execute("ShowPlayers")
				if err == nil {
					if strings.Trim(response, "\n") != "name,playeruid,steamid" { // TODO MAYBE, customize this for non-PalWorld servers
						s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: "Not shutting down due to players being online in this server. Force shut down with `skip-players-check` set to `True`.",
							},
						})
						return
					}
				} else {
					rconNoReply = true
				}
			} else {
				rconNoReply = true
			}
			// Should still shutdown since RCON did not reply, assuming RCON service is down
		}

		content := "Attempting shutting down... "
		if rconNoReply {
			content = "RCON check failed, will shutdown anyway assuming RCON server is offline... "
		}

		// Sending shutdown to TRUENAS
		fullurl := ""
		if isHTTPS, _ := strconv.ParseBool(data.ConfigDatabase.TruenasIsHTTPS); isHTTPS {
			fullurl += "https://"
		} else {
			fullurl += "http://"
		}
		fullurl += data.ConfigDatabase.TruenasHost + "/api/v2.0/system/shutdown"
		nativeUrl, err := url.Parse(fullurl)
		if err != nil {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: content + "URL malformed!",
				},
			})
			return
		}

		insecureMode, _ := strconv.ParseBool(data.ConfigDatabase.TruenasIsMismatchCertForTLS)

		httpClient := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: insecureMode,
				},
			},
		}

		headers := http.Header{}
		headers.Add("Authorization", "Bearer "+data.ConfigDatabase.TruenasAPIKey)

		resp, err := httpClient.Do(&http.Request{
			URL:    nativeUrl,
			Header: headers,
			Method: http.MethodPost,
		})
		if err != nil {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: content + "Failed to send shutdown request!",
				},
			})
			return
		}

		if resp.StatusCode != http.StatusOK {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: content + "Failed shutdown!",
				},
			})
			return
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: content + "Successful shutdown.",
			},
		})
	},
	ApplicationCommand: &discordgo.ApplicationCommand{
		Name:        "poweroff",
		Description: "Power off while checking if there are no more online players",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Name:        "skip-players-check",
				Description: "Skip player check and shutdown anyway",
			},
		},
	},
}

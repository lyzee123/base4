package cmd

import (
  x "mywabot/system"

  "fmt"
  "regexp"
  "strings"
)

func init() {
  x.NewCmd(&x.ICmd{
    Name:   "(channelinfo|ci)",
    Cmd:    []string{"channelinfo"},
    Tags:   "info",
    Prefix: true,
    IsQuery: true,
    Exec: func(sock *x.Nc, m *x.IMsg) {
       m.React("⏱️")

    pattern := regexp.MustCompile(`https?://whatsapp.com/channel/`)
    if !pattern.MatchString(m.Query) {
      m.Reply("Url Invalid")
      return
    }

    key, err := sock.WA.GetNewsletterInfoWithInvite(strings.Split(m.Query, "/")[4])
    if err != nil {
      m.Reply("i don't know")
      return
    }

    m.Reply(fmt.Sprintf("*Channel Information*\n*Link:* %s\n*ID:* %s\n*Name:* %v\n*Followers:* %v\n\n*Description:* %v\n*Create At:* %v", m.Query, key.ID, key.ThreadMeta.Name.Text, key.ThreadMeta.SubscriberCount, key.ThreadMeta.Description.Text, key.ThreadMeta.CreationTime))

      m.React("✅")
    },
  })
}

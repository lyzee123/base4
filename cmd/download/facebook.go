package cmd

import (
  x "mywabot/system"

  "fmt"
  "strings"
)

func init() {
  x.NewCmd(&x.ICmd{
    Name:   "(fb|facebook|fbdl)",
    Cmd:    []string{"facebook"},
    Tags:   "download",
    Prefix: true,
    IsQuery: true,
    Exec: func(sock *x.Nc, m *x.IMsg) {
       m.React("⏱️")

    if !strings.Contains(m.Query, "facebook") {
        m.Reply("Itu bukan link facebook")
      return
    }

    resp, err := x.Fb(m.Query)
    if err != nil {
      fmt.Println("Error:", err)
      return
    }

    if resp != nil {
      sock.SendVideo(m.From, resp.HD, resp.Desc, *m)
    }

      m.React("✅")
    },
  })
}

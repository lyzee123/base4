package cmd

import (
  x "mywabot/system"

  "fmt"
)

func init() {
  x.NewCmd(&x.ICmd{
    Name:   "(short|shorturl)",
    Cmd:    []string{"shorturl"},
    Tags:   "tools",
    Prefix: true,
    IsQuery: true,
    Exec: func(sock *x.Nc, m *x.IMsg) {
       m.React("⏱️")

      res, err := x.ShortUrl(m.Query)
      if err != nil {
        fmt.Println("Error:", err)
        return
      }
         m.Reply(res)

      m.React("✅")
    },
  })
}
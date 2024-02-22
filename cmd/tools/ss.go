package cmd

import (
  x "mywabot/system"

  //"fmt"
  "net/url"
)

func init() {
  x.NewCmd(&x.ICmd{
    Name:   "(ss|screnshoot)",
    Cmd:    []string{"ss"},
    Tags:   "tools",
    Prefix: true,
    IsQuery: true,
    Exec: func(sock *x.Nc, m *x.IMsg) {
       m.React("⏱️")

      res := "https://api.apiflash.com/v1/urltoimage?access_key=185eff3aa9fe4e3c8e30bda63b1fb9cf&wait_until=page_loaded&url=" + url.QueryEscape(m.Query)
      bytes, err := x.ToByte(res)
      if err != nil {
        m.Reply(err.Error())
        return
      }
      sock.SendImage(m.From, bytes, "Screenshot Web", *m)

      m.React("✅")
    },
  })
}
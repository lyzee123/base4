package cmd

import (
  x "mywabot/system"

)

func init() {
  x.NewCmd(&x.ICmd{
    Name:   "(Bot|bot)",
    Cmd:    []string{"bot"},
    Tags:   "main",
    Prefix: false,
    Exec: func(sock *x.Nc, m *x.IMsg) {

      m.Reply("y")
    },
  })
}

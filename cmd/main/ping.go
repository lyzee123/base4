package cmd

import (
  x "mywabot/system"

  "fmt"
  "time"
)

func init() {
  x.NewCmd(&x.ICmd{
    Name:   "ping",
    Cmd:    []string{"ping"},
    Tags:   "main",
    Prefix: true,
    Exec: func(sock *x.Nc, m *x.IMsg) {

      now := time.Now()
      mdate := time.Unix(m.Timestamp.Unix(), 0)
      mtime := now.Sub(mdate)
      ms := mtime.Seconds()
      txt := fmt.Sprintf("%.3f seconds", ms)
      m.Reply(txt)

    },
  })
}
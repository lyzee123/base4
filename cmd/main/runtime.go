package cmd

import (
  x "mywabot/system"

  "fmt"
  "time"
)

func init() {
  x.NewCmd(&x.ICmd{
    Name:   "runtime",
    Cmd:    []string{"runtime"},
    Tags:   "main",
    Prefix: true,
    Exec: func(sock *x.Nc, m *x.IMsg) {

      elapsed := time.Since(x.GetUptime())
      days := int(elapsed.Hours()) / 24
      hours := int(elapsed.Hours()) % 24
      minutes := int(elapsed.Minutes()) % 60
      seconds := int(elapsed.Seconds()) % 60
      m.Reply(fmt.Sprintf("Bot aktif selama: \n%02d Day(s) %02d Hour(s) %02d Minute(s) and %02d Seconds!", days, hours, minutes, seconds))

    },
  })
}

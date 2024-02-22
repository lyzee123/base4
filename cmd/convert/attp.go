package cmd

import (
  x "mywabot/system"

  "fmt"
  "net/url"
)

func init() {
  x.NewCmd(&x.ICmd{
    Name: "attp",
    Cmd:  []string{"attp"},
    Tags: "convert",
    Prefix: true,
    IsQuery: true,
    Exec: func(sock *x.Nc, m *x.IMsg) {
      m.React("⏱️")


      bytes, err := x.ToByte("https://aemt.me/attp?text="+url.QueryEscape(m.Query))
      if err != nil {
         fmt.Println("Error:", err)
        return
      }

  s := x.StickerApi(&x.Sticker{
    File: bytes,
    Tipe: func() x.MediaType {
    if m.IsImage || m.IsQuotedImage || m.IsQuotedSticker {
      return x.IMAGE
    } else if m.IsVideo || m.IsQuotedVideo {
      return x.VIDEO
    } else {
      return x.TEKS
    }
    }(),
  }, &x.MetadataSticker{
    Author:    m.PushName,
    Pack:      "https://s.id/ryuubot",
    KeepScale: true,
    Removebg:  "true",
    Circle: func() bool {
      if m.Query == "-c" {
        return true
      } else {
        return false
      }
    }(),
  })

  sock.SendSticker(m.From, s.Build(), *m)

        m.React("✅")
    },
  })
}


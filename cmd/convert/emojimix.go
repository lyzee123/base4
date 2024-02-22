package cmd

import (
  x "mywabot/system"

  "fmt"
  "strings"
  "net/url"
  "net/http"
  "encoding/json"
)

func init() {
  x.NewCmd(&x.ICmd{
    Name: "(emojimix|emoji)",
    Cmd:  []string{"emojimix"},
    Tags: "convert",
    Prefix: true,
    IsQuery: true,
    Exec: func(sock *x.Nc, m *x.IMsg) {
      m.React("⏱️")

    emojis := strings.Split(m.Query, "+")
if len(emojis) != 2 {
    m.Reply("Example: .emojimix😅+🤔")
    return
}

emoji1 := emojis[0]
emoji2 := emojis[1]

url := fmt.Sprintf("https://tenor.googleapis.com/v2/featured?key=AIzaSyAyimkuYQYF_FXVALexPuGQctUWRURdCYQ&contentfilter=high&media_filter=png_transparent&component=proactive&collection=emoji_kitchen_v5&q=%s_%s", url.QueryEscape(emoji1), url.QueryEscape(emoji2))

resp, err := http.Get(url)
if err != nil {
    fmt.Println("Error:", err)
    return
}
defer resp.Body.Close()

var data struct {
    Results []struct {
        URL string `json:"url"`
    } `json:"results"`
}

err = json.NewDecoder(resp.Body).Decode(&data)
if err != nil {
    fmt.Println("Error:", err)
    return
}

for _, res := range data.Results {
bytes, err := x.ToByte(res.URL)
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
}
        m.React("✅")
    },
  })
}


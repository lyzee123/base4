package cmd

import (
  x "mywabot/system"

  "fmt"
  "net/http"
  "encoding/json"
  "io/ioutil"
  "net/url"
)

func init() {
  x.NewCmd(&x.ICmd{
    Name:   "(ttsearch|tiktoksearch)",
    Cmd:    []string{"ttsearch"},
    Tags:   "search",
    Prefix: true,
    IsQuery: true,
    Exec: func(sock *x.Nc, m *x.IMsg) {
       m.React("⏱️")

      url := "https://skizo.tech/api/ttsearch?search=" + url.QueryEscape(m.Query) + "&apikey=batu"
      resp, err := http.Get(url)
      if err != nil {
        fmt.Println("Error:", err)
        return
      }
      defer resp.Body.Close()

      body, err := ioutil.ReadAll(resp.Body)
      if err != nil {
        fmt.Println("Error:", err)
        return
      }

      var res struct {
        Title      string `json:"title"`
        Region     string `json:"region"`
        Music      string `json:"music"`
        MusicInfo struct {
          Title string `json:"title"`
          Play   string `json:"play"`
          Author string `json:"author"`
        } `json:"music_info"`
        Play string `json:"play"`
      }

      err = json.Unmarshal(body, &res)
      if err != nil {
        fmt.Println("Error:", err)
        m.React("❌")
        return
      }

      caption := fmt.Sprintf(`*TIKTOK SEARCH*

*𖦹 Judul:* %s
*𖦹 Region:* %s
*𖦹 Musik:* %s
*- Musik Info:*
  *• Judul:* %s
  *• Link:* %s
  *• Author:* %s
`, res.Title, res.Region, res.Music, res.MusicInfo.Title, res.MusicInfo.Play, res.MusicInfo.Author)

      sock.SendVideo(m.From, res.Play, caption, *m)

      m.React("✅")
    },
  })
}

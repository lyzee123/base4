package cmd

import (
  x "mywabot/system"

  "fmt"
  "encoding/json"
  "regexp"
)

func init() {
  x.NewCmd(&x.ICmd{
    Name:   "(capcut|capcutdl|cpdl)",
    Cmd:    []string{"capcut"},
    Tags:   "download",
    Prefix: true,
    IsQuery: true,
    Exec: func(sock *x.Nc, m *x.IMsg) {
       m.React("⏱️")

      type Result struct {
        Code        int    `json:"code"`
        Title       string `json:"title"`
        Description string `json:"description"`
        Usage       string `json:"usage"`
        OriginalVideoUrl string `json:"originalVideoUrl"`
        CoverUrl    string `json:"coverUrl"`
        AuthorUrl   string `json:"authorUrl"`
      }

      regex := regexp.MustCompile(`(https?:\/\/[^\s]+)`)
       newLink := regex.FindStringSubmatch(m.Query) 
      
      data, err := x.Capcutdl(newLink[0])
      if err != nil {
        fmt.Println("Error:", err)
        return
      }

      var result Result
      err = json.Unmarshal(data, &result)
      if err != nil {
        fmt.Println(err)
      }

      teks := `*CAPCUT DOWNLOADER*

      𖦹 *Title:* ` + result.Title + `
      𖦹 *Description:* ` + result.Description + `
      𖦹 *Usage:* ` + result.Usage + `

      Reply/balas video ini dengan ketik *.toaudio* untuk menjadikan video ke audio`

      sock.SendVideo(m.From, "https://ssscap.net"+result.OriginalVideoUrl, teks, *m)
      m.React("✅")
    },
  })
}

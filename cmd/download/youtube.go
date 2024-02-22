package cmd

import (
  x "mywabot/system"

  "fmt"
 
  "regexp"
  "time"

  "github.com/Pauloo27/searchtube"
)

func init() {
  x.NewCmd(&x.ICmd{
    Name:   "(ytmp4|ytmp3|play)",
    Cmd:    []string{"ytmp4", "ytmp3", "play"},
    Tags:   "download",
    Prefix: true,
    IsQuery: true,
    Exec: func(sock *x.Nc, m *x.IMsg) {
       m.React("⏱️")

      var url string
        if x.IsYoutubeURL(m.Query) {
          url = m.Query
        } else {
          ser, _ := searchtube.Search(m.Query, 10)
          if len(ser) == 0 {
            m.Reply("Not Found")
            return
          }
          for _, v := range ser {
            d, _ := v.GetDuration()
            if d < 8*time.Minute {
              url = v.URL
              break
            }
          }
        }
        yt, err := x.YoutubeDL(url)
        if err != nil {
          m.Reply(err.Error())
          return
        }

        caption := fmt.Sprintf("*Title*: %s\n*Author*: %s", yt.Info.Title, yt.Info.Author)

        if reg, _ := regexp.MatchString(`(ytmp3|play)`, m.Text); reg {
          build, err := yt.Link.Audio[0].Url()
          if err != nil {
            m.Reply(err.Error())
            return
          }

           m.Reply(caption)
          sock.SendAudio(m.From, build, false, *m)
      
        } else {
          build, err := yt.Link.Video[0].Url()
          if err != nil {
            m.Reply(err.Error())
            return
          }
          
          sock.SendVideo(m.From, build, caption, *m)
        }

      m.React("✅")
    },
  })
}

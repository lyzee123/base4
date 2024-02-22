package cmd

import (
  x "mywabot/system"

  "encoding/json"
  "fmt"
  "net/http"
  "os"
   "io/ioutil"
)

func init() {
  x.NewCmd(&x.ICmd{
    Name:   "(carimusik|whatmusic|carimusic|carijudulmusik)",
    Cmd:    []string{"carijudulmusik"},
    Tags:   "search",
    Prefix: true,
    IsQuery: false,
    IsMedia: false,
    Exec: func(sock *x.Nc, m *x.IMsg) {
       m.React("⏱️")

      type Result struct {
        Result string `json:"result"`
      }

      // quoted video
      if m.IsQuotedVideo {
        conjp := "./tmp/" + m.ID + ".mp4"
        byte, _ := sock.WA.Download(m.Quoted.QuotedMessage.VideoMessage)
        err := os.WriteFile(conjp, byte, 0644)
        if err != nil {
          fmt.Println("Failed saved mp4")
          return
        }
        url, err := x.Upload(conjp)
          if err != nil {
             fmt.Println(err)
              return
          }
       
        os.Remove(conjp)

        type Result struct {
          Result string `json:"result"`
        }

        resp, err := http.Get("https://aemt.me/whatmusic?url="+url)
          if err != nil {
            fmt.Println("Error sending request:", err)
            return
          }
          defer resp.Body.Close()

          body, err := ioutil.ReadAll(resp.Body)
          if err != nil {
            fmt.Println("Error reading response body:", err)
            return
          }

          var result Result
          err = json.Unmarshal(body, &result)
        if err != nil {
          fmt.Println("Error:", err)
          return
        }
         m.Reply(result.Result)
        
        m.React("✅")
      } else {
      
      // from audio
      conjp := "./tmp/" + m.ID + ".mp3"
      byte, _ := sock.WA.Download(m.Quoted.QuotedMessage.AudioMessage)
      err := os.WriteFile(conjp, byte, 0644)
      if err != nil {
        fmt.Println("Failed saved mp3")
        return
      }
      url, err := x.Upload(conjp)
        if err != nil {
           fmt.Println(err)
            return
        }
      
      os.Remove(conjp)
      
      resp, err := http.Get("https://aemt.me/whatmusic?url="+url)
      if err != nil {
        fmt.Println("Error sending request:", err)
        return
      }
      defer resp.Body.Close()

      body, err := ioutil.ReadAll(resp.Body)
      if err != nil {
        fmt.Println("Error reading response body:", err)
        return
      }

      var result Result
      err = json.Unmarshal(body, &result)
    if err != nil {
      fmt.Println("Error:", err)
      return
    }
     m.Reply(result.Result)

      m.React("✅")
        }
    },
  })
}

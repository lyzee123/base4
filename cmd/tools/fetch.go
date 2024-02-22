package cmd

import (
  x "mywabot/system"

  "fmt"
  "net/http"
  "strings"
  "io/ioutil"
  "os"
  "bytes"
  "encoding/json"
  
)

func init() {
  x.NewCmd(&x.ICmd{
    Name:   "(get|fetch)",
    Cmd:    []string{"fetch"},
    Tags:   "tools",
    Prefix: true,
    IsQuery: true,
    Exec: func(sock *x.Nc, m *x.IMsg) {
       m.React("⏱️")

      resp, err := http.Get(m.Query)
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

        mime := resp.Header.Get("Content-Type")

        if mime == "" {
          m.Reply("No Content-Type")
        }
      fmt.Println(mime)

        if strings.Contains(mime, "video") {
          sock.SendVideo(m.From, body, m.Query, *m)   
        } else if strings.Contains(mime, "webp") {

          conjp := "./tmp/" + m.ID + ".jpg"
          conwp := "./tmp/" + m.ID + ".webp"
          err = os.WriteFile(conjp, body, 0644)
          if err != nil {
            fmt.Println("Failed saved jpg")
            return
          }
          x.ImgToWebp(conjp, conwp)
          sock.StickerPath(m.From, conwp, *m)
          os.Remove(conwp)
          os.Remove(conjp)
          
        } else if strings.Contains(mime, "image") {
          sock.SendImage(m.From, body, m.Query, *m)
        } else if strings.Contains(mime, "audio") {
          sock.SendAudio(m.From, body, false, *m)
        } else if strings.Contains(mime, "json") {
          formattedJSON := formatJSON(body)
          m.Reply(formattedJSON)
        } else if strings.Contains(mime, "text") {
          m.Reply(string(body))
        }

      m.React("✅")
    },
  })
}

func formatJSON(data []byte) string {
  var prettyJSON bytes.Buffer
  err := json.Indent(&prettyJSON, data, "", "  ")
  if err != nil {
    return "Error formatting JSON"
  }
  return prettyJSON.String()
}
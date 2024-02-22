package cmd

import (
  "fmt"
  x "mywabot/system"
  "os"
 // "os/exec"
  "io/ioutil"
)

func init() {
  x.NewCmd(&x.ICmd{
    Name: "toaudio",
    Cmd:  []string{"toaudio"},
    Tags: "convert",
    IsMedia: true,
    Prefix: true,
    Exec: func(sock *x.Nc, m *x.IMsg) {
      m.React("⏱️")

      

      // quoted video
      if m.IsQuotedVideo {
        mp3 := "./tmp/" + m.ID + ".mp3"
        byte, _ := sock.WA.Download(m.Media)
        whatsappAudio, err := x.ToAudio(byte, "mp4")
          if err != nil {
            fmt.Println("Error:", err)
            return
          }

          err = ioutil.WriteFile(mp3, whatsappAudio, 0644)
          if err != nil {
            fmt.Println("Error:", err)
            return
          }
        
        url, err := x.Upload(mp3)
        if err != nil {
            fmt.Println("Error:", err)
            return
        }

        sock.SendAudio(m.From, url, false, *m)
       
        os.Remove(mp3)
        m.React("✅")
      }

      // from video
      if m.IsVideo {
        mp3 := "./tmp/" + m.ID + ".mp3"
        byte, _ := sock.WA.Download(m.Media)
        whatsappAudio, err := x.ToAudio(byte, "mp4")
          if err != nil {
            fmt.Println("Error:", err)
            return
          }

          err = ioutil.WriteFile(mp3, whatsappAudio, 0644)
          if err != nil {
            fmt.Println("Error:", err)
            return
          }

        url, err := x.Upload(mp3)
        if err != nil {
            fmt.Println("Error:", err)
            return
        }

        sock.SendAudio(m.From, url, false, *m)

        os.Remove(mp3)
        m.React("✅")
      }

    },
  })
}

package cmd

import (
  "fmt"
  x "mywabot/system"
  "os"
  "os/exec"
  "bytes"
)

func init() {
  x.NewCmd(&x.ICmd{
    Name: "(toimage|toimg)",
    Cmd:  []string{"toimage"},
    Tags: "convert",
    IsMedia: true,
    Prefix: true,
    Exec: func(sock *x.Nc, m *x.IMsg) {
      m.React("⏱️")


      if m.IsQuotedSticker {
        conjp := "./tmp/" + m.ID + ".jpg"
        conwb := "./tmp/" + m.ID + ".webp"
        byte, _ := sock.WA.Download(m.Quoted.QuotedMessage.StickerMessage)
        err := os.WriteFile(conwb, byte, 0644)
        if err != nil {
          fmt.Println("Failed saved webp")
          return
        }

        cmd := exec.Command("ffmpeg", "-i", conwb, conjp)
        var out bytes.Buffer
        var stderr bytes.Buffer
        cmd.Stdout = &out
        cmd.Stderr = &stderr
        err = cmd.Run()

        // Check error
        if err != nil {
          fmt.Println("Error:", err)
          return
        }
        
        url, err := x.Upload(conjp)
        if err != nil {
            fmt.Println("Error:", err)
            return
        }
       
        sock.SendImage(m.From, url, "", *m)
        
        os.Remove(conwb)
        os.Remove(conjp)
        m.React("✅")
      }
    },
  })
}

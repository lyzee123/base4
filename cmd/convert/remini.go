package cmd

import (
  x "mywabot/system"

  "fmt"
  "os"
  "os/exec"
  "bytes"
)

func init() {
  x.NewCmd(&x.ICmd{
    Name: "(hd|remini|tohd)",
    Cmd:  []string{"remini"},
    Tags: "convert",
    IsMedia: true,
    Prefix: true,
    Exec: func(sock *x.Nc, m *x.IMsg) {
      m.React("⏱️")


          // quoted sticker
      if m.IsQuotedSticker {
        byte, _ := sock.WA.Download(m.Quoted.QuotedMessage.StickerMessage)
        res, err := x.Remini(byte, "enhance")
        if err != nil {
          fmt.Println("Failed saved jpg")
          return
        }

        conjp1 := "./tmp/" + m.ID + ".png"
        conjp2 := "./tmp/" + m.ID + ".jpg"
        err = os.WriteFile(conjp1, res, 0644)
        if err != nil {
          fmt.Println("Failed saved jpg")
          return
        }

        cmd := exec.Command("ffmpeg", "-i", conjp1, conjp2)
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

        url, err := x.Upload(conjp2)
        if err != nil {
            fmt.Println("Error:", err)
            return
        }

        sock.SendImage(m.From, url, "", *m)
        os.Remove(conjp1)
        os.Remove(conjp2)    
        m.React("✅")
      }

      
      // quoted image
      if m.IsQuotedImage {
        byte, _ := sock.WA.Download(m.Quoted.QuotedMessage.ImageMessage)
        res, err := x.Remini(byte, "enhance")
        if err != nil {
          fmt.Println("Failed saved jpg")
          return
        }

        conjp1 := "./tmp/" + m.ID + ".png"
        conjp2 := "./tmp/" + m.ID + ".jpg"
        err = os.WriteFile(conjp1, res, 0644)
        if err != nil {
          fmt.Println("Failed remini")
          return
        }

        cmd := exec.Command("ffmpeg", "-i", conjp1, conjp2)
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
        
        url, err := x.Upload(conjp2)
        if err != nil {
            fmt.Println("Error:", err)
            return
        }
        
        sock.SendImage(m.From, url, "", *m)
        os.Remove(conjp1)
        os.Remove(conjp2)
        m.React("✅")
      }


      // from image
      if m.IsImage {
        byte, _ := sock.WA.Download(m.Media)
        res, err := x.Remini(byte, "enhance")
        if err != nil {
          fmt.Println("Failed remini")
          return
        }

        conjp1 := "./tmp/" + m.ID + ".png"
        conjp2 := "./tmp/" + m.ID + ".jpg"
        err = os.WriteFile(conjp1, res, 0644)
        if err != nil {
          fmt.Println("Failed saved jpg")
          return
        }

        cmd := exec.Command("ffmpeg", "-i", conjp1, conjp2)
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

        url, err := x.Upload(conjp2)
        if err != nil {
            fmt.Println("Error:", err)
            return
        }

        sock.SendImage(m.From, url, "", *m)
        os.Remove(conjp1)
        os.Remove(conjp2)
        m.React("✅")
      }

    },
  })
}
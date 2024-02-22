package cmd

import (
  x "mywabot/system"

  "fmt"
  "os"
)

func init() {
  x.NewCmd(&x.ICmd{
    Name:   "(up|upload)",
    Cmd:    []string{"upload"},
    Tags:   "tools",
    Prefix: true,
    //IsQuery: true,
    Exec: func(sock *x.Nc, m *x.IMsg) {
       m.React("⏱️")

      // quoted sticker
      if m.IsQuotedSticker {
        conjp := "./tmp/" + m.ID + ".webp"
        byte, _ := sock.WA.Download(m.Quoted.QuotedMessage.StickerMessage)
        err := os.WriteFile(conjp, byte, 0644)
        if err != nil {
          fmt.Println("Failed saved webp")
          return
        }
        url, err := x.Upload(conjp)
          if err != nil {
             fmt.Println(err)
              return
          }
        m.Reply(url)
        os.Remove(conjp)
        m.React("✅")
      }

      // quoted image
      if m.IsQuotedImage {
        conjp := "./tmp/" + m.ID + ".jpg"
        byte, _ := sock.WA.Download(m.Quoted.QuotedMessage.ImageMessage)
        err := os.WriteFile(conjp, byte, 0644)
        if err != nil {
          fmt.Println("Failed saved jpg")
          return
        }
        url, err := x.Upload(conjp)
          if err != nil {
             fmt.Println(err)
              return
          }
        m.Reply(url)
        os.Remove(conjp)
        m.React("✅")
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
        m.Reply(url)
        os.Remove(conjp)
        m.React("✅")
      }

      // from video
      if m.IsVideo {
        conjp := "./tmp/" + m.ID + ".mp4"
        byte, _ := sock.WA.Download(m.Media)
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
        m.Reply(url)
        os.Remove(conjp)
        m.React("✅")
      }

      // from image
      if m.IsImage {
        conjp := "./tmp/" + m.ID + ".jpg"
        byte, _ := sock.WA.Download(m.Media)
        err := os.WriteFile(conjp, byte, 0644)
        if err != nil {
          fmt.Println("Failed saved jpg")
          return
        }
        url, err := x.Upload(conjp)
          if err != nil {
             fmt.Println(err)
              return
          }
        m.Reply(url)
        os.Remove(conjp)
        m.React("✅")
      } else {
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
        m.Reply(url)
        os.Remove(conjp)
        m.React("✅")
      }

      
    },
  })
}
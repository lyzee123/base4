package cmd

import (
  "fmt"
  x "mywabot/system"
  "os"
 // "os/exec"
)

func init() {
  x.NewCmd(&x.ICmd{
    Name: "(sm|smeme)",
    Cmd:  []string{"smeme"},
    Tags: "convert",
    IsMedia: true,
    Prefix: true,
    IsQuery: true,
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
            fmt.Println("Error:", err)
            return
        }
        res := "https://api.memegen.link/images/custom/-/" +m.Query+ ".webp?background="+url
        bytes, err := x.ToByte(res)
        if err != nil {
           fmt.Println("Error:", err)
          return
        }

        s := x.StickerApi(&x.Sticker{
          File: bytes,
          Tipe: func() x.MediaType {
            if m.IsImage || m.IsQuotedImage || m.IsQuotedSticker {
              return x.IMAGE
            } else {
              return x.VIDEO
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

        /*
        conwp := "./tmp/" + m.ID + ".webp"
        err = os.WriteFile(conwp, bytes, 0644)
        if err != nil {
          fmt.Println("Failed saved webp")
          return
        }
        
        createExif := fmt.Sprintf("webpmux -set exif %s %s -o %s", "tmp/exif/mywabot.exif", conwp, conwp)
        cmd := exec.Command("bash", "-c", createExif)
        err = cmd.Run()
        if err != nil {
          fmt.Println("Failed to set webp metadata", err)
        }
        sock.StickerPath(m.From, conwp, *m)

        os.Remove(conwp)
        */
        os.Remove(conjp)      
        m.React("✅")
      }

      
      // quoted image
      if m.IsQuotedImage {
        conjp1 := "./tmp/" + m.ID + ".jpg"
        byte, _ := sock.WA.Download(m.Quoted.QuotedMessage.ImageMessage)
        err := os.WriteFile(conjp1, byte, 0644)
        if err != nil {
          fmt.Println("Failed saved jpg")
          return
        }
        
        url, err := x.Upload(conjp1)
        if err != nil {
            fmt.Println("Error:", err)
            return
        }
        res := "https://api.memegen.link/images/custom/-/" +m.Query+ ".jpg?background="+url
        bytes, err := x.ToByte(res)
        if err != nil {
           fmt.Println("Error:", err)
          return
        }

        s := x.StickerApi(&x.Sticker{
          File: bytes,
          Tipe: func() x.MediaType {
            if m.IsImage || m.IsQuotedImage || m.IsQuotedSticker {
              return x.IMAGE
            } else {
              return x.VIDEO
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
        
/*
        conjp := "./tmp/" + m.ID + ".jpg"
        conwp := "./tmp/" + m.ID + ".webp"
        err = os.WriteFile(conjp, bytes, 0644)
        if err != nil {
          fmt.Println("Failed saved jpg")
          return
        }
        x.ImgToWebp(conjp, conwp)
        sock.StickerPath(m.From, conwp, *m)
        os.Remove(conwp)
        os.Remove(conjp)
        */
        os.Remove(conjp1)
        m.React("✅")
      }


      // from image
      if m.IsImage {
        conjp1 := "./tmp/" + m.ID + ".jpg"
        byte, _ := sock.WA.Download(m.Media)
        err := os.WriteFile(conjp1, byte, 0644)
        if err != nil {
          fmt.Println("Failed saved jpg")
          return
        }

        url, err := x.Upload(conjp1)
        if err != nil {
            fmt.Println("Error:", err)
            return
        }
        res := "https://api.memegen.link/images/custom/-/" +m.Query+ ".jpg?background="+url
        bytes, err := x.ToByte(res)
        if err != nil {
           fmt.Println("Error:", err)
          return
        }

        s := x.StickerApi(&x.Sticker{
          File: bytes,
          Tipe: func() x.MediaType {
            if m.IsImage || m.IsQuotedImage || m.IsQuotedSticker {
              return x.IMAGE
            } else {
              return x.VIDEO
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

        /*
        conjp := "./tmp/" + m.ID + ".jpg"
        conwp := "./tmp/" + m.ID + ".webp"
        err = os.WriteFile(conjp, bytes, 0644)
        if err != nil {
          fmt.Println("Failed saved jpg")
          return
        }
        x.ImgToWebp(conjp, conwp)
        sock.StickerPath(m.From, conwp, *m)
        os.Remove(conwp)
        os.Remove(conjp)
        */
        os.Remove(conjp1)

        m.React("✅")
      }

    },
  })
}

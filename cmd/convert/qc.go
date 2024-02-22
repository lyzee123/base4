package cmd

import (
  x "mywabot/system"

  //"fmt"
 // "os"
    "math/rand"
    "github.com/disintegration/imaging"
    "github.com/fogleman/gg"
    "net/http"
    "strings"
    "bytes"
    "encoding/json"
    "encoding/base64"
    "strconv"
  
)

func init() {
  x.NewCmd(&x.ICmd{
    Name: "qc",
    Cmd:  []string{"qc"},
    Tags: "convert",
    //IsMedia: true,
    Prefix: true,
    Exec: func(sock *x.Nc, m *x.IMsg) {
      m.React("⏱️")

      type From struct {
        ID   int    `json:"id"`
        Name string `json:"name"`
        Photo struct {
          URL string `json:"url"`
        } `json:"photo"`
      }

      type JsonResponse struct {
        Ok    bool   `json:"ok"`
        Result struct {
          Image string `json:"image"`
        } `json:"result"`
      }

      type Message struct {
        Entities []interface{} `json:"entities"`
        Media    struct {
          URL string `json:"url"`
        } `json:"media"`
        Avatar   bool      `json:"avatar"`
        From     From      `json:"from"`
        Text     string    `json:"text"`
        ReplyMessage interface{} `json:"replyMessage"`
      }

    
  
      teks := m.Query
      name := m.PushName
      avatar := "https://telegra.ph/file/89c1638d9620584e6e140.png"
      id := pickRandom([]int{0, 4, 5, 3, 2, 7, 5, 9, 8, 1, 6, 10, 9, 7, 5, 3, 1, 2, 4, 6, 8, 0, 10})
      
      
      jsonStr := `{
        "type": "quote",
        "format": "png",
        "backgroundColor": "#FFFFFF",
        "width": 512,
        "height": 768,
        "scale": 2,
        "messages": [
          {
            "entities": [],
            "avatar": true,
            "from": {
              "id": ` + strconv.Itoa(id) + `,
              "name": "` + name + `",
              "photo": {
                "url": "` + avatar + `"
              }
            },
            "text": "` + teks + `",
            "replyMessage": {}
          }
        ]
      }`

      var jsonData JsonResponse
      resp, err := http.Post("https://quoteapi-ld81.onrender.com/generate", "application/json", strings.NewReader(jsonStr))
      if err != nil {
        // handle error
      }
      defer resp.Body.Close()

      err = json.NewDecoder(resp.Body).Decode(&jsonData)
      if err != nil {
        // handle error
      }

      if !jsonData.Ok {
        // handle error
      }

      buffer, err := base64.StdEncoding.DecodeString(jsonData.Result.Image)
      if err != nil {
        // handle error
      }

      dc := gg.NewContext(512, 768)
      img, err := imaging.Decode(bytes.NewReader(buffer))
      if err != nil {
        // handle error
      }
      dc.DrawImage(img, 0, 0)

      s := x.StickerApi(&x.Sticker{
        File: buffer,
        Tipe: func() x.MediaType {
        if m.IsImage || m.IsQuotedImage || m.IsQuotedSticker {
          return x.IMAGE
        } else if m.IsVideo || m.IsQuotedVideo {
          return x.VIDEO
        } else {
          return x.TEKS
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
      err = os.WriteFile(conjp, buffer, 0644)
      if err != nil {
        fmt.Println("Failed saved jpg")
        return
      }
      x.ImgToWebp(conjp, conwp)
      sock.StickerPath(m.From, conwp, *m)
      os.Remove(conwp)
      os.Remove(conjp
      */

        m.React("✅")
        
    },
  })
}

func pickRandom(arr []int) int {
  return arr[rand.Intn(len(arr))]
}
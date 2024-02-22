package cmd

import (
  x "mywabot/system"

  "net/http"
  "encoding/json"
  "net/url"
  "fmt"
  "strings"

  //"io/ioutil"
//  "os"
  //  "os/exec"
  //"bytes"
)

func init() {
  x.NewCmd(&x.ICmd{
    Name:   "(ig|instagram|igdown|igdl)",
    Cmd:    []string{"instagram"},
    Tags:   "download",
    Prefix: true,
    IsQuery: true,
    Exec: func(sock *x.Nc, m *x.IMsg) {
       m.React("⏱️")

    if !strings.Contains(m.Query, "instagram") {
        m.Reply("Itu bukan link instagram")
      return
    }


      type Result struct {
        WM       string `json:"wm"`
        Thumbnail string `json:"thumbnail"`
        URL       string `json:"url"`
      }

      type Response struct {
        Status  bool     `json:"status"`
        Code    int      `json:"code"`
        Creator string   `json:"creator"`
        Result  []Result `json:"result"`
      }

      url := "https://aemt.me/download/igdl?url="+url.QueryEscape(m.Query)
      resp, err := http.Get(url)
      if err != nil {
        fmt.Println("Error:", err)
        return
      }
      defer resp.Body.Close()

      var response Response
      err = json.NewDecoder(resp.Body).Decode(&response)
      if err != nil {
        fmt.Println("Error:", err)
        return
      }

      if response.Status {
        for _, result := range response.Result {
          
          resp, err := http.Get(result.URL)
          if err != nil {
            fmt.Println("Error:", err)
            return
          }

            defer resp.Body.Close()

            mime := resp.Header.Get("Content-Type")

            if mime == "" {
              m.Reply("No Content-Type")
            }

            if strings.Contains(mime, "video") {
              sock.SendVideo(m.From, result.URL, "", *m)
            } else if strings.Contains(mime, "application/octet-stream") {
               sock.SendVideo(m.From, result.URL, "", *m)
              } else {
               sock.SendImage(m.From, result.URL, "", *m)
            }
          
        }
      } else {
        m.Reply(string(response.Code))
      }
/*
      //screp1
      type Response struct {
        Status  int `json:"status"`
        Creator string `json:"creator"`
        Result struct {
          Status int `json:"status"`
          Media []string `json:"media"`
        } `json:"result"`
      }

      respp, err := http.Get("https://api.arifzyn.tech/download/instagram?url="+url.QueryEscape(m.Query)+"&apikey=Danukiding")
      if err != nil {
        fmt.Println("Error:", err)
        return
      }
      defer respp.Body.Close()

      var response Response
      err = json.NewDecoder(respp.Body).Decode(&response)
      if err != nil {
        fmt.Println("Error:", err)
        return
      }

      for _, media := range response.Result.Media {
        
        resp, err := http.Get(media)
        if err != nil {
          fmt.Println("Error:", err)
          return
        }

          defer resp.Body.Close()

          mime := resp.Header.Get("Content-Type")

        fmt.Println(mime)
          if mime == "" {
            m.Reply("No Content-Type")
          }

          if strings.Contains(mime, "video") {
            //fmt.Println("video")
           sock.SendVideo(m.From, media, "", *m)
          } else if strings.Contains(mime, "image") {
            //fmt.Println("image")
             sock.SendImage(m.From, media, "", *m)
            } else {

        conjp := "./tmp/" + m.ID + ".mp4"
        conwb := "./tmp/" + m.ID + ".mkv"
       // byte, _ := sock.WA.Download(m.Quoted.QuotedMessage.StickerMessage)
          resp, err := http.Get(media)
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
            
        err = os.WriteFile(conwb, body, 0644)
        if err != nil {
          fmt.Println("Failed saved mkv")
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
            
        os.Remove(conwb)
        os.Remove(conjp)
            
        sock.SendVideo(m.From, url, "", *m)
          }
      }

      
      //screp2
      result, err := x.Instagram(m.Query)
      if err != nil {
        fmt.Println("Error:", err)
        return
      }

      fmt.Println(result)

      types := []string{}
      image := []string{}
      urls := []string{}

      for _, value := range result {
          typess := value["type"]
          types = append(types, typess)

          if typess == "video" {
              url := value["url"]
              urls = append(urls, url)
            
            sock.SendVideo(m.From, urls[0], "", *m)
            
          } else if typess == "image" {
              img := value["url"]
              image = append(image, img)
                     
            sock.SendImage(m.From, image[0], "", *m)
          }
      }


      //screp3
      resp, err := http.Get("https://skizo.tech/api/igdl?url="+url.QueryEscape(m.Query)+"&apikey=batu")

      if strings.Contains(m.Query, "https://www.instagram.com/reel/") {
      type respon struct {
        Caption string   `json:"caption"`
        Media   []string `json:"media"`
      }
      if err != nil {
          fmt.Println("Error:", err)
          return
        }
        defer resp.Body.Close()
        var data respon
        err = json.NewDecoder(resp.Body).Decode(&data)
        if err != nil {
          fmt.Println("Error:", err)
          return
        }
        
        // Mengambil media
          caption := data.Caption
        media := data.Media
        for _, url := range media {  
          sock.SendVideo(m.From, url, caption, *m)
        }

       } else if strings.Contains(m.Query, "https://www.instagram.com/p/") {
        type respon struct {
        Caption string   `json:"caption"`
        Media   []string `json:"media"`
      }
      if err != nil {
          fmt.Println("Error:", err)
          return
        }
        defer resp.Body.Close()
        var data respon
        err = json.NewDecoder(resp.Body).Decode(&data)
        if err != nil {
          m.Reply(err.Error())
          return
        }
        // Mengambil media
          caption := data.Caption
        media := data.Media
        for _, ur := range media {

          resp, err := http.Get(ur)
          if err != nil {
            fmt.Println("Error:", err)
            return
          }

            defer resp.Body.Close()
          
            mime := resp.Header.Get("Content-Type")

            if mime == "" {
              m.Reply("No Content-Type")
            }

            if strings.Contains(mime, "video") {
              sock.SendVideo(m.From, ur, caption, *m)
            } else {
               sock.SendImage(m.From, ur, caption, *m)
            }
        }   

      } else if strings.Contains(m.Query, "https://www.instagram.com/stories/") {
        type respon struct {
          Caption string   `json:"caption"`
          Media   []string `json:"media"`
        }
        if err != nil {
            fmt.Println("Error:", err)
            return
          }
          defer resp.Body.Close()
          var data respon
          err = json.NewDecoder(resp.Body).Decode(&data)
          if err != nil {
            m.Reply(err.Error())
            return
          }
          // Mengambil media
            caption := data.Caption
          media := data.Media
          for _, ur := range media {

            resp, err := http.Get(ur)
              if err != nil {
                fmt.Println("Error:", err)
                return
              }

                defer resp.Body.Close()

                mime := resp.Header.Get("Content-Type")

                if mime == "" {
                  m.Reply("No Content-Type")
                }

                if strings.Contains(mime, "video") {
                  sock.SendVideo(m.From, ur, caption, *m)
                } else {
                   sock.SendImage(m.From, ur, caption, *m)
                }
            }   
          }   
          */

      m.React("✅")
    },
  })
}

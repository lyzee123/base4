package cmd

import (
  x "mywabot/system"
    "fmt"
    "net/http"
    "net/url"
    "time"
     "encoding/json"
    "io/ioutil"
    "strconv"
    "strings"
    //"os"
    "regexp"
)

func init() {
  x.NewCmd(&x.ICmd{
    Name:   "(tt|ttnowm|tiktoknowm|tiktok)",
    Cmd:    []string{"tiktok"},
    Tags:   "download",
    Prefix: true,
    IsQuery: true,
    Exec: func(sock *x.Nc, m *x.IMsg) {
       m.React("⏱️")

      if !strings.Contains(m.Query, "tiktok") {
        m.Reply("Itu bukan link tiktok")
      return
      }  


      type TikTokData struct {
      Creator       string `json:"creator"`
      Code          int    `json:"code"`
      Msg           string `json:"msg"`
      ProcessedTime float64 `json:"processed_time"`
      Data          struct {
        ID              string `json:"id"`
        Region          string `json:"region"`
        Title           string `json:"title"`
        Cover           string `json:"cover"`
        OriginCover     string `json:"origin_cover"`
        Duration        int    `json:"duration"`
        Play            string `json:"play"`
        WmPlay          string `json:"wmplay"`
        HdPlay          string `json:"hdplay"`
        Size            int    `json:"size"`
        WmSize          int    `json:"wm_size"`
        HdSize          int    `json:"hd_size"`
        Music           string `json:"music"`
        MusicInfo       struct {
          ID       string `json:"id"`
          Title    string `json:"title"`
          Play     string `json:"play"`
          Cover    string `json:"cover"`
          Author   string `json:"author"`
          Original bool   `json:"original"`
          Duration int    `json:"duration"`
          Album    string `json:"album"`
        } `json:"music_info"`
        PlayCount     int `json:"play_count"`
        DiggCount     int `json:"digg_count"`
        CommentCount  int `json:"comment_count"`
        ShareCount    int `json:"share_count"`
        DownloadCount int `json:"download_count"`
        CollectCount  int `json:"collect_count"`
        CreateTime    int `json:"create_time"`

        Author              struct {
          ID        string `json:"id"`
          UniqueID  string `json:"unique_id"`
          Nickname  string `json:"nickname"`
          Avatar    string `json:"avatar"`
        } `json:"author"`
        Images    []string `json:"images"`
      } `json:"data"`
      }

      regex := regexp.MustCompile(`(https?:\/\/[^\s]+)`)
      newLink := regex.FindStringSubmatch(m.Query) 

      url := "https://skizo.tech/api/tiktok?url="+url.QueryEscape(newLink[0])+"&apikey=batu"//+os.Getenv("KEY")

      response, err := http.Get(url)
      if err != nil {
      fmt.Println("Error:", err)
      return
      }
      defer response.Body.Close()


      body, err := ioutil.ReadAll(response.Body)
      if err != nil {
      fmt.Println("Error:", err)
      return
      }


      var tiktokData TikTokData
      err = json.Unmarshal(body, &tiktokData)
      if err != nil {
      fmt.Println("Error:", err)
      return
      }


      if tiktokData.Data.Duration == 0 {
      for _, i := range tiktokData.Data.Images {
        x.Sleep(2 * time.Second)

        bytes, err := x.ToByte(i)
        if err != nil {
          m.Reply(err.Error())
          return
        }
        
        sock.SendImage(m.From, bytes, "", *m) 
      }

      } else { 

        teks := `*TIKTOK NO WATERMARK*

𖦹 *ID:* ` + tiktokData.Data.ID + `
𖦹 *Author:* ` + tiktokData.Data.Author.UniqueID + `
𖦹 *Region:* ` + tiktokData.Data.Region + `
𖦹 *Judul:* ` + tiktokData.Data.Title + `
𖦹 *Durasi:* ` + strconv.Itoa(tiktokData.Data.Duration) + `
𖦹 *Music:* ` + tiktokData.Data.Music + `
𖦹 *Info Musik:*
  - *Judul:* ` + tiktokData.Data.MusicInfo.Title + `
  - *Author:* ` + tiktokData.Data.MusicInfo.Author + `
𖦹 *Jumlah Komentar:* ` + strconv.Itoa(tiktokData.Data.CommentCount) + `
𖦹 *Jumlah Share:* ` + strconv.Itoa(tiktokData.Data.ShareCount) + `
𖦹 *Didownload:* ` + strconv.Itoa(tiktokData.Data.DownloadCount) + ` kali`

      bytes, err := x.ToByte(tiktokData.Data.Play)
      if err != nil {
      m.Reply(err.Error())
      return
      }
      sock.SendVideo(m.From, bytes, teks, *m)
      }
      
      m.React("✅")
    },
  })
}

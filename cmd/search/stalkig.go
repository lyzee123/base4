package cmd

import (
  x "mywabot/system"

  "fmt"
  "net/http"
   "encoding/json"
  "io/ioutil"
  "strconv"
)

func init() {
  x.NewCmd(&x.ICmd{
    Name:   "(igstalk|stalkig)",
    Cmd:    []string{"igstalk"},
    Tags:   "search",
    Prefix: true,
    IsQuery: true,
    Exec: func(sock *x.Nc, m *x.IMsg) {
       m.React("⏱️")

      type User struct {
        PhotoProfile string `json:"photo_profile"`
        Username     string `json:"username"`
        Fullname     string `json:"fullname"`
        Posts        int    `json:"posts"`
        Followers    int    `json:"followers"`
        Following    int    `json:"following"`
        Bio          string `json:"bio"`
      }

      resp, err := http.Get("https://skizo.tech/api/igstalk?user="+ m.Query +"&apikey=batu")
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

      var user User
      err = json.Unmarshal(body, &user)
      if err != nil {
        fmt.Println("Error:", err)
        return
      }

      teks := `*INSTAGRAM STALK*

𖦹 *Username:* ` + user.Username + `
𖦹 *Fullname:* ` + user.Fullname + `
𖦹 *Posts:* ` + strconv.Itoa(user.Posts) + `
𖦹 *Followers:* ` + strconv.Itoa(user.Followers) + `
𖦹 *Following:* ` + strconv.Itoa(user.Following) + `
𖦹 *Bio:* ` + user.Bio + `
      `

      sock.SendImage(m.From, user.PhotoProfile, teks, *m)

      m.React("✅")
    },
  })
}

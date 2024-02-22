package cmd

import (
  x "mywabot/system"

  "fmt"
  "encoding/json"
  "net/http"
  "net/url"
)

func init() {
  x.NewCmd(&x.ICmd{
    Name:   "(lirik|searchlirik)",
    Cmd:    []string{"lirik"},
    Tags:   "search",
    Prefix: true,
    IsQuery: true,
    Exec: func(sock *x.Nc, m *x.IMsg) {
       m.React("⏱️")


      type Response struct {
        Status  int `json:"status"`
        Creator string `json:"creator"`
        Result struct {
          URL   string `json:"url"`
          Title string `json:"title"`
          Thumb string `json:"thumb"`
          Lyrics string `json:"lyrics"`
        } `json:"result"`
      }

      resp, err := http.Get("https://api.arifzyn.tech/search/musixmatch?query="+url.QueryEscape(m.Query)+"&apikey=Danukiding")
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

     
     m.Reply(response.Result.Lyrics)

      m.React("✅")
    },
  })
}

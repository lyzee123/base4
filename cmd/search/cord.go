package cmd

import (
  x "mywabot/system"

  "encoding/json"
  "fmt"
  "net/http"
  "net/url"
)

func init() {
  x.NewCmd(&x.ICmd{
    Name:   "(cord|kuncilagu|kuncigitar)",
    Cmd:    []string{"cord"},
    Tags:   "search",
    Prefix: true,
    IsQuery: true,
    Exec: func(sock *x.Nc, m *x.IMsg) {
       m.React("⏱️")



    type Response struct {
      Status  bool   `json:"status"`
      Creator string `json:"creator"`
      Result  struct {
        Chord string `json:"chord"`
      } `json:"result"`
    }

      resp, err := http.Get("https://aemt.me/chord?query="+url.QueryEscape(m.Query))
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


     m.Reply(response.Result.Chord)

      m.React("✅")
    },
  })
}

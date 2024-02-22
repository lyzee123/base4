package cmd

import (
  x "mywabot/system"
  
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "net/url"
)

func init() {
  x.NewCmd(&x.ICmd{
    Name:   "bard",
    Cmd:    []string{"bard"},
    Tags:   "ai",
    IsQuery:  true,
     Prefix: true,
    ValueQ: ".bard siap kamu?",
    Exec: func(sock *x.Nc, m *x.IMsg) {
      m.React("⏱️")


      type Response struct {
        Status  bool   `json:"status"`
        Creator string `json:"creator"`
        Result  string `json:"result"`
      }
      
      response, err := http.Get("https://aemt.me/bard?text=" + url.QueryEscape(m.Query))
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

      var result Response
      err = json.Unmarshal(body, &result)
      if err != nil {
        fmt.Println("Error:", err)
        return
      }
       m.Reply(result.Result)     
      m.React("✅")
    },
  })
}

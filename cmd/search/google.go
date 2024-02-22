package cmd

import (
  x "mywabot/system"

  "encoding/json"
  "fmt"
  "io/ioutil"
  "math/rand"
  "net/http"
  "net/url"
  "time"
)

func init() {
  x.NewCmd(&x.ICmd{
    Name:   "(google|googleimage)",
    Cmd:    []string{"google"},
    Tags:   "search",
    Prefix: true,
    IsQuery: true,
    Exec: func(sock *x.Nc, m *x.IMsg) {
       m.React("⏱️")

    type Result struct {
      Creator string   `json:"creator"`
      Result  []string `json:"result"`
    }

    url := "https://aemt.me/googleimage?query="+url.QueryEscape(m.Query)
    resp, err := http.Get(url)
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

    var result Result
    err = json.Unmarshal(body, &result)
    if err != nil {
      fmt.Println("Error:", err)
      return
    }

    rand.Seed(time.Now().UnixNano())
    randomIndex := rand.Intn(len(result.Result))
    //fmt.Println("Random Result:", result.Result[randomIndex])

      sock.SendImage(m.From, result.Result[randomIndex], "", *m)
     
      m.React("✅")
    },
  })
}

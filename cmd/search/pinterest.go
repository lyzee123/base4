package cmd

import (
  x "mywabot/system"

  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "math/rand"
  "time"
  "net/url"
)

func init() {
  x.NewCmd(&x.ICmd{
    Name:   "(pinterest|pin)",
    Cmd:    []string{"pinterest"},
    Tags:   "search",
    Prefix: true,
    IsQuery: true,
    Exec: func(sock *x.Nc, m *x.IMsg) {
       m.React("⏱️")

      apiUrl := "https://skizo.tech/api/pinterest?search="+url.QueryEscape(m.Query)+"&apikey=batu"

        res, err := fetchJson(apiUrl)
        if err != nil {
         fmt.Println("Error:", err)
           m.React("❌")
          return
        }

        var urls []string
        for _, result := range res {
          urls = append(urls, result.Media.URL)
        }

        randomUrl := pickRandom(urls)
        fmt.Println("Random URL:", randomUrl)
      
      

      sock.SendImage(m.From, randomUrl, "", *m)

      m.React("✅")
    },
  })
}

type Media struct {
  URL string `json:"url"`
}

type Result struct {
  Media Media `json:"media"`
}

func fetchJson(url string) ([]Result, error) {
  resp, err := http.Get(url)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  var results []Result
  err = json.Unmarshal(body, &results)
  if err != nil {
    return nil, err
  }

  return results, nil
}

func pickRandom(urlList []string) string {
  rand.Seed(time.Now().Unix())
  return urlList[rand.Intn(len(urlList))]
}
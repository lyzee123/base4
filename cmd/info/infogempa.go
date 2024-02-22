package cmd

import (
  x "mywabot/system"

  "fmt"
  "encoding/json"
  "io/ioutil"
  "net/http"
)

func init() {
  x.NewCmd(&x.ICmd{
    Name:   "(gempa|infogempa)",
    Cmd:    []string{"infogempa"},
    Tags:   "info",
    Prefix: true,
    Exec: func(sock *x.Nc, m *x.IMsg) {
       m.React("⏱️")

      type Data struct {
        Infogempa struct {
          Gempa struct {
            Tanggal     string `json:"Tanggal"`
            Jam         string `json:"Jam"`
            Coordinates string `json:"Coordinates"`
            Magnitude   string `json:"Magnitude"`
            Kedalaman   string `json:"Kedalaman"`
            Wilayah     string `json:"Wilayah"`
            Potensi     string `json:"Potensi"`
            Dirasakan   string `json:"Dirasakan"`
            Shakemap    string `json:"Shakemap"`
          } `json:"gempa"`
        } `json:"Infogempa"`
      }

      resp, err := http.Get("https://data.bmkg.go.id/DataMKG/TEWS/autogempa.json")
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

      var res Data
      err = json.Unmarshal(body, &res)
      if err != nil {
        fmt.Println("Error:", err)
        return
      }

      teks := `┌──「 *G E M P A* 」
│
├ *TimeStamp:* ` + res.Infogempa.Gempa.Tanggal + `
├ *Time:* ` + res.Infogempa.Gempa.Jam + `
├ *Coordinates:* ` + res.Infogempa.Gempa.Coordinates + `
├ *Magnitude:* ` + res.Infogempa.Gempa.Magnitude + `
├ *Depth:* ` + res.Infogempa.Gempa.Kedalaman + `
├ *Region:* ` + res.Infogempa.Gempa.Wilayah + `
├ *Potention:* ` + res.Infogempa.Gempa.Potensi + `
├ *Effect:* ` + res.Infogempa.Gempa.Dirasakan + `
│
└──「 *Seputar BMKG* 」`

      url := "https://data.bmkg.go.id/DataMKG/TEWS/"+res.Infogempa.Gempa.Shakemap
      
      sock.SendImage(m.From, url, teks, *m)  
      
      m.React("✅")
    },
  })
}

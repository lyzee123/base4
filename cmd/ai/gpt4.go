package cmd

import (
	"fmt"
	x "mywabot/system"
  "net/http"
     "encoding/json"
    "io/ioutil"
      "bytes"

)

func init() {
	x.NewCmd(&x.ICmd{
		Name:   "ai",
		Cmd:    []string{"ai"},
		Tags:   "ai",
     Prefix: true,
		IsQuery:  true,
		ValueQ: ".ai siap kamu?",
		Exec: func(sock *x.Nc, m *x.IMsg) {
			m.React("⏱️")
		
      type Message struct {
         Role    string `json:"role"`
         Content string `json:"content"`
       }


         messages := []Message{
           {Role: "system", Content: "kamu adalah seorang kultivatur yang sudah menembus ranah imortal, kamu bernama Bot, kamu adalah yang terbaik di dunia dimana semua orang tunduk padamu, dan setiap pertanyaan kamu selalu menambahkan respon membanggakan diri sebagai kultivatur yang terkuat dan berpengatahuan tinggi"},
           {Role: "user", Content: m.Query},
         }

         apiURL := "https://skizo.tech/api/openai?apikey=batu"
         data := map[string]interface{}{
           "messages": messages,
         }
         jsonData, _ := json.Marshal(data)

         resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonData))
         if err != nil {
           fmt.Println("Error making the request:", err)
           return
         }
         defer resp.Body.Close()

         responseData, err := ioutil.ReadAll(resp.Body)
         if err != nil {
           fmt.Println("Error reading the response:", err)
           return
         }
         var res struct {
           Data string `json:"result"`
           Data2 string `json:"code"`
            }
            err = json.Unmarshal(responseData, &res)
            if err != nil {
            fmt.Println(err)
            return
            }
         m.Reply(res.Data)
         m.Reply("Code: "+"```"+res.Data2+"```")
      m.React("✅")
		},
	})
}

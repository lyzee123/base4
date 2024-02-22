package cmd

import (
  x "mywabot/system"
  
  "fmt"
  "io/ioutil"
  "os"
  "os/exec"
  "syscall"
  "time"
  "bytes"
 
)

func init() {
  x.NewCmd(&x.ICmd{
    Name:    `backup`,
    Cmd:     []string{"backup"},
    Tags:    "owner",
    Prefix:  true,
    IsOwner: true,
    Exec: func(client *x.Nc, m *x.IMsg) {
      m.React("⏱️")

      files, err := ioutil.ReadDir(".")
      if err != nil {
         fmt.Println(err)
      }

      var filteredFiles []string
      for _, file := range files {
        if file.Name() != "main" && file.Name() != ".cache" && file.Name() != ".git" {
          filteredFiles = append(filteredFiles, file.Name())
        }
      }

      zipCmd := exec.Command("zip", append([]string{"-r", "backup.zip"}, filteredFiles...)...)
      var out bytes.Buffer
      zipCmd.Stdout = &out
      err = zipCmd.Run()
      if err != nil {
        if exitErr, ok := err.(*exec.ExitError); ok {
          if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
            fmt.Println(status.ExitStatus())
          }
        }
         fmt.Println(err)
      }

      time.Sleep(3 * time.Second)

      file, err := os.Open("backup.zip")
      if err != nil {
         fmt.Println(err)
      }
      defer file.Close()

      bytes, err := ioutil.ReadFile("backup.zip")
      if err != nil {
        fmt.Println("Error reading file:", err)
        return
      }
      client.SendDocument(m.From, bytes, fmt.Sprintf("%s.zip", "backup"), " ", *m)

      err = os.Remove("backup.zip")
      if err != nil {
       fmt.Println(err)
      }      
      
      m.React("✅")
      
    },
  })
}

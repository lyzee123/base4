package cmd

import (
  x "mywabot/system"

  "fmt"
  "os/exec"
  "os"
  "syscall"
)

func init(){
  x.NewCmd(&x.ICmd{
    Name: "(restart|r)",
    Cmd: []string{"restart"},
    Tags: "owner",
    Prefix: true,
    IsOwner: true,
    Exec: func(sock *x.Nc, m *x.IMsg){
      
      dir, err := os.Getwd()
      if err != nil {
        fmt.Println(err)
      }

         m.Reply("sukses restarting")
      cmd := exec.Command("go", "run", dir+"/main.go")
      cmd.Stdout = os.Stdout
      cmd.Stderr = os.Stderr
      cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
      err = cmd.Start()
      if err != nil {
        fmt.Println(err)
      }

      // Menunggu proses selesai
      err = cmd.Wait()
      if err != nil {
        fmt.Println(err)
      }
    },
  })
}
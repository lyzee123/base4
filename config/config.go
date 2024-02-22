package config

import "sync"

var (
	mu    sync.Mutex
	Name  = "waSocket Bot"
	Login = "code"
	Bot   = "6283873541589" //ryu 6283873541589
	Owner = []string{"628388024064"}
	Self  = false
	LolSite = "https://api.lolhuman.xyz/"
	LolKey = "5f38494f3555283d0446abdf"
)

func SetName(newName string) {
	mu.Lock()
	defer mu.Unlock()
	Name = newName
}

func SetSelf(new bool) {
	mu.Lock()
	defer mu.Unlock()
	Self = new
}

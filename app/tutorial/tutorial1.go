package tutorial

import (
	"fmt"

	"gopkg.in/ini.v1"
)

func Tutorial1() {
	s, err := ini.Load("setting.ini")
	if err != nil {
		fmt.Println(err)
	}
	name := s.Section("User").Key("SiteName").String()
	fmt.Println(name)
}

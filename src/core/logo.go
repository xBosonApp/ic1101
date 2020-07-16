package core

import (
	"fmt"
)

const GVersion = "1.0.2"
const GAppName = "IOT-IC1101 物联网中台/组态/智能数据中心"

var LOGO = `
.___ ________ ___________      .___ _________  ____  ____ _______   ____ 
|   |\_____  \\__    ___/      |   |\_   ___ \/_   |/_   |\   _  \ /_   |
|   | /   |   \ |    |  ______ |   |/    \  \/ |   | |   |/  /_\  \ |   |
|   |/    |    \|    | /_____/ |   |\     \____|   | |   |\  \_/   \|   |
|___|\_______  /|____|         |___| \______  /|___| |___| \_____  /|___|
             \/                             \/                   \/      
QQ: 412475540 / Email: yanming-sohu@sohu.com
--------------------------------------------
`

func init() {
	fmt.Print(LOGO)
	fmt.Print()
}

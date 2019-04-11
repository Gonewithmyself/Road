package main

import "C"
import (
	"fmt"

	_ "misc/db"
)

func main() {

}

//export gotest
func gotest(s *C.char) {
	defer func() {
		fmt.Println("i am in defer")
	}()
	fmt.Println("this is go ", C.GoString(s))
}

/*
import requests


'''
POST /api/v1/users/account/is_bind HTTP/1.1
Content-Type: application/json
Content-Length: 51
app_key: com.dragon.radiumme
nonce: 227
timestamp: 1554986099
check_sum: 963ff5d614952724be39e27d3da6d6fd98cec7c5
User-Agent: Dalvik/2.1.0 (Linux; U; Android 6.0.1; NX549J Build/MMB29M)
Host: account.dragonbattle.cn
Accept-Encoding: gzip
Connection: keep-alive

{"deviceid":"c5835123-69cf-4581-a292-69a60ae0cd14"}
'''
url = 'http://account.dragonbattle.cn/api/v1/users/account/is_bind'
data = {"deviceid":"c5835123-69cf-4581-a292-69a60ae0cd14"}
headers = {"app_key": "com.dragon.radiumme",
            "nonce": "227",
            "timestamp": '1554986099',
            "check_sum": "963ff5d614952724be39e27d3da6d6fd98cec7c5"}
resp=requests.post(url, json=data, headers=headers)
print(resp, resp.text)


'''
POST /api/v1/users/account/login HTTP/1.1
Content-Type: application/json
Content-Length: 310
app_key: com.dragon.radiumme
nonce: 550
timestamp: 1554986099
check_sum: b457161bfd28a39b6ba57e9fb5b650d720eb7647
User-Agent: Dalvik/2.1.0 (Linux; U; Android 6.0.1; NX549J Build/MMB29M)
Host: account.dragonbattle.cn
Accept-Encoding: gzip
Connection: keep-alive

{"acc_type":"CHANNEL_TOKEN",
"access_token":"A843GoQ60R1N/1UzMXKMY92LEk03U4/ZpXX11XYXVPoxKjKpj9+MmLuKXRavkR3WbAPfMbavFjNFASjFyOBC0FUyuzRn4S9WsqQuLLkWgRqZZk8HMeqGmrqr97DGpH2XYQo/+A4SYsqcJFJiffx3iAjCB0AuqH4Is7dEC0fuFXA=",
"device_id":"c5835123-69cf-4581-a292-69a60ae0cd14",
"passport":"rad_1593ffd40d35013c10012c6"}
'''
url = 'http://account.dragonbattle.cn/api/v1/users/account/login'
data = {"acc_type":"CHANNEL_TOKEN",
"access_token":"A843GoQ60R1N/1UzMXKMY92LEk03U4/ZpXX11XYXVPoxKjKpj9+MmLuKXRavkR3WbAPfMbavFjNFASjFyOBC0FUyuzRn4S9WsqQuLLkWgRqZZk8HMeqGmrqr97DGpH2XYQo/+A4SYsqcJFJiffx3iAjCB0AuqH4Is7dEC0fuFXA=",
"device_id":"c5835123-69cf-4581-a292-69a60ae0cd14",
"passport":"rad_1593ffd40d35013c10012c6"}
headers = {"app_key": "com.dragon.radiumme",
            "nonce": "550",
            "timestamp": '1554986099',
            "check_sum": "b457161bfd28a39b6ba57e9fb5b650d720eb7647"}
resp=requests.post(url, json=data, headers=headers)
print(resp, resp.text)
*/

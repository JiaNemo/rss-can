package fn_test

import (
	"testing"

	"github.com/soulteary/RSS-Can/internal/fn"
)

func TestIsDomTagName(t *testing.T) {
	for _, tag := range fn.DomList {
		ret := fn.IsDomTagName(tag)
		if ret != true {
			t.Fatal("An error occurred while checking whether the element is a DOM Tag")
		}
	}

	ret := fn.IsDomTagName("hhhh")
	if ret == true {
		t.Fatal("An error occurred while checking whether the element is a DOM Tag")
	}
}

func TestIsCssSelector(t *testing.T) {
	ret := fn.IsCssSelector(".a")
	if !ret {
		t.Fatal("Error checking if element is a CSS Selector")
	}

	ret = fn.IsCssSelector("h1 a")
	if !ret {
		t.Fatal("Error checking if element is a CSS Selector")
	}

	ret = fn.IsCssSelector("#id")
	if !ret {
		t.Fatal("Error checking if element is a CSS Selector")
	}

	ret = fn.IsCssSelector("> nth-child(1)")
	if !ret {
		t.Fatal("Error checking if element is a CSS Selector")
	}
}

func TestIsStrInArray(t *testing.T) {
	ret := fn.IsStrInArray(fn.DomList, "a")
	if !ret {
		t.Fatal("Checking string array contains data failed")
	}

	ret = fn.IsStrInArray(fn.DomList, "a!!")
	if ret {
		t.Fatal("Checking string array contains data failed")
	}
}

func TestIsVaildPortRange(t *testing.T) {
	ret := fn.IsVaildPortRange(0)
	if ret {
		t.Fatal("IsVaildPortRange failed")
	}

	ret = fn.IsVaildPortRange(-1)
	if ret {
		t.Fatal("IsVaildPortRange failed")
	}

	ret = fn.IsVaildPortRange(1000000)
	if ret {
		t.Fatal("IsVaildPortRange failed")
	}

	ret = fn.IsVaildPortRange(3000)
	if !ret {
		t.Fatal("IsVaildPortRange failed")
	}
}

func TestIsNotEmptyAndNotDefaultString(t *testing.T) {
	ret := fn.IsNotEmptyAndNotDefaultString("", "d")
	if ret {
		t.Fatal("IsNotEmptyAndNotDefaultString failed")
	}

	ret = fn.IsNotEmptyAndNotDefaultString("d", "d")
	if ret {
		t.Fatal("IsNotEmptyAndNotDefaultString failed")
	}
}

func TestIsVaildLogLevel(t *testing.T) {
	ret := fn.IsVaildLogLevel("info")
	if !ret {
		t.Fatal("IsVaildLogLevel failed")
	}
	ret = fn.IsVaildLogLevel("error")
	if !ret {
		t.Fatal("IsVaildLogLevel failed")
	}
	ret = fn.IsVaildLogLevel("warn")
	if !ret {
		t.Fatal("IsVaildLogLevel failed")
	}
	ret = fn.IsVaildLogLevel("debug")
	if !ret {
		t.Fatal("IsVaildLogLevel failed")
	}
	ret = fn.IsVaildLogLevel("not-vaild")
	if ret {
		t.Fatal("IsVaildLogLevel failed")
	}
}

func TestIsBoolString(t *testing.T) {
	ret := fn.IsBoolString("true")
	if !ret {
		t.Fatal("IsBoolString failed")
	}
	ret = fn.IsBoolString("on")
	if !ret {
		t.Fatal("IsBoolString failed")
	}
	ret = fn.IsBoolString("1")
	if !ret {
		t.Fatal("IsBoolString failed")
	}
	ret = fn.IsBoolString("ON")
	if !ret {
		t.Fatal("IsBoolString failed")
	}

	ret = fn.IsBoolString("not-vaild")
	if ret {
		t.Fatal("IsBoolString failed")
	}
}

func TestIsVaildAddr(t *testing.T) {
	ips := []string{"10.10.10.10"}
	for _, ip := range ips {
		if !fn.IsVaildAddr(ip) {
			t.Fatal("IsVaildAddr failed", ip)
		}
	}

	ips = []string{"10.10.10.300", "1234.1234.1234.1234"}
	for _, ip := range ips {
		if fn.IsVaildAddr(ip) {
			t.Fatal("IsVaildAddr failed", ip)
		}
	}

	hosts := []string{"abc", "abc.cc", "a.b.c.d", "ajaja-11.a1.cn", "1-2.c.s", "1_2-a.1.a"}
	for _, host := range hosts {
		if !fn.IsVaildAddr(host) {
			t.Fatal("IsVaildAddr failed", host)
		}
	}

	hosts = []string{"1-a#.s.a", "1@ss.c.c", "1_2-a.1.", "!", "1.2.3"}
	for _, host := range hosts {
		if fn.IsVaildAddr(host) {
			t.Fatal("IsVaildAddr failed", host)
		}
	}

	addrs := []string{"10.10.10.10:1024", "127.0.0.1:6379"}
	for _, addr := range addrs {
		if !fn.IsVaildAddr(addr) {
			t.Fatal("IsVaildAddr faild", addr)
		}
	}

	addrs = []string{"10.10.10.10:", "127.0.0.1:637900", "127.0.0.1:-1"}
	for _, addr := range addrs {
		if fn.IsVaildAddr(addr) {
			t.Fatal("IsVaildAddr faild", addr)
		}
	}
}

func TestIsVaildIPAddr(t *testing.T) {
	ips := []string{"10.10.10.10"}
	for _, ip := range ips {
		if !fn.IsVaildIPAddr(ip) {
			t.Fatal("TestIsVaildIPAddr failed", ip)
		}
	}

	ips = []string{"10.10.10.300", "1234.1234.1234.1234"}
	for _, ip := range ips {
		if fn.IsVaildIPAddr(ip) {
			t.Fatal("TestIsVaildIPAddr failed", ip)
		}
	}

	ret := fn.IsVaildIPAddr("a.b.c.d")
	if ret {
		t.Fatal("TestIsVaildIPAddr failed")
	}
}

func TestIsVaildAddrWithHttpProtocol(t *testing.T) {
	urls := []string{"http://localhost", "https://localhost", "https://10.11.12.13", "http://127.0.0.1:1234", "http://a.b.c.s", "https://a.c.v:233"}
	for _, url := range urls {
		if !fn.IsVaildAddrWithHttpProtocol(url) {
			t.Fatal("Test IsVaildAddrWithHttpProtocol failed", url)
		}
	}

	urls = []string{"http://localhost:", "https://localhost:111111", "https://a.!.a", "no://10.12.12.12"}
	for _, url := range urls {
		if fn.IsVaildAddrWithHttpProtocol(url) {
			t.Fatal("Test IsVaildAddrWithHttpProtocol failed", url)
		}
	}
}

func TestIsVaildAddrWithWsProtocol(t *testing.T) {
	urls := []string{"ws://localhost", "wss://localhost", "wss://10.11.12.13", "wss://127.0.0.1:1234", "ws://a.b.c.s", "wss://a.c.v:233"}
	for _, url := range urls {
		if !fn.IsVaildAddrWithWsProtocol(url) {
			t.Fatal("IsVaildAddrWithWsProtocol failed", url)
		}
	}

	urls = []string{"ws://localhost:", "wss://localhost:111111", "wss://a.!.a", "no://a.b.c"}
	for _, url := range urls {
		if fn.IsVaildAddrWithWsProtocol(url) {
			t.Fatal("IsVaildAddrWithWsProtocol failed", url)
		}
	}
}

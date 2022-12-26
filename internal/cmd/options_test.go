package cmd_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/soulteary/RSS-Can/internal/cmd"
	"github.com/soulteary/RSS-Can/internal/define"
)

func TestSantizeFeedPath(t *testing.T) {
	feedPath := cmd.SantizeFeedPath("////feedpath")
	if feedPath != "/feedpath" {
		t.Fatal("TestSantizeFeedPath failed")
	}

	feedPath = cmd.SantizeFeedPath("feedpath//")
	if feedPath != "/feedpath" {
		t.Fatal("TestSantizeFeedPath failed")
	}

	feedPath = cmd.SantizeFeedPath("////feedpath///")
	if feedPath != "/feedpath" {
		t.Fatal("TestSantizeFeedPath failed")
	}

	feedPath = cmd.SantizeFeedPath("//// feedpath ///")
	if feedPath != "/feedpath" {
		t.Fatal("TestSantizeFeedPath failed")
	}

	feedPath = cmd.SantizeFeedPath("//// feed path ///")
	fmt.Println(feedPath)
	if feedPath != "/feed path" {
		t.Fatal("TestSantizeFeedPath failed")
	}

	feedPath = cmd.SantizeFeedPath("//// fee$%^&*d p!ath /!//")
	fmt.Println(feedPath)
	if feedPath != define.DEFAULT_HTTP_FEED_PATH {
		t.Fatal("TestSantizeFeedPath failed")
	}
}

func TestUpdateBoolOption(t *testing.T) {

	// env: empty, args: false, default: false
	ret := cmd.UpdateBoolOption("TEST_KEY", false, false)
	if ret {
		t.Fatal("UpdateBoolOption failed")
	}
	// env: empty, args: false, default: true
	ret = cmd.UpdateBoolOption("TEST_KEY", false, true)
	if ret {
		t.Fatal("UpdateBoolOption failed")
	}
	// env: empty, args: true, default: true
	ret = cmd.UpdateBoolOption("TEST_KEY", true, true)
	if !ret {
		t.Fatal("UpdateBoolOption failed")
	}
	// env: empty, args: false, default: true
	ret = cmd.UpdateBoolOption("TEST_KEY", false, true)
	if ret {
		t.Fatal("UpdateBoolOption failed")
	}

	// env: on, args: false, default: false
	os.Setenv("TEST_KEY", "on")
	ret = cmd.UpdateBoolOption("TEST_KEY", false, false)
	if !ret {
		t.Fatal("UpdateBoolOption failed")
	}
	// env: on, args: true, default: false
	os.Setenv("TEST_KEY", "on")
	ret = cmd.UpdateBoolOption("TEST_KEY", true, false)
	if !ret {
		t.Fatal("UpdateBoolOption failed")
	}
	// env: on, args: true, default: true
	os.Setenv("TEST_KEY", "on")
	ret = cmd.UpdateBoolOption("TEST_KEY", true, true)
	if !ret {
		t.Fatal("UpdateBoolOption failed")
	}
	// env: on, args: true, default: false
	os.Setenv("TEST_KEY", "on")
	ret = cmd.UpdateBoolOption("TEST_KEY", true, false)
	if !ret {
		t.Fatal("UpdateBoolOption failed")
	}

	// env: off, args: false, default: false
	os.Setenv("TEST_KEY", "off")
	ret = cmd.UpdateBoolOption("TEST_KEY", false, false)
	if ret {
		t.Fatal("UpdateBoolOption failed")
	}
	// env: off, args: true, default: false
	os.Setenv("TEST_KEY", "on")
	ret = cmd.UpdateBoolOption("TEST_KEY", true, false)
	if !ret {
		t.Fatal("UpdateBoolOption failed")
	}
	// env: off, args: true, default: true
	os.Setenv("TEST_KEY", "on")
	ret = cmd.UpdateBoolOption("TEST_KEY", true, true)
	if !ret {
		t.Fatal("UpdateBoolOption failed")
	}
	// env: on, args: true, default: false
	os.Setenv("TEST_KEY", "on")
	ret = cmd.UpdateBoolOption("TEST_KEY", true, false)
	if !ret {
		t.Fatal("UpdateBoolOption failed")
	}

	os.Setenv("TEST_KEY", "")
}

func TestUpdateNumberOption(t *testing.T) {
	// env: empty, args:1, default:0, allowZero:true
	ret := cmd.UpdateNumberOption("TEST_KEY", 1, 0, true)
	if ret != 1 {
		t.Fatal("UpdateNumberOption failed")
	}
	// env: empty, args:0, default:0, allowZero:true
	ret = cmd.UpdateNumberOption("TEST_KEY", 0, 0, true)
	if ret != 0 {
		t.Fatal("UpdateNumberOption failed")
	}
	// env: empty, args:0, default:1, allowZero:true
	ret = cmd.UpdateNumberOption("TEST_KEY", 0, 1, true)
	if ret != 0 {
		t.Fatal("UpdateNumberOption failed")
	}
	// env: empty, args:0, default:1, allowZero:false
	ret = cmd.UpdateNumberOption("TEST_KEY", 0, 1, false)
	if ret != 1 {
		t.Fatal("UpdateNumberOption failed")
	}

	os.Setenv("TEST_KEY", "2")
	// env: "2", args:1, default:0, allowZero:true
	ret = cmd.UpdateNumberOption("TEST_KEY", 1, 0, true)
	fmt.Println(ret)
	if ret != 1 {
		t.Fatal("UpdateNumberOption failed")
	}
	// env: "2", args:0, default:0, allowZero:true
	ret = cmd.UpdateNumberOption("TEST_KEY", 0, 0, true)
	if ret != 2 {
		t.Fatal("UpdateNumberOption failed")
	}
	// env: "2", args:0, default:1, allowZero:true
	ret = cmd.UpdateNumberOption("TEST_KEY", 0, 1, true)
	if ret != 0 {
		t.Fatal("UpdateNumberOption failed")
	}
	// env: "2", args:0, default:1, allowZero:false
	ret = cmd.UpdateNumberOption("TEST_KEY", 0, 1, false)
	if ret != 2 {
		t.Fatal("UpdateNumberOption failed")
	}
	// env: "2", args:3, default:1, allowZero:false
	ret = cmd.UpdateNumberOption("TEST_KEY", 3, 1, false)
	if ret != 3 {
		t.Fatal("UpdateNumberOption failed")
	}
}

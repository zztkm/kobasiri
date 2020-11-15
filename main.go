package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/pkg/browser"
)

const (
	name     = "kobasiri"
	version  = "0.1.0"
	revision = "HEAD"
)

var printVersion = flag.Bool("version", false, "print version")

func fatal(format string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, format, err)
	} else {
		fmt.Fprint(os.Stderr, format)
	}
	os.Exit(1)
}

func usage() {
	fmt.Println(`
kobasiri <command> [arguments]
	[twitter, t] こばしり。さんの公式Twitterにぶっ飛ぶ
		$ kobasiri twitter
	[instagram, i] こばしり。さんの公式Instagramにぶっ飛ぶ
		$ kobasiri instagram
	[makey, m] こばしり。さんの公式Makeyにぶっ飛ぶ
		$ kobasiri makey
	[k, kimamalabo] KIMAMALaboのショップサイトにぶっ飛ぶ
		$ kobasiri kimamalabo
	[youtube, y] こばしり。さんの公式YouTubeにぶっ飛ぶ
		$ kobasiri youtube
	[game, g] 「こばしりはゲームがしたい。」チャンネル(YouTube)にぶっ飛ぶ
		$ kobasiri game
	`)
	os.Exit(1)
}

func main() {
	flag.Parse()

	if *printVersion {
		fmt.Printf("%s %s (rev: %s/%s)\n", name, version, revision, runtime.Version())
		return
	}

	if flag.NArg() < 1 {
		usage()
	}

	switch flag.Arg(0) {
	case "t", "twitter":
		url := "https://twitter.com/KIMAMALabo"
		browser.OpenURL(url)
	case "i", "instagram":
		url := "https://www.instagram.com/lovetomato_24/"
		browser.OpenURL(url)
	case "m", "makey":
		url := "https://makey.asia/lady.php?id=61198"
		browser.OpenURL(url)
	case "k", "kimamalabo":
		url := "https://kimamalabo.shop/"
		browser.OpenURL(url)
	case "y", "youtube":
		url := "https://www.youtube.com/channel/UC4GMLPKC8TgfW1BtJ-0vf-Q"
		browser.OpenURL(url)
	case "g", "game":
		url := "https://www.youtube.com/channel/UCjWthOL2qhSJghqkmBYXmbA"
		fmt.Println("ゲーム動画もまた見たいな～")
		browser.OpenURL(url)
	default:
		usage()
	}
}

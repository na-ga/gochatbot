package main

import (
	"fmt"

	"github.com/na-ga/gochatbot"
)

func main() {
	client, _ := gochatbot.NewClient()
	message := "Hello"
	fmt.Printf("==> %s\n", message)
	for i := 0; i < 30; i++ {
		res, err := client.Chat(message)
		if err != nil {
			fmt.Printf("error occurred. error=%s", err.Error())
			return
		}
		if !res.IsSuccess() || res.IsEmptyResult() {
			fmt.Printf("error occurred. status=%s, result=%s\n", res.Status, res.Result)
			return
		}
		message = res.Result
		fmt.Printf("<== %s\n", message)
	}

	// ==> Hello
	// <== 何度も何度もありがとう！
	// <== すごーっ
	// <== あれれ？
	// <== なーん
	// <== えーと...
	// <== ふーん？
	// <== え？
	// <== 同じこと何回聞くのー
	// <== えーw
	// <== ほほほー
	// <== 呼んだ?
	// <== 呼んでないー
	// <== なんて呼ぶの？
	// <== 呼んじゃおっかなー笑
	// <== にゃにかね？
	// <== にゃんでも …
	// <== ふふふ
	// <== ふふーっ
	// <== かわいいっ
	// <== でしょでしょー！！
	// <== うん
	// <== ん?
	// <== んー？
	// <== 全然わからないw
	// <== はーい
	// <== (´・ω・｀)
	// <== なでなで
	// <== して？
	// <== なんで？
	// <== なにがー？
}

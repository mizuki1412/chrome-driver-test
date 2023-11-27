package main

import (
	"context"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

func main() {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false), // 开启图像界面
		//chromedp.ProxyServer("http://10.10.1.1:21869"), // 设置代理访问
		chromedp.Flag("mute-audio", false), // 关闭声音
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// navigate to a page, wait for an element, click
	var example string
	var nodes []*cdp.Node
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.baidu.com`),
		// wait for footer element is visible (ie, page is loaded)
		chromedp.WaitVisible(`body > #wrapper`),
		// find and click "Example" link
		//chromedp.Click(`#example-After`, chromedp.NodeVisible),
		// retrieve the text of the textarea
		chromedp.Nodes(`#s-top-left`, &nodes),
	)
	if err != nil {
		log.Fatal("error: ", err)
	}
	for _, e := range nodes {
		log.Println(e.Name, e.Value)
	}
	log.Println(example)
}

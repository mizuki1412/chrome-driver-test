package main

import (
	"context"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"log"
)

func main() {
	// 取消headless
	opts := append(
		chromedp.DefaultExecAllocatorOptions[3:],
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
		//chromedp.ProxyServer("http://10.10.1.1:21869"), // 设置代理访问
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// create a timeout
	//ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	//defer cancel()

	// navigate to a page, wait for an element, click
	var example string
	var nodes []*cdp.Node
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.baidu.com`),
		// wait for footer element is visible (ie, page is loaded)
		chromedp.WaitVisible(`#kw`, chromedp.ByID),
		// find and click "Example" link
		//chromedp.Click(`#example-After`, chromedp.NodeVisible),
		// retrieve the text of the textarea
		//chromedp.Nodes(`#s-top-left`, &nodes),
		chromedp.InnerHTML("#s-top-left", &example),
		//chromedp.Sleep(time.Duration(3)*time.Second),
		chromedp.SendKeys("#kw", "test", chromedp.ByID),
		//chromedp.Click("#su"),
		chromedp.Submit("#form", chromedp.ByID),
	)
	if err != nil {
		log.Fatal("error: ", err)
	}
	for _, e := range nodes {
		log.Println(len(e.Children))
		log.Println(e.Attributes)
	}
	log.Println(example)
	select {}
}

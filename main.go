package main

import (
	"log/slog"
	"os"
)

var (
	itemList map[string]int
	tlogger  *slog.Logger
	jlogger  *slog.Logger
)

type MyCart struct {
	Items []string
}

func init() {
	itemList = map[string]int{
		"Apple":  100,
		"Orange": 50,
		"Grape":  500,
	}

	th := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}
			return a
		},
	})
	tlogger = slog.New(th)
	jh := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}
			return a
		},
	})
	jlogger = slog.New(jh)
}

func main() {
	tlogger.Info("init item cart")
	myCart := &MyCart{}
	tlogger.Info("my cart", "contents", myCart)
	myCart.addItem("Apple")
	myCart.addItem("Orange")
	myCart.addItem("Meet")
	tlogger.Info("my cart", "contents", myCart)

	jlogger.Info("total price", "price", myCart.getTotalPrice())
}

func (c *MyCart) addItem(item string) {
	tlogger.Info("ask to add item", "item", item)
	if _, ok := itemList[item]; !ok {
		tlogger.Error("not carried", "item", item)
	}
	c.Items = append(c.Items, item)
}

func (c *MyCart) getTotalPrice() int {
	var total int
	for _, i := range c.Items {
		total += itemList[i]
	}
	return total
}

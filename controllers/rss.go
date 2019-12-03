package controllers

import (
	"fmt"

	"github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/feeds"
	"github.com/antracewalking/antblog/helpers"
	"github.com/antracewalking/antblog/models"
	"github.com/antracewalking/antblog/system"
)

func RssGet(c *gin.Context) {
	now := helpers.GetCurrentTime()
	domain := system.GetConfiguration().Domain
	feed := &feeds.Feed{
		Title:       "Wblog",
		Link:        &feeds.Link{Href: domain},
		Description: "Wblog,talk about golang,java and so on.",
		Author:      &feeds.Author{Name: "Wangsongyan", Email: "antracewalkinglove@163.com"},
		Created:     now,
	}

	feed.Items = make([]*feeds.Item, 0)
	posts, err := models.ListPublishedPost("", 0, 0)
	if err != nil {
		seelog.Error(err)
		return
	}

	for _, post := range posts {
		item := &feeds.Item{
			Id:          fmt.Sprintf("%s/post/%d", domain, post.ID),
			Title:       post.Title,
			Link:        &feeds.Link{Href: fmt.Sprintf("%s/post/%d", domain, post.ID)},
			Description: string(post.Excerpt()),
			Created:     now,
		}
		feed.Items = append(feed.Items, item)
	}
	rss, err := feed.ToRss()
	if err != nil {
		seelog.Error(err)
		return
	}
	c.Writer.WriteString(rss)
}

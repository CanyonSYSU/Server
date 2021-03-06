package service

import (
	"github.com/CanyonSYSU/Server/entity"
	"github.com/CanyonSYSU/Server/loghelper"
	//"fmt"
	"github.com/CanyonSYSU/Server/db"
	//simplejson "github.com/bitly/go-simplejson"
)

func CommentRegister(order_id, rating_star int, rate_at,
	username, tags, buddha_src, client_text, merchant_text string) (bool, error) {
	var v entity.Comment
	v.Order_ID = order_id
	v.Rating_star = rating_star
	v.Rate_at = rate_at
	v.Username = username
	v.Merchant_text = merchant_text
	v.Client_text = client_text
	v.Buddha_src = buddha_src
	v.Tags = tags
	err := db.CreateComment(&v)
	if err != nil {
		loghelper.Error.Println(err)
		return false, err
	}
	return true, nil
}

func ListAllComments() []entity.Comment {
	return db.QueryComment(func(u *entity.Comment) bool {
		return true
	})
}

func GetCommentCountByTag(tag string) int {
	return db.QueryCommentCountsByTag(tag)
}

func ListAllTags() []entity.Tags {
	tag_str := db.QueryTag()
	if len(tag_str) == 0 {
		return nil
	}
	var tag_map map[string]int
	tag_map = make(map[string]int)
	for _, v := range tag_str {
		tag_map[v]++
	}
	var tags []entity.Tags
	for k, v := range tag_map {
		tags = append(tags, entity.Tags{Tag: k, Count: v})
	}
	return tags
}

func ListCommentsByCount(begin, offset int) []entity.Comment {
	return db.QueryCommentByCount(begin, offset)
}

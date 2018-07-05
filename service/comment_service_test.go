package service

import (
	"reflect"
	"testing"
)

func TestCommentRegister(t *testing.T) {
	type args struct {
		order_id      int
		rating_star   int
		rate_at       string
		username      string
		tags          string
		buddha_src    string
		client_text   string
		merchant_text string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"新增评论0", args{0, 5, "20180630", "huangjm", "good", "touxiang.jpg", "very nice", "thanks"}, true, false},
		{"新增评论1", args{1, 2, "20180701", "mohk", "good", "re23r.jpg", "I like it", "welcome"}, true, false},
		{"新增评论2", args{2, 1, "20180704", "huangzhe", "bad", "dadads.jpg", "what the hell food", "fuck u"}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CommentRegister(tt.args.order_id, tt.args.rating_star, tt.args.rate_at, tt.args.username, tt.args.tags, tt.args.buddha_src, tt.args.client_text, tt.args.merchant_text)
			if (err != nil) != tt.wantErr {
				t.Errorf("CommentRegister() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CommentRegister() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListAllComments(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"获取全部评论", 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := len(ListAllComments()); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAllComments() totally = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCommentCountByTag(t *testing.T) {
	type args struct {
		tag string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"获取标签为'good'的一类评论数量", args{"good"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCommentCountByTag(tt.args.tag); got != tt.want {
				t.Errorf("GetCommentCountByTag() totally = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListAllTags(t *testing.T) { //ListAllTags()需要额外返回bool
	tests := []struct {
		name string
		want int
	}{
		{"目前只有“good”与“bad”两类评论标签", 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := len(ListAllTags()); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAllTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListCommentsByCount(t *testing.T) {
	type args struct {
		begin  int
		offset int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"获取开头2条评论", args{0, 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := len(ListCommentsByCount(tt.args.begin, tt.args.offset)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListCommentsByCount() totally = %v, want %v", got, tt.want)
			}
		})
	}
}

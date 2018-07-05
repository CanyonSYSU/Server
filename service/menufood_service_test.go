package service

import (
	"reflect"
	"testing"
)

func TestMenufoodRegister(t *testing.T) {
	type args struct {
		name          string
		price         float64
		restaurant_id int
		categorys     string
		detail        string
		src           string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"添加可乐", args{"cola", 1.5, 1, "juice", "ice cola", "adwww.jpg"}, true, false},
		{"添加雪碧", args{"sprit", 2.5, 1, "juice", "ice sprit", "32dsd.jpg"}, true, false},
		{"添加健力宝", args{"Jianlibao", 3.5, 1, "juice", "ice Jianlibao", "fwef.jpg"}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MenufoodRegister(tt.args.name, tt.args.price, tt.args.restaurant_id, tt.args.categorys, tt.args.detail, tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("MenufoodRegister() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MenufoodRegister() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListAllMenufoods(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"查看目前三个商品可乐和雪碧和健力宝", 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := len(ListAllMenufoods()); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAllMenufoods() totally = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListAllMenufoodsThroughTags(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"目前只有juice一类菜品", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := len(ListAllMenufoodsThroughTags()); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAllMenufoodsThroughTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMenufoodByName(t *testing.T) {
	type args struct {
		rname string
	}
	tests := []struct {
		name  string
		args  args
		exist bool
	}{
		{"获取可乐", args{"cola"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := (GetMenufoodByName(tt.args.rname) != nil); !reflect.DeepEqual(got, tt.exist) {
				t.Errorf("GetMenufoodByName() = %v, want %v", got, tt.exist)
			}
		})
	}
}

func TestUpdateMenufood(t *testing.T) {
	type args struct {
		id        int
		src       string
		name      string
		price     float64
		detail    string
		categorys string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"根据错误id更新可乐的价格", args{0, "adwww.jpg", "cola", 3.5, "ice cola", "juice"}, 0},
		{"根据正确id更新可乐的价格", args{1, "adwww.jpg", "cola", 3.5, "ice cola", "juice"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateMenufood(tt.args.id, tt.args.src, tt.args.name, tt.args.price, tt.args.detail, tt.args.categorys); got != tt.want {
				t.Errorf("UpdateMenufood() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteMenufood(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"根据错误id删除健力宝", args{5}, 0},
		{"根据正确id删除健力宝", args{3}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteMenufood(tt.args.id); got != tt.want {
				t.Errorf("DeleteMenufood() totally = %v, want %v", got, tt.want)
			}
		})
	}
}

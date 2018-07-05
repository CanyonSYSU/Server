package service

import (
	"reflect"
	"testing"
)

func TestRestaurantRegister(t *testing.T) {
	type args struct {
		name         string
		address      string
		certificates string
		servertime   string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"创建新商店", args{"canyon", "higer city", "zhengshu", "9:00-20:00"}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RestaurantRegister(tt.args.name, tt.args.address, tt.args.certificates, tt.args.servertime)
			if (err != nil) != tt.wantErr {
				t.Errorf("RestaurantRegister() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RestaurantRegister() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListAllRestaurants(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"获取目前商家列表，只有一个", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := len(ListAllRestaurants()); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAllRestaurants() totally = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRestaurantByName(t *testing.T) {
	type args struct {
		rname string
	}
	tests := []struct {
		name  string
		args  args
		exist bool
	}{
		{"获取名为canyon的商店", args{"canyon"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := (GetRestaurantByName(tt.args.rname) != nil); !reflect.DeepEqual(got, tt.exist) {
				t.Errorf("GetRestaurantByName() = %v, want %v", got, tt.exist)
			}
		})
	}
}

func TestUpdateRestaurant(t *testing.T) {
	type args struct {
		name         string
		address      string
		servertime   string
		certificates string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"更新canyon商家的地址", args{"canyon", "sysu shen6", "8:00-20:00", "zhengshu"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateRestaurant(tt.args.name, tt.args.address, tt.args.servertime, tt.args.certificates); got != tt.want {
				t.Errorf("UpdateRestaurant() = %v, want %v", got, tt.want)
			}
		})
	}
}

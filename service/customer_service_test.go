package service

import (
	"fmt"
	"testing"
)

//测试插入用户
func Test_CustomerRegister(t *testing.T) {

	shouldSuccess := []struct {
		name          string
		password      string
		restaurant_id int
		phone         string
		expected      bool
	}{
		{"hzw12", "123", 1, "13719179588", true},
		{"hzw23", "123", 1, "", false},
	}
	for _, ts := range shouldSuccess {
		if real, _ := CustomerRegister(ts.name, ts.password, ts.restaurant_id, ts.phone); real != ts.expected {
			t.Errorf("fail in " + ts.name + ts.phone)
		} else {
			t.Log("success in " + ts.name + ts.phone)
		}
	}

}

//在插入的基础上测试查询所有用户
func Test_ListAllCustomers(t *testing.T) {
	real := ListAllCustomers()
	for _, item := range real {
		fmt.Println(item.Name)
	}
	if len(real) == 2 {
		t.Log("success in ")
	} else {
		t.Errorf("fail in ")
	}
}

//在插入的基础上测试通过Name查询用户
func Test_GetCustomerByName(t *testing.T) {
	shouldSuccess := []struct {
		cname    string
		expected bool
	}{
		{"hzw12", true},
		{"hzw2", false},
	}
	for _, ts := range shouldSuccess {
		if real := GetCustomerByName(ts.cname); real != nil && ts.expected == true ||
			real == nil && ts.expected == false {
			t.Logf("success in " + ts.cname)
		} else {
			t.Errorf("fail in " + ts.cname)
		}

	}

}

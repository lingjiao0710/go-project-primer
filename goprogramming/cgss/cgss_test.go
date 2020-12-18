package main

import (
	"reflect"
	"testing"
)

func TestGetCommandHandlers(t *testing.T) {
	tests := []struct {
		name string
		want map[string]func(args []string) int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCommandHandlers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCommandHandlers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHelp(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Help(tt.args.args); got != tt.want {
				t.Errorf("Help() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListPlayer(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListPlayer(tt.args.args); got != tt.want {
				t.Errorf("ListPlayer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogin(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Login(tt.args.args); got != tt.want {
				t.Errorf("Login() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogout(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Logout(tt.args.args); got != tt.want {
				t.Errorf("Logout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuit(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Quit(tt.args.args); got != tt.want {
				t.Errorf("Quit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSend(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Send(tt.args.args); got != tt.want {
				t.Errorf("Send() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_startCenterService(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := startCenterService(); (err != nil) != tt.wantErr {
				t.Errorf("startCenterService() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

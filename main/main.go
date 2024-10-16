package main

import (
	"context"
	"fmt"
	timeweb "github.com/nil-raysh/go-timeweb"
)

func main() {
	ctx := context.Background()

	conf := timeweb.NewConfiguration()
	conf.AddDefaultHeader("Authorization", "Bearer eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCIsImtpZCI6IjFrYnhacFJNQGJSI0tSbE1xS1lqIn0.eyJ1c2VyIjoiY281Mjc3NCIsInR5cGUiOiJhcGlfa2V5IiwicG9ydGFsX3Rva2VuIjoiNjM5MDdkZjYtYTBiZS00NTY0LWIwMjktMmQxMjVlMWIyMmQwIiwiYXBpX2tleV9pZCI6IjYwZmNlYThhLTBhZTItNDgxZS05M2M2LTAyOGQxOTQxNmFmOSIsImlhdCI6MTcyOTA5MDA5NX0.elU3IOulnfbUZhHGMmw-_79ngk03O5N71Gmhk3uvY3lgvgppgpUpRuqLHzc1V6RF0a030mINImW5vk0QNGLrGDo_LMYBSqUM6N1A2h2h-DOvYm4NEliB4SWQE7TtEtE8j0rj1kpOi6aLEqXIzRKMA5pWJ_IgDP6wWMfmrHMOQME6nTdVEaoE4YOpD_XE1R12sTnOif-Pz5OAnwHmUcImT5AJ8JgsFkfbTqOXiY1PDZYeX0qqPQBhXI7yEo17UDW-tHamVwtIXORjmcAyS-OV2dWgQM8QD8QF_yFBzr4d1ldD88sERTMkJGbdlUKVQ9i_SOwdyT0nKPBoK7SkvHOPR_YgGQe_39k-9-qAUNjJnzDacWj4d4lkgQBGs2WXL9xD2WD5b00ASz-8MzUIs3ISovXYYFporOiFW0iXAZnCC3EgW_sUhFXGfFW7lH47T1OVXw20x6VX8iMqGW0zn4DWGzU2JC8pmxepTESG_Vgau4cJiS4LB5Ovt6GosMOKshsM")
	timewebClient := timeweb.NewAPIClient(conf)

	servers, _, _ := timewebClient.ServersAPI.GetServers(ctx).Execute()
	_ = servers
	fmt.Println(servers)
}

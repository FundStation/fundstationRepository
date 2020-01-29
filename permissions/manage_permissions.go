package permission

import (
	"strings"
)

type permission struct {
	roles   []string
	methods []string
}

type authority map[string]permission

var authorities = authority{
	"/donor": permission{
		roles:   []string{"DONOR"},
		methods: []string{"GET", "POST"},
	},
	"/recipient": permission{
		roles:   []string{"RECIPIENT"},
		methods: []string{"GET", "POST"},
	},
	"/home":permission{
		roles:[]string{"RECIPIENT","DONOR","ADMIN"},
		methods:[]string{"GET","POST"},
	},
	"/admin":permission{
		roles:[]string{"ADMIN"},
		methods:[]string{"GET","POST"},
	},



}


func HasPermission(path string, role string, method string) bool {
	if strings.HasPrefix(path, "/donor") {
		path = "/donor"
	}
	if strings.HasPrefix(path, "/recipient") {
		path = "/recipient/signup"
	}
	if strings.HasPrefix(path, "/admin") {
		path = "/admin"
	}
	perm := authorities[path]
	checkedRole := checkRole(role, perm.roles)
	checkedMethod := checkMethod(method, perm.methods)
	if !checkedRole || !checkedMethod {
		return false
	}
	return true
}

func checkRole(role string, roles []string) bool {
	for _, r := range roles {
		if strings.ToUpper(r) == strings.ToUpper(role) {
			return true
		}
	}
	return false
}

func checkMethod(method string, methods []string) bool {
	for _, m := range methods {
		if strings.ToUpper(m) == strings.ToUpper(method) {
			return true
		}
	}
	return false
}

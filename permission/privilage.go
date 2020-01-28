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
	"/contact": permission{
		roles:   []string{"Tourist"},
		methods: []string{"GET", "POST"},
	},
	"/login": permission{
		roles:   []string{"Tourist"},
		methods: []string{"GET", "POST"},
	},
	"/about": permission{
		roles:   []string{"Tourist"},
		methods: []string{"GET"},
	},
	"/logout": permission{
		roles:   []string{"Tourist"},
		methods: []string{"POST"},
	},
	"/signup": permission{
		roles:   []string{"Tourist"},
		methods: []string{"GET", "POST"},
	},
	"/menu": permission{
		roles:   []string{"Tourist"},
		methods: []string{"GET"},
	},
	"/order": permission{
		roles:   []string{"Tourist"},
		methods: []string{"GET", "POST"},
	},
	"/LandLord": permission{
		roles:   []string{"LandLord"},
		methods: []string{"GET", "POST"},
	},
}

// HasPermission checks if a given role has permission to access a given route for a given method
func HasPermission(path string, role string, method string) bool {
	if strings.HasPrefix(path, "/LandLord") {
		path = "/LandLord"
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

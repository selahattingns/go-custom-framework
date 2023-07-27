package Routes

type Route struct {
	Url        string
	Middleware string
	Module     string
	Controller string
	Function   string
}

func Api() []Route {

	return []Route{
		{Url: "Home", Module: "/", Controller: "", Function: "homeHandler", Middleware: ""},
	}
}

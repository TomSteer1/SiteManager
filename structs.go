package main

type Portfolio struct {
	Id         string           `json:"id"`
	Name       string           `json:"name"`
	Title      string           `json:"title"`
	Projects   []Project        `json:"projects"`
	Categories []Category       `json:"categories"`
	Skills     map[string]Skill `json:"skills"`
	Education  []Education      `json:"education"`
	Contacts   []Contact        `json:"contacts"`
	Pages      Pages            `json:"pages"`
}

type Project struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Skills      []string `json:"skills"`
	Image       string   `json:"image"`
	Link        string   `json:"link"`
}

type Category struct {
	Id     string   `json:"id"`
	Name   string   `json:"name"`
	Skills []string `json:"skills"`
}

type Skill struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
	Fg    string `json:"fg"`
	Bg    string `json:"bg"`
}

type Education struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	File string `json:"file"`
}

type Contact struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Link     string `json:"link"`
	Image    string `json:"icon"`
}

type Pages struct {
	Home      bool `json:"home"`
	Skills    bool `json:"skills"`
	Projects  bool `json:"projects"`
	Education bool `json:"education"`
	Posts     bool `json:"posts"`
	Contact   bool `json:"contact"`
}

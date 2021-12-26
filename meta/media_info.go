package meta

type Audio struct {
	Title    string    `json:"title"`
	Artist   string    `json:"artist"`
	Album    string    `json:"album"`
	Resource *Resource `json:"resource"`
}

type Resource struct {
	Url     string            `json:"url"`
	Headers map[string]string `json:"-"`
}

type MediaMeta struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

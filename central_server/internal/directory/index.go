package directory

//To define a better structure for the directory index.
type Index struct{
	Username string `json:"username,omitempty"`
	Files []string `json:"files"`
}
package models

type Dist struct {
	Shasum  string `json:"shahum"`
	Tarball string `json:"tarball"`
}

type PkgInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Dist    Dist   `json:"dist"`
}

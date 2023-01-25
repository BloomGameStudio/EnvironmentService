package controllers

type MeshesVertices struct {
	Meshes []Mesh `json:"meshes"`
}

type Mesh struct {
	Vertices  []Position `json:"vertices"`
	Triangles []int64    `json:"triangles"`
	Name      string     `json:"name"`
	Layer     string     `json:"layer"`
	Transform Transform  `json:"transform"`
}

type Transform struct {
	Position Position `json:"position"`
	Scale    Position `json:"scale"`
	Rotation Position `json:"rotation"`
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

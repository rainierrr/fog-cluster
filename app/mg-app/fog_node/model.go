package fogNode

// User is user models property
type FogNode struct {
	Id    uint   `db:"id, primarykey, autoincrement" json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
	Tag   string `json:"tag"`
	Ip    string `json:"ip"`
}

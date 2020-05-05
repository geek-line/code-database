package structs

//Knowledge knowledgesテーブルのエンティティ
type Knowledge struct {
	ID          int
	Title       string
	Content     string
	CreatedAt   string
	UpdatedAt   string
	Likes       int
	EyecatchSrc string
	IsPublished bool
}

//Tag tagsテーブルのエンティティ
type Tag struct {
	ID        int
	Name      string
	CreatedAt string
	UpdatedAt string
}

//KnowledgesTags knowledges_tagsテーブルのエンティティ
type KnowledgesTags struct {
	ID          int
	KnowledgeID int
	TagID       int
	CreatedAt   string
	UpdatedAt   string
}

//Eyecatch eyecatchesテーブルのエンティティ
type Eyecatch struct {
	ID   int
	Name string
	Src  string
}

//AdminUser admin_userテーブルのエンティティ
type AdminUser struct {
	ID       int
	Email    string
	Password string
}

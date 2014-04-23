package main

/**
 * gitlab user object
 *
 * see http://doc.gitlab.com/ce/api/session.html
 */
type User struct {
	Id               uint32 `json:id`
	Username         string `json:username`
	Email            string `json:email`
	ColorSchemeId    int32  `json:"color_scheme_id"`
	State            string `json:state`
	AvatarUrl        string `json:"avatar_url"`
	Provider         string `json:provider"`
	Name             string `json:name`
	PrivateToken     string `json:"private_token"`
	Blocked          bool   `json:blocked`
	CreatedAt        string `json:"created_at"`
	Bio              string `json:bio`
	Skype            string `json:skype`
	Linkedin         string `json:linkedin`
	Twitter          string `json:twitter`
	ExternUid        string `json:"extern_uid"`
	WebsiteUrl       string `json:"website_url"`
	DarkScheme       bool   `json:"dark_scheme"`
	ThemeId          uint32 `json:"theme_id"`
	IsAdmin          bool   `json:"is_admin"`
	CanCreateGroup   bool   `json:"can_create_group"`
	CanCreateTeam    bool   `json:"can_create_team"`
	CanCreateProject bool   `json:"can_create_project"`
}

/**
 * users' SSH keys struct
 * see http://doc.gitlab.com/ce/api/users.html
 */
type SSHKeyItem struct {
	Id    uint32 `json:id`
	Title string `json:title`
	Key   string `json:key`
}

/**
 * group struct
 * see http://doc.gitlab.com/ce/api/groups.html
 */
type GroupItem struct {
	Id      uint32 `json:id`
	Name    string `json:name`
	Path    string `json:path`
	OwnerId uint32 `json:"owner_id"`
}

/**
 * group member stuct
 * see http://doc.gitlab.com/ce/api/groups.html
 */

type GroupMemberItem struct {
	Id          uint32 `json:id`
	Username    string `json:username`
	Email       string `json:email`
	State       string `json:state`
	CreatedAt   string `json:"created_at"`
	AccessLevel uint8  `json:access_level`
}

/**
 * one commit struct
 * see http://doc.gitlab.com/ce/api/commits.html
 */
type CommitItem struct {
	Id             string   `json:id`
	ShortId        string   `json:"short_id"`
	Title          string   `json:title`
	AuthorName     string   `json:author_name`
	AuthorEmail    string   `json:author_email`
	CreatedAt      string   `json:"created_at"`
	CommiittedDate string   `json:"committed_date"`
	AuthoredDate   string   `json:"authored_date"`
	ParentIds      []string `json:"parent_ids"`
}

/**
 * commits diff struct
 * see http://doc.gitlab.com/ce/api/commits.html
 */
type CommitsDiffItem struct {
	Diff        string `json:diff`
	NewPath     string `json:"new_path"`
	OldPath     string `json:"old_path"`
	AMode       string `json:"a_mode"`
	BMode       string `json:"b_mode"`
	NewFile     bool   `json:"new_file"`
	RenamedFile bool   `json:"renamed_file"`
	DeletedFile bool   `json:"deleted_file"`
}

/**
 * merge struct
 * @see http://doc.gitlab.com/ce/api/merge_requests.html
 */
type MergeItem struct {
	Id           uint64 `json:id`
	Iid          uint64 `json:iid`
	TargetBranch string `json:"target_branch"`
	SourceBranch string `json:"source_branch"`
	ProjectId    uint64 `json:"project_id"`
	Title        string `json:title`
	State        string `json:state`
	Upvotes      uint32 `json:upvotes`
	Downvotes    uint32 `json:downvotes`
	Author       *User  `json:author`
	Assignee     *User  `json:assignee`
}

/**
 * merge struct
 * @see http://doc.gitlab.com/ce/api/merge_requests.html
 */
type MergeCommentItem struct {
	Id     uint64 `json:id`
	Note   string `json:note`
	Author *User  `json:author`
}

/**
 * repository file struct
 * see http://doc.gitlab.com/ce/api/repository_files.html
 */
type RepoFileItem struct {
	FileName   string `json:"file_name"`
	FilePath   string `json:"file_path"`
	Size       string `json:size`
	BranchName string `json:"branch_name"`
	Encoding   string `json:encoding`
	Content    string `json:content`
	Ref        string `json:ref`
	BlobId     string `json:"blob_id"`
	CommitId   string `json:"commit_id"`
}

/**
 * project milestone struct
 * http://doc.gitlab.com/ce/api/milestones.html
 */
type MileStoneItem struct {
	Id          uint32 `json:id`
	Iid         uint32 `json:iid`
	ProjectId   uint32 `json:"project_id"`
	Title       string `json:title`
	Description string `json:description`
	DueDate     string `json:"due_date"`
	State       string `json:state`
	UpdatedAt   string `json:"updated_at"`
	CreatedAt   string `json:"created_at"`
}

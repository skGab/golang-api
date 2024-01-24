package valueObjects

type UserVo struct {
	ID       string
	Email    int
	Password string
	User     string
	Tasks    []TasksVo
}

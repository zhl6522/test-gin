package model

type Student struct {
	Name	string
	Score	float64
}
type student struct {
	Name	string
	score	float64		//隐私字段，不对外公开展示
}

func Newstudent(n string, s float64) *student {
	return &student{
		Name:n,
		score:s,
	}
}

//如果score字段首字母小写，则在其他包不可以直接使用，我们提供一个方法
func (s *student) GetScore() float64 {
	return s.score
}

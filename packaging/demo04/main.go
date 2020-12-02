package model

type student struct {
	Name  string
	Score float64
}

//因为student的首字母小写，只能在本包使用,其他包不能使用
//通过工厂模式解决

func NewStudent(n string, s float64) *student {
	return &student{
		Name:  n,
		Score: s,
	}
}

//如果socre首字母小写，则在其他包不能直接访问
func (s *student) GetScore() float64 {
	return s.Score
}

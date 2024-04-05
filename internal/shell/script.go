package shell

type Script struct {
	Content []byte
}

func (s *Script) AddLine(line []byte) {
	s.Content = append(s.Content, line...)
	s.Content = append(s.Content, "\n"...)
}
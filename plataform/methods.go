package plataform

// Cuando el metodo inicia con mayuscula es publico y si es minuscula es privado
func (course *Course) GetTitle() string {
	return course.Title
}

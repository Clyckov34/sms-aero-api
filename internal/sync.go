package internal

//setting настройка группы горутины.
//	number - Укажите кол-во групп gorutino
func (m *WaitGroup) setting(number int) {
	m.Wg.Add(number)
}

//wait ожидание всех горутин
func (m *WaitGroup) wait() {
	m.Wg.Wait()
}

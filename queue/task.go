package queue

type Task struct{
	URL string
	Status string
}

// Methods Get Status/Change the Status?
func (task *Task) ChangeStatus(status string){
	task.Status = status
}
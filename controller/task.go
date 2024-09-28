package controller

// func CreateTask(c *gin.Context) {
// 	var task model.Task
// 	db := db.NewDB()
// 	err := c.Bind(&task)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, err)
// 	}
// 	storedUser := model.User{}
// 	err = db.Where("id=?", task.UserRefer).First(&storedUser).Error
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, err)
// 	}
// 	newTask := model.Task{Title: task.Title, Detail: task.Detail, Status: task.Status}
// 	newTask.UserRefer = storedUser.ID
// 	err = db.Create(&newTask).Error
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, err)
// 	}
// 	c.JSON(http.StatusCreated, gin.H{"task": task})
// }
//
// func GetTasks(c *gin.Context) {
// 	user := model.User{}
// 	c.Bind(&user)
// 	c.JSON(200, user)
// }

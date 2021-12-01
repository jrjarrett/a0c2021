package application


type App struct {

}


func New() *App {
	application := App{}
	return &application
}

func (a *App) Run() {

}

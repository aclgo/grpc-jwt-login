package user

type UserRepoDatabase interface {
	Add()
	Find()
	FindByID()
	FindByEmail()
	FindAll()
	Update()
	Delete()
}

type UserRepoCache interface {
	Set()
	Get()
	Del()
}

type UserUC interface {
	Add()
	Find()
	FindByID()
	FindByEmail()
	FindAll()
	Update()
	Delete()
}

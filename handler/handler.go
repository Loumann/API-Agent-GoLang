package handler

import "awesomeProject/repos"

type Handler struct {
	r *repos.Repository
}

func GetHadler(repo *repos.Repository) *Handler {
	return &Handler{r: repo}
}

package handler

/*
type TreeController struct {
	repo domain.TreeRepository
}

func NewTreeController(repo domain.TreeRepository) *TreeController {
	return &TreeController{
		repo: repo,
	}
}

func (c TreeController) GetAll(w http.ResponseWriter, r *http.Request) {

	tree, err := c.repo.FindAll()
	if err != nil {
		httpmiddleware.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	httpmiddleware.RespondWithJSON(w, http.StatusOK, tree)
}

func (c TreeController) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := uuid.Parse(params["id"])
	if err != nil {
		httpmiddleware.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	tree, err := c.repo.FindByID(ID)
	if err != nil {
		httpmiddleware.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	httpmiddleware.RespondWithJSON(w, http.StatusOK, tree)
}

func (c TreeController) New(w http.ResponseWriter, r *http.Request) {
	var tree *domain.Tree

	if err := httpmiddleware.DecodeJSONBody(w, r, tree); err != nil {
		httpmiddleware.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	if err := c.repo.Persist(tree); err != nil {
		httpmiddleware.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	httpmiddleware.RespondWithJSON(w, http.StatusCreated, tree)
}

func (c TreeController) Update(w http.ResponseWriter, r *http.Request) {
	var tree *domain.Tree

	if err := httpmiddleware.DecodeJSONBody(w, r, tree); err != nil {
		httpmiddleware.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	params := mux.Vars(r)
	ID, err := uuid.Parse(params["id"])
	if err != nil {
		httpmiddleware.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	tree.ID = ID

	err = c.repo.Persist(tree)
	if err != nil {
		httpmiddleware.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	httpmiddleware.RespondWithJSON(w, http.StatusOK, tree)
}

func (c TreeController) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := uuid.Parse(params["id"])
	if err != nil {
		httpmiddleware.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	if err = c.repo.Delete(ID); err != nil {
		httpmiddleware.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	httpmiddleware.RespondWithJSON(w, http.StatusNoContent, nil)
}
*/

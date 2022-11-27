package rest_test

/*
type RouterTestSuite struct {
	suite.Suite
}

func TestRouter(t *testing.T) {
	suite.Run(t, new(RouterTestSuite))
}

func (s *RouterTestSuite) TestGet() {
	svc := rest.NewService(rest.Config{})
	svc.Get("/test", func(_ http.ResponseWriter, _ *http.Request) {})

	req := httptest.NewRequest("GET", "/test", nil)
	res := httptest.NewRecorder()
	svc.ServeHTTP(res, req)

	s.Equal(http.StatusOK, res.Code)
}

func (s *RouterTestSuite) TestGetNotFound() {
	svc := rest.NewService(rest.Config{})

	req := httptest.NewRequest("GET", "/test", nil)
	res := httptest.NewRecorder()
	svc.ServeHTTP(res, req)

	s.Equal(http.StatusNotFound, res.Code)
}

func (s *RouterTestSuite) TestGetNotAllowed() {
	svc := rest.NewService(rest.Config{})
	svc.Post("/test", func(_ http.ResponseWriter, _ *http.Request) {})

	req := httptest.NewRequest("GET", "/test", nil)
	res := httptest.NewRecorder()
	svc.ServeHTTP(res, req)

	s.Equal(http.StatusMethodNotAllowed, res.Code)
}

func (s *RouterTestSuite) TestPost() {
	svc := rest.NewService(rest.Config{})
	svc.Post("/test", func(_ http.ResponseWriter, _ *http.Request) {})

	req := httptest.NewRequest("POST", "/test", nil)
	res := httptest.NewRecorder()
	svc.ServeHTTP(res, req)

	s.Equal(http.StatusOK, res.Code)
}

func (s *RouterTestSuite) TestPostNotAllowed() {
	svc := rest.NewService(rest.Config{})
	svc.Get("/test", func(_ http.ResponseWriter, _ *http.Request) {})

	req := httptest.NewRequest("POST", "/test", nil)
	res := httptest.NewRecorder()
	svc.ServeHTTP(res, req)

	s.Equal(http.StatusMethodNotAllowed, res.Code)
}
*/

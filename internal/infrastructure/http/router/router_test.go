package router_test

import (
	"net/http/httptest"
	"testing"

	"github.com/joffrua/go-famtree/internal/infrastructure/http/router"

	"github.com/stretchr/testify/suite"
)

type RouterTestSuite struct {
	suite.Suite
}

func TestRouter(t *testing.T) {
	suite.Run(t, new(RouterTestSuite))
}

func (s *RouterTestSuite) Test1() {
	router := router.New(router.Config{})
	server := httptest.NewServer(router.Handler())
	defer server.Close()

}

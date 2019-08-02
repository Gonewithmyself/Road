package main

import (
	"log"
	"testing"

	gocheck "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	gocheck.TestingT(t)
}

type mySuite struct{}

var f = 10

func (s *mySuite) SetUpSuite(c *gocheck.C) {
	// panic(123)
	f = 20
	log.Println(f)
}

func (s *mySuite) TestGo(c *gocheck.C) {
	c.Assert(123, gocheck.Equals, 123)
	log.Println(f)
}

func (s *mySuite) BenchmarkLogic(c *gocheck.C) {
	for i := 0; i < c.N; i++ {
		join()
	}
}

func (s *mySuite) BenchmarkLogic1(c *gocheck.C) {
	for i := 0; i < c.N; i++ {
		join1()
	}
}

func (s *mySuite) BenchmarkLogic2(c *gocheck.C) {
	for i := 0; i < c.N; i++ {
		join2()
	}
}

func (s *mySuite) BenchmarkLogic3(c *gocheck.C) {
	for i := 0; i < c.N; i++ {
		join3()
	}
}

var _ = gocheck.Suite(&mySuite{})

func Test_join(t *testing.T) {

}

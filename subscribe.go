package conexus

import (
	"github.com/kierdavis/argo"
	"github.com/linkeddata/gold"
	"log"
)

var (
	foaf = argo.NewNamespace("http://xmlns.com/foaf/0.1/")
	synd = argo.NewNamespace("http://ns.rww.io/synd#")
)

func (user *User) Subscribe(uri string) bool {
	g := gold.NewGraph(uri)
	err := g.LoadURI(uri)
	if err != nil {
		log.Printf("%+v\n", err)
		return false
	}

	sources := g.All(argo.NewResource(uri), synd.Get("source"), nil)

	if !userExists(user.Uri) {
		err = user.addUser()
		if err != nil {
			log.Printf("%+v\n", err)
			return false
		}
	}

	for _, s := range sources {
		log.Printf("%+v\n", s.Object)
	}

	return true
}

func userExists(webid string) bool {
	res := make([]User, 0)
	err := db.Cols("uri").Where("uri LIKE ?", `%`+webid+`%`).Find(&res)
	if err == nil && len(res) > 0 && len(res[0].Uri) > 0 {
		return true
	}
	return false
}

func (user *User) addUser() error {
	//TODO also add feeds from user
	_, err := db.InsertOne(user)
	if err != nil {
		return err
	}
	return nil
}

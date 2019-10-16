package lungo

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

var alternativeTypes = map[string]string{
	"*mongo.Client":     "lungo.IClient",
	"*mongo.Database":   "lungo.IDatabase",
	"*mongo.Collection": "lungo.ICollection",
	"*mongo.Cursor":     "lungo.ICursor",
}

func TestClientInterface(t *testing.T) {
	a := reflect.TypeOf((*IClient)(nil)).Elem()
	b := reflect.TypeOf(&mongo.Client{})
	assert.Equal(t, listMethods(a, false), listMethods(b, true))
}

func TestDatabaseInterface(t *testing.T) {
	a := reflect.TypeOf((*IDatabase)(nil)).Elem()
	b := reflect.TypeOf(&mongo.Database{})
	assert.Equal(t, listMethods(a, false), listMethods(b, true))
}

func TestCollectionInterface(t *testing.T) {
	a := reflect.TypeOf((*ICollection)(nil)).Elem()
	b := reflect.TypeOf(&mongo.Collection{})
	assert.Equal(t, listMethods(a, false), listMethods(b, true))
}

func TestCursorInterface(t *testing.T) {
	a := reflect.TypeOf((*ICursor)(nil)).Elem()
	b := reflect.TypeOf(&mongo.Cursor{})
	assert.Equal(t, listMethods(a, false), listMethods(b, true))
}

func listMethods(t reflect.Type, original bool) string {
	var list []string
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		f := m.Type.String()[4:]

		if original {
			c := strings.Index(f, ",")
			if c >= 0 {
				f = "(" + f[c+2:]
			} else {
				c = strings.Index(f, ")")
				f = "(" + f[c:]
			}

			for a, b := range alternativeTypes {
				f = strings.ReplaceAll(f, a, b)
			}
		}

		list = append(list, m.Name+f)
	}

	return strings.Join(list, "\n")
}

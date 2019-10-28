package lungo

import (
	"strings"

	"github.com/256dpi/lungo/bsonkit"
)

type NS [2]string

func (ns NS) String() string {
	return strings.Join(ns[:], ".")
}

type Data struct {
	Namespaces map[NS]*Namespace `bson:"namespaces"`
}

func NewData() *Data {
	return (&Data{}).Prepare()
}

func (d *Data) Prepare() *Data {
	// ensure namespaces
	if d.Namespaces == nil {
		d.Namespaces = make(map[NS]*Namespace)
	}

	// init namespaces
	for _, namespace := range d.Namespaces {
		namespace.Prepare()
	}

	return d
}

func (d *Data) Clone() *Data {
	// create clone
	clone := &Data{
		Namespaces: map[NS]*Namespace{},
	}

	// copy namespaces
	for name, namespace := range d.Namespaces {
		clone.Namespaces[name] = namespace
	}

	return clone
}

type Namespace struct {
	Documents *bsonkit.Set              `bson:"documents"`
	Indexes   map[string]*bsonkit.Index `bson:"indexes"`
}

func NewNamespace(name NS) *Namespace {
	return (&Namespace{
		Documents: bsonkit.NewSet(nil),
		Indexes: map[string]*bsonkit.Index{
			"_id_": bsonkit.NewIndex(true, []bsonkit.Column{
				{Path: "_id"},
			}),
		},
	}).Prepare()
}

func (n *Namespace) Prepare() *Namespace {
	// prepare indexes
	for _, index := range n.Indexes {
		index.Prepare(n.Documents.List)
	}

	return n
}

func (n *Namespace) Clone() *Namespace {
	// create new namespace
	clone := &Namespace{
		Documents: n.Documents.Clone(),
		Indexes:   map[string]*bsonkit.Index{},
	}

	// clone indexes
	for name, index := range n.Indexes {
		clone.Indexes[name] = index.Clone()
	}

	return clone
}

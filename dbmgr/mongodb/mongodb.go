package mongodb

import (
	"../../dbmgr"
	"gopkg.in/mgo.v2"
)

// DBMgr Data Base Manager
type DBMgr struct {
	mpSession   *mgo.Session
	mCollection map[string]map[string]*mgo.Collection
}

// InitializeInValue input value
type InitializeInValue struct {
	IP string
}

// InputQuery input query
type InputQuery struct {
	DBName         string
	CollectionName string

	FindQuery   interface{}
	InsertQuery interface{}
	UpdateQuery interface{}
}

// NewIDB is Constructor
func (p *DBMgr) NewIDB() {
	p.mpSession = nil
	p.mCollection = make(map[string]map[string]*mgo.Collection)
}

// Initialize 最初に実行してください.InitializeInValueを入力してください
func (p *DBMgr) Initialize(inputraw interface{}) error {
	input, ok := inputraw.(InitializeInValue)
	if ok != true {
		return ErrorNotMatchInValue
	}

	// create session
	session, err := mgo.Dial(input.IP)

	if err != nil {
		return err
	}

	p.mpSession = session
	return nil
}

// Find find method
func (p *DBMgr) Find(inputraw interface{}) *dbmgr.ReturnValue {
	input, ok := inputraw.(InputQuery)
	if !ok {
		return &dbmgr.ReturnValue{Error: ErrorNotMatchInValue}
	}
	collection := p.getCollection(input)

	var results interface{} //[]bson.M
	err := collection.Find(input.FindQuery).All(&results)
	res := results.([]byte)

	return &dbmgr.ReturnValue{Error: err, Data: res}
}

// Insert insert method
func (p *DBMgr) Insert(inputraw interface{}) *dbmgr.ReturnValue {
	input, ok := inputraw.(InputQuery)
	if !ok {
		return &dbmgr.ReturnValue{Error: ErrorNotMatchInValue}
	}
	collection := p.getCollection(input)

	err := collection.Insert(input.InsertQuery)
	return &dbmgr.ReturnValue{Error: err}
}

// Update update method
func (p *DBMgr) Update(inputraw interface{}) *dbmgr.ReturnValue {
	input, ok := inputraw.(InputQuery)
	if !ok {
		return &dbmgr.ReturnValue{Error: ErrorNotMatchInValue}
	}
	collection := p.getCollection(input)

	err := collection.Update(input.FindQuery, input.UpdateQuery)
	return &dbmgr.ReturnValue{Error: err}
}

//
func (p *DBMgr) getCollection(input InputQuery) *mgo.Collection {
	return p.getCollectionImpl(input.DBName, input.CollectionName)
}

func (p *DBMgr) getCollectionImpl(dbname, collectionname string) *mgo.Collection {
	if p.mCollection[dbname] == nil {
		p.mCollection[dbname] = map[string]*mgo.Collection{}
	}
	if p.mCollection[dbname][collectionname] == nil {
		p.mCollection[dbname][collectionname] = p.mpSession.DB(dbname).C(collectionname)
	}
	return p.mCollection[dbname][collectionname]
}

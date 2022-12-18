package AwesomeDB

type OurDB struct {
	Name string
}

type OurDBInterface interface {
	GetData() string
	WriteData(data string)
	Close() error
}

func (db *OurDB) GetData() string {

}

func (db *OurDB) WriteData(data string) {

}
func (db *OurDB) Close() error {

}

//ourDb의 메서드 GetData, WriteData, Close 메서드 포함

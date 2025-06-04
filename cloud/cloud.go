package cloud

type CloudDB struct {
	url string
}

func (cloud CloudDB) Read() ([]byte, error) {
	return []byte{}, nil
}

func (cloud CloudDB) Write(content []byte) {
	return
}

func NewCloudDB(url string) *CloudDB {
	return &CloudDB{
		url: url,
	}
}

func Write(db CloudDB) ([]byte, error) {
	return []byte{}, nil
}
func Read(db CloudDB) (content []byte) {
	return nil
}

package persist

import (
	"encoding/json"
	"fmt"
	"key-value-db/storage"
	"log"
	"os"
)

type Persist struct {
}

func NewPersist() *Persist {
	return &Persist{}
}

type PersistService interface {
	SaveAtServer(remoteAddr string, mp map[string]storage.Value) error
}

func (p *Persist) SaveAtServer(remoteAddr string, mp map[string]storage.Value) error {
	fileName := fmt.Sprintf("%s.%s", remoteAddr, "json")

	log.Println("Filename is ",fileName)

	file, err := os.Create(fileName)
	if err != nil {
		log.Printf("Error in creating the json file %s: %v", fileName, err)
		return err
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			log.Printf("Error in closing the file %s: %v", fileName, cerr)
		}
	}()

	jsonData, err := json.Marshal(mp)
	if err != nil {
		log.Printf("Error in converting to json: %v", err)
		return err
	}

	_, err = file.Write(jsonData)
	if err != nil {
		log.Printf("Error in saving as json to server %s: %v", fileName, err)
		return err
	}

	if err = file.Sync(); err != nil {
		log.Printf("Error in syncing file %s: %v", fileName, err)
		return err
	}

	log.Printf("Data successfully saved to %s", fileName)
	return nil
}

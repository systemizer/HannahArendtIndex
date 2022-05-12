package ocr

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	vision "cloud.google.com/go/vision/apiv1"
	"github.com/systemizer/ArendtArchives/backend/store"
	"github.com/systemizer/ArendtArchives/backend/store/sqlite"
)

type OCR struct {
	store store.Store
}

func New() (*OCR, error) {
	s, err := sqlite.New()
	if err != nil {
		return nil, err
	}

	return &OCR{
		store: s,
	}, nil
}

func (o *OCR) Start() error {
	ctx := context.Background()
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	cis, err := o.store.ListCollectionItems()
	if err != nil {
		return err
	}

	totalFailures := 0
	for _, ci := range cis {

		if totalFailures >= 10 {
			fmt.Println("Exiting due to too many failures")
			break
		}
		// Only run detection on CollectionItems without
		// existing text detection
		if ci.OCRResult != nil {
			continue
		}

		fmt.Printf("Detecting %s\n", ci.ImageURL)

		//image := vision.NewImageFromURI(ci.ImageURL)

		res, err := http.Get(ci.ImageURL)
		if err != nil {
			totalFailures += 1
			time.Sleep(10 * time.Second)
			continue
		}
		defer res.Body.Close()
		image, err := vision.NewImageFromReader(res.Body)
		if err != nil {
			totalFailures += 1
			time.Sleep(10 * time.Second)
			continue
		}

		fmt.Println("Running Text Detection")
		annonation, err := client.DetectDocumentText(ctx, image, nil)
		if err != nil {
			totalFailures += 1
			time.Sleep(10 * time.Second)
			continue
		}

		if annonation == nil {
			fmt.Println("No annotation found")
			continue
		} else {
			fmt.Println(annonation.Text)
		}

		ci.OCRResult = &annonation.Text

		err = o.store.SaveCollectionItem(&ci)
		if err != nil {
			return err
		}

		time.Sleep(time.Second * 1)
	}

	return nil
}

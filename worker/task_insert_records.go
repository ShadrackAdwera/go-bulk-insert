package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	db "github.com/ShadrackAdwera/go-bulk-insert/db/sqlc"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

type CovidData struct {
	DateRep               string `json:"date_rep"`
	Day                   int16  `json:"day"`
	Month                 int16  `json:"month"`
	Year                  int64  `json:"year"`
	Cases                 int64  `json:"cases"`
	Deaths                int64  `json:"deaths"`
	CountriesAndTerritory string `json:"countries_and_territories"`
	GeoId                 string `json:"geo_id"`
	CountryTerritoryCode  string `json:"country_territory_code"`
	ContinentExp          string `json:"continent_exp"`
	LoadDate              string `json:"load_date"`
	IsoCountry            string `json:"iso_country"`
}

const TaskInsertData = "task:insert_data"
const TaskInsertDataV2 = "task:insert_data:v2"

func (distro *DataTaskDistributor) DistributeData(ctx context.Context, payload *[]CovidData, QueueKey string, opts ...asynq.Option) error {
	jsonPayload, err := json.Marshal(payload)

	if err != nil {
		return fmt.Errorf("failed to marshall json body")
	}

	taskInsert := asynq.NewTask(QueueKey, jsonPayload, opts...)

	info, err := distro.client.EnqueueContext(ctx, taskInsert)

	if err != nil {
		return fmt.Errorf("unable to enqueue task context : %w", err)
	}

	log.Info().
		Str("task_type", info.Type).
		Str("task_id", info.ID).
		Str("queue", info.Queue).
		Int("max_retries", info.MaxRetry).
		Msg("task enqueued")

	return nil
}

func (processor *DataTaskProcessor) TaskProcessData(ctx context.Context, task *asynq.Task) error {

	var payload []CovidData

	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshall json %w", asynq.SkipRetry)
	}

	batchSize := 1000 // Number of records to insert in each batch
	totalRecords := len(payload)
	currentBatch := 0

	c, cancel := context.WithTimeout(ctx, time.Hour) // Set a timeout of 60 minutes
	defer cancel()

	for currentBatch < totalRecords {
		end := currentBatch + batchSize
		if end > totalRecords {
			end = totalRecords
		}
		batch := payload[currentBatch:end]

		// Execute the prepared statement for each record in the batch
		for _, record := range batch {
			_, err := processor.store.CreateCaseTx(c, db.CreateCaseParams{
				DateRep:                 record.DateRep,
				Day:                     int32(record.Day),
				Month:                   int32(record.Month),
				Year:                    int32(record.Year),
				Cases:                   record.Cases,
				Deaths:                  record.Deaths,
				CountriesAndTerritories: record.CountriesAndTerritory,
				GeoID:                   record.GeoId,
				CountryTerritoryCode:    record.CountryTerritoryCode,
				ContinentExp:            record.ContinentExp,
				LoadDate:                record.LoadDate,
				IsoCountry:              record.IsoCountry,
			})
			if err != nil {
				log.Err(err).Msg(err.Error())
				return err
			}
		}

		currentBatch += batchSize
		fmt.Printf("Type: %s, Inserted %d records out of %d\n", task.Type(), currentBatch, totalRecords)
	}
	return nil
}

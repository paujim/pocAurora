package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rdsdataservice"
)

const (
	RACE_ID = iota
	RACE_NAME
	RACE_NUMBER
	MEETING_ID
	MEETING_NAME
	CATEGORY_ID
	ADVERTISED_START
)

// DataServiceAPI ...
type DataServiceAPI interface {
	ExecuteStatement(input *rdsdataservice.ExecuteStatementInput) (*rdsdataservice.ExecuteStatementOutput, error)
}

// SQLClient ...
type SQLClient struct {
	client    DataServiceAPI
	auroraArn *string
	secretArn *string
}

// GetNextRacesByCategory ...
func (h *SQLClient) GetNextRacesByCategory(count int, categories []string) (map[string]interface{}, map[string]interface{}, error) {
	// sql := "SELECT c.category_id, c.race_id, c.race_name, c.race_number, c.meeting_id, c.meeting_name, c.advertised_start LIMIT(5) FROM Categories as c WHERE category_id in [categories]  JOIN Races as r ON r.race_id == c.race_id "
	sql := "SELECT race_id, race_name, race_number, meeting_id, meeting_name, category_id, advertised_start FROM Racing.Races WHERE category_id IN ("
	last := len(categories) - 1
	for i, cat := range categories {
		sql += fmt.Sprintf(" '%s'", cat)
		if i != last {
			sql += ","
		}
	}
	sql += fmt.Sprintf(") LIMIT %d;", count)
	log.Printf("SQL: %s\n", sql)

	params := &rdsdataservice.ExecuteStatementInput{
		ResourceArn: h.auroraArn,
		SecretArn:   h.secretArn,
		Sql:         aws.String(sql),
	}
	resp, err := h.client.ExecuteStatement(params)
	if err != nil {
		log.Printf("Error fetching races: %s", err)
		return nil, nil, err
	}

	categoryRaceMap := map[string]interface{}{}
	raceSummaries := map[string]interface{}{}
	data := []map[string]interface{}{}
	for _, record := range resp.Records {
		obj := map[string]interface{}{
			"race_id":          *record[RACE_ID].StringValue,
			"race_name":        *record[RACE_NAME].StringValue,
			"race_number":      *record[RACE_NUMBER].LongValue,
			"meeting_id":       *record[MEETING_ID].StringValue,
			"meeting_name":     *record[MEETING_NAME].StringValue,
			"category_id":      *record[CATEGORY_ID].StringValue,
			"advertised_start": map[string]string{"seconds": *record[ADVERTISED_START].StringValue},
		}

		log.Printf("%v", obj)
		data = append(data, obj)

		if _, ok := categoryRaceMap[*record[CATEGORY_ID].StringValue]; !ok {
			categoryRaceMap[*record[CATEGORY_ID].StringValue] = map[string]interface{}{
				"race_ids": []string{*record[RACE_ID].StringValue},
			}
		} else {
			categories := categoryRaceMap[*record[CATEGORY_ID].StringValue].(map[string]interface{})
			races := categories["race_ids"].([]string)
			races = append(races, *record[RACE_ID].StringValue)
			categories["race_ids"] = races
		}

		raceSummaries[*record[RACE_ID].StringValue] = obj
	}
	return categoryRaceMap, raceSummaries, nil
}

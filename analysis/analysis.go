package analysis

import (
	"errors"
	"fmt"
	"github.com/FindMyProfessors/backend/graph/model"
	"time"
)

const (
	OptimisticNumberOfChartPoints = 7
)

type Year struct {
	Months []*MonthReviews
}

// GetMonthReviews time.Month is an integer type where January = 1, February = 2, ... December = 12
func (y *Year) GetMonthReviews(month time.Month) *MonthReviews {
	// subtract 1 to 0 index month
	return y.Months[month-1]
}

func (y *Year) AddMonthReviews(month time.Month, reviews *MonthReviews) {
	y.Months[month-1] = reviews
}

type MonthReviews struct {
	Reviews       []*model.Review
	DifficultySum float64
	QualitySum    float64
}

func (r *MonthReviews) DifficultyAverage() float64 {
	return r.DifficultySum / float64(len(r.Reviews))
}

func (r *MonthReviews) QualitySumAverage() float64 {
	return r.QualitySum / float64(len(r.Reviews))
}

// BeginAnalysis reviews must be ordered descending by time
func BeginAnalysis(reviews []*model.Review) (*model.ProfessorAnalysis, error) {
	values, err := calculateChartValues(reviews)
	if err != nil {
		return nil, err
	}
	return &model.ProfessorAnalysis{
		TagAmount:           getTags(reviews),
		AverageRatingValues: values,
	}, nil
}

func getTags(reviews []*model.Review) []*model.TagAmount {
	tagMap := map[model.Tag]int{}

	// get all tags for all reviews and insert into map
	for _, review := range reviews {
		for _, tag := range review.Tags {
			if count, ok := tagMap[tag]; ok {
				tagMap[tag] = count + 1
			} else {
				tagMap[tag] = 1
			}
		}
	}

	// create a slice with capacity of tagMap
	tags := make([]*model.TagAmount, 0, len(tagMap))

	// insert map contents into slice
	for tag, count := range tagMap {
		tags = append(tags, &model.TagAmount{
			Tag:    tag,
			Amount: count,
		})
	}

	return tags
}

func calculateChartValues(reviews []*model.Review) (values []*model.ChartValue, err error) {
	years := map[int]Year{}

	for _, review := range reviews {
		t := review.Time
		year, ok := years[t.Year()]
		if !ok {
			year = Year{
				Months: make([]*MonthReviews, 12, 12),
			}
		}

		monthReviews := year.Months[t.Month()-1]
		if monthReviews == nil {
			monthReviews = &MonthReviews{
				Reviews: []*model.Review{},
			}
		}
		monthReviews.Reviews = append(monthReviews.Reviews, review)

		monthReviews.DifficultySum += review.Difficulty
		monthReviews.DifficultySum += review.Quality

		// TODO: check if this is necessary
		year.Months[t.Month()-1] = monthReviews
	}

	amountOfYears := len(years)
	if amountOfYears == 1 {
		var year Year
		var yearInt int
		// Only iterates once
		for i, y := range years {
			yearInt = i
			year = y
			break
		}
		return calculateOneYear(values, year, yearInt)
	}

	return calculate(reviews, OptimisticNumberOfChartPoints), nil
}

func calculateOneYear(values []*model.ChartValue, year Year, yearInt int) ([]*model.ChartValue, error) {
	months := year.Months

	amountOfMonths := len(months)
	if amountOfMonths < 5 {
		return nil, errors.New("this professor has too few reviews across time to provide a relevant chart")
	}

	// gets the most recent reviews
	var fixedMonth int
	if amountOfMonths > OptimisticNumberOfChartPoints {
		fixedMonth = OptimisticNumberOfChartPoints
	} else {
		fixedMonth = amountOfMonths
	}

	for i := fixedMonth; i >= 0; i++ {
		monthReviews := months[i]
		englishMonth := time.Month(i + 1).String()

		chartValue := &model.ChartValue{
			Value: monthReviews.QualitySum,
			Month: englishMonth,
			Year:  yearInt,
		}
		values = append(values, chartValue)
	}
	return values, nil
}

func calculate(reviews []*model.Review, numberOfPoints int) (chartValues []*model.ChartValue) {
	if len(reviews) == 0 {
		return nil
	}
	chunkSize := len(reviews) / numberOfPoints

	if chunkSize == 0 {
		for _, review := range reviews {
			chartValues = append(chartValues, &model.ChartValue{
				Value: review.Quality,
				Month: review.Time.Month().String(),
				Year:  review.Time.Year(),
			})
		}
		return chartValues
	}

	fmt.Printf("chunkSize=%d\n", chunkSize)

	completed := 0

	for numberOfPoints > 0 {
		currentSum := 0.0
		var timeMillisSum int64 = 0

		for i := 0; i < chunkSize; i++ {
			fmt.Printf("i=%d\n", i)
			review := reviews[completed]
			timeMillisSum += review.Time.UnixMilli()

			currentSum += review.Quality
			completed++
			fmt.Printf("completed=%d\n", completed)
		}

		average := currentSum / float64(chunkSize)

		t := time.UnixMilli(timeMillisSum / int64(chunkSize))
		chartValues = append(chartValues, &model.ChartValue{
			Value: average,
			Month: t.Month().String(),
			Year:  t.Year(),
		})

		numberOfPoints--
	}

	return chartValues
}

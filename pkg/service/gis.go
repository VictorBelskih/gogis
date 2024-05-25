package service

import (
	"math"
	"sort"

	"github.com/VictorBelskih/gogis"
	"github.com/VictorBelskih/gogis/pkg/repository"
)

// передача интерфейсов
type GisService struct {
	repo repository.Gis
}

// создание интерфейса интерфейсов
func NewGisService(repo repository.Gis) *GisService {
	return &GisService{repo: repo}
}

func (s *GisService) GetFieldByUser(id int, role int) (gogis.GeoJSON, error) {
	return s.repo.GetFieldByUser(id, role)
}

func (s *GisService) GetField() (gogis.GeoJSON, error) {
	return s.repo.GetField()
}

func (s *GisService) CreateFarm(farm gogis.Farm) error {
	return s.repo.CreateFarm(farm)
}

func (s *GisService) GetFarmByID(id int) (*gogis.Farm, error) {
	return s.repo.GetFarmByID(id)
}

func (s *GisService) DeleteFarm(id int) error {
	return s.repo.DeleteFarm(id)
}

func (s *GisService) UpdateFarm(farm gogis.Farm) error {
	return s.repo.UpdateFarm(farm)
}

func (s *GisService) GetFieldData(id int, role int) ([]gogis.Field, error) {
	return s.repo.GetFieldData(id, role)
}

func (s *GisService) GetCult() ([]gogis.Cult, error) {
	return s.repo.GetCult()
}

func (s *GisService) GetFarm() ([]gogis.Farm, error) {
	return s.repo.GetFarm()
}

func (s *GisService) GetDistrict() ([]gogis.District, error) {
	return s.repo.GetDistrict()
}

func (s *GisService) GetCultByID(id int) (*gogis.Cult, error) {
	return s.repo.GetCultByID(id)
}

func (s *GisService) UpdateCult(cult gogis.Cult) error {
	return s.repo.UpdateCult(cult)
}

func (s *GisService) CreateCult(cult gogis.Cult) error {
	return s.repo.CreateCult(cult)
}

func (s *GisService) DeleteCult(id int) error {
	return s.repo.DeleteCult(id)
}

func (s *GisService) CalculateTotalAreaByFieldType(id int, role int) (map[string]float64, error) {
	fields, err := s.repo.GetFieldData(id, role)
	if err != nil {
		return nil, err
	}

	totalAreaByFieldType := make(map[string]float64)
	totalArea := 0.0

	for _, field := range fields {
		totalArea += field.Area_f
		switch field.Tlu {
		case 102:
			totalAreaByFieldType["пашня"] += field.Area_f
		case 300:
			totalAreaByFieldType["сенокос"] += field.Area_f
		case 200:
			totalAreaByFieldType["пастбище"] += field.Area_f
		case 500:
			totalAreaByFieldType["залежь"] += field.Area_f
		}
	}

	for key, value := range totalAreaByFieldType {
		totalAreaByFieldType[key] = math.Round(value*100) / 100
	}

	return totalAreaByFieldType, nil
}

func (s *GisService) TotalArea(id int, role int) (float64, error) {
	fields, err := s.repo.GetFieldData(id, role)
	if err != nil {
		return 0, err
	}

	totalArea := 0.0

	for _, field := range fields {
		totalArea += field.Area_f
	}

	totalArea = math.RoundToEven(totalArea*100) / 100

	return totalArea, nil
}

type HumusData struct {
	ID           int
	Class        string
	AverageValue float64
	TotalArea    float64
	Percentage   float64 // Общая площадь для класса
}

func (s *GisService) CalculateAverageHumusByClass(id int, role int) ([]HumusData, error) {
	fields, err := s.repo.GetFieldData(id, role)
	if err != nil {
		return nil, err
	}

	humusData := make(map[string]float64)
	areaByClass := make(map[string]float64)
	totalArea := 0.0

	for _, field := range fields {
		humusData[field.Humus_class] += field.Organic * field.Area_f
		areaByClass[field.Humus_class] += field.Area_f
		totalArea += field.Area_f
	}

	var result []HumusData

	order := map[string]int{
		"Очень низкое":  1,
		"Низкое":        2,
		"Среднее":       3,
		"Повышенное":    4,
		"Высокое":       5,
		"Очень высокое": 6,
	}

	ids := 1
	for class, totalOrganic := range humusData {
		classArea := areaByClass[class]
		average := totalOrganic / classArea
		percentage := (classArea / totalArea) * 100
		result = append(result, HumusData{ID: ids, Class: class, AverageValue: average, TotalArea: classArea, Percentage: percentage})
		ids++
	}

	// Сортируем результат по порядку, определенному в order
	sort.Slice(result, func(i, j int) bool {
		return order[result[i].Class] < order[result[j].Class]
	})

	return result, nil
}

type RadionuclideSummary struct {
	AverageCesium    float64
	AverageStrontium float64
	MaxCesium        float64
	MaxStrontium     float64
	AvgDensityCs137  float64 // Средняя плотность загрязнения цезием
	AvgDensitySr90   float64 // Средняя плотность загрязнения стронцием
	MaxDensityCs137  float64 // Максимальная плотность загрязнения цезием
	MaxDensitySr90   float64 // Максимальная плотность загрязнения стронцием
}

func (s *GisService) CalculateRadionuclideSummary(id int, role int) (RadionuclideSummary, error) {
	fields, err := s.repo.GetFieldData(id, role)
	if err != nil {
		return RadionuclideSummary{}, err
	}

	var totalCesium, totalStrontium, maxCesium, maxStrontium float64
	var totalDensityCs137, totalDensitySr90, maxDensityCs137, maxDensitySr90 float64
	for _, field := range fields {
		totalCesium += field.S_cs137
		totalStrontium += field.S_sr90
		totalDensityCs137 += field.Cs137 // Предполагаем, что Cs137 - это плотность загрязнения цезием
		totalDensitySr90 += field.Sr90   // Предполагаем, что Sr90 - это плотность загрязнения стронцием

		if field.S_cs137 > maxCesium {
			maxCesium = field.S_cs137
		}
		if field.S_sr90 > maxStrontium {
			maxStrontium = field.S_sr90
		}
		if field.Cs137 > maxDensityCs137 {
			maxDensityCs137 = field.Cs137
		}
		if field.Sr90 > maxDensitySr90 {
			maxDensitySr90 = field.Sr90
		}
	}

	count := float64(len(fields))
	averageCesium := totalCesium / count
	averageStrontium := totalStrontium / count
	avgDensityCs137 := totalDensityCs137 / count // Рассчитываем среднюю плотность загрязнения цезием
	avgDensitySr90 := totalDensitySr90 / count   // Рассчитываем среднюю плотность загрязнения стронцием

	return RadionuclideSummary{
		AverageCesium:    averageCesium,
		AverageStrontium: averageStrontium,
		MaxCesium:        maxCesium,
		MaxStrontium:     maxStrontium,
		AvgDensityCs137:  avgDensityCs137,
		AvgDensitySr90:   avgDensitySr90,
		MaxDensityCs137:  maxDensityCs137,
		MaxDensitySr90:   maxDensitySr90,
	}, nil
}

type NutrientData struct {
	ID           int
	Class        string
	AverageValue float64
	TotalArea    float64
	Percentage   float64
}

func (s *GisService) AvgPotassiumByClass(id int, role int) ([]NutrientData, error) {
	fields, err := s.repo.GetFieldData(id, role)
	if err != nil {
		return nil, err
	}

	potassiumData := make(map[string]float64)
	areaByClass := make(map[string]float64)
	totalArea := 0.0

	for _, field := range fields {
		potassiumData[field.Class_k] += field.El_k * field.Area_f
		areaByClass[field.Class_k] += field.Area_f
		totalArea += field.Area_f
	}

	var result []NutrientData

	order := map[string]int{
		"Очень низкая":  1,
		"Низкая":        2,
		"Средняя":       3,
		"Повышенная":    4,
		"Высокая":       5,
		"Очень высокая": 6,
	}

	ids := 1
	for class, totalPotassium := range potassiumData {
		classArea := areaByClass[class]
		average := totalPotassium / classArea
		percentage := (classArea / totalArea) * 100
		result = append(result, NutrientData{ID: ids, Class: class, AverageValue: average, TotalArea: classArea, Percentage: percentage})
		ids++
	}

	// Сортируем результат по порядку, определенному в order
	sort.Slice(result, func(i, j int) bool {
		return order[result[i].Class] < order[result[j].Class]
	})

	return result, nil
}

func (s *GisService) AvgPhosphorByClass(id int, role int) ([]NutrientData, error) {
	fields, err := s.repo.GetFieldData(id, role)
	if err != nil {
		return nil, err
	}

	phosphorusData := make(map[string]float64)
	areaByClass := make(map[string]float64)
	totalArea := 0.0

	for _, field := range fields {
		phosphorusData[field.Class_p] += field.El_p * field.Area_f
		areaByClass[field.Class_p] += field.Area_f
		totalArea += field.Area_f
	}

	var result []NutrientData

	order := map[string]int{
		"Очень низкая":  1,
		"Низкая":        2,
		"Средняя":       3,
		"Повышенная":    4,
		"Высокая":       5,
		"Очень высокая": 6,
	}

	ids := 1
	for class, totalPhosphorus := range phosphorusData {
		classArea := areaByClass[class]
		average := totalPhosphorus / classArea
		percentage := (classArea / totalArea) * 100
		result = append(result, NutrientData{ID: ids, Class: class, AverageValue: average, TotalArea: classArea, Percentage: percentage})
		ids++
	}

	// Сортируем результат по порядку, определенному в order
	sort.Slice(result, func(i, j int) bool {
		return order[result[i].Class] < order[result[j].Class]
	})

	return result, nil
}

// func (s *GisService) CalculateAverageHumusByClass() (map[string]float64, error) {
// 	fields, err := s.repo.GetFieldData()
// 	if err != nil {
// 		return nil, err
// 	}

// 	totalHumusByClass := make(map[string]float64)
// 	countByClass := make(map[string]int)

// 	for _, field := range fields {
// 		totalHumusByClass[field.Humus_class] += field.Organic
// 		countByClass[field.Humus_class]++
// 	}

// 	averageHumusByClass := make(map[string]float64)
// 	for class, total := range totalHumusByClass {
// 		count := float64(countByClass[class])
// 		averageHumusByClass[class] = total / count
// 	}

// 	return averageHumusByClass, nil
// }

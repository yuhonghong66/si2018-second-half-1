package entities

import (
	"math"
	"strconv"
	"time"

	"github.com/eure/si2018-second-half-1/models"
	"github.com/go-openapi/strfmt"
)

type UserStats struct {
	UserID          int64           `xorm:"user_id"`
	Birthday        float64         `xorm:"birthday"`
	HomeStateX      float64         `xorm:"home_state_x"`
	HomeStateY      float64         `xorm:"home_state_y"`
	ResidenceStateX float64         `xorm:"residence_state_x"`
	ResidenceStateY float64         `xorm:"residence_state_y"`
	Education       float64         `xorm:"education"`
	AnnualIncome    float64         `xorm:"annual_income"`
	Height          float64         `xorm:"height"`
	BodyBuild       float64         `xorm:"body_build"`
	Smoking         float64         `xorm:"Smoking"`
	Drinking        float64         `xorm:"drinking"`
	HolidayWeekday  float64         `xorm:"holiday_weekday"`
	HolidayWeekend  float64         `xorm:"holiday_weekend"`
	HolidayRandom   float64         `xorm:"holiday_random"`
	HolidayOthers   float64         `xorm:"holiday_others"`
	JobDoctor       float64         `xorm:"job_doctor"`
	JobOffice       float64         `xorm:"job_office"`
	JobClerk        float64         `xorm:"job_clerk"`
	CreatedAt       strfmt.DateTime `xorm:"created_at"`
	UpdatedAt       strfmt.DateTime `xorm:"updated_at"`
}

type Coordinate struct {
	Latitude  float64 // 緯度
	Longitude float64 // 経度
}

var DrinkingChoices = map[string]float64{"飲まない": 0, "ときどき飲む": 1.0 / 2, "飲む": 2.0 / 2}
var EducationChoices = map[string]float64{"その他": 0, "高校卒": 1.0 / 4, "短大/専門学校": 2.0 / 4, "大学卒": 3.0 / 4, "大学院卒": 4.0 / 4}
var BodyBuildChoices = map[string]float64{"スリム": 0, "やや細め": 1.0 / 6, "普通": 2.0 / 6, "グラマー": 3.0 / 6, "筋肉質": 4.0 / 6, "ややぽっちゃり": 5.0 / 6, "ぽっちゃり": 6.0 / 6}
var SmokingChoices = map[string]float64{"吸わない": 0, "非喫煙者の前では吸わない": 1.0 / 5, "相手が嫌ならやめる": 2.0 / 5, "ときどき吸う": 3.0 / 5, "吸う（電子タバコ）": 4.0 / 5, "吸う": 5.0 / 5}
var CoordinateMap = map[string]Coordinate{
	"北海道": {43.06417, 141.34694},
	"青森":  {40.82444, 140.74},
	"岩手":  {39.70361, 141.1525},
	"宮城":  {38.26889, 140.87194},
	"秋田":  {39.71861, 140.1025},
	"山形":  {38.24056, 140.36333},
	"福島":  {37.75, 140.46778},
	"茨城":  {36.34139, 140.44667},
	"栃木":  {36.56583, 139.88361},
	"群馬":  {36.39111, 139.06083},
	"埼玉":  {35.85694, 139.64889},
	"千葉":  {35.60472, 140.12333},
	"東京":  {35.68944, 139.69167},
	"神奈川": {35.44778, 139.6425},
	"新潟":  {37.90222, 139.02361},
	"富山":  {36.69528, 137.21139},
	"石川":  {36.59444, 136.62556},
	"福井":  {36.06528, 136.22194},
	"山梨":  {35.66389, 138.56833},
	"長野":  {36.65139, 138.18111},
	"岐阜":  {35.39111, 136.72222},
	"静岡":  {34.97694, 138.38306},
	"愛知":  {35.18028, 136.90667},
	"三重":  {34.73028, 136.50861},
	"滋賀":  {35.00444, 135.86833},
	"京都":  {35.02139, 135.75556},
	"大阪":  {34.68639, 135.52},
	"兵庫":  {34.69139, 135.18306},
	"奈良":  {34.68528, 135.83278},
	"和歌山": {34.22611, 135.1675},
	"鳥取":  {35.50361, 134.23833},
	"島根":  {35.47222, 133.05056},
	"岡山":  {34.66167, 133.935},
	"広島":  {34.39639, 132.45944},
	"山口":  {34.18583, 131.47139},
	"徳島":  {34.06583, 134.55944},
	"香川":  {34.34028, 134.04333},
	"愛媛":  {33.84167, 132.76611},
	"高知":  {33.55972, 133.53111},
	"福岡":  {33.60639, 130.41806},
	"佐賀":  {33.24944, 130.29889},
	"長崎":  {32.74472, 129.87361},
	"熊本":  {32.78972, 130.74167},
	"大分":  {33.23806, 131.6125},
	"宮崎":  {31.91111, 131.42389},
	"鹿児島": {31.56028, 130.55806},
	"沖縄":  {26.2125, 127.68111},
}

func getNearChoices(average float64, choices map[string]float64) []string {
	return []string{}
}

type Range struct {
	From int
	To   int
}

func getRoundedRange(average float64, lower, upper, unit int64) Range {
	return Range{From: 0, To: 0}
}

func getNearHeight(average float64) models.IdealTypeHeight {
	r := getRoundedRange(average, 135, 200, 5)
	return models.IdealTypeHeight{From: strconv.Itoa(r.From) + "cm", To: strconv.Itoa(r.To) + "cm"}
}

func getNearState(x, y float64) []string {
	return []string{"東京"}
}

func round(f float64) float64 {
	return math.Floor(f + .5)
}

func getNearAnnualIncome(average float64) models.IdealTypeAnnualIncome {
	return models.IdealTypeAnnualIncome{From: "500万円", To: "500万円"}
}

func getNearAge(average float64) models.IdealTypeAge {
	now := time.Now()
	birth := time.Unix(int64(round(average)), 0)
	span := now.Sub(birth)
	years := span.Hours() / (24 * 365)
	r := getRoundedRange(years, 18, 65, 1)
	return models.IdealTypeAge{From: strconv.Itoa(r.From) + "歳", To: strconv.Itoa(r.To) + "歳"}
}

func getModeJob([]float64) []string {
	return []string{"会社員"}
}

func getModeHoliday([]float64) []string {
	return []string{"土日"}
}

func (u UserStats) Build() models.IdealType {
	age := getNearAge(u.Birthday)
	height := getNearHeight(u.Height)
	income := getNearAnnualIncome(u.AnnualIncome)
	return models.IdealType{
		Drinking:       getNearChoices(u.Drinking, DrinkingChoices),
		Education:      getNearChoices(u.Education, EducationChoices),
		BodyBuild:      getNearChoices(u.BodyBuild, BodyBuildChoices),
		Smoking:        getNearChoices(u.Smoking, SmokingChoices),
		Age:            &age,
		Height:         &height,
		AnnualIncome:   &income,
		HomeState:      getNearState(u.HomeStateX, u.HomeStateY),
		ResidenceState: getNearState(u.ResidenceStateX, u.ResidenceStateY),
		Job:            getModeJob([]float64{u.JobClerk, u.JobDoctor, u.JobOffice}),
		Holiday:        getModeHoliday([]float64{u.HolidayWeekend, u.HolidayWeekday, u.HolidayRandom, u.HolidayOthers})}
}

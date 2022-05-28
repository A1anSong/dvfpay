package dvfpay

type Statistics struct {
	Value int    `json:"value"`
	Name  string `json:"name"`
}
type Statisticians []Statistics

func (s Statisticians) Len() int { return len(s) }
func (s Statisticians) Less(i, j int) bool {
	statisticiansMap := map[string]int{
		"成功": 0,
		"待定": 1,
		"结算": 2,
	}
	return statisticiansMap[s[i].Name] < statisticiansMap[s[j].Name]
}
func (s Statisticians) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

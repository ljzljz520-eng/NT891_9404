package catalog

import "traveldeck/domain"

func Seed() []domain.Destination {
	return []domain.Destination{
		{"d001", "Lijiang", "Yunnan", "Ancient canals beneath Jade Dragon Snow Mountain", "spring", domain.Mountain, true},
		{"d002", "Xiamen", "Fujian", "Island lanes and sea breeze", "autumn", domain.Coastal, true},
		{"d003", "Chengdu", "Sichuan", "Tea houses and slow city mornings", "winter", domain.City, true},
		{"d004", "Wuyuan", "Jiangxi", "Golden fields and white-walled villages", "spring", domain.Countryside, true},
		{"d005", "Dali", "Yunnan", "Erhai cycling with Cangshan views", "summer", domain.Coastal, true},
		{"d006", "Zhangjiajie", "Hunan", "Quartz peaks in a sea of clouds", "autumn", domain.Mountain, true},
		{"d007", "Hangzhou", "Zhejiang", "West Lake gardens and tea hills", "spring", domain.City, true},
		{"d008", "Moganshan", "Zhejiang", "Bamboo forests and quiet cabins", "summer", domain.Countryside, true},
	}
}

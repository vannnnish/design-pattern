package main

import (
	"fmt"
	"pattern/structural/filter/filter"
)

func main() {
	rub := make([]filter.Rubbish, 0)
	rub = append(rub, filter.Rubbish{
		Name:       "果壳",
		IsHarm:     false,
		IsRecycled: false,
		IsDry:      true,
		IsWet:      false,
	})
	rub = append(rub, filter.Rubbish{"陶瓷", false, false, true, false})
	rub = append(rub, filter.Rubbish{"菜根菜叶", false, false, false, true})
	rub = append(rub, filter.Rubbish{"果皮", false, false, false, true})
	rub = append(rub, filter.Rubbish{"水银温度计", true, false, false, false})
	rub = append(rub, filter.Rubbish{"电池", true, false, false, false})
	rub = append(rub, filter.Rubbish{"灯泡", true, false, false, false})
	rub = append(rub, filter.Rubbish{"废纸塑料", false, true, false, false})
	rub = append(rub, filter.Rubbish{"金属和布料", false, true, false, false})
	rub = append(rub, filter.Rubbish{"玻璃", false, true, false, false})

	dryFilter := filter.DryRubbishCriteria{}
	wetFilter := filter.WetRubbishCriteria{}
	harmFilter := filter.HarmfulRubbishCriteria{}
	recyFilter := filter.RecycledRubbishCriteria{}
	// 打印四种过滤器过滤的结果
	fmt.Println(dryFilter.RubbishFilter(rub))
	fmt.Println(wetFilter.RubbishFilter(rub))
	fmt.Println(harmFilter.RubbishFilter(rub))
	fmt.Println(recyFilter.RubbishFilter(rub))
}

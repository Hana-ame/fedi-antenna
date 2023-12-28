package utils

func MergeMaps(maps ...map[string]string) map[string]string {

	// delete nils
	for p0, p1 := 0, 0; ; p1++ {
		if p1 >= len(maps) {
			maps = maps[:p0]
			break
		}
		if maps[p0] == nil {
			maps[p0], maps[p1] = maps[p1], maps[p0]
		}
		if maps[p0] != nil {
			p0++
		}
	}

	if len(maps) == 0 {
		return nil
	}

	if len(maps) == 1 {
		return maps[0]
	}

	for i, m := range maps {
		if i == 0 {
			continue
		}
		for k, v := range m {
			maps[0][k] = v
		}
	}

	return maps[0]
}

package array_string

func CanPlaceFlowers(flowerbed []int, n int) bool {
	for i := range flowerbed {
		if flowerbed[i] == 0 {
			if i == 0 || flowerbed[i-1] == 0 {
				if i == (len(flowerbed)-1) || flowerbed[i+1] == 0 {
					flowerbed[i] = 1
					n--
					if n == 0 {
						return true
					}

				}
			}
		}
	}

	return false

}

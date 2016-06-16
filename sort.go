package utils

func SortMapListById(id string, mapList []map[string]string) {
    length := len(mapList)
    for i := 0; i < length; i++ {
        for j := i + 1; j < length; j++ {
            if mapList[i][id] > mapList[j][id] {
                tmp := mapList[i]
                mapList[i] = mapList[j]
                mapList[j] = tmp
            }
        }
    }
}

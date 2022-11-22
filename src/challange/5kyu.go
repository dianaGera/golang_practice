package challange


type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func PickPeaks(array []int) PosPeaks {
	// Example: pickPeaks([3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3]) should return 
	// {pos: [3, 7], peaks: [6, 3]} (or equivalent in other languages)
	// {1, 2, 2, 2, 1} --> {Pos: {1}, Peaks: {2}}
 	res := PosPeaks{}
	heap := [4]int{0, 0, 0, 0}
	if len(array) > 1 {
		for i, v := range array[1:len(array)-1] {
			if heap[1] != array[i] {
				heap[0] = array[i]
			} 
			if array[i+2] == v {
				heap[2] += 1
			} else {
				heap[3] = heap[2]
				heap[2] = 0
			}
			
			heap[1] = array[i+2]
			
			if heap[0] < v && v > heap[1] {
				res.Pos = append(res.Pos, i-heap[3]+1)
				res.Peaks = append(res.Peaks, v) 
			}
			
		}
	}
	return res
}


func DirReduc(arr []string) []string {
	// improve direction
	// In ["NORTH", "SOUTH", "EAST", "WEST"], the direction "NORTH" + "SOUTH" 
	// is going north and coming back right away.
	// The path becomes ["EAST", "WEST"], now "EAST" and "WEST" annihilate each other, 
	// therefore, the final result is [] (nil in Clojure).
	dir := map[string]string{
	 "NORTH": "SOUTH",
	 "SOUTH": "NORTH",
	 "WEST": "EAST",
	 "EAST": "WEST",
   }
   res := []string{"START"}
   for _, v := range arr {
	 if res[len(res)-1] == dir[v] {
	   res = res[0:len(res)-1]
	 } else {
	   res = append(res, v)
	 }
   }
   return res[1:]
 }



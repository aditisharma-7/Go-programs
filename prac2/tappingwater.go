package main ()

var arr={3,0,0,2,0,4}
var maxHeight int

func min(x,y int)int {             //to calculate min height because water wont rise above min height
       if (x<y){
		   return x
	   }
	   return y
}

func volume(h int) int{
	return h //h*1*1

}

func main(){
     maxHeight= min(arr[0],arr[len(arr)-1])
	 for index,value :=range arr{
		 if index >0 && index <len(arr){
			 heightleft := arr[index-1]
			 heightright :=arr[index +1]

			if heightleft > maxHeight{
				heightleft =maxHeight     
			}
			if heightright > maxHeight{
				heightright =maxHeight     
			}
			watercaptured += volume(min(heightleft,heightright))
		 }
	 }
}
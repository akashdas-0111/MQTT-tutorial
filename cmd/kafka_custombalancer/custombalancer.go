package main
import(
	"akash-mqtttut/internal/balancer"
	"fmt"
)
func main(){
	balancer.Balancee("kafkatest",5)
	fmt.Println("done")
}
 
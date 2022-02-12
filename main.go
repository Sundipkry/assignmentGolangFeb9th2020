package main

import (
    "fmt"
    "strings"
    )

func getdish(code string) (dish string, rate int){

	switch code{
	    
	    case "C":
	    return "Cofee", 40
	    
	    case "D":
	    return "Dosa", 80
	    
	    case "T":
	    return "Tomato Soup", 20
	    
	    case "I":
	    return "Idli", 20
	    
	    case "V":
	    return "Vada", 25
	    
	    case "B":
	    return "Bhature & Chhole", 30
	    
	    case "P":
	    return "Paneer Pakoda", 80
	    
	    case "M":
	    return "Manchurian", 90
	    
	    case "H":
	    return "Hakka Noodle", 70
	    
	    case "F":
	    return "French Fries", 30
	    
	    case "J":
	    return "Jalebi", 30
	    
	    case "L":
	    return "Lemonade", 15
	    
	    case "S":
	    return "Spring roll", 20
	    
	    default:
	    return "nil", 0
	}
}



func main(){
    
    var str string
    var total int
    
    for str!="END"{
        
        var tOrder []string
        var tQuant []int
        sum:=0
        
        fmt.Println("Please enter the order")
        var order string
        var quantity int
        for order !="NEXT"{
            
            fmt.Scanln(&quantity, &order)
            order=strings.ToUpper(order)
            
            if order=="NEXT"{
                
                for i:=0; i<len(tOrder); i++{
                    
                    dish,rate:=getdish(tOrder[i])
                    tPrice:= rate * tQuant[i]
                    fmt.Println("Your order:", dish , "Rate:", rate ,"Quantity:", tQuant[i], "Total:", tPrice)
                    sum=sum+tPrice
                    }
                    break
                }
                
            tQuant=append(tQuant,quantity)
            tOrder=append(tOrder,order)
            fmt.Println("Enter '0 NEXT' for next customer or write the order to continue for same customer")
            
        }
        fmt.Println("Your total is:", sum)
        total=total+sum
        
        fmt.Println("Enter End to end the day or anything to continue to next customer")
        fmt.Scanln(&str)
        str=strings.ToUpper(str)
        if str=="END"{
            
            fmt.Println("Your total sale for the day is", total)
            break
        }
    }
}

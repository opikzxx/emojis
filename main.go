package main 
  
 import ( 
 	"context" 
 	"fmt" 
  
 	"github.com/ServiceWeaver/weaver" 
 ) 
  
 func main() { 
 	if err := weaver.Run(context.Background(), run); err != nil { 
 		panic(err) 
 	} 
 } 
  
 // app is the main component of our application. 
 type app struct { 
 	weaver.Implements[weaver.Main] 
 } 
  
 // run implements the application main. 
 func run(ctx context.Context, a *app) error { 
 	fmt.Println("Hello, World!") 
 	return nil 
 } 
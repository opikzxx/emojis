package main 
  
 import ( 
 	"context" 
 	"sort" 
 	"strings" 
  
 	"github.com/ServiceWeaver/weaver" 
 	"golang.org/x/exp/slices" 
 ) 
  
 type Searcher interface { 

 	Search(context.Context, string) ([]string, error) 
 } 
  
 type searcher struct { 
 	weaver.Implements[Searcher] 
 } 
  
 func (s *searcher) Search(ctx context.Context, query string) ([]string, error) { 

 	words := strings.Fields(strings.ToLower(query)) 
 	var results []string 
 	for emoji, labels := range emojis { 
 		if matches(labels, words) { 
 			results = append(results, emoji) 
 		} 
 	} 
 	sort.Strings(results) 
 	return results, nil 
 } 
  
 func matches(labels, words []string) bool { 
 	for _, word := range words { 
 		if !slices.Contains(labels, word) { 
 			return false 
 		} 
 	} 
 	return true 
 } 
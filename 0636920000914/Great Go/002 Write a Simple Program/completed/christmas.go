package main

import "fmt"

func main() {
  days := [12]string{ "1st", "2nd", "3rd", "4th", "5th", "6th",
                      "7th", "8th", "9th", "10th", "11th", "12th" };
  gifts := [12]string{ "partridge in a pear tree",
                       "turtle doves",
                       "French hens",
                       "calling birds",
                       "golden rings",
                       "geese a-laying",
                       "swans a-swimming",
                       "maids a-milking",
                       "ladies dancing",
                       "lords a-leaping",
                       "pipers piping",
                       "drummers drumming" };

  for i := 0; i < 12; i++ {
    fmt.Printf("On the %s day of Christmas, my true love gave to me\n",
               days[i]);
    if (i == 0) {
      fmt.Printf("\tA %s.\n", gifts[i]);
    } else {
      for j := i; j > 0; j-- {
        fmt.Printf("\t%d %s,\n", j+1, gifts[j]);
      }
      fmt.Printf("\tAnd a %s.\n", gifts[0]);
    }
  }
  
}
package main

import "fmt"

func main() {
    x, y, z, n := 1, 2, 6, 32
    FindSolution(x, y, z, n)
}

func FindSolution(x int, y int , z int, n int ) {
    total := x * 10 + y * 5 + z * 2
    if total < n {
        fmt.Println("不能刚好凑齐", n, "元")
    } else if total == n {
        fmt.Println("需要", x, "张十元纸币，", y, "张五元纸币，", z, "张两元纸币，", "刚好可凑齐", n, "元")
    } else {
        if n < 2 {
            fmt.Println("不能刚好凑齐", n, "元")
        } else if n >= 2 && n < 5 {
            if n % 2 == 0 && n/2 <= z {
                fmt.Println("需要0张十元纸币，", "0张五元纸币，", n/2, "张两元纸币，", "刚好可凑齐", n, "元")
            }
        } else if n >= 5 && n < 10 {
            if n % 5 == 0 && n/5 <= y {
                fmt.Println("需要0张十元纸币，", n/5, "张五元纸币，", "0张两元纸币，", "刚好可凑齐", n, "元")
            } else if n % 2 == 0 && n/2 <= z {
                fmt.Println("需要0张十元纸币，", "0张五元纸币，", n/2, "张两元纸币，", "刚好可凑齐", n, "元")
            } else if (n-5) % 2 == 0 && y >= 1 && (n-5)/2 <= z {
                fmt.Println("需要0张十元纸币，", "1张五元纸币，", (n-5)/2, "张两元纸币，", "刚好可凑齐", n, "元")
            }
        } else if n >= 10 {
            if n % 10 == 0 {
                if n <= 10*x {
                    fmt.Println("需要", n/10, "张十元纸币，", "0张五元纸币，", "0张两元纸币，", "刚好可凑齐", n, "元")
                } else if n > 10*x && n <= 10*x + 5*y {
                    fmt.Println("需要", x, "张十元纸币，", (n-10*x)/5, "张五元纸币，", "0张两元纸币，", "刚好可凑齐", n, "元")
                } else {
                    fmt.Println("需要", x, "张十元纸币，", y, "张五元纸币，", (n-10*x-5*y)/2, "张两元纸币，", "刚好可凑齐", n, "元")
                }
            } else if n % 5 == 0 && y >= 1 {
                if n/10 <= x {
                    fmt.Println("需要", n/10, "张十元纸币，", "1张五元纸币，", "0张两元纸币，", "刚好可凑齐", n, "元")
                } else if n <= 10*x + 5*y {
                    fmt.Println("需要", x, "张十元纸币，", (n-10*x)/5, "张五元纸币，", "0张两元纸币，", "刚好可凑齐", n, "元")
                }
            } else if (n % 5) % 2 == 0 && z >= 1 {
                if n/10 <= x && (n % 10)/5 <= y && (n % 5)/2 <= z {
                    fmt.Println("需要", n/10, "张十元纸币，", (n % 10)/5, "张五元纸币，", (n % 5)/2, "张两元纸币，", "刚好可凑齐", n, "元")
                } else if n/10 > x && (n - 10*x)/5 <= y && n%5/2 <= z {
                    fmt.Println("需要", x, "张十元纸币，", (n-10*x)/5, "张五元纸币，", (n % 5)/2, "张两元纸币，", "刚好可凑齐", n, "元")
                } else if n/10 > x && (n - 10*x)/5 > y {
                    fmt.Println("需要", x, "张十元纸币，", y, "张五元纸币，", (n - 10*x - 5*y)/2, "张两元纸币，", "刚好可凑齐", n, "元")
                }
            }
        } else {
            fmt.Println("不能刚好凑齐", n, "元")
        }
    }
}

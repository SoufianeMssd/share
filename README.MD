# Share CMD App

## Introduction
`Share` app is command-line based app used to split the bill  
So basically :
* You add an event
* You add participant(ppl you want to split the bill with)
* Finally, You split the bill of the event
> *There is a help(share -h) flag in the app in order to know your options*

## How to Run
* **Clone the repo**  
    ```sh  
    git clone https://github.com/SoufianeMssd/share.git
    ```
* **Build the project**  
    ```go
    go mod tidy
    ```  
    ```go
    go build
    ```
* **Change directory**  
    ```sh
    cd share
    ```
* **Execute**  
    ```sh
    ./share
    ```
    > Or add exec *'share'* to your path & execute it anywhere

* **Ouput**  

    <span style="color:green">Event : Birthday</span>
    | Name                   | Paid                     | CreatedAt             |
    |------------------------| ------------------------ |----------------------:|
    | Alex                   | 500                      | Jan  8 18:02:30       |
    | Norman                 | 300                      | Mar  8 18:02:30       |
    | Mohammed               | 20                       | Feb  8 18:02:30       |

    > *Ouput is assuming that you created an event `Birthday` & added 3 Participants*
## About Go  
I'm still beginner at Go & this is my first mini side project, done mainly because me & friends always struggle when trying to split the bill
```go
package main

func main() {
    fmt.println("Hello World!")
}
```
* **Ouput**  
    <span style="color:gray">Hello World!</span>
### Links
* [Official Doc](https://go.dev/)
* [Golang for Beginners](https://www.youtube.com/watch?v=YS4e4q9oBaU "FreeCodeCamp Course")
* [Golang build projects](https://www.youtube.com/watch?v=jFfo23yIWac "FreeCodeCamp Course")
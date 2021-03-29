# go-snake
🐍 A cli snake game in Go

## How I created it
It took me quite some time to figure out how I should do it.

The one thing bothering me was `gamestate` and how should I handle it.
At first I've created a **Live** like snake game which move on its own:

- The snake goes to the same direction unless we change its direction.
But using a 2D array for `gamestate` has no use case here. (IMO)
  
So I've changed the way how to code it (current branch `master`).

In my interpretation,
- `gamestate` as it is a 2D array
- So creating a game based on the 2D array will be very interesting
- So each item in the array is `tile` 
- So whenever I want to persist something in the `gamestate`,
- I just create a tile for it.
    - See [Cell](https://github.com/theArtechnology/go-snake/blob/master/src/game/game.go#L18)
    
- Using the Cell, we can construct the entire state of the game with numbers
 - Since it is a flat surface (2D), we can't have the same item at the same position
- The move of the user should be decided by the user, so the best case scenario is to make it sequential
 and wait for the user input.
## How long it took
- It took me about 10 hours, 5 hours in the wrong direction see [wip](https://github.com/theArtechnology/go-snake/tree/wip) branch, and 5 additional hours for this version.

## How to run it
Simply use the Makefile, or run `go run ./cmd/main.go --size N` where N is the size of the Board (Width and Height)




# Battleship Ideal Strategy Simulator

A Go module for playing the ideal Battleship strategy.

Originally Inspired by The Battleship Algorithm video from Vsauce2 on YouTube: <https://www.youtube.com/watch?v=LbALFZoRrw8>

Further inspired by the original analysis of Battleship by DataGenetics here: <http://www.datagenetics.com/blog/december32011/>

Also inspired by C. Liam Brown's implementation and methodology here: <https://cliambrown.com/battleship/>

## Introduction

This module is a pet project I am making for the purposes of learning Go, and having a bit of fun in the process. I was inspired to try my own hand at making the algorithm mentioned in the aforementioned YouTube video, which pointed me to the above DataGenetics page, which details various analyses of Battleship strategies, and attempts to come up with an ideal solution.

As I started reading a bit more into this, I came across an implementation of this algorithm available online from C. Liam Brown, who developed a PHP-based algorithm for ideal gameplay. After reading into his analysis a bit, I decided I wanted to try my own implementation of this algorithm.

## My Implmentation

My implementation of the Battleship algorithm has multiple goals:

* Provide meaningful practice to assist me further learning Go
* Develop the application myself without using code/pseudocode from others
* Develop a feature-rich command-line program for playing Battleship with others
* Develop a simulator for testing the competitiveness of the algorithm
* Test the algorithm against different strategies for placing ships
* Utilize TDD (test-driven development) and complete a full test suite
* Use my program to trick some friends as I play Battleship online

The application and the underlying packages will be purposefully over-engineered, as I wish to get good practice in developing a full Go module, complete with packaging, proper organization, documentation, and tests. However, the overall goal is to have fun and to challenge myself in completing this idea. I also intend to play a few games of Battleship with my command-line tool to fool some unknowing friends into losing some games rather quickly. :-)

### Rule Assumptions

The following assumptions are made for the rules of gameplay:

* The gameplay involves the standard Milton-Bradley rules
* 5 pieces are used, one length 5, one length 4, two length 3 and one length 2
* Only one shot is allowed per turn per player (no salvo rules)
* During a shot, the defending player must announce a hit or a miss
* If a ship is sunk, the defending player must announce which ship was sunk
* Gameplay alternates until one player loses all 5 ships

### Approach

The approach I am taking to implement this algorithm is to have two different modes the algorithm will use to track the current board state, `seek` mode and `destroy` mode. It will use `seek` mode to hunt out individual ships on the game board, and switch to `destroy` mode any time there are hits on the board that are not part of a sunk ship.

#### Game Board

The algorithm will keep track of the game board which will act as the source of truth during gameplay. At the start of the game, a Go slice (a flexible array) will be generated for each game piece length (2-5 squares) and tally up all possible locations for a ship of the given length. This will be used to generate heatmaps for `seek` mode.

#### Seek Mode

For each turn in `seek` mode, the algorithm will total up the number of possible ship locations for each ship still in the game and generate a total of total potential ship orientations per square. This will be generated into a heatmap, which will be used to recommend the next square to choose. The highest score in the heatmap (or any ties) will be recommended for gameplay. During gameplay, misses will rule out possible squares for every ship type. When a miss is found, all possibilities in each slice will be removed that includes the chosen square. The next turn will then recalculate the heat map based on what possibilities are left.

This may seem inefficient, but each turn will only be adding up a potential 600 ship locations into the heatmap, rather than calculating every potential location based on what squares are still available. In addition, populating the heatmap only requires a single pass of each slice, and will result in a total possible 1,880 additions to the board each turn, fewer as squares are picked and ruled out and as ships are sunk. Also, since the Destroyer and Submarine both have three squares, adding values to the board is the same computational cost, as it only needs to multiply the 3 square slice by 2 if both ships are in play.

#### Destroy Mode

When a hit is detected, the game will then go into `destroy` mode. Hits will be added to the game board as a `generic` hit. Each turn will then find the highest heatmap score for each square adjacent to any hit square until a sink is detected. Once a ship is sunk and the ship type is declared, the algorithm will determine the orientation of the ship, update the game squares to reflect that exact ship that was sunk (e.g. a Destroyer), and remove all of those squares from each slice still in use. For example, if the Cruiser is sunk, the 5-square slice will be discarded and no longer used.

After a ship is sunk, the algorithm will see if any unaccounted `generic` hit squares are still pending. If so, it will resume `destroy` mode as before, otherwise it will continue with `seek` gameplay as before. This will repeat until all five ships are sunk, and the game is won.

## Usage

Currently, the application is still in the prelimiary development phase. However, I am adding unit tests for every module I create, so `go test` can be run inside each package to see if unit tests are passing.

# Concepts

- *Board*: holds state of the game describes what you're looking at
- *Tile*: described by a letter and a score and occupies the board, hand or bag
- *Placement*: describes a tile and it's position on the board
- *Word*: tiles arranged left to right, or up to down on the board that are found in the dictionary
- *Disputed word*: A word that players think should or should not be in the game
- *Wildcard tile*: blank tile with no score which can count towards a word
- *Bag*: Distributes tiles to players
- *Turn*: Contains placements and increments score by summing tile scores of each word found against those placements
- *Turn history*: Details about turns made by a player, contains words and score gained from each word
- *Hand*: 7 tiles held by the player to be used in a turn, replenished at the end of each turn from the bag
- *Player*: Entity which drives turns in the game
- *Player score*: A sum of scores made from turns, used to judge a winner
- *Game*: Contains 2 or more players
- *Game over*: Decrements score by tile scores left over in players hands, happens when players give up

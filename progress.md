## Rules
* All state of the game should live on-chain. State includes open games, games currently in progress and completed games.

* Any user can submit a transaction to the network to invite others to start a game (i.e. create an open game).

* Other users may submit transactions to accept invitations. When an invitation is accepted, the game starts.

* The roles of “X” and “O” are decided as follows. The user's public keys are concatenated and the result is hashed. If the first bit of the output is 0, then the game's initiator (whoever posted the invitation) plays "O" and the second player plays "X" and vice versa. “X” has the first move.

* Both users submit transactions to the network to make their moves until the game is complete.

* The game needs to support multiple concurrent games sessions/players.

## ToDo

### Basics
- [x] Scaffold new blockchain
- [x] Create game logic
  - [x] Copy from https://golangbyexample.com/tic-tac-toe-program-golang/
  - [ ] Adjust game logic to my needs

### Actions
- [x] Create-Game
  - [x] Player invites other player
  - [x] Create new game with status `open`
  - [x] Errors: 
    - [x] Invalid player address
- [x] Accept-Invite
  - [x] Invited player accepts game invitation
  - [x] Game status changes to `in-progress`
  - [x] Errors:
    - [x] Game does not exist
    - [x] Game is not open
    - [x] Player is not invited
- [ ] Decline-Invite <b>(not necessary for this exercise)</b>
  - [ ] Invited player declines game invitation
  - [ ] Game status changes to `ended`
  - [ ] Errors:
    - [ ] Game does not exist
    - [ ] Game is not open
    - [ ] Player is not invited
- [x] Play-Turn
  - [x] If winning move, finish game with status `ended` and set winner to current player
  - [x] If draw, finish game with status `ended` and set winner to `draw`
  - [x] Errors:
    - [x] Game does not exist
    - [x] Game is not in progress
    - [x] Player is not allowed to play
    - [x] Invalid move

### Queries
- [x] List-Games
  - [x] Return all games
  - [ ] Filter for game status
  - [ ] Filter for player
- [x] Show-Game
  - [x] Show info of single game

### Storage of information
- [x] Store information about all games in `SystemInfo`
  ```go
  nextGameId uint32
  ```
- [x] Store game information in `StoredGame`
    ```go
    id uint32
    status GameStatus   // open, in-progress, ended
    player1 string
    player2 string
    turn uint32         // 1 or 2
    winner WinnerStatus // none, player1, player2, draw
    board [3][3]byte   // [...],[xxx],[ooo]
    ```

### Other considerations
- How do I keep track of all the games?</br>
I can use a counter that increments after a new game has been created.
- What happens if a user stops to play?</br>
There could be a time limit for to play a turn. 
If a player does not play in given timeframe he loses.</br>
<b>(not necessary for this exercise)</b>
- What happens if an invited user never accepts?</br>
There could also be a time limit for accepting an invitation.</br>
<b>(not necessary for this exercise)</b>
- There could be a query to check if a move can be played.</br>
<b>(not necessary for this exercise)</b>
- Should a player be able to abandon a game?</br>
<b>(not necessary for this exercise)</b>

### Testing
- [ ] Test game logic
- [ ] Integration tests
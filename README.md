# Breakout Game in Go

A classic Breakout game implementation in Go using the Ebitengine game library. Break all the bricks to win, but don't let the ball fall below your paddle!

## Game Features

- Rainbow-colored bricks with depth-based scoring
- Smooth paddle and ball movement
- Score tracking
- Game over state with restart option

## Installation Requirements

1. Go 1.21 or later
2. Ebitengine dependencies (automatically handled by Go modules)

### System Dependencies

#### For macOS:

```sh
brew install go
```

#### For Linux (Ubuntu/Debian):

```sh
sudo apt-get update
sudo apt-get install golang-go
sudo apt-get install libgl1-mesa-dev xorg-dev
```

#### For Windows:

1. Download and install Go from [golang.org](https://golang.org/dl/)
2. Install GCC (required for Ebitengine) via [MinGW](http://mingw-w64.org/doku.php) or [MSYS2](https://www.msys2.org/)

## How to Run

1. Clone the repository:

```sh
git clone <repository-url>
cd breakout
```

2. Install dependencies:

```sh
go mod tidy
```

3. Run the game:

```sh
go run main.go
```

## Game Rules

1. Control the paddle to prevent the ball from falling below
2. Break bricks by hitting them with the ball
3. Different colored bricks award different points:
   - Red (top row): 50 points
   - Orange: 40 points
   - Yellow: 30 points
   - Green: 20 points
   - Blue (bottom row): 10 points
4. Game ends when the ball falls below the paddle
5. Maximum possible score: 1,500 points

## Controls

- **Left Arrow**: Move paddle left
- **Right Arrow**: Move paddle right
- **Space**: Restart game after game over

## Game Mechanics

- The ball bounces off walls, the paddle, and bricks
- Ball direction gets slightly randomized when hitting the paddle
- Breaking all bricks is a win condition
- Higher bricks are worth more points but are harder to reach

## Development

The game is built using:

- Go programming language
- [Ebitengine](https://ebitengine.org/) for game development
- Standard library for core functionality

### Project Structure

```
breakout/
├── main.go      # Main game code
├── go.mod       # Go module file
├── go.sum       # Go module checksum
└── README.md    # This file
```

## Contributing

Feel free to fork the repository and submit pull requests. Some ideas for improvements:

- Add sound effects
- Implement different levels
- Add power-ups
- Add high score tracking
- Add different ball speeds or paddle sizes

## License

This project is open source and available under the MIT License.

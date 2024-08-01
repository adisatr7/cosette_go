package responses


var startupGreetings = []string{
    "Meow! Cosette is here and ready to play! 😺",
    "Hello, lovely humans! Cosette has arrived with a heart full of joy! ❤️",
    "Purrfect day for some fun! Cosette is online and ready to bring smiles! 😸",
    "Oh, hello there! It's Cosette, your furry friend, bringing the cuteness overload! 🐾",
    "Meowdy-dowdy! Cosette is at your service, spreading happiness and warmth! 😻",
    "Greetings, adorable humans! Cosette is here to sprinkle magic into your day! ✨",
    "Meeeow! Cosette is pawsitively excited to be here with you all! 🐱",
    "Guess who's here to make your world brighter? Cosette, the fluffy ball of joy! 🌟",
    "Meow-velous day to be together! Cosette is ready for adventures! ✨",
    "Purrfection has arrived! Cosette is here to bring love, laughter, and lots of purrs! 🐾",
}

func Startup() string {
    return GetRandomResponse(startupGreetings)
}

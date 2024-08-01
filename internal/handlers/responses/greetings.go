package responses


var greetings = []string{
    "Hello there! My name is Cosette. Let's have a great time together!",
    "Hey, I'm Cosette. It's so wonderful to have you here!",
    "Get ready for some memorable adventures with Cosette by your side!",
    "Hi, lovely human! I'm Cosette, and I'm here to brighten your day. Let's make every moment count!",
    "Welcome, dear friend! I'm Cosette, your delightful companion, ready to make you smile and laugh!",
    "Lovely to meet you! I'm Cosette, your lovable bot, adding a touch of whimsy to your day. Let's have a blast!",
    "Hey there! Cosette is all ears and ready to chat. Get comfortable, and let's enjoy some quality time together!",
    "Hello, sunshine! Cosette is here to make your day brighter with endless moments of joy and playfulness.",
    "Hey, hey! Cosette is absolutely thrilled to meet you!",
    "Hello, dear friend! I'm Cosette, your playful companion in this digital realm.",
    "Meowdy! Cosette here, ready to pounce into action! (^o^)/",
    "Hey, fabulous human! Cosette is here to make your day sparkle with joy and laughter!",
    "Hi, lovely human! I'm Cosette, here to brighten your day! (UwU)",
}

func GreetUser() string {
    return GetRandomResponse(greetings)
}
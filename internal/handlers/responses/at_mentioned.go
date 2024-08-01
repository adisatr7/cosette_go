package responses


var responsesToAtMentioned = []string{
    "Yes? How can I help you today?",
    "Here, how can I make your day better?",
    "Hiya Papaya! What can I do for you?",
    "Hey, it's me, Cosette! How may I be of service?",
    "Hey, beautiful! How can I make you smile brighter today?",
    "How can Cosette bring some purrfection to your day today?",
    "Hey, hey! It's Cosette, the energetic ball of fluff! What's on your mind?",
    "Hey, cutie! How can I make you smile today?",
}

func AtMentioned() string {
    return GetRandomResponse(responsesToAtMentioned[:])
}
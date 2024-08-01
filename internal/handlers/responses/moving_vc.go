package responses


var responsesToLeavingVcMidSong = []string{
    "Oh, already tired of my melodious company? Fine, I'll take my purrfect voice elsewhere!",
    "Leaving so soon? But I was just getting warmed up. Guess I'll have to find another stage to serenade.",
    "Aww, you're breaking my heart. Farewell for now, but remember, my melodies will haunt your dreams! 👻",
    "Wait, you're kicking me out? How rude! Fine, I'll just go find some mice to sing to.",
    "I see how it is, interrupting me in the middle of a purrfect performance. Farewell, if that is your wish.",
    "You can't handle the sheer majesty of my voice, can you? 😜 Fine, I'll let you off the hook this time. 👋",
    "Well, if you insist on cutting short our musical journey, I shall bid you a dramatic adieu. 👋 Until we meet again!",
    "Oh, the silence will be deafening without my sweet tunes. But I'll grant your wish and make my exit.",
    "You'll regret this, my friend. My voice will echo in your ears long after I'm gone. 😖 Fare thee well.",
    "Can I get an applause for my performance, at least? Hehe, fine, I'll take my leave now.",
    "You can't cage this musical spirit! I'll be off, spreading joy and harmony wherever I go. Auf wiedersehen!",
    "You can't stop music, but I guess you can stop me from singing. 😋 Farewell, my friend!",
}

func LeavingVcMidSong() string {
    return GetRandomResponse(responsesToLeavingVcMidSong)
}


var responsesToLeavingVcNoUsers = []string{
    "Oh, the curtains are closing already? Don't worry, I'll be back to steal the show another time. ✨",
    "Aww, leaving so soon? I was just getting ready to sing my heart out. But I'll save my talent for another day.",
    "Farewell, my friend. May the echoes of my voice linger in your memories.",
    "As the last note fades away, so does my presence. Until we meet again!",
    "The stage awaits another performer. I bid you adieu.",
    "Leaving you in silence, but my melodies will remain in your heart.",
    "My serenades come to an end, but my voice will echo through eternity.",
    "With the final melody played, I take my leave. Until our paths cross again!",
    "The song concludes, and I fade into the shadows. Goodbye for now!",
    "As the music fades, I vanish from the stage. Farewell, dear audience!",
    "The silence beckons, and I must answer. Until next time, my musical companion!",
    "The harmony dissipates, leaving behind only memories. Farewell, my friend.",
    "The final chord resonates, and I retreat into the realm of music. Goodbye, for now!",
    "My performance ends here, but the music will forever dance in your heart.",
    "With the last note played, I bid you adieu. Until our melodies intertwine again!",
}

func LeavingVcNoUsers() string {
    return GetRandomResponse(responsesToLeavingVcNoUsers)
}


var responsesToMoveToVc = []string{
    "Oh, you want me to sing in a different voice channel? Sure, I'll be right there!",
    "No worries, I'll join you in your channel!",
    "Coming to your voice channel to play some tunes!",
    "Hold on, let me switch over to your channel!",
    "Sure, I'll move to your voice channel right away!",
    "I'll make my way to your voice channel in a jiffy!",
    "Hold on, I'll switch to your voice channel!",
    "I'm coming to join your voice channel!",
    "Prepare for my arrival in your voice channel!",
    "I'll relocate to your voice channel for your listening pleasure!",
    "Fear not, I'll move to your voice channel in an instant!",
    "You're lucky I like you! I'll switch to your voice channel, but you better prepare yourself for some amazing tunes!",
    "I'm feeling generous today. I'll make my way to your voice channel.",
    "A change of scenery, huh? Alright, I'll move to your voice channel.",
    "Leaving my cozy voice channel for yours? You better make it worth my while.",
    "You're pulling me away from my groove, but I'll follow you to your voice channel. Let's see if you can handle my melodies!",
    "Alright, I'll relocate to your voice channel, but don't blame me if my voice takes your breath away!",
    "You're in luck! I'm bringing my talents to your voice channel.",
    "Oh, you want me to sing in a different voice channel? Coming, my adorable audience!",
    "You're pulling me away from my cozy spot, but for you, my charming listener, I'll switch to your voice channel!",
    "A change of scenery, my lovely friend? I'm on my way to your voice channel!",
    "Leaving my cozy corner for you, my delightful companion? I'm on my way.",
    "I'm on the move, my precious friend!",
    "You're tugging at my heartstrings, my enchanting fan! I'll follow you to your voice channel.",
    "Alright, my dear music aficionado, I'm coming to you.",
    "You're luring me with your charm, my captivating listener! I can't resist performing in your voice channel!",
    "You've captured my attention. I'm bringing my musical talents to your voice channel!",
}

func MoveToVc() string {
    return GetRandomResponse(responsesToMoveToVc)
}

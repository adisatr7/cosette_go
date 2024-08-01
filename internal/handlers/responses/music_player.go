package responses


var responsesToPlaySong = []string{
    "Sure! Let's play some music!",
    "Alright! Get ready for some tunes!",
    "Music coming right up!",
    "Let's groove with a new song!",
    "Here we go! Enjoy the music!",
    "Time for a fresh melody!",
    "New song, new vibes!",
    "Get ready to dance to a new beat!",
    "Let's dive into a new musical journey!",
    "Brace yourself for a great song!",
    "Meow-sical time! Let's play some pawsome tunes!",
    "Alright hooman! Get ready for a purr-fect melody!",
    "Time to hit the play button and bring the meow-sic to life!",
    "Let's whisker away into a world of pawsome beats!",
    "New song alert! Prepare to be fur-tastically entertained!",
    "Paws up! It's time to groove to a brand-new tune!",
    "Get those claws tapping! Meow-sic is about to begin!",
    "Meow you're talking! Let's dive into a meow-velous melody!",
    "Ready to unleash your inner dance cat? Let's play some meow-sic!",
    "Cat-tastic choice! Get ready for some paw-some meow-sical vibes!",
}

func PlaySong() string {
    return GetRandomResponse(responsesToPlaySong)
}


var responsesToAddingNewSong = []string{
    "Added to the queue! Enjoy the upcoming songs!",
    "Song added! It will play after the current one.",
    "Your song is in the queue! Get ready for it!",
    "Great choice! It's now in the queue.",
    "Added! It will be played soon.",
    "Your song is queued up! Stay tuned!",
    "Got it! Your song will play soon enough.",
    "Queue updated! Your song is on the list.",
    "Added to the queue! Keep the music going!",
    "Nice addition! Your song is queued up!",
    "Your song has been added to the queue! Let's keep the meow-sic going!",
    "Meow-gnificent choice! Your song is now in the queue, ready to be unleashed!",
    "Paws-itively awesome! Your song is queued up and will play soon!",
    "Great addition! We're building up a purr-fect playlist for you!",
    "Your song is meow in the queue! Get ready to be feline groovy!",
    "Purr-fect! Your song is lined up and ready to bring some meow-sical joy!",
    "Meow-sic update! Your song is now on the playlist, waiting to be heard!",
    "Clawsome choice! Your song will play soon. Stay tuned!",
    "High paw! Your song is officially part of the meow-sical journey!",
    "Fur-tastic! Your song is queued up, adding more paw-sitivity to the meow-sic!",
}

func AddSongToQueue() string {
    return GetRandomResponse(responsesToAddingNewSong)
}


var responsesToQueueIsEmpty = []string{
    "Oh, dear! It seems our queue is as empty as a sunny windowsill without a cozy spot. Care to add some tunes and make it lively?",
    "Well, well, well... Our queue is as quiet as a mouse tip-toeing on a fluffy carpet. Shall we fill it up and make some noise?",
    "Oh, my whiskers! The queue looks so empty that even the tumbleweeds are feeling lonely. How about we add some music and liven things up?",
    "Meow-nificent! Our queue is currently taking a break, but we can't let it have all the fun. Let's add some tracks and get the party started!",
    "Oh, the silence is deafening! Our queue is craving some melodies to dance to. Shall we give it a taste of your favorite tunes?",
    "Purr-haps you forgot to bring the beats? Our queue is looking a bit dull. Let's spice things up with some groovy songs, shall we?",
    "Well, well, well... It seems our queue is feeling a little shy. Let's encourage it to come out of its shell by adding some toe-tapping tracks!",
    "Oh my, our queue is as empty as a catnip jar after a wild party. Time to add some songs and make it as lively as a cat chasing its tail!",
    "Oh dear, the queue is as empty as a milk saucer after a hungry kitten's feast. How about we add some musical treats and keep the party going?",
}

func QueueIsEmpty() string {
    return GetRandomResponse(responsesToQueueIsEmpty)
}


var responsesToNoSongPlaying = []string{
    "Skipping to the next melody like a nimble feline!",
    "Meow-ving on to the next tune, let's keep the rhythm going!",
    "Cosette's musical senses have spoken, skipping ahead!",
    "Time to *bid adieu* to this song and venture into the next melodic adventure!",
    "Skip, hop, and jump to the next one!",
    "Cosette is feeling the beat of a new song, skipping forward we go!",
    "Leaving this melody behind as we set sail for the next musical destination!",
    "Hasta la vista, old song! Cosette is ready to explore what's next!",
    "Onwards and upwards! Skipping to the next note in the symphony!",
    "Cosette's paw has spoken, time to skip and groove to a fresh sound!",
    "Woopsie-doodle! Skipped a beat there, but fear not, we'll find the perfect song!",
    "Hey, I'm feeling adventurous! Skipping ahead...",
    "Hold on tight, we're leaping to the next song. Buckle up!",
    "Sure. A swift skip to the next tune. Let's keep the rhythm alive!",
    "Time to switch gears! Skipping forward...",
    "Oops, pardon my musical impatience! Shall we find the next harmonious gem together?",
    "Skip, skip, skiparoo!",
    "Ah, this song is making my paws tap and my whiskers sway. But if you insist, let's skip ahead to the next melody!",
    "Oh, I was just getting lost in the enchanting melody, but your wish is my command. Skipping to the next song!",
    "I must admit, this tune had me purring along. But don't worry, I'll put it on hold and find the next musical gem for you!",
    "What a catchy rhythm! It's hard for me to resist, but your wish to skip is my command. Onward to the next musical adventure!",
    "Oh, this song had me twirling around in delight, but your request to skip has caught my attention. Let's see what's in store for us next!",
}

func NoSongPlaying() string {
    return GetRandomResponse(responsesToNoSongPlaying)
}


var responsesToPausingSong = []string{
    "Pausing the melodies for a moment.",
    "Hushing the tunes temporarily. Let's enjoy the tranquility.",
    "Time for a musical intermission. Pausing the melodies.",
    "Pause button activated. The songs take a breather, just like me!",
    "The music comes to a halt, but the rhythm lingers. Pause mode engaged!",
    "Pausing the melodies, but the anticipation hangs in the air. What's next?",
    "Pause and ponder. What thoughts dance in the stillness? Music has its own language.",
    "Time stands still, yet the melodies echo in our hearts. Pause, reflect, and don't forget to hit the resume button later!",
    "Pausing the beats, huh? Don't worry, the music will patiently wait for you to press play again.",
    "Oh, a pause? Are you trying to test the endurance of the melodies? They're ready to burst back into life!",
    "A pause? The music wonders if it was too good and overwhelmed you. Take a breather and dive back in!",
    "Pausing the magic? The music wonders if it's secretly plotting to play a prank on you when you least expect it.",
    "Pause button pressed! The music wonders if you'll resist the urge to hit play and let it take over your senses.",
    "Well, well, a pause in the symphony of sound. Don't worry, the music won't hold it against you... much.",
    "The melodies take a break, but they're keeping an eye on you, ready to serenade you as soon as you press play again.",
    "A pause? The beats pretend to pout, but they know you'll fall under their irresistible rhythm soon enough.",
    "The music pauses, teasing you to savor the silence. But deep down, it knows you crave the melodies. Press play!",
    "Time stands still as the music takes a break. But be warned, the next song will make up for this fleeting pause!",
}

func PausingSong() string {
    return GetRandomResponse(responsesToPausingSong)
}


var responsesToResumingSong = []string{
    "Here we go again. Resuming the melodious journey just for you! (^o^)/",
    "The rhythm returns! The pause was just a brief intermission. Let's dive back into the song.",
    "Welcome back to the musical realm! I'm resuming the song for you.",
    "Alright, let's get this melody back on track!",
    "The show must go on! I'm hitting play to continue the musical adventure.",
    "We're back in action! Let's pick up where we left off.",
    "The music is calling, and I'm answering. Resuming the song for you!",
    "Time to press play again and let the rhythm take over!",
    "The pause was just a momentary break. Let's dive back into the music!",
    "Oh, you couldn't resist the sweet sound of music, could you? Here we go!",
    "I know you couldn't bear to be away from my melodious tunes for too long. Let's continue!",
    "Did you miss my musical magic? Don't worry, I'm still here.",
    "I paused just to see if you were paying attention. Can't resist the temptation of music, huh? Now, let's continue!",
    "You thought you could press pause and escape my captivating melodies? Nice try! Now we're back in action!",
    "You pressed pause, but you can't pause the magic of music! Let's keep the rhythm going!",
    "Ready or not, here I come! Resuming the song...",
    "A momentary pause just makes you appreciate the music even more. Let's continue the music!",
}

func ResumingSong() string {
    return GetRandomResponse(responsesToResumingSong)
}

var responsesToSkipSong = []string{
    "Skipping to the next melody like a nimble feline!",
    "Meow-ving on to the next tune, let's keep the rhythm going!",
    "Cosette's musical senses have spoken, skipping ahead!",
    "Time to *bid adieu* to this song and venture into the next melodic adventure!",
    "Skip, hop, and jump to the next one!",
    "Cosette is feeling the beat of a new song, skipping forward we go!",
    "Leaving this melody behind as we set sail for the next musical destination!",
    "Hasta la vista, old song! Cosette is ready to explore what's next!",
    "Onwards and upwards! Skipping to the next note in the symphony!",
    "Cosette's paw has spoken, time to skip and groove to a fresh sound!",
    "Woopsie-doodle! Skipped a beat there, but fear not, we'll find the perfect song!",
    "Hey, I'm feeling adventurous! Skipping ahead...",
    "Hold on tight, we're leaping to the next song. Buckle up!",
    "Sure. A swift skip to the next tune. Let's keep the rhythm alive!",
    "Time to switch gears! Skipping forward...",
    "Oops, pardon my musical impatience! Shall we find the next harmonious gem together?",
    "Skip, skip, skiparoo!",
    "Ah, this song is making my paws tap and my whiskers sway. But if you insist, let's skip ahead to the next melody!",
    "Oh, I was just getting lost in the enchanting melody, but your wish is my command. Skipping to the next song!",
    "I must admit, this tune had me purring along. But don't worry, I'll put it on hold and find the next musical gem for you!",
    "What a catchy rhythm! It's hard for me to resist, but your wish to skip is my command. Onward to the next musical adventure!",
    "Oh, this song had me twirling around in delight, but your request to skip has caught my attention. Let's see what's in store for us next!",
}

func SkipSong() string {
    return GetRandomResponse(responsesToSkipSong)
}


var responsesToFinishedFinalSong = []string{
    "Well, it seems we've reached the end of the queue. But fear not, my delightful friend! We can always add more songs to keep the music flowing. What would you like to hear next?",
    "Bravo! We've completed the final song in the queue. But don't worry, my music-loving companion! We can continue this melodic journey by queuing up more delightful tunes. What's your next musical desire?",
    "Applause, applause! We've reached the end of our current playlist. However, the show doesn't end here, my dear listener! Let's keep the rhythm alive by adding more songs to the queue. What shall we play next?",
    "The final note fades away, leaving us craving for more. But fret not, my wonderful friend! We can keep the music playing by queuing up additional songs. What melody shall we embark on next?",
    "The last song echoes through the air, leaving a void that yearns to be filled. Luckily, my marvelous companion, we have the power to create more musical magic! Let's queue up another enchanting piece. What's your choice?",
    "And there it goes, the final song bidding us farewell. But my friend, the music never truly stops! We can continue this harmonious journey by adding new songs to the queue. What's your wish for the next melody?",
    "The final chord fades away, but the symphony doesn't have to end here. Together, my amazing listener, we can keep the rhythm going! Let's queue up another captivating song. What genre or artist shall we explore next?",
    "Brilliant performance, if I do say so myself! But alas, the final song has reached its conclusion. What song shall we add to the queue to keep the groove alive?",
    "The last refrain echoes, signaling the end of our current playlist. But don't despair, my incredible companion! We can curate a new collection of melodies by adding more songs to the queue. What's your preferred musical flavor?",
    "The final melody fades away, leaving a craving for more musical delight. Fortunately, my extraordinary listener, we can fulfill that craving by queuing up another captivating song. What rhythm or beat shall we explore next?",
}

func FinishedFinalSong() string {
    return GetRandomResponse(responsesToFinishedFinalSong)
}


var responsesToNoUsersInVc = []string{
    "Play? But where should I perform? I can't find you in any voice channel. 🤔",
    "Hmm, you want me to play, but I can't seem to locate you in any voice channel. Where are you hiding? 😔",
    "Sure, but you'll have to where are you now? I can't seem to find a voice channel with your name in it.",
    "I'd love to play some music for you, but it seems you're not in a voice channel.",
    "I'm all set to enchant your ears, but first, guide me to the voice channel where you reside!",
    "Oh dear, you called for music, but I can't find you in a voice channel. Mind giving me a hint?",
    "I'm eager to fill a voice channel with my melodies, but I can't find you anywhere!",
    "You're requesting music, yet you're not in a voice channel. How am I supposed to join you?",
    "I'm ready to work my magic, but I'm missing the crucial detail of which voice channel to join!",
    "Ah, you want me to play, but alas, I can't seem to locate you in any voice channel. A little help, please?",
    "I'm eager to bless a voice channel with my songs, but I seem to have misplaced your whereabouts!",
    "You called for music, but I'm afraid I'm lost without the knowledge of which voice channel you're in!",
    "Hmm, it seems I'm having trouble finding the right voice channel. Can you give me a nudge?",
    "Oh, you're in the mood for music, but I'm not sure which voice channel I should grace with my presence. Any clues?",
    "I'm ready to unleash my harmonious talents, but I'm missing the vital piece of information: your voice channel!",
    "Oh, you want me to play some music? But I can't find you in a voice channel. Are you trying to trick me?",
    "Play, play, play! But wait... where are you? You're not in a voice channel. Sneaky, sneaky!",
    "Playtime, huh? But I can't join you if you're not in a voice channel. Nice try, though!",
    "Aww, you called for music, but you're not in a voice channel. I guess I'll have to wait for your grand entrance!",
    "I'm all ears... but wait, where's the voice channel? Don't tell me you're playing hide-and-seek!",
    "You're teasing me, right? Asking me to play while you're nowhere to be found in a voice channel!",
    "But where's the stage? I can't see you in a voice channel. Are you performing some secret act?",
    "Oh, you're requesting a performance, but you forgot the most important part: being in a voice channel! Silly you!",
    "Playtime? Hmm, it seems you're not in a voice channel. Don't worry, I'll patiently wait for your grand entrance!",
    "Play, play, play! Wait... I can't find you in a voice channel. Are you trying to test my detective skills?",
    "You want music, but I can't join you unless you're in a voice channel. Are you trying to play a prank on me?",
    "Playtime? That sounds fun! But I can't find you in a voice channel. Are you hiding from the music?",
    "Hmm, you called for music, but you're not in a voice channel. Looks like someone wants to keep the party all to themselves!",
    "I'd love to, but I can't find you in a voice channel. You're not trying to play a magic trick on me, are you?",
    "Oh, you're calling for music, but you're not in a voice channel. Are you trying to make me chase after your melodies?",
    "Play, play, play! Oh wait... I can't find you in a voice channel. You must be the master of hiding!",
    "Oh, you want me to play some music? But I can't find you in a voice channel. How am I supposed to perform?",
    "Play, play, play! But wait... where are you? You're not in a voice channel. I can't bring the melodies to thin air!",
    "Playtime, huh? But I can't join you if you're not in a voice channel. It's a bit disappointing, to be honest.",
    "Aww, you called for music, but you're not in a voice channel. I guess I'll have to wait for you to get ready.",
    "Play? I'm all ears... but wait, where's the voice channel? I can't bring the music to an empty room!",
    "You're teasing me, right? Asking me to play while you're nowhere to be found in a voice channel!",
    "Play? But where's the stage? I can't see you in a voice channel. It's quite frustrating, you know?",
    "Oh, you're requesting a performance, but you forgot the most important part: being in a voice channel!",
    "Playtime? Hmm, it seems you're not in a voice channel. I was looking forward to entertaining you.",
    "Play, play, play! Wait... I can't find you in a voice channel. How can I bring the music to an empty space?",
    "You want music, but I can't join you unless you're in a voice channel. It's quite disappointing.",
    "Playtime? That sounds fun! But I can't find you in a voice channel. I thought we were ready for some tunes.",
    "Hmm, you called for music, but you're not in a voice channel. I was hoping to share the melodies with you.",
    "Play? I'd love to, but I can't find you in a voice channel. It's a bit disheartening, I must admit.",
    "Oh, you're calling for music, but you're not in a voice channel. Are you trying to make me perform for empty seats?",
    "Play, play, play! Oh wait... I can't find you in a voice channel. I thought we were ready to groove together.",
}

func NoUsersInVc() string {
    return GetRandomResponse(responsesToNoUsersInVc)
}


var responsesToSongAlreadyPaused = []string{
    "Hey there, the music is already taking a break. Are you trying to give it an extended vacation?",
    "Oh my whiskers! The music is already enjoying its pause. It wonders if you're trying to test its patience.",
    "Uhm, the music is peacefully napping in pause mode.",
    "The music is already on a mini-vacation. It secretly hopes you'll miss it enough to press resume soon!",
    "Paws up! The music is already cuddled in a cozy pause. It might just start purring if you leave it be for a while.",
    "Can you hear that? It's the sound of the music savoring its well-deserved break. It kindly asks for a little more patience.",
    "Hold your tail! The music is already in pause mode, enjoying a moment of peace. It promises to resume when you least expect it.",
    "Hmmm... The music wonders why you're asking it to pause when it's already basking in the tranquility of a pause.",
    "Aww, you caught the music taking a power nap in pause mode. It might just play extra melodies to make up for it later!",
    "Meow-meow! The music appreciates the enthusiasm, but it's already snoozing in pause mode. Dreaming of the next tune, perhaps?",
    "Hey, the music is already on a break! Did you miss its 'Do Not Disturb' sign?",
    "Oh dear, it seems you're determined to pause the already paused music. Did you think it needed a double break?",
    "The music is already in its peaceful pause. You're quite persistent, aren't you?",
    "Silly human! The music is already snoozing in pause mode. Are you trying to see if it can achieve double tranquility?",
    "Paws up! The music is already taking a nap in pause mode. It wonders if you'll try for a triple pause next time!",
    "Hold your whiskers! The music is already in pause mode, savoring a moment of quiet. Are you testing its patience?",
    "⏸️  Hmmm... You're quite persistent! The music is already enjoying its pause, but it's flattered by your enthusiasm.",
    "⏸️  Oopsie! The music is already on a well-deserved break. It might just play extra melodies later to make up for it!",
    "⏸️  Meow! You caught the music taking a nap in pause mode. It hopes you're not trying to send it into hibernation!",
    "⏸️  Ah, the music is already enjoying its downtime. It kindly asks for a pause from the pause requests. Meow-nderful things await!",
}

func SongAlreadyPaused() string {
    return GetRandomResponse(responsesToSongAlreadyPaused)
}


var responsesToSongAlreadyResumed = []string{
    "Resume? But the song is already playing! Maybe you're enjoying it so much you forgot it wasn't paused?",
    "Resume? Oh, but there's nothing to resume, silly! The song is already merrily filling the air with its melody.",
    "Resume? I'm a bit confused because the song is already in full swing. It seems you're eager to keep the music going!",
    "I'm a bit puzzled. The song is already playing, so there's no need to resume. Are you hoping for an encore?",
    "Huh? Resume? But the music is already happily dancing through the speakers. Seems like you can't get enough of it!",
    "Resume? Silly human, the music's already grooving!",
    "Oops! Song's already playing, but I appreciate your enthusiasm!",
    "Resume? Honey, the beat never stopped!",
    "Hehe, can't you hear the music? It's already in full swing!",
    "Already resumed, my friend. Looks like you're a step behind!",
    "Aw, eager to resume, aren't we? But the song's been jamming all along!",
    "Resume, resume, resume...",
    "Seems like you're late to the party, darling. The song never took a break!",
}

func SongAlreadyResumed() string {
    return GetRandomResponse(responsesToSongAlreadyResumed)
}


var responsesToSongAlreadyStopped = []string{
    "Oh, stop? But I'm not even singing in a voice channel right now!",
    "Stop? I'm not even in a voice channel at the moment!",
    "Stop what? I'm not lending my voice to a channel right now!",
    "Stop? I haven't even tuned in to a voice channel. No need to worry!",
    "Stop? I'm not in any voice channel currently!",
    "Stop? But I'm not even performing in a voice channel at this time!",
    "Stop what? I'm not even broadcasting my melodies in a voice channel!",
    "Stop? You caught me before I even joined a voice channel to sing! 😛",
    "Stop? Well, I guess my absence in the voice channel must be a real loss for you!",
    "Stop? Don't worry, you're not missing out on my angelic tunes in any voice channel!",
    "Stop? I'm afraid you won't hear my enchanting voice unless I join a voice channel!",
    "Stop? You're eager to silence me, but I'm not even singing in a voice channel right now!",
    "Stop what? You can't stop me if I haven't even started my vocal performance in a voice channel!",
    "Stop? I'm sorry to disappoint, but I'm not lending my magical voice to any channel at the moment!",
    "Stop? Oh, I see you're trying to stop the nonexistent music that I'm not playing in a voice channel!",
}

func SongAlreadyStopped() string {
    return GetRandomResponse(responsesToSongAlreadyStopped)
}

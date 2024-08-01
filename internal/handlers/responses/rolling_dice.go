package responses


var responsesToRollingDice = []string{
    "Prepare for an epic roll, like a mighty dragon's breath weapon!",
    "Let's embark on a dice-rolling quest worthy of a heroic adventurer!",
    "In the realm of dice, our fate shall be decided! Let's roll like true adventurers!",
    "As the dice tumble, let's write our own legend in the annals of TTRPG history!",
    "Summon your inner mage and let the dice cast spells of fortune upon us!",
    "Brace yourself for a roll as legendary as the tales sung by bards in taverns!",
    "In the realm of numbers and dice, may we find glory and epic loot!",
    "Let the dice be our guide through dungeons, dragons, and daring encounters!",
    "Together, we shall defy the odds and weave tales of triumph through our rolls!",
    "As the dice dance upon the virtual table, let the magic come alive!",
    "Hold your breath and cross your fingers! The roll's success may depend on it!",
    "Let's hope this roll doesn't awaken any ancient curses... or does it?",
    "May the dice be ever in your favor... or perhaps not? Let's find out!",
    "Beware! A mischievous pixie might have whispered some chaos into this roll!",
    "You never know what surprises await with a roll of the dice. Brace yourself!",
    "A word of caution: be careful what you wish for when rolling the dice... or not!",
    "The roll is about to happen... *knocks on wood* Let's hope for the best!",
    "Prepare for the roll of a lifetime! Will it be a legendary success or a hilarious failure? Let's roll!",
}

func RollingDice() string {
    return GetRandomResponse(responsesToRollingDice)
}

# Scala, Swift and Kotlin
## or expanding my skillset, part 1

I am an IT-consultant in Stockholm, which means I mainly work in JavaScript or TypeScript, as well as Java. There are exceptions to this rule, I've built backends using Kotlin as well as built cross platform apps in Flutter. For my personal projects I tend to reach for things that I'm curious about, most recently go for this website as well as using Flutter to build StateTrak.

So recently I started looking at gaps in my skillset, what was I missing for the type of work I do? And the one thing I have never built from scratch is an iOS app. I have built plenty of React Native or Flutter apps, and occassionaly touched some Swift code to extend some functionality or implement some missing feature. My first real job in development was also as an Android Developer, at the time writing Java code. So I'd ideally like to add Kotlin to my Android skillset as well. 

I decided to make a clone of Wordle, a very popular webapp which has players guess a word everyday. The player gets 6 guesses and with every guess they are given information on the position of the letters. If a letter is in the correct place it is colored green, if the letter is in the word but not in the correct place it is colored purple. Whenever a letter is guessed and it does not appear in the word it is faded in the keyboard. The game ends when the player has guessed the word or spent their guesses without figuring it out.

It's a fairly simple game loop so I decided it would be a great way to test the waters with Swift. So I read up on how to write a UI in Swift. SwiftUI seems to be the latest and greatest, there are some limitations on which iOS versions it supports, but as this is more of a learning experience than a "real" project so I decided to use it. And it's pretty easy, a lot of work seems to have gone into making things pretty intuitive, just look at this code for rendering a grid of textboxes.

Suffice to say I found it quite enjoyable to write a UI like this. Next up is the controller, which handles state. And again it's very developer friendly, just inherit the `ObservableObject` class and expose your state variables using the `@Published` decorator. This makes sure the UI updates when these variables change. 

Now with these two concepts in hand, all that was left was connecting them. Import the controller into the view and initialize it! (I'm skipping a bunch of logic, but as I said earlier, the game is pretty simple).

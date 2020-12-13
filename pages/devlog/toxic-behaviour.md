# Toxic behaviour
## or being an asshole

You have just closed your computer, you start making dinner when your phone vibrates. You see it's a message in Slack and you can't resist a peek. It's a message in a group chat for the developers of the project you are currently working on. It's a message commenting on a line of code you wrote a month ago, saying that "this endpoint doesn't conform to the standard we set up". 
"Weird" you think and look in the file, and it looks like it does the same things all the other endpoints. So you ask what standard, and get an explanitation which doesn't make sense when you look at the file. So you pull the latest code from the Repo, and you see that a bunch of changes were made yesterday. There is indeed a new flow and standard for the API controllers, but noone changed the one you made.
So you fix it the day after. And try to let go of the fact that they decided on a new flow without you, and mad it a point to lecture you about not adhering to it. 
Fast forward to a few months later, you spend a few hours debugging and fixing a small issue caused by some generated code. Quick explaination is that a non nullable type was null, crashing out of the call. So you switch out the type for string and parse it where you needed the type. Another dev screengrabs the specific part of the code here you parse the type, saying: "This is why quick fixes are bad, you don't understand the underlying problem, that this type is a string". 
You answer that you switched the type to a string, explaining that the other type was causing an issue when the value was null. You ask for specific feedback on the complaint, you get none and the dev spends 3 hours trying to "fix" your code. You look into the commit messages and sees, he first just dumped your code, only to realize that there was a need for parsing, and adds it in such a complex way that noone would realize where it happens.
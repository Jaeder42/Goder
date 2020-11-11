# Flutter > React Native

## or why 1.0 matters

Recently I was assigned to a new client, where they needed someone to create their new app. I was excited, time to get back to my roots with some app development! The first thing I had to do was choose the tech stack for the app, and I got even more excited. Do I want to build 2 apps, one for each OS? Or do I use a cross platform framework?

That choice was pretty easy, a small company with me as the sole frontend developer. Cross platform it is! So now comes the second choice, do I use React Native or Flutter? (Yes I know there are others, but come on)

I've used React Native professionally before, there is a strong community and Facebook is backing it. It is very close to React, meaning that most React developers have an easy time learning it. It is JavaScript which means that there are a lot of developers who know the language. All in all React Native would be a good choice.

If not for one factor, which is reflected in their versioning. It is still not 1.0, the team has been bumping it slowly up to 0.61. And when that version bumps, it is not uncommon for some of your dependencies to start breaking. Meaning you either stick with the version you got, or spend a day updating all your broken dependencies.

I've only used flutter for hobby projects. There is a stronger community for it now than ever. Google is backing it. It is written in Dart, an application language developed by Google. It is not that foreign to anyone who has written in Java, Kotlin or even Javascript. And it is beyond version 1.0. And I've yet to face an issue where updating Flutter breaks anything for me, not saying it can't happen but I have more faith in a framework who've committed to being released.

Now where I saw JavaScript as a positive for React Native, I see Dart as an even bigger positive for Flutter, at least if you are the one writing the code. Dart is typed, actually typed. Even using TypeScript in React Native leaves you with number as your only number type. Dart has int, double, float the works! (I mean, as well as a bunch more primitives etc. but numbers in JavaScript is a real sore point for me). One downside with Dart is that not a lot of people have used it.

So given all this, type safety, major version and good community support, I choose Flutter. You may not agree, you might yell at me that Google already proved they can't be trusted with frameworks, looking at you angular. You may be an expert in JavaScript and don't think it's worth the extra time learning a new language, which is fair. You may even say that I'm dumb from the start and the only way to build an app is to do it natively.

To these arguments:

1. I think google learnt their lesson on angular, and Flutter seems to be fairly self sufficient.
2. If there's no time, there's no time. But time spent learning isn't time wasted.
3. In this use case, where there is basically no special functionality that can be gained from native apps it is not worth the time to write 2 apps.

Hopefully you found this brain dump helpful if you're ever in the same situation :)

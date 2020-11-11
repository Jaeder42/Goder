# My (almost) year as a startup employee

## or why I left,

This is a long one, I've been saving and writing on this since I quit, partly to write it as well (as I can) as possible, and to look at the story with some distance. Honestly it's still a pretty weird story so read on!

In 2016 I was working as an Android developer in Norrköping, a town in Sweden . I liked my job, liked my colleagues and was planning on staying in Norrköping at least a couple of years.

Then at the start of 2017, because of a family situation, I started looking for jobs in Stockholm. I had some leads, but I told a friend I was looking. He almost immediately connected me with the CTO of the company he was working at, a startup in the center of Stockholm. We quickly set up a meeting, a voice call in Discord, a chat service for gaming. The meeting went well, he browsed my GitHub repos and told me that it “looked like I knew NodeJS”, and asked me about React, I told him I was working in Angular for the current project I was assigned to. He told me he would talk to the CEO and that I should expect a call shortly.

Shortly after that the CEO called me, told me the general goal of the company and a salary proposal. I told him it sounded good and he told me they were interviewing one more person the day after, but that unless he was a “guru” at React, the job would be mine. He told me he would be in touch during the next few days.

I waited a few days and heard nothing, so I sent a message: “What is the status?”. The response came in a call, where the CEO explained that the other developer had indeed been a “guru” at React, and that he could not afford not to hire him. He told me that he wanted to double the development team some time during the summer and that he would be in touch then. A little disappointed I went back to the job search.

### Short wait

Not a month later I get a message from the CTO again. He wants to meet in Stockholm to talk about a job as a backend developer. I accept and we set a date to meet.

I got up early on a Saturday to catch a train to Stockholm. When I arrive at the central station, I meet the CTO and we go to a nearby coffe shop. He says he would like to hire me with the same offer as the month before. We talk about some of the technologies used on the backend, NodeJS in this case. He tells me that the CEO will call me soon again with an official offer.

Indeed I get an official offer with a contract that he tells me to sign and send to him. I read it and after some altering, we agree and I tell him that I will resign from my current position as soon as possible. I send an email to my current boss that night, explaining that I have received an offer and that I am planning on accepting. We set a meeting and she tells me that it is a shame to lose me, but that it is good that I found a job in Stockholm. We set a final date for me, with enough time for me to leave my current projects in a state that someone can take over. I was sent off with well wishing and an expected amount of friendly teasing of working in finance, startups and Stockholm.

### First day

The weekend before I start the new job, I move all my things home to my mother in Stockholm where I will live until we both can find apartments in town. We then visit my grandfather in Norway who had fallen and broken his hip a few days earlier. We get back the day before I start.

The first day I take the bus from the house outside the city, to the subway. This will be my daily routine for some time, about a 40 minute journey if I time the bus and subway just right. I am excited to start the new job and to work with my friend.

I get to the office and meet the CEO in person for the first time. He welcomes me and I get a laptop and a desk, and I start setting up. Greeting new coworkers as I wait for the first standup.

Although being hired as a backend developer, my first assignments are in the frontend. The reason is that there are some things that need to be finished there. I work on some frontend stuff for a few weeks before being set to work in the backend.

### A few months in

We are supposed to go into beta before the summer. We have already pushed it from “during the spring”. There has been continuing miscommunication between CTO, product owner and CEO about the state of our product. As summer vacation approaches for everyone, we realize that we are not reaching a beta state, and we are also without one of the licenses we need. So we plan everyones vacations so that development never stands still.

During the summer I am reporting to the CTO and CEO on a weekly basis. Our backlog is almost empty and as I leave for my vacation I feel good about what we’ve accomplished. I am gone for two weeks.

### Back from vacation

When I return, the backlog is full again, new features that are essential for the product to work and be profitable. It is up to the product owner and CEO to set the scope of the product so we all go along. I am assigned to a service previously built by external consultants. The service is responsible for fetching data from a provider and has been crashing while fetching the data, which has caused some other services to work on old data.

I get to work and realise that is because of a combination of bad streaming and large files. Basically there is a stream created but it is only paused for one object in a set of objects, instead of pausing for each individual write to the database. I do not immediately realise this as the code base is large and confusing. I do however see that the reason for the crash is that the javascript heap is running out of memory. So I explain that the logic in the fetcher needs a rewrite. I am met by a short “the consultants already solved this”. I try to explain that something is obviously eating up the memory.

I finally figure out the exact point of failure and rewrite the stream and database interface logic, which fixes the crashing. This causes the effect that I am now in charge of this service and all new data sources has now become my responsibility. The service becomes a monolith, that handles our custom data as well as third parties’.

### Launch

It is now autumn 2017, we finally have a consumer product, that we launch. We take out ads in some print newspapers, which is a little strange but whatever it’s exposure. It is not enough to drive traffic to our product though. The CEO is unfazed though as we are already dreaming of the next step, making money.

The company shifts focus from consumers to organisations. A potentially more profitable market. The problem being that we had ignored that platform in favour of the consumer product, until now. We have to rewrite all the organisational logic, currently a monolith, to micro services.

### First man out

My friend who got me the job in the first place is looking for other jobs now. He has a few offers and finally accepts one and gives his intent to leave.

The following morning standup, the company is informed that he is leaving. Some people are surprised others not so much. The CEO tells me he wants to sit down and talk to me. We sit down and he tells me he is concerned that I will leave as well because my friend is quitting. I take some offence to the idea that I would leave a job on the sole basis that my friend is no longer working there. I get his usual routine of telling me I’m a great guy and talented. He offers lunch to talk, I accept, thinking it’s just the two of us . The entire development team is taken to lunch, and we are asked to give suggestions on improvements to our work flow.

We talk about needing more clear planning and less micro-managing. We are assured that planning will be taken more seriously from now on, with our product owner removing himself from coding and focusing on product management. There is an initial improvement, but everyone soon falls back to old habits.

### Quitting

One day I am drinking with an old friend and start complaining about my current employment, that we don’t have a plan, we lack structure when working and that we do a lot of work that we ultimately have to redo or just delete because someone has changed their mind. He tells another friend that then messages me a few days later, asking if I would like him to find out if the consultancy company he works for is hiring, I accept. After 10 minutes a recruiter has messaged me, asking if I would like to come and hear more about the company. I accept.

Fast-forward to me having an offer on the table from the new company. I have to tell my current boss, the CEO, that I am quitting. I gather my courage at the end of the day and we sit down and I try to quit. Try, because he asks that I take the weekend to mull it over, which I do, though I have already made up my mind.

That Monday, I think that the CEO will want to talk to me at the end of the day, but he asks if we can talk right after the stand up. We sit down and he asks me if I’ve made up my mind. I says yes, I am accepting the offer from the new company. What follows is a manipulative effort to get me to “decide” to change my mind, “being a consultant sucks”, “all the consultants that I know are trying to work for real companies”, “I think you’re making a huge mistake”. I was trying not to be rude, but the mention that I was making a mistake got me a little riled up. I give my honest opinion about the way we work and the new product we’re building. This triggers him to once more say that I am making a mistake and the new product is going to make so much money, but then he accepts it.

### Not alone

The day after I tell the rest of the company at the morning standup. Some people are surprised, most of the development team is not, some knew and some understand why. This is followed by the CEO giving a speech about how everyone who works there now is going to make an excellent career at the company. Might be coincidence, but it feels directed at me.

Later that day the person who sits next to me, the person that got the original job, he tells me he is also going to quit. He then tells the CEO and now there are 3 people working off their notice time, the person who got me the job, the person who first got that job and me.

### Home stretch

The last 3 months at the company are strange. I am trying to finish up my projects before I leave so I am working hard. One day it is announced that one of the lawyers have an issue with some wording on our site, a problem we had pointed out, but that was ignored. At lunch that day I was discussing with the senior backend developer how to improve the workflow. We finish talking as the CEO and the other developer who’s leaving come into the kitchen. The CEO starts yelling at me and the other developer leaving, that we don’t have respect for the company and that he won’t stand for it even if we’re leaving. A bit confused and angry at the accusation, I ask calmly “Are you thinking of something particular?”. It turns out he has perceived that we reacted to the changes required in a negative way, no one had. Another reason is the way I resigned, that I told him that I did not believe in the new focus of the product or the way we work, he thought that was disrespectful. He then sits down and eats lunch with us, talking about apartment prices in Stockholm. I eat up and leave.

### Professionalism

The last full week I work, I have finished up a service and I need it deployed. The product owner takes on the task of pushing it to production. He then told me it was not working. Strange I thought, as it worked both on my machine (insert joke here) and in our staging environment. Nonetheless I got to debugging the error. As I could not find anything wrong I told him as much. The message that followed was the least professional message I have ever received. Paraphrasing, “you really need to make sure that things work, you have to take responsibility and owership for you services, some of the issues are honestly like handing in an essay without proofreading”. Preceded by a list of unnecessary code changes. I get the sense that he wasn’t going to be reasonable, so I make the changes and resubmit. The only meaningful change I made was to add an error message to an error that could only be caused if a) there was an error in the deployment configurations or b) the service HE had just created used for authentication wasn’t running properly. Some time later the service was up without comment except a tap on the shoulder and a “Good job, it works now”.

### Hindsight

Remember the Vasa, a term used by Bjarne Stroustrup talking about the creation of C++, comparing it to how the regal ship Vasa sank because there were so many changes that the ship ended up topheavy and it capsized. I was reminded of this term during the entire summer and autumn, and also said it aloud more than once. We changed something on a whim, started new development within 5 minutes of the CEO or product owner thinking of it and any due dates we set were either pushed back or lead to some of us working 14 hour days while getting paid for 8.

I wish all my coworkers the best, and I genuinely hope they find a way to both deliver the product they have promised and find a way to do it without going insane.

Do I wish I hadn’t taken the job? No. It was a learning experience, as I told my father when we were talking about this, it wasn’t a mistake to take the job, but it would’ve been to stay.

# No JS
## Can I build a website with no js?
Short answer yes, you're here aren't you? But this just looks like html and css, I hear you telling me from the other side of your the screen. And sure, it is just that.
The real meat of this comes from the path we are taking to end up with "just" html and css. 
I set up some rules for myself in this: It should have no JS and I cant just cheat and write only HTML files. So step one, get a server running, and I can't just grab NodeJS and spin up an express app. I have played around with GoLang before and want to learn more so I grabbed it and read up on runnng a server. As it turns out, it's very easy.
Go also includes built in html templating, which felt too easy in this case, so I'm using the string templates and just loading a html file which includes some escapes for what I want to insert.
The next step was to decide how I want to injest pages. I decided on md files as it already includes some rules for formatting that I could implement. So I read the file and format it to HTML and inject it in my template.


# Node.js heap out of memory
## or the drowning mailman
Let’s talk about a bug I once had the pleasure of debugging. 

There was an automated script that downloaded a 4GB file and wrote the contents to a database. But for some reason it kept crashing. I was tasked with solving the bug.
So let’s break down what was happening. Imagine a mail room, the download is a conveyor belt. There is a man who receives the packages and stops the belt while he reads the package and hands it over to a second man who sorts the package into a shelf. 

The man stopping the belt and reading the package is the stream and the man sorting the packages is the database handler. If the man sorting falls behind the packages stack up and he takes care of them one at a time. 

% mailman.png

Now imagine the man on the conveyor belt can’t see the man sorting the boxes, he only drops the packages into a chute not caring about what tempo he was keeping. If the sorter isn’t as fast as this he will pretty soon be overwhelmed by boxes, until he can’t work any more.

% mailman2.png

This is what was happening in the script I was debugging. Instead of pausing the stream until the database handler had written to the database, the stream only paused to send it to the handler. This meant that everything just piled up in memory until we used up the javascript heap. 
So how come this script was ever even there? Because it worked for most files, that had a max size of around 200-300MB. The size in memory can’t hit the magic 1024MB limit and thus the only time it crashed was when this huge file was downloaded. 
The fix was easy once I figured out what was actually happening.
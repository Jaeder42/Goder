# HTTP 524

## or how to DDoS yourself

A simple deploy, just some frontend stuff, nothing big.

And then chaos, we saw major spikes in our dashboard, one micro service was being requested 900 times per second. wat.
So what's happening? Can anyone see anything else that's behaving weirdly? Ok queues in AWS are piling up and not being handled, a clue! What did we change? Shit, we removed our handling for one of the messages, no worries add it back, redeploy.

Temporary calm, but we are still spiking on the one micro service. Quick check in the code for handling sessions. Shit, we call the micro service basically every time we get any request on any micro service. That might explain why we're spiking, but not why it's happening now and not 2 hours ago.

Write a fix and deploy it, only problem is that we then need to update every micro service. Fine, do it. Okay the spikes are gone but there is still increased traffic to the site, causing strange problems. Fuck, a third party integration says we're killing their DB. Add caching to the request that we send to them, solves that issue at least.

Clean up some of the code, check the dashboard. Everything looks ok, get some sleep.

Wake up, look at the dashboard, everything still looks fine.

Check the database and socket connections, realise that there are people spamming us with about 100 open tabs at a time, sending around 300 requests per second, block IP's that abuse.

Finally everything looks fine.

Nervously monitor the dashboard.

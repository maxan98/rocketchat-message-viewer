# rocketchat-message-viewer
Simple tool which increases admin **functionality** and brings ability to view rocketchat messages from CLI or you can expose your API and use WEB frontend.

##Build
To build backend just cd to the repo and run `go build .`

I don't know how to build and run frontend cuz all frontend is a piece of +hit. And I hope i'll never heard "JS" again in my life.

##Run
- Build
- Edit `run.sh` to point to your VM with RC instance. Use IP or DNS it just a SHH tunnel.
- run ./run.sh {cli-parameters}
  - OR Make ssh tunnel by yourself and jusr exec ./goreadmongo
- Once again: I don't give a hack how to run frontend

Script uses basic name of RC DB with no protection. In case you 4 some reason use different name or password - you'll need to adapt this tool.

Btw to make frontend work properly e.g show message attachments don't forget to switch off triggers for "Protect uploaded files" and "File Upload Json Web Token Secret" in your RC instance.
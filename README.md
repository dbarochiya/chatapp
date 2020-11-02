# ChatApp 

This is a realtime chatbox implementation where multiple users join in a single chat pool. I have worked on this with a motive to learn these cool techs - 

  - backend : [GO](https://golang.org/)
  - frontend : [ReactJs](https://reactjs.org/)
  - communication protocol : [Websocket](https://en.wikipedia.org/wiki/WebSocket#:~:text=WebSocket%20is%20a%20computer%20communications,being%20standardized%20by%20the%20W3C.)

# Reference

- [more about me](https://dhruvbarochiya.com)
- [golang tutorial](https://www.youtube.com/watch?v=YS4e4q9oBaU)
- reference for this project : [tutotial blog](https://tutorialedge.net/projects/chat-system-in-go-and-react/)

# How to run

## backend

Follow these commands to spin up a go server which listens on port `:8080` and broadcasts messages to clients using websocket.

```
  (install go if you haven't already)
  git clone https://github.com/dbarochiya/chatapp.git
  cd chatapp/backend
  go run main.go
``` 

## frontend

This will spin up frontend react server running on `:3000` 
``` 
 (install npm if you haven't already)
  git clone https://github.com/dbarochiya/chatapp.git
  cd chatapp/frontend
  npm install
  npm start 
```

Now go to [localhost:3000](localhost:3000) in browser and TADA !! 

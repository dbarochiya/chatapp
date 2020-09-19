import React, { Component } from 'react';
import './App.css';
import { connect, sendMsg } from './api';
import Header from './components/header/Header';
import ChatHistory from './components/chathistory/ChatHistory'
import ChatInput from './components/chatinput/ChatInput'
import Connect from './components/connect/Connect'

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      chatHistory: [],
    }
  }
  
  // setUsername = (name) =>{
  //   this.setState(prevState => ({
  //     userName : name
  //   }))
  // }
  // connect = (msg) => {
  //   console.log('my test');
  //   this.setState(prevState => ({
  //     chatHistory: [...this.state.chatHistory, msg],
  //   }))
  // }

  componentDidMount() {
    connect((msg) => {
      this.setState(prevState => ({
        chatHistory: [...this.state.chatHistory, msg]
      }))
    });
  }

  send(event) {
    if(event.keyCode === 13) {
      sendMsg(event.target.value);
      event.target.value = "";
    }
  }

  render() {
    return (
      <div className="App">
        <Header/>
        <Connect connect={this.connect} />
        <ChatHistory chatHistory={this.state.chatHistory} />
        <ChatInput send={this.send} />
      </div>
    );
  }
}


export default App;

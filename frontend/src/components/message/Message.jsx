import React, { Component } from 'react'

import "./Message.css";

export class Message extends Component {
  constructor(props) {
    super(props);
    let temp = JSON.parse(this.props.message);
    this.state = {
      message: temp
    };
  }

  render() {
    return (
      <div className="messageBox">
        <div className="userID">{this.state.message.sender} : </div> 
        <div className={(this.state.message.type === 1)? "specialMessage": "message"}>
          {this.state.message.body}</div>
      </div>
    )
  }
}

export default Message;
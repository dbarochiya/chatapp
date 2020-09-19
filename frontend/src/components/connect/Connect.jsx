import React, { Component } from 'react'
import './Connect.css'

export class Connect extends Component {
  constructor(props) {
    super(props);
    this.state = {
      value: '',
      disabled : false};

    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChange(event) {
    this.setState({value: event.target.value});
  }

  handleSubmit(event) {
    console.log('A name was submitted: ' + this.state.value);
    this.setState( {disabled: !this.state.disabled} )
    event.preventDefault();
  }

  render() {
    return (
      <form onSubmit={this.handleSubmit}>
        <label>
          Name:
          <input 
            type="text" 
            value={this.state.value} 
            placeholder = "what's your name"
            onChange={this.handleChange} 
            disabled = {(this.state.disabled)? "disabled" : ""}/>
        </label>
        <input 
          type="submit" 
          value="hit me up ;)"
          disabled = {(this.state.disabled)? "disabled" : ""}/>
      </form>
    );
  }
}

export default Connect

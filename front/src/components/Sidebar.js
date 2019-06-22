import React, { Component } from 'react'
import ListaChats from './ListaChats';
import PropTypes from 'prop-types';

export class Sidebar extends Component {
  getDivStyle = () => {
    return {
      width: '30%',
      minWidth: '175px',
      height: '100%',
      background: 'white',
      zIndex: '0',
      float: 'left',
      position: 'relative',
    }
  }

  render() {
    return (
      <div style={this.getDivStyle()}>
        <ListaChats chats={this.props.chats} myID={this.props.myID} getChat={this.props.getChat} />
      </div>
    )
  }
}

Sidebar.propTypes = {
  myID: PropTypes.number.isRequired,
  getChat: PropTypes.func.isRequired,
}

export default Sidebar

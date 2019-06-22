import React, { Component } from 'react'
import ListaChats from './ListaChats';
import PropTypes from 'prop-types';
import NovoChat from './NovoChat';

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
        <NovoChat chats={this.props.chats} addChat={this.props.addChat} myID={this.props.myID} 
          usuariosAtivos={this.props.usuariosAtivos} />
      </div>
    )
  }
}

Sidebar.propTypes = {
  usuariosAtivos: PropTypes.array.isRequired,
  myID: PropTypes.number.isRequired,
  chats: PropTypes.array.isRequired,
  getChat: PropTypes.func.isRequired,
  addChat: PropTypes.func.isRequired,
}

export default Sidebar

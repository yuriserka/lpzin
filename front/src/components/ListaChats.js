import React from 'react'
import PropTypes from 'prop-types'
import Chat from './Chat';

export class ListaChats extends React.Component {
  render() {
    return this.props.chats.map(
      (chat) => (
        <Chat key={chat.ID} chat={chat} myID={this.props.myID} getChat={this.props.getChat} />
      )
    )
  }
}

ListaChats.propTypes = {
  chats: PropTypes.array.isRequired,
  myID: PropTypes.number.isRequired,
}

export default ListaChats

import React from 'react'
import PropTypes from 'prop-types'
import Chat from './Chat';

export class ListaChats extends React.Component {
  render() {
    if (this.props.chats === null) {
      return <></>
    }
    return this.props.chats.map(
      (chat) => (
        <Chat key={chat.ID} chat={chat} myID={this.props.myID} getChat={this.props.getChat} />
      )
    )
  }
}

ListaChats.propTypes = {
  chats: PropTypes.array,
  myID: PropTypes.number.isRequired,
  getChat: PropTypes.func.isRequired,
}

export default ListaChats

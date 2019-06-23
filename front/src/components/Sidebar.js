import React, { Component } from 'react'
import ListaChats from './ListaChats';
import PropTypes from 'prop-types';
import NovoChat from './NovoChat';
import styled from 'styled-components'

const SideBarDiv = styled.div`
  width: 30%;
  min-width: 175px;
  height: 100%;
  background: white;
  z-index: 0;
  float: left;
  position: relative;

  ::-webkit-scrollbar {
    width: 6px !important;
    height: 6px !important;
  }
  ::-webkit-scrollbar-thumb {
    background-color: lightgray;
  }
  overflow-x: hidden;
  overflow-y: scroll;
`

export class Sidebar extends Component {
  render() {
    return (
      <SideBarDiv>
        <ListaChats chats={this.props.chats} myID={this.props.myID} getChat={this.props.getChat} />
        <NovoChat chats={this.props.chats} addChat={this.props.addChat} myID={this.props.myID}
          usuariosAtivos={this.props.usuariosAtivos} />
      </SideBarDiv>
    )
  }
}

Sidebar.propTypes = {
  usuariosAtivos: PropTypes.array.isRequired,
  myID: PropTypes.number.isRequired,
  chats: PropTypes.array,
  getChat: PropTypes.func.isRequired,
  addChat: PropTypes.func.isRequired,
}

export default Sidebar

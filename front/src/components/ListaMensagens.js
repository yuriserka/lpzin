import React from 'react';
import Mensagem from './Mensagem';
import PropTypes from 'prop-types'
import DefaultPage from './DefaultPage';
import styled from 'styled-components';

const SideBarDiv = styled.div`
  ::-webkit-scrollbar {
    width: 6px !important;
    height: 6px !important;
  }
  ::-webkit-scrollbar-thumb {
    background-color: #92186f;
  }
  height: 500px;
  overflow-x: hidden;
  overflow-y: scroll;
  float: left;
  width: 70%;
`

class ListaMensagens extends React.Component {
  getHSL = (u) => {
    const rnd = (min, max) => {
      return Math.floor(Math.random() * (max - min + 1) + min);
    }

    const gerarCor = () => {
      return {
        hue: rnd(0, 360),
        sat: rnd(70, 100),
        lum: rnd(25, 35),
      }
    }

    u.Cor = gerarCor()
  }

  getMensagens = () => {
    if (this.props.chatAtual === null) {
      return <DefaultPage />
    }
    this.props.chatAtual.Usuarios.map((u) => {
      this.getHSL(u)
      return u
    })

    return this.props.chatAtual.Mensagens.map(
      (msg) => (
        <Mensagem key={msg.ID} chatAtual={this.props.chatAtual} mensagem={msg}
          myID={this.props.myID} ehGrupo={this.props.ehGrupo} />
      )
    )
  }

  render() {
    return (
      <SideBarDiv>
        {this.getMensagens()}
      </SideBarDiv>
    )
  }
}

ListaMensagens.propTypes = {
  myID: PropTypes.number.isRequired,
  chatAtual: PropTypes.object,
  ehGrupo: PropTypes.bool.isRequired,
}

export default ListaMensagens;
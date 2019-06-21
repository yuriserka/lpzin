import React from 'react'
import PropTypes from 'prop-types'
import styled from 'styled-components'

const LeftHeader = styled.div`
  height: 100%;
  float: left;
  width: 31%;
  &:hover {
    background: #5a0e27;
  }
`;

const RightHeader = styled.div`
  height: 100%;
  float: left;
  width: 69%;
  &:hover {
    background: #5a0e27;
  }
`;

export class Header extends React.Component {
  getChatInfo = () => {
    const antiTextoLongo = () => {
      return {
        display: 'inline-block',
        flexGrow: '1',
        position: 'relative',
        textOverflow: 'ellipsis',
        whiteSpace: 'nowrap',
        overflow: 'hidden',
        width: '100%',
        marginLeft: '20px',
        marginTop: '5px',
      }
    }
    const size = this.props.chatAtual.Usuarios.length
    const ehGrupo = size > 2 ? true : false
    if (ehGrupo) {
      return (
        <div>
          <span style={antiTextoLongo()}>
            {this.props.chatAtual.Nome + ' '}
          </span>
          <span style={{ fontSize: '12px', marginLeft: '15px', }}> {size} membros</span>
        </div>
      )
    }
    
    const usrs = this.props.chatAtual.Usuarios
    const outroUsuario = usrs.find(usuario => usuario.ID !== this.props.usuarioAtual.ID)
    return (
      <div>
        <span style={antiTextoLongo()}> {this.props.chatAtual.Nome} </span>
        <span style={{ fontSize: '12px', marginLeft: '15px', }}> {outroUsuario.UltimaVez}</span>
      </div>
    )
  }

  renderizarInfoConversa = () => {
    if (this.props.chatAtual === null) {
      return
    }
    return (
      <RightHeader>
        {this.getChatInfo()}
      </RightHeader>
    )
  }

  render() {
    return (
      <div style={{ color: 'white', background: '#921840', height: '48px', margin: '0' }}>
        <LeftHeader>
          <div>
            <span style={{ marginLeft: '10px' }}>
              Telezap is real
          </span>
          </div>
        </LeftHeader>
        {this.renderizarInfoConversa()}
      </div>
    )
  }
}

Header.propTypes = {
  chatAtual: PropTypes.object,
  usuarioAtual: PropTypes.object,
}

export default Header

import React from 'react'
import PropTypes from 'prop-types'
import styled from 'styled-components'

const LeftHeader = styled.div`
  height: 100%;
  float: left;
  width: 30%;
  &:hover {
    background: #5a0e27;
  }
`;

const RightHeader = styled.div`
  height: 100%;
  float: left;
  width: 70%;
  &:hover {
    background: #5a0e27;
  }
`;

export class Header extends React.Component {
  antiTextoLongoStyle = () => {
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

  grupoInfo = () => {
    const size = this.props.chatAtual.Usuarios.length
    return (
      <div>
        <span style={this.antiTextoLongoStyle()}>
          {this.props.chatAtual.Nome + ' '}
        </span>
        <span style={{ fontSize: '12px', marginLeft: '15px', }}> {size} membros </span>
      </div>
    )
  }

  outroUsuarioInfo = () => {
    const usrs = this.props.chatAtual.Usuarios
    const outroUsuario = usrs.find(usuario => usuario.ID !== this.props.usuarioAtual.ID)
    return (
      <div>
        <span style={this.antiTextoLongoStyle()}> {this.props.chatAtual.Nome} </span>
        <span style={{ fontSize: '12px', marginLeft: '15px', }}> {outroUsuario.UltimaVez} </span>
      </div>
    )
  }

  getChatInfo = () => {
    return this.props.ehGrupo === true ? this.grupoInfo() : this.outroUsuarioInfo()
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
      <div style={{ color: 'white', background: '#921840', height: '48px', cursor: 'pointer' }}>
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
  usuarioAtual: PropTypes.object.isRequired,
  ehGrupo: PropTypes.bool,
}

export default Header

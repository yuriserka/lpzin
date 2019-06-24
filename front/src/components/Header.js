import React from 'react'
import PropTypes from 'prop-types'
import styled from 'styled-components'
import PopUp from 'reactjs-popup'
import ImagemPerfil from './ImagemPerfil';

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

  getUsuarioDivStyle = () => {
    return {
      background: 'lightgray',
      borderRadius: '10px',
      height: '60px',
      margin: '5px',
      // width: '50%',
      // display: 'flex',
      // justifyContent: 'center'
    }
  }

  renderizarInfoConversa = () => {
    if (this.props.chatAtual === null) {
      return
    }
    return (
      <PopUp modal closeOnEscape closeOnDocumentClick trigger={
        <RightHeader>
          {this.getChatInfo()}
        </RightHeader>
      } contentStyle={{ margin: '8px 8px 0 75.5%', padding: '0', height: '97.1%', }}>
        <>
          <div style={{ display: 'flex', flexWrap: 'wrap', justifyContent: 'center' }}>
            <ImagemPerfil encoded={true} obj={this.props.chatAtual} h={100} w={100} ehGrupo={this.props.ehGrupo}/>
          </div>
          <h2 style={{ color: 'black', textAlign: 'center', background: 'whitesmoke' }}>
            {this.props.chatAtual.Nome}
          </h2>
          <h3 style={{ color: 'black', textAlign: 'center' }}>Membros:</h3>
          {
            this.props.chatAtual.Usuarios.map(ua => (
              <div key={ua.ID} style={this.getUsuarioDivStyle()}>
                <div style={{ display: 'flex' }}>
                  <ImagemPerfil obj={ua} h={40} w={40} encoded={true} ehGrupo={this.props.ehGrupo} />
                  <span style={{ padding: '22px 15px 10px', color: 'black' }}>
                    {ua.ID === this.props.usuarioAtual.ID ? 'Você' : ua.Nome}
                  </span>
                </div>
              </div>
            ))
          }
        </>
      </PopUp>
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

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
  getNameStyle = () => {
    return {
      display: 'inline-block',
      flexGrow: '1',
      position: 'relative',
      textOverflow: 'ellipsis',
      whiteSpace: 'nowrap',
      overflow: 'hidden',
      width: '100%',
    }
  }
  
  getChatInfo = () => {
    const ehGrupo = this.props.chatAtual.Usuarios.length > 2 ? true : false
    if (ehGrupo) {
      return (
        <div>
          <span style={this.getNameStyle()}>
            {this.props.chatAtual.Nome + ' '}
            <span style={{fontSize: '12px'}}> {this.props.chatAtual.Usuarios.length} membros</span>
          </span>
        </div>
      )
    }
    return <span> {this.props.chatAtual.Nome} </span>
  }

  renderizarInfoConversa = () => {
    if (this.props.chatAtual === null) {
      return
    }
    return (
      <div>
        <RightHeader>
          {this.getChatInfo()}
        </RightHeader>
      </div>
    ) 
  }

  render() {
    return (
      <div style={ { color: 'white', background: '#921840', height: '48px', margin: '0' } }>
        <LeftHeader>
          Telezap is real
        </LeftHeader>
        {this.renderizarInfoConversa()}  
      </div>
    )
  }
}

Header.propTypes = {
  chatAtual: PropTypes.object,
}

export default Header
